package storage

import (
	"database/sql"
)

func StmtByRole(db *sql.DB, clientId *string) (*sql.Stmt, error) {
	if clientId != nil {
		stmt, errStmt := db.Prepare(QueryListByEmployee)
		if errStmt != nil {
			return nil, errStmt
		}

		return stmt, nil
	} else {
		stmt, errStmt := db.Prepare(QueryListByAdmin)
		if errStmt != nil {
			return nil, errStmt
		}

		return stmt, nil
	}
}
