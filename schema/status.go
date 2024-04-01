package schema

type Status string

const (
	New        Status = "NEW"
	Rejected   Status = "REJECTED"
	InProccess Status = "IN_PROCESS"
	Done       Status = "DONE"
)
