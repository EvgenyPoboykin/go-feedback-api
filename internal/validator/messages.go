package validator

const (
	SERVICE_PARSE     = "SERVICE_PARSE"
	FIELD_DESCRIPTION = "FIELD_DESCRIPTION"
	FIELD_URI         = "FIELD_URI"

	ResponseMessage_ServiceBodyParser = "Service can't parse request body!"
	ResponseMessage_DescriptionError  = "Description is required field!"
	ResponseMessage_UriError          = "Uri is required field!"
	ResponseMessage_StatusError       = "Status is required field!"
	ResponseMessage_ParamsIdError     = "Id is not found!"
)
