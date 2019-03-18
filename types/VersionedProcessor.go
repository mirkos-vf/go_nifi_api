package types

type VersionedProcessor struct {
	Identifier                       string   `json:"identifier"`
	Name                             string   `json:"name"`
	Comments                         string   `json:"comments"`
	Position                         Position `json:"position"`
	Bundle                           Bundle   `json:"bundle"`
	Style                            object   `json:"style"`
	Type                             string   `json:"type"`
	Properties                       object   `json:"properties"`
	PropertyDescriptors              object   `json:"propertydescriptors"`
	AnnotationData                   string   `json:"annotationdata"`
	SchedulingPeriod                 string   `json:"schedulingperiod"`
	SchedulingStrategy               string   `json:"schedulingstrategy"`
	ExecutionNode                    string   `json:"executionnode"`
	PenaltyDuration                  string   `json:"penaltyduration"`
	YieldDuration                    string   `json:"yieldduration"`
	BulletinLevel                    string   `json:"bulletinlevel"`
	RunDurationMillis                int64    `json:"rundurationmillis"`
	ConcurrentlySchedulableTaskCount int64    `json:"concurrentlyschedulabletaskcount"`
	AutoTerminatedRelationships      []string `json:"autoterminatedrelationships"`
	ComponentType                    string   `json:"componenttype"`
	GroupIdentifier                  string   `json:"groupidentifier"`
}
