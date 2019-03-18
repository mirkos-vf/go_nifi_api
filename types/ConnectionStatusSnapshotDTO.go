package types

type ConnectionStatusSnapshotDTO struct {
	Id              string `json:"id"`
	GroupId         string `json:"groupid"`
	Name            string `json:"name"`
	SourceId        string `json:"sourceid"`
	SourceName      string `json:"sourcename"`
	DestinationId   string `json:"destinationid"`
	DestinationName string `json:"destinationname"`
	FlowFilesIn     int32  `json:"flowfilesin"`
	BytesIn         int64  `json:"bytesin"`
	Input           string `json:"input"`
	FlowFilesOut    int32  `json:"flowfilesout"`
	BytesOut        int64  `json:"bytesout"`
	Output          string `json:"output"`
	FlowFilesQueued int32  `json:"flowfilesqueued"`
	BytesQueued     int64  `json:"bytesqueued"`
	Queued          string `json:"queued"`
	QueuedSize      string `json:"queuedsize"`
	QueuedCount     string `json:"queuedcount"`
	PercentUseCount int32  `json:"percentusecount"`
	PercentUseBytes int32  `json:"percentusebytes"`
}
