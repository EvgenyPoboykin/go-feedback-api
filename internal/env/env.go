package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DSN          string
	DbName       string
	DbUser       string
	DbPassword   string
	DbPort       string
	DbSSL        string
	AppPort      string
	Secret       string
	AdminRole    string
	EmployeeRole string
}

var Enviroment *Env

func NewEnv() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env config not found!")
	}

	dbHost := checkEnv("DB_HOST")
	dbName := checkEnv("DB_NAME")
	dbUser := checkEnv("DB_USER")
	dbPassword := checkEnv("DB_PASSWORD")
	dbPort := checkEnv("DB_PORT")
	dbSSL := checkEnv("DB_SSL")

	appPort := checkEnv("APP_PORT")

	secret := checkEnv("SUPPORT_SERVICE_SECRET_KEY")
	adminRole := checkEnv("SUPPORT_SERVICE_ROLE_ADMIN")
	employeeRole := checkEnv("SUPPORT_SERVICE_ROLE_USER")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSL)

	Enviroment = &Env{
		DSN:        psqlInfo,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbPort:     dbPort,
		DbSSL:      dbSSL,
		AppPort:    ":" + appPort,

		Secret:       secret,
		AdminRole:    adminRole,
		EmployeeRole: employeeRole,
	}

	return Enviroment
}

func checkEnv(value string) string {
	value = os.Getenv(value)

	if value == "" {
		log.Fatalf("in .env config not found require key %s!", value)
	}

	return value
}
