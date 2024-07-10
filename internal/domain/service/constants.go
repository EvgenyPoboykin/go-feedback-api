package service

const (
	NoCredential       = "NO_CREDENTIAL"
	ServiceCreateIssue = "SERVICE_CREATE_ISSUE"
	BadRequest         = "BAD_REQUEST"
	ServiceReturn      = "SERVICE_RETURN"
	ServiceValidate    = "SERVICE_VALIDATE"

	ResponseMessage_AccessDenied       = "Access denied!"
	ResponseMessage_NotCreateIssue     = "Something went wrong! Issue not created!"
	ResponseMessage_ListError          = "Service can't return list of issues!"
	ResponseMessage_OptionsError       = "Service can't return options!"
	ResponseMessage_NotFoundIssue      = "Not found issue with id="
	ResponseMessage_UpdateStatusError  = "Not supported status!"
	ResponseMessage_ServiceUpdateError = "Service can't updated issue!"
)
