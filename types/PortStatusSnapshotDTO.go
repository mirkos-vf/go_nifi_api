package main

type PortStatusSnapshotDTO struct {
	Id                string `json:"id"`
	GroupId           string `json:"groupid"`
	Name              string `json:"name"`
	ActiveThreadCount int32  `json:"activethreadcount"`
	FlowFilesIn       int32  `json:"flowfilesin"`
	BytesIn           int64  `json:"bytesin"`
	Input             string `json:"input"`
	FlowFilesOut      int32  `json:"flowfilesout"`
	BytesOut          int64  `json:"bytesout"`
	Output            string `json:"output"`
	Transmitting      bool   `json:"transmitting"`
	RunStatus         string `json:"runstatus"`
}
