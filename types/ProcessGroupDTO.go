package main

type ProcessGroupDTO struct {
	Id                           string                       `json:"id"`
	VersionedComponentId         string                       `json:"versionedcomponentid"`
	ParentGroupId                string                       `json:"parentgroupid"`
	Position                     PositionDTO                  `json:"position"`
	Name                         string                       `json:"name"`
	Comments                     string                       `json:"comments"`
	Variables                    object                       `json:"variables"`
	VersionControlInformation    VersionControlInformationDTO `json:"versioncontrolinformation"`
	RunningCount                 int32                        `json:"runningcount"`
	StoppedCount                 int32                        `json:"stoppedcount"`
	InvalidCount                 int32                        `json:"invalidcount"`
	DisabledCount                int32                        `json:"disabledcount"`
	InactiveRemotePortCount      int32                        `json:"inactiveremoteportcount"`
	UpToDateCount                int32                        `json:"uptodatecount"`
	LocallyModifiedCount         int32                        `json:"locallymodifiedcount"`
	StaleCount                   int32                        `json:"stalecount"`
	LocallyModifiedAndStaleCount int32                        `json:"locallymodifiedandstalecount"`
	SyncFailureCount             int32                        `json:"syncfailurecount"`
	InputPortCount               int32                        `json:"inputportcount"`
	OutputPortCount              int32                        `json:"outputportcount"`
	Contents                     FlowSnippetDTO               `json:"contents"`
}
