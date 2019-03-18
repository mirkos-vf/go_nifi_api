package main

type RemoteProcessGroupPortDTO struct {
	Id                               string           `json:"id"`
	TargetId                         string           `json:"targetid"`
	VersionedComponentId             string           `json:"versionedcomponentid"`
	GroupId                          string           `json:"groupid"`
	Name                             string           `json:"name"`
	Comments                         string           `json:"comments"`
	ConcurrentlySchedulableTaskCount int32            `json:"concurrentlyschedulabletaskcount"`
	Transmitting                     bool             `json:"transmitting"`
	Usecompression                   bool             `json:"usecompression"`
	Exists                           bool             `json:"exists"`
	TargetRunning                    bool             `json:"targetrunning"`
	Connected                        bool             `json:"connected"`
	BatchSettings                    BatchSettingsDTO `json:"batchsettings"`
}
