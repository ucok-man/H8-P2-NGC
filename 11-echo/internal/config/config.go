package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/structs"
	"gorm.io/gorm"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	dbconn *gorm.DB

	// address
	ApiPort int `mapstructure:"API_PORT"`

	// database
	DB_Name     string `mapstructure:"DB_NAME"`
	DB_User     string `mapstructure:"DB_USER"`
	DB_Password string `mapstructure:"DB_PASSWORD"`
	DB_Url      string `mapstructure:"DB_URL"`
	DB_Port     string `mapstructure:"DB_PORT"`

	// DB pooling
	DbMaxOpenConn int    `mapstructure:"DB_MAX_OPEN_CONN"`
	DbMaxIdleConn int    `mapstructure:"DB_MAX_IDLE_CONN"`
	DbMaxIdleTime string `mapstructure:"DB_MAX_IDLE_TIME"`

	// jwt
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

func (cfg *Config) GetDBConn() *gorm.DB {
	return cfg.dbconn
}

func ShowHelp(listenv []string, errors []error) {
	for _, err := range errors {
		fmt.Fprintf(flag.CommandLine.Output(), "%v\n", err)
	}

	pflag.Usage()

	fmt.Fprintln(flag.CommandLine.Output(), "\nLIST ENVIRONMENT: (COMMAND LINE WILL OVERRIDE ENV)")
	for _, env := range listenv {
		fmt.Fprintln(flag.CommandLine.Output(), env)
	}
	fmt.Fprintln(flag.CommandLine.Output())
}

func New() (*Config, error) {
	viper.SetDefault("API_PORT", 8080)
	viper.SetDefault("DB_MAX_OPEN_CONN", 25)
	viper.SetDefault("DB_MAX_IDLE_CONN", 25)
	viper.SetDefault("DB_MAX_IDLE_TIME", "15m")

	if err := scanEnv(); err != nil {
		return nil, err
	}

	if err := tryCommandLine(); err != nil {
		return nil, err

	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err

	}

	if listenv, errors := check(&cfg); errors != nil {
		return nil, ErrorMissingValue{Listenv: listenv, Errors: errors}
	}

	// open db conn
	dbconn, err := cfg.openConn()
	if err != nil {
		return nil, err
	}
	cfg.dbconn = dbconn

	return &cfg, nil
}

func scanEnv() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return nil
		default:
			return err
		}
	}
	return nil
}

func tryCommandLine() error {
	pflag.Int("API_PORT", 8080, "API Server port")

	pflag.String("DB_NAME", "", "Database name")
	pflag.String("DB_USER", "", "Database user")
	pflag.String("DB_PASSWORD", "", "Database password")
	pflag.String("DB_URL", "", "Database url")
	pflag.String("DB_PORT", "", "Database port")

	pflag.Int("DB_MAX_OPEN_CONN", 25, "Database max open connections")
	pflag.Int("DB_MAX_IDLE_CONN", 25, "Database max idle connections")
	pflag.String("DB_MAX_IDLE_TIME", "15m", "Database max connection idle time")

	pflag.String("JWT_SECRET", "", "JWT signature secret")
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return err
	}
	return nil
}

func check(cfg *Config) ([]string, []error) {
	if structs.HasZero(cfg) {
		var listenv []string
		var errors []error

		for _, field := range structs.Fields(cfg) {
			env := field.Tag("mapstructure")
			listenv = append(listenv, env)

			if env == "" {
				continue
			}

			if field.IsZero() {
				errors = append(errors, fmt.Errorf("ERROR MISING VALUE: %v", env))
			}
		}

		if len(errors) > 0 {
			return listenv, errors
		}
	}

	return nil, nil
}
