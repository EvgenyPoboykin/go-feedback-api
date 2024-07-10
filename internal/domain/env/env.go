package env

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DbName       string
	DbUser       string
	DbPassword   string
	DbHost       string
	DbPort       string
	DbSSL        string
	AppPort      string
	Secret       string
	AdminRole    string
	EmployeeRole string
}

var Environment *Env

func NewEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return errors.New("env config not found")
	}

	dbHost, err := checkEnv("DB_HOST")
	if err != nil {
		return err
	}

	dbName, err := checkEnv("DB_NAME")
	if err != nil {
		return err
	}

	dbUser, err := checkEnv("DB_USER")
	if err != nil {
		return err
	}

	dbPassword, err := checkEnv("DB_PASSWORD")
	if err != nil {
		return err
	}

	dbPort, err := checkEnv("DB_PORT")
	if err != nil {
		return err
	}

	dbSSL, err := checkEnv("DB_SSL")
	if err != nil {
		return err
	}

	appPort, err := checkEnv("APP_PORT")
	if err != nil {
		return err
	}

	secret, err := checkEnv("SUPPORT_SERVICE_SECRET_KEY")
	if err != nil {
		return err
	}

	adminRole, err := checkEnv("SUPPORT_SERVICE_ROLE_ADMIN")
	if err != nil {
		return err
	}

	employeeRole, err := checkEnv("SUPPORT_SERVICE_ROLE_USER")
	if err != nil {
		return err
	}

	Environment = &Env{
		DbName:       dbName,
		DbUser:       dbUser,
		DbPassword:   dbPassword,
		DbPort:       dbPort,
		DbHost:       dbHost,
		DbSSL:        dbSSL,
		AppPort:      ":" + appPort,
		Secret:       secret,
		AdminRole:    adminRole,
		EmployeeRole: employeeRole,
	}

	return nil
}

func checkEnv(value string) (string, error) {
	value = os.Getenv(value)

	if value == "" {
		return "", errors.New(fmt.Sprintf("in .env config not found require key %s!", value))
	}

	return value, nil
}
