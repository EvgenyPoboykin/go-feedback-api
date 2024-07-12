package postgres

import (
	"fmt"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/env"
)

func PostgresDsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", env.Environment.DbUser, env.Environment.DbPassword, env.Environment.DbHost, env.Environment.DbPort, env.Environment.DbName, env.Environment.DbSSL)
}
