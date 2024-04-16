package handlers

const (
	NO_CREDENTIAL        = "NO_CREDENTIAL"
	SERVICE_CREATE_ISSEU = "SERVICE_CREATE_ISSEU"
	BAD_REQUEST          = "BAD_REQUEST"
	SERVICE_PARSE        = "SERVICE_PARSE"
	SERVICE_RETURN       = "SERVICE_RETURN"
	SERVICE_VAliDATE     = "SERVICE_VAliDATE"

	ResponseMessage_AccessDenied       = "Access denied!"
	ResponseMessage_NotCreateIsseu     = "Something went wrong! Isseu not created!"
	ResponseMessage_ServerReturnError  = "Something went wrong!"
	ResponseMessage_ListError          = "Service can't return list of isseus!"
	ResponseMessage_OptionsError       = "Service can't return options!"
	ResponseMessage_ServiceBodyParser  = "Service can't parse request body!"
	ResponseMessage_NotFoundIsseu      = "Not found isseu with id="
	ResponseMessage_UpdateStatusError  = "Not supported status!"
	ResponseMessage_ServiceUpdateError = "Service can't updated isseu!"
)
