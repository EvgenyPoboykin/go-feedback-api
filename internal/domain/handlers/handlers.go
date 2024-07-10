package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/models"
	"github.com/eugenepoboykin/go-feedback-api/internal/errors"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"

	"github.com/eugenepoboykin/go-feedback-api/internal/endpoints"
	"github.com/go-chi/chi/v5"
)

type Service interface {
	Create(ctx context.Context, role string, body io.ReadCloser) (*models.ServiceIssue, *errors.ErrorResponse)
	IssueById(ctx context.Context, role string, id string) (*models.ServiceIssue, *errors.ErrorResponse)
	ListAdmin(ctx context.Context, role string, body io.ReadCloser) (*models.ServiceIssuesList, *errors.ErrorResponse)
	ListEmployee(ctx context.Context, role string, clientId string, body io.ReadCloser) (*models.ServiceIssuesList, *errors.ErrorResponse)
	Options(ctx context.Context, role string) (*[]models.ServiceOption, *errors.ErrorResponse)
	Update(ctx context.Context, role string, id string, body io.ReadCloser) (*models.ServiceIssue, *errors.ErrorResponse)
}

type Handlers struct {
	ctx     context.Context
	service Service
}

func NewHandlers(ctx context.Context, service Service) *Handlers {
	return &Handlers{
		ctx:     ctx,
		service: service,
	}
}

func (h Handlers) RegisterRoutes(r *chi.Mux) {
	endpointsApi := endpoints.NewEndpoints("v1")

	r.Post(endpointsApi.ListAdmin(), h.ListAdmin)
	r.Post(endpointsApi.ListEmployee(), h.ListEmployee)
	r.Get(endpointsApi.Item(), h.IssueById)
	r.Put(endpointsApi.Item(), h.Update)
	r.Post(endpointsApi.Create(), h.Create)
	r.Get(endpointsApi.Options(), h.Options)
}

func (h Handlers) ListAdmin(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("oauth.role").(string)

	res, err := h.service.ListAdmin(h.ctx, role, r.Body)

	defer func() {
		err := r.Body.Close()
		if err != nil {
			return
		}
	}()

	if err != nil {
		response.ErrorResponse(w, http.StatusBadGateway, *err)

		return
	}

	response.Response(w, res)
}

func (h Handlers) ListEmployee(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("oauth.role").(string)
	clientId := r.Context().Value("oauth.clientId").(string)

	res, err := h.service.ListEmployee(h.ctx, role, clientId, r.Body)

	defer func() {
		err := r.Body.Close()
		if err != nil {
			return
		}
	}()

	if err != nil {
		response.ErrorResponse(w, http.StatusBadGateway, *err)

		return
	}

	response.Response(w, res)
}

func (h Handlers) IssueById(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("oauth.role").(string)
	id := chi.URLParam(r, "id")

	res, err := h.service.IssueById(h.ctx, role, id)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadGateway, *err)

		return
	}

	response.Response(w, res)
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("oauth.role").(string)

	res, err := h.service.Create(h.ctx, role, r.Body)

	defer func() {
		err := r.Body.Close()
		if err != nil {
			return
		}
	}()

	if err != nil {
		response.ErrorResponse(w, http.StatusBadGateway, *err)

		return
	}

	response.Response(w, res)
}

func (h Handlers) Update(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("oauth.role").(string)

	id := chi.URLParam(r, "id")

	res, err := h.service.Update(h.ctx, role, id, r.Body)

	defer func() {
		err := r.Body.Close()
		if err != nil {
			return
		}
	}()

	if err != nil {
		response.ErrorResponse(w, http.StatusBadGateway, *err)

		return
	}

	response.Response(w, res)
}

func (h Handlers) Options(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("oauth.role").(string)

	res, err := h.service.Options(h.ctx, role)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadGateway, *err)

		return
	}

	response.Response(w, res)
}
