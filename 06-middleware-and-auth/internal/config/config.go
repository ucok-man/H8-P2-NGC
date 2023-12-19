package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

type CtxKey string

type configuration struct {
	Port       int
	DBdsn      string
	JWTSecreet string
	context    struct {
		userkey CtxKey
	}
}

func NewConfig() configuration {
	var cfg configuration
	flag.IntVar(&cfg.Port, "port", 8080, "API Server port")
	flag.StringVar(&cfg.DBdsn, "dsn", "", "Database source connection")
	flag.StringVar(&cfg.JWTSecreet, "secret", "", "JWT secret token")
	flag.Parse()

	// if env is set then use it
	if dsn := os.Getenv("DB_DSN"); dsn != "" {
		cfg.DBdsn = dsn
	}

	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		cfg.JWTSecreet = jwtSecret
	}

	// check
	if cfg.DBdsn == "" {
		fmt.Fprintf(os.Stderr, "err: you must provide db source connection through ENV (DB_DSN as name) or Command Line Flag\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if cfg.JWTSecreet == "" {
		fmt.Fprintf(os.Stderr, "err: you must provide jwt secret token through ENV (JWT_SECRET as name) or Command Line Flag\n\n")
		flag.Usage()
		os.Exit(1)
	}
	return cfg
}

func (cfg *configuration) GetUserCtxKey() CtxKey {
	return cfg.context.userkey
}
