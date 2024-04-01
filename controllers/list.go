package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/schema"
	"github.com/eugenepoboykin/go-feedback-api/services"
)

func List(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("oauth.role").(string)
	clientId := r.Context().Value("oauth.clientId").(string)

	var body schema.IssuesListArgs

	var empty = make([]schema.Issue, 0)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		helpers.ErrorResponse(w, http.StatusServiceUnavailable, constant.SERVICE_PARSE, constant.ResponseMessage_ServiceBodyParser)

		return
	}

	var page schema.IssuesList

	page.Page = body.Page
	page.PageSize = body.PageSize
	page.Status = body.Status

	fmt.Println("-----", page.Status, &page.Status, "------")

	if string(body.Status) != `` {
		_, errorStatus := services.GetOptionByValue(string(body.Status))

		if errorStatus != nil {
			helpers.ErrorResponse(w, http.StatusServiceUnavailable, constant.SERVICE_VAliDATE, constant.ResponseMessage_UpdateStatusError)

			return
		}
	}

	if role == constant.Employee {
		isseus, err := services.GetListEmployee(clientId)

		if err != nil {
			helpers.ErrorResponse(w, http.StatusServiceUnavailable, constant.SERVICE_RETURN, constant.ResponseMessage_ListError)

			return
		}

		if isseus != nil {
			isseusSlice := helpers.IssuePerPage(isseus, body)

			page.TotalCount = len(isseus)
			page.Issues = isseusSlice
		} else {
			page.TotalCount = 0
			page.Issues = empty
		}

		helpers.Response[schema.IssuesList](w, &page)

		return
	}

	isseus, err := services.GetListAll()

	if err != nil {
		helpers.ErrorResponse(w, http.StatusServiceUnavailable, constant.SERVICE_RETURN, constant.ResponseMessage_ListError)

		return
	}

	if isseus != nil {
		isseusSlice := helpers.IssuePerPage(isseus, body)

		page.TotalCount = len(isseus)
		page.Issues = isseusSlice
	} else {
		page.TotalCount = 0
		page.Issues = empty
	}

	helpers.Response[schema.IssuesList](w, &page)
}
