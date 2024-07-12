package storage

import (
	"context"
	"fmt"
	"reflect"
	"time"

	globalErrors "errors"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) (*Repository, error) {
	if pool == nil {
		return nil, globalErrors.New("no db connection")
	}

	return &Repository{
		Pool: pool,
	}, nil
}

func (r *Repository) ListAdmin(ctx context.Context) (*[]models.StorageIssue, error) {
	var issues []models.StorageIssue

	rows, err := r.Pool.Query(ctx, QueryListByAdmin)
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

func (r *Repository) ListEmployee(ctx context.Context, clientId *string) (*[]models.StorageIssue, error) {
	var issues []models.StorageIssue

	fmt.Print("i`m here", r.Pool)
	rows, err := r.Pool.Query(ctx, QueryListByEmployee, clientId)
	fmt.Println(reflect.TypeOf(rows))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var issue models.StorageIssue

		err := rows.Scan(&issue.Id, &issue.Uri)
		if err != nil {
			return nil, err
		}

		issues = append(issues, issue)
	}

	return &issues, nil
}

func (r *Repository) IssueById(ctx context.Context, issueId string) (*models.StorageIssue, error) {
	var issue models.StorageIssue

	res := r.Pool.QueryRow(ctx, QueryGetById, issueId)

	err := res.Scan(&issue.Id, &issue.Uri, &issue.Image64, &issue.Description, &issue.Comment, &issue.Status, &issue.Created, &issue.Updated, &issue.ClientId, &issue.ClientName)
	if err != nil {
		return nil, err
	}

	return &issue, nil
}

func (r *Repository) IssueByParams(ctx context.Context, params models.StorageAddIssueDTO) (*models.StorageIssue, error) {
	var issue models.StorageIssue

	res := r.Pool.QueryRow(ctx, QueryFindIssueByParams, params.Uri, params.Image64, params.Description, params.ClientId, params.ClientName)

	err := res.Scan(&issue.Id, &issue.Uri, &issue.Image64, &issue.Description, &issue.Comment, &issue.Status, &issue.Created, &issue.Updated, &issue.ClientId, &issue.ClientName)
	if err != nil {
		return nil, err
	}

	return &issue, nil
}

func (r *Repository) Create(ctx context.Context, params models.StorageAddIssueDTO) (*models.StorageIssue, error) {
	_, err := r.Pool.Exec(ctx, QueryAddIssueItem,
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
	_, err := r.Pool.Exec(ctx, QueryUpdateIssueItem,
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

	rows, err := r.Pool.Query(ctx, QueryGetOptions)
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

	res := r.Pool.QueryRow(ctx, QueryGetOptionByValue, value)

	err := res.Scan(&option.Value, &option.Label)
	if err != nil {
		return nil, err
	}

	return &option, nil
}
