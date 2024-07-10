package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/models"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) List(ctx context.Context, clientId *string) (*[]models.StorageIssue, error) {
	var issues []models.StorageIssue

	stmt, errStmt := StmtByRole(r.DB, clientId)
	if errStmt != nil {
		return nil, errStmt
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var issue models.StorageIssue

		err := rows.Scan(&issue)
		if err != nil {
			return nil, err
		}

		issues = append(issues, issue)
	}

	return &issues, nil
}

func (r *Repository) IssueById(ctx context.Context, issueId string) (*models.StorageIssue, error) {
	var issue models.StorageIssue

	stmt, errStmt := r.DB.Prepare(QueryGetById)
	if errStmt != nil {
		return nil, errStmt
	}

	res := stmt.QueryRowContext(ctx, issueId)

	err := res.Scan(&issue.Id, &issue.Uri, &issue.Image64, &issue.Description, &issue.Comment, &issue.Status, &issue.Created, &issue.Updated, &issue.ClientId, &issue.ClientName)
	if err != nil {
		return nil, err
	}

	return &issue, nil
}

func (r *Repository) IssueByParams(ctx context.Context, params models.StorageAddIssueDTO) (*models.StorageIssue, error) {
	var issue models.StorageIssue

	stmt, errStmt := r.DB.Prepare(QueryFindIssueByParams)
	if errStmt != nil {
		return nil, errStmt
	}

	res := stmt.QueryRowContext(ctx, params.Uri, params.Image64, params.Description, params.ClientId, params.ClientName)

	err := res.Scan(&issue.Id, &issue.Uri, &issue.Image64, &issue.Description, &issue.Comment, &issue.Status, &issue.Created, &issue.Updated, &issue.ClientId, &issue.ClientName)
	if err != nil {
		return nil, err
	}

	return &issue, nil
}

func (r *Repository) Create(ctx context.Context, params models.StorageAddIssueDTO) (*models.StorageIssue, error) {
	stmt, errStmt := r.DB.Prepare(QueryAddIssueItem)
	if errStmt != nil {
		return nil, errStmt
	}

	_, err := stmt.ExecContext(ctx,
		params.Uri,
		params.Image64,
		params.Description,
		time.Now(),
		time.Now(),
		params.ClientId,
		params.ClientName,
	)
	if err != nil {
		return nil, err
	}

	issue, err := r.IssueByParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

func (r *Repository) Update(ctx context.Context, params models.StorageUpdateIssueDTO) (*models.StorageIssue, error) {
	stmt, errStmt := r.DB.Prepare(QueryUpdateIssueItem)
	if errStmt != nil {
		return nil, errStmt
	}

	_, err := stmt.ExecContext(ctx,
		params.Comment,
		params.Status,
		time.Now(),
		params.Id,
	)
	if err != nil {
		return nil, err
	}

	issue, err := r.IssueById(ctx, params.Id)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

func (r *Repository) Options(ctx context.Context) (*[]models.StorageOption, error) {
	var options []models.StorageOption

	stmt, errStmt := r.DB.Prepare(QueryGetOptions)
	if errStmt != nil {
		return nil, errStmt
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var option models.StorageOption

		err := rows.Scan(
			&option.Value,
			&option.Label,
		)
		if err != nil {
			return nil, err
		}

		options = append(options, option)
	}

	return &options, nil
}

func (r *Repository) OptionByValue(ctx context.Context, value string) (*models.StorageOption, error) {
	var option models.StorageOption

	stmt, errStmt := r.DB.Prepare(QueryGetOptionByValue)
	if errStmt != nil {
		return nil, errStmt
	}

	res := stmt.QueryRowContext(ctx, value)

	err := res.Scan(&option.Value, &option.Label)
	if err != nil {
		return nil, err
	}

	return &option, nil
}
