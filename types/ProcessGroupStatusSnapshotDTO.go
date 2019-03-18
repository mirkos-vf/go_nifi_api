package types

type ProcessGroupStatusSnapshotDTO struct {
	Id                                string                                   `json:"id"`
	Name                              string                                   `json:"name"`
	ConnectionStatusSnapshots         []ConnectionStatusSnapshotEntity         `json:"connectionstatussnapshots"`
	ProcessorStatusSnapshots          []ProcessorStatusSnapshotEntity          `json:"processorstatussnapshots"`
	ProcessGroupStatusSnapshots       []ProcessGroupStatusSnapshotEntity       `json:"processgroupstatussnapshots"`
	RemoteProcessGroupStatusSnapshots []RemoteProcessGroupStatusSnapshotEntity `json:"remoteprocessgroupstatussnapshots"`
	InputPortStatusSnapshots          []PortStatusSnapshotEntity               `json:"inputportstatussnapshots"`
	OutputPortStatusSnapshots         []PortStatusSnapshotEntity               `json:"outputportstatussnapshots"`
	VersionedFlowState                string                                   `json:"versionedflowstate"`
	FlowFilesIn                       string                                   `json:"flowfilesin"`
	BytesIn                           int64                                    `json:"bytesin"`
	Input                             string                                   `json:"input"`
	FlowFilesQueued                   int32                                    `json:"flowfilesqueued"`
	BytesQueued                       int64                                    `json:"bytesqueued"`
	Queued                            string                                   `json:"queued"`
	QueuedCount                       string                                   `json:"queuedcount"`
	QueuedSize                        string                                   `json:"queuedsize"`
	BytesRead                         int64                                    `json:"bytesread"`
	Read                              string                                   `json:"read"`
	BytesWritten                      int64                                    `json:"byteswritten"`
	Written                           string                                   `json:"written"`
	FlowFilesOut                      int32                                    `json:"flowfilesout"`
	BytesOut                          int64                                    `json:"bytesout"`
	Output                            string                                   `json:"output"`
	FlowFilesTransferred              int32                                    `json:"flowfilestransferred"`
	BytesTransferred                  int64                                    `json:"bytestransferred"`
	Transferred                       string                                   `json:"transferred"`
	BytesReceived                     int64                                    `json:"bytesreceived"`
	FlowFilesReceived                 int32                                    `json:"flowfilesreceived"`
	Received                          string                                   `json:"received"`
	BytesSent                         int64                                    `json:"bytessent"`
	FlowFilesSent                     int32                                    `json:"flowfilessent"`
	Sent                              string                                   `json:"sent"`
	ActiveThreadCount                 int32                                    `json:"activethreadcount"`
	TerminatedThreadCount             int32                                    `json:"terminatedthreadcount"`
}
