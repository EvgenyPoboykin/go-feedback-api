package endpoints

import "fmt"

type Endpoints struct {
	BaseUrl string
}

func NewEndpoints(version string) *Endpoints {
	return &Endpoints{
		BaseUrl: fmt.Sprintf("/api/%s", version),
	}
}

func (a *Endpoints) ListAdmin() string {
	return a.BaseUrl + "/list/admin"
}

func (a *Endpoints) ListEmployee() string {
	return a.BaseUrl + "/list/employee"
}

func (a *Endpoints) Item() string {
	return a.BaseUrl + "/list/issue/{id}"
}

func (a *Endpoints) Create() string {
	return a.BaseUrl + "/list/issue/create"
}

func (a *Endpoints) Options() string {
	return a.BaseUrl + "/options"
}
