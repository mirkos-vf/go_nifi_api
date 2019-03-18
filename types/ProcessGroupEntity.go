package types

type ProcessGroupEntity struct {
	Revision                     RevisionDTO           `json:"revision"`
	Id                           string                `json:"id"`
	Url                          string                `json:"url"`
	Position                     PositionDTO           `json:"position"`
	Permissions                  PermissionsDTO        `json:"permissions"`
	Bulletins                    []BulletinEntity      `json:"bulletins"`
	DisconnectedNodeAcknowledged bool                  `json:"disconnectednodeacknowledged"`
	Component                    ProcessGroupDTO       `json:"component"`
	Status                       ProcessGroupStatusDTO `json:"status"`
	VersionedFlowSnapshot        versionedFlowSnapshot `json:"versionedflowsnapshot"`
	RunningCount                 int32                 `json:"runningcount"`
	StoppedCount                 int32                 `json:"stoppedcount"`
	InvalidCount                 int32                 `json:"invalidcount"`
	DisabledCount                int32                 `json:"disabledcount"`
	ActiveRemotePortCount        int32                 `json:"activeremoteportcount"`
	InactiveRemotePortCount      int32                 `json:"inactiveremoteportcount"`
	VersionedFlowState           string                `json:"versionedflowstate"`
	UpToDateCount                int32                 `json:"uptodatecount"`
	LocallyModifiedCount         int32                 `json:"locallymodifiedcount"`
	StaleCount                   int32                 `json:"stalecount"`
	LocallyModifiedAndStaleCount int32                 `json:"locallymodifiedandstalecount"`
	SyncFailureCount             int32                 `json:"syncfailurecount"`
	InputPortCount               int32                 `json:"inputportcount"`
	OutputPortCount              int32                 `json:"outputportcount"`
}
