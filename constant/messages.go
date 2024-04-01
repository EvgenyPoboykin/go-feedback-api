package constant

const (
	Log_Ping               = "*** Ping database successfuly! ***"
	Log_ErrorResponse      = "*** Error:"
	Log_ErrorNoEnv         = "*** No .env file found! ***"
	Log_ErrorInitialConfig = "*** Error initial configs: %s ***"
	Log_ErrorInsert        = "*** Error db insert: %s ***"
	Log_ErrorSelect        = "*** Error db select: %s ***"
	Log_ErrorUpdate        = "*** Error db update: %s ***"

	ResponseMessage_AuthHeaderError        = "Authorization header is not provided"
	ResponseMessage_AuthFormatHeaderError  = "Authorization header is not provided"
	ResponseMessage_AuthUnsupportTypeError = "Unsupported authorization type "
	ResponseMessage_NotAuthorizedError     = "You not authorized!"

	ResponseMessage_AccessDenied       = "Access denied!"
	ResponseMessage_NotCreateIsseu     = "Something went wrong! Isseu not created!"
	ResponseMessage_ServerReturnError  = "Something went wrong!"
	ResponseMessage_ListError          = "Service can't return list of isseus!"
	ResponseMessage_OptionsError       = "Service can't return options!"
	ResponseMessage_ServiceBodyParser  = "Service can't parse request body!"
	ResponseMessage_NotFoundIsseu      = "Not found isseu with id="
	ResponseMessage_UpdateStatusError  = "Not supported status!"
	ResponseMessage_ServiceUpdateError = "Service can't updated isseu!"
	ResponseMessage_DescriptionError   = "Description is required field!"
	ResponseMessage_UriError           = "Uri is required field!"
)
