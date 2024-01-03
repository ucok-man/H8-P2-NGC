package config

import (
	"flag"
	"fmt"
	"os"

	"gorm.io/gorm"
)

const (
	userkey string = "user"
)

type Config struct {
	Addr string
	env  struct {
		dbdsn     string
		jwtsecret string
	}
	context struct {
		userkey string
	}
	dbconn *gorm.DB
}

func New() *Config {
	var cfg Config
	cfg.context.userkey = userkey

	flag.StringVar(&cfg.Addr, "addr", ":8080", "API Server port")
	flag.StringVar(&cfg.env.dbdsn, "dsn", "", "Database source connection")
	flag.StringVar(&cfg.env.jwtsecret, "secret", "", "JWT secret token")
	flag.Parse()

	// if env is set then use it
	if dsn := os.Getenv("DB_DSN"); dsn != "" {
		cfg.env.dbdsn = dsn
	}

	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		cfg.env.jwtsecret = jwtSecret
	}
	fmt.Println(cfg)

	return check(&cfg)
}

func (cfg *Config) GetUserCtxKey() string {
	return cfg.context.userkey
}

func (cfg *Config) GetServerAddr() string {
	return cfg.env.dbdsn
}

func (cfg *Config) GetJwtSecret() string {
	return cfg.env.jwtsecret
}

func (cfg *Config) GetDBConn() *gorm.DB {
	if cfg.dbconn == nil {
		cfg.dbconn = OpenConn(cfg.env.dbdsn)
	}
	return cfg.dbconn
}

func check(cfg *Config) *Config {
	if cfg.env.dbdsn == "" {
		fmt.Fprintf(os.Stderr, "err: you must provide db source connection through ENV (DB_DSN as name) or Command Line Flag\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if cfg.env.jwtsecret == "" {
		fmt.Fprintf(os.Stderr, "err: you must provide jwt secret token through ENV (JWT_SECRET as name) or Command Line Flag\n\n")
		flag.Usage()
		os.Exit(1)
	}
	return cfg
}
