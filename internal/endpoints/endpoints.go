package endpoints

type Api struct {
	BaseUrl string
}

func NewApiVersion(version string) *Api {
	return &Api{
		BaseUrl: "/" + "api" + "/" + version,
	}
}

func (a *Api) ListAdmin() string {
	return a.BaseUrl + "/list/admin"
}

func (a *Api) ListEmployee() string {
	return a.BaseUrl + "/list/employee"
}

func (a *Api) Item() string {
	return a.BaseUrl + "/list/issue/{id}"
}

func (a *Api) Create() string {
	return a.BaseUrl + "/list/issue/create"
}

func (a *Api) Options() string {
	return a.BaseUrl + "/options"
}
