package converter

import (
	"encoding/json"

	"github.com/eugenepoboykin/go-feedback-api/internal/domain/models"
)

func ToListArgsFromService(params models.StorageListDTO) models.ServiceListDTO {
	return models.ServiceListDTO{
		Status:   params.Status,
		Page:     params.Page,
		PageSize: params.PageSize,
	}
}

func FromDbToServiceMap[A any, O any](i A) (*O, error) {
	var output O

	b, err := json.Marshal(&i)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
