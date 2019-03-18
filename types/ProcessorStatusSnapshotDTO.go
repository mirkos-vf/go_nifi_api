package types

type ProcessorStatusSnapshotDTO struct {
	Id                    string `json:"id"`
	GroupId               string `json:"groupid"`
	Name                  string `json:"name"`
	Type                  string `json:"type"`
	RunStatus             string `json:"runstatus"`
	ExecutionNode         string `json:"executionnode"`
	BytesRead             int64  `json:"bytesread"`
	BytesWritten          int64  `json:"byteswritten"`
	Read                  string `json:"read"`
	Written               string `json:"written"`
	FlowFilesIn           int32  `json:"flowfilesin"`
	BytesIn               int64  `json:"bytesin"`
	Input                 string `json:"input"`
	FlowFilesOut          int32  `json:"flowfilesout"`
	BytesOut              int64  `json:"bytesout"`
	Output                string `json:"output"`
	TaskCount             int32  `json:"taskcount"`
	TasksDurationNanos    int64  `json:"tasksdurationnanos"`
	Tasks                 string `json:"tasks"`
	TasksDuration         string `json:"tasksduration"`
	ActiveThreadCount     int32  `json:"activethreadcount"`
	TerminatedThreadCount int32  `json:"terminatedthreadcount"`
}
