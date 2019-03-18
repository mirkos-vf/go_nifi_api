package types

type RemoteProcessGroupStatusSnapshotDTO struct {
	Id                 string `json:"id"`
	GroupId            string `json:"groupid"`
	Name               string `json:"name"`
	TargetUri          string `json:"targeturi"`
	TransmissionStatus string `json:"transmissionstatus"`
	ActiveThreadCount  int    `json:"activethreadcount"`
	FlowFilesSent      int    `json:"flowfilessent"`
	BytesSent          int    `json:"bytessent"`
	Sent               string `json:"sent"`
	FlowFilesReceived  int    `json:"flowfilesreceived"`
	BytesReceived      int    `json:"bytesreceived"`
	Received           string `json:"received"`
}
