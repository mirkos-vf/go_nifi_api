package main

type RevisionDTO struct {
	ClientId     string `json:"clientid"`
	Version      int    `json:"version"`
	LastModifier string `json:"lastmodifier"`
}

type PositionDTO struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type AccessConfigurationDTO struct {
	SupportsLogin bool `json:"supportslogin"`
}

type AccessStatusDTO struct {
	Identity string `json:"identity"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

type AccessStatusEntity struct {
	AccessStatus AccessStatusDTO `json:"accessstatus"`
}

type PermissionsDTO struct {
	CanRead  bool `json:"canread"`
	CanWrite bool `json:"canwrite"`
}

type BulletinDTO struct {
	Id          int64  `json:"id"`
	NodeAddress string `json:"nodeaddress"`
	Category    string `json:"category"`
	GroupId     string `json:"groupid"`
	SourceId    string `json:"sourceid"`
	SourceName  string `json:"sourcename"`
	Level       string `json:"level"`
	Message     string `json:"message"`
	Timestamp   string `json:"TIMESTAMP"`
}

type BulletinEntity struct {
	Id          int64       `json:"id"`
	GroupId     string      `json:"groupid"`
	SourceId    string      `json:"sourceid"`
	Timestamp   string      `json:"timestamp"`
	NodeAddress string      `json:"nodeaddress"`
	CanRead     bool        `json:"canread"`
	Bulletin    BulletinDTO `json:"bulletin"`
}

type object struct {
	Values map[string]string `json:"values"`
}

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

type RemoteProcessGroupDTO struct {
	Id                            string                        `json:"id"`
	VersionedComponentId          string                        `json:"versionedcomponentid"`
	ParentGroupId                 string                        `json:"parentgroupid"`
	Position                      PositionDTO                   `json:"position"`
	TargetUri                     string                        `json:"targeturi"`
	TargetUris                    string                        `json:"targeturis"`
	TargetSecure                  bool                          `json:"targetsecure"`
	Name                          string                        `json:"name"`
	Comments                      string                        `json:"comments"`
	CommunicationsTimeout         string                        `json:"communicationstimeout"`
	YieldDuration                 string                        `json:"yieldduration"`
	TransportProtocol             string                        `json:"transportprotocol"`
	LocalNetworkInterface         string                        `json:"localnetworkinterface"`
	ProxyHost                     string                        `json:"proxyhost"`
	ProxyPort                     int32                         `json:"proxyport"`
	ProxyUser                     string                        `json:"proxyuser"`
	ProxyPassword                 string                        `json:"proxypassword"`
	AuthorizationIssues           []string                      `json:"authorizationissues"`
	ValidationErrors              []string                      `json:"validationerrors"`
	Transmitting                  bool                          `json:"transmitting"`
	InputPortCount                int32                         `json:"inputportcount"`
	OutputPortCount               int32                         `json:"outputportcount"`
	ActiveRemoteInputPortCount    int32                         `json:"activeremoteinputportcount"`
	InactiveRemoteInputPortCount  int32                         `json:"inactiveremoteinputportcount"`
	ActiveRemoteOutputPortCount   int32                         `json:"activeremoteoutputportcount"`
	InactiveRemoteOutputPortCount int32                         `json:"inactiveremoteoutputportcount"`
	FlowRefreshed                 string                        `json:"flowrefreshed"`
	Contents                      RemoteProcessGroupContentsDTO `json:"contents"`
}

type RemoteProcessGroupContentsDTO struct {
	InputPorts  []RemoteProcessGroupPortDTO `json:"inputports"`
	OutputPorts []RemoteProcessGroupPortDTO `json:"outputports"`
}

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

type BatchSettingsDTO struct {
	Count    int32  `json:"count"`
	Size     string `json:"size"`
	Duration string `json:"duration"`
}

type ProcessorDTO struct {
	Id                         string            `json:"id"`
	VersionedComponentId       string            `json:"versionedcomponentid"`
	ParentGroupId              string            `json:"parentgroupid"`
	Position                   PositionDTO       `json:"position"`
	Name                       string            `json:"name"`
	Type                       string            `json:"type"`
	Bundle                     BundleDTO         `json:"bundle"`
	State                      string            `json:"state"`
	Style                      object            `json:"style"`
	Relationships              []RelationshipDTO `json:"relationships"`
	Description                string            `json:"description"`
	SupportsParallelProcessing bool              `json:"supportsparallelprocessing"`
	SupportsEventDriven        bool              `json:"supportseventdriven"`
	SupportsBatching           bool              `json:"supportsbatching"`
	PersistsState              bool              `json:"persistsstate"`
	InputRequirement           string            `json:"inputrequirement"`
	Config                     string            `json:"config"`
	ValidationErrors           []string          `json:"validationerrors"`
}

type RelationshipDTO struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	AutoTerminate bool   `json:"autoterminate"`
}

type BundleDTO struct {
	Group    string `json:"group"`
	Artifact string `json:"artifact"`
	Version  string `json:"version"`
}

type PortDTO struct {
	Id                               string      `json:"id"`
	VersionedComponentId             string      `json:"versionedcomponentid"`
	ParentGroupId                    string      `json:"parentgroupid"`
	Position                         PositionDTO `json:"position"`
	Name                             string      `json:"name"`
	Comments                         string      `json:"comments"`
	State                            string      `json:"state"`
	Type                             string      `json:"type"`
	Transmitting                     bool        `json:"transmitting"`
	ConcurrentlySchedulableTaskCount int32       `json:"concurrentlyschedulabletaskcount"`
	UserAccessControl                []string    `json:"useraccesscontrol"`
	GroupAccessControl               []string    `json:"groupaccesscontrol"`
	ValidationErrors                 []string    `json:"validationerrors"`
}

type ConnectionDTO struct {
	Id                            string         `json:"id"`
	VersionedComponentId          string         `json:"versionedcomponentid"`
	ParentGroupId                 string         `json:"parentgroupid"`
	Position                      PositionDTO    `json:"position"`
	Source                        ConnectableDTO `json:"source"`
	Destination                   ConnectableDTO `json:"destination"`
	Name                          string         `json:"name"`
	LabelIndex                    int32          `json:"labelindex"`
	GetzIndex                     int32          `json:"getzindex"`
	SelectedRelationships         []string       `json:"selectedrelationships"`
	AvailableRelationships        []string       `json:"availablerelationships"`
	BackPressureObjectThreshold   int32          `json:"backpressureobjectthreshold"`
	BackPressureDataSizeThreshold string         `json:"backpressuredatasizethreshold"`
	FlowFileExpiration            string         `json:"flowfileexpiration"`
	Prioritizers                  []string       `json:"prioritizers"`
	Bends                         []PositionDTO  `json:"bends"`
	LoadBalanceStrategy           string         `json:"loadbalancestrategy"`
	LoadBalancePartitionAttribute string         `json:"loadbalancepartitionattribute"`
	LoadBalanceCompression        string         `json:"loadbalancecompression"`
	LoadBalanceStatus             string         `json:"loadbalancestatus"`
}

type ConnectableDTO struct {
	Id                   string `json:"id"`
	VersionedComponentId string `json:"versionedcomponentid"`
	Type                 string `json:"type"`
	GroupId              string `json:"groupid"`
	Name                 string `json:"name"`
	Running              bool   `json:"running"`
	Transmitting         bool   `json:"transmitting"`
	Exists               bool   `json:"exists"`
	Comments             string `json:"сomments"`
}

type LabelDTO struct {
	Id                   string      `json:"id"`
	VersionedComponentId string      `json:"versionedcomponentid"`
	ParentGroupId        string      `json:"parentgroupid"`
	Position             PositionDTO `json:"position"`
	Label                string      `json:"label"`
	Width                float32     `json:"width"`
	Height               float32     `json:"height"`
	Style                object      `json:"style"`
}

type FunnelDTO struct {
	Id                   string      `json:"id"`
	VersionedComponentId string      `json:"versionedcomponentid"`
	ParentGroupId        string      `json:"parentgroupid"`
	Position             PositionDTO `json:"position"`
}

type ControllerServiceDTO struct {
	Id                        string                                        `json:"id"`
	VersionedComponentId      string                                        `json:"versionedcomponentid"`
	ParentGroupId             string                                        `json:"parentgroupid"`
	Position                  PositionDTO                                   `json:"position"`
	Name                      string                                        `json:"name"`
	Type                      string                                        `json:"type"`
	Bundle                    BundleDTO                                     `json:"bundle"`
	ControllerServiceApis     []ControllerServiceApiDTO                     `json:"controllerserviceapis"`
	Comments                  string                                        `json:"comments"`
	State                     string                                        `json:"state"`
	PersistsState             bool                                          `json:"persistsstate"`
	Restricted                bool                                          `json:"restricted"`
	Deprecated                bool                                          `json:"deprecated"`
	MultipleVersionsAvailable bool                                          `json:"multipleversionsavailable"`
	Properties                object                                        `json:"properties"`
	Descriptors               object                                        `json:"descriptors"`
	CustomUiUrl               string                                        `json:"customuiurl"`
	AnnotationData            string                                        `json:"annotationdata"`
	ReferencingComponents     []ControllerServiceReferencingComponentEntity `json:"referencingcomponents"`
	ValidationErrors          []string                                      `json:"validationerrors"`
	ValidationStatus          string                                        `json:"validationstatus"`
	ExtensionMissing          bool                                          `json:"extensionmissing"`
}

type ControllerServiceApiDTO struct {
	Type   string    `json:"type"`
	Bundle BundleDTO `json:"bundle"`
}

type ControllerServiceReferencingComponentEntity struct {
	Revision                     RevisionDTO                              `json:"revision"`
	Id                           string                                   `json:"id"`
	Uri                          string                                   `json:"uri"`
	Position                     PositionDTO                              `json:"position"`
	Permissions                  PermissionsDTO                           `json:"permissions"`
	Bulletins                    []BulletinEntity                         `json:"bulletins"`
	DisconnectedNodeAcknowledged bool                                     `json:"disconnectednodeacknowledged"`
	Component                    ControllerServiceReferencingComponentDTO `json:"component"`
	OperatePermissions           PermissionsDTO                           `json:"operatepermissions"`
}

type ControllerServiceReferencingComponentDTO struct {
	GroupId               string                                        `json:"groupid"`
	Id                    string                                        `json:"id"`
	Name                  string                                        `json:"name"`
	Type                  string                                        `json:"type"`
	State                 string                                        `json:"state"`
	Properties            object                                        `json:"properties"`
	Descriptors           object                                        `json:"descriptors"`
	ValidationErrors      []string                                      `json:"validationerrors"`
	ReferenceType         string                                        `json:"referencetype"`
	ActiveThreadCount     int32                                         `json:"activethreadcount"`
	ReferenceCycle        bool                                          `json:"referencecycle"`
	ReferencingComponents []ControllerServiceReferencingComponentEntity `json:"referencingcomponents"`
}

type FlowSnippetDTO struct {
	ProcessGroups       []ProcessGroupDTO       `json:"processgroups"`
	RemoteProcessGroups []RemoteProcessGroupDTO `json:"remoteprocessgroups"`
	Processors          []ProcessorDTO          `json:"processors"`
	InputPorts          []PortDTO               `json:"inputports"`
	OutputPorts         []PortDTO               `json:"outputports"`
	Connections         []ConnectionDTO         `json:"connections"`
	Labels              []LabelDTO              `json:"labels"`
	Funnels             []FunnelDTO             `json:"funnels"`
	ControllerServices  []ControllerServiceDTO  `json:"controllerservices"`
}

type VersionControlInformationDTO struct {
	GroupId          string `json:"groupid"`
	RegistryId       string `json:"registryid"`
	RegistryName     string `json:"registryname"`
	BucketId         string `json:"bucketid"`
	BucketName       string `json:"bucketname"`
	FlowId           string `json:"flowid"`
	FlowName         string `json:"flowname"`
	FlowDescription  string `json:"flowdescription"`
	Version          int32  `json:"version"`
	State            string `json:"state"`
	StateExplanation string `json:"stateexplanation"`
}

type ProcessGroupStatusDTO struct {
	Id                 string                              `json:"id"`
	Name               string                              `json:"name"`
	StatsLastRefreshed string                              `json:"statslastrefreshed"`
	AggregateSnapshot  ProcessGroupStatusSnapshotDTO       `json:"aggregatesnapshot"`
	NodeSnapshots      []NodeProcessGroupStatusSnapshotDTO `json:"nodesnapshots"`
}

type Link struct {
	Type       string     `json:"type"`
	Rels       []string   `json:"rels"`
	Rel        string     `json:"rel"`
	Uri        string     `json:"uri"`
	Params     object     `json:"params"`
	Title      string     `json:"title"`
	UriBuilder UriBuilder `json:"uribuilder"`
}

type UriBuilder struct {
	// TODO: UriBuilder
}

type VersionedFlowSnapshotMetadata struct {
	Link             Link   `json:"link"`
	BucketIdentifier string `json:"bucketidentifier"`
	FlowIdentifier   string `json:"flowidentifier"`
	Version          int32  `json:"version"`
	Timestamp        int64  `json:"timestamp"`
	Author           string `json:"author"`
	Comments         string `json:"comments"`
}

type VersionedRemoteProcessGroup struct {
	Identifier            string        `json:"identifier"`
	Name                  string        `json:"name"`
	Comments              string        `json:"comments"`
	Position              Position      `json:"position"`
	TargetUri             string        `json:"targeturi"`
	TargetUris            string        `json:"targeturis"`
	CommunicationsTimeout string        `json:"communicationstimeout"`
	YieldDuration         string        `json:"yieldduration"`
	TransportProtocol     string        `json:"transportprotocol"`
	LocalNetworkInterface string        `json:"localnetworkinterface"`
	ProxyHost             string        `json:"proxyhost"`
	ProxyPort             int32         `json:"proxyport"`
	ProxyUser             string        `json:"proxyuser"`
	InputPorts            []InputPorts  `json:"inputports"`
	OutputPorts           []OutputPorts `json:"outputports"`
	ComponentType         string        `json:"componenttype"`
	GroupIdentifier       string        `json:"groupidentifier"`
}

type InputPorts struct {
	// TODO: InputPorts in RUN
}

type OutputPorts struct {
	// TODO: OutputPorts
}

type UsersEntity struct {
	Generated string       `json:"generated"`
	Users     []UserEntity `json:"users"`
}

type UserEntity struct {
	Revision                     RevisionDTO      `json:"revision"`
	Id                           string           `json:"id"`
	Uri                          string           `json:"uri"`
	Position                     PositionDTO      `json:"position"`
	Permissions                  PermissionsDTO   `json:"permissions"`
	Bulletins                    []BulletinEntity `json:"bulletins"`
	DisconnectedNodeAcknowledged bool             `json:"disconnectednodeacknowledged"`
	Component                    UserDTO          `json:"component"`
}

type UserDTO struct {
	Id                   string                      `json:"id"`
	VersionedComponentId string                      `json:"versionedcomponentid"`
	ParentGroupId        string                      `json:"permissionsdto"`
	Position             PositionDTO                 `json:"position"`
	Identity             string                      `json:"identity"`
	Configurable         bool                        `json:"configurable"`
	UserGroups           []TenantEntity              `json:"usergroups"`
	AccessPolicies       []AccessPolicySummaryEntity `json:"accesspolicies"`
}

type AccessPolicySummaryEntity struct {
	// TODO: AccessPolicySummaryEntity
}

type UserGroupEntity struct {
	Revision                     RevisionDTO      `json:"revision"`
	Id                           string           `json:"id"`
	Uri                          string           `json:"uri"`
	Position                     PositionDTO      `json:"position"`
	Permissions                  PermissionsDTO   `json:"permissions"`
	Bulletins                    []BulletinEntity `json:"bulletins"`
	DisconnectedNodeAcknowledged bool             `json:"disconnectednodeacknowledged"`
	Component                    UserGroupDTO     `json:"component"`
}

type UserGroupDTO struct {
	Id                   string               `json:"id"`
	VersionedComponentId string               `json:"versionedcomponentid"`
	ParentGroupId        string               `json:"parentgroupid"`
	Position             PositionDTO          `json:"position"`
	Identity             string               `json:"identity"`
	Configurable         bool                 `json:"configurable"`
	Users                []TenantEntity       `json:"users"`
	AccessPolicies       []AccessPolicyEntity `json:"accesspolicies"`
}

type TenantEntity struct {
	// TODO: TenantEntity
}

type AccessPolicyEntity struct {
	// TODO: AccessPolicyEntity
}

type UserGroupsEntity struct {
	UserGroups []UserGroupEntity `json:"usergroups"`
}

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

type VersionedPort struct {
	Identifier                       string   `json:"identifier"`
	Name                             string   `json:"name"`
	Comments                         string   `json:"comments"`
	Position                         Position `json:"position"`
	Type                             string   `json:"type"`
	ConcurrentlySchedulableTaskCount int32    `json:"concurrentlyschedulabletaskcount"`
	ComponentType                    string   `json:"componenttype"`
	GroupIdentifier                  string   `json:"groupidentifier"`
}

type VersionedConnection struct {
	Identifier                    string               `json:"identifier"`
	Name                          string               `json:"name"`
	Comments                      string               `json:"comments"`
	Position                      Position             `json:"position"`
	Source                        ConnectableComponent `json:"source"`
	Destination                   ConnectableComponent `json:"destination"`
	LabelIndex                    int32                `json:"labelindex"`
	ZIndex                        int64                `json:"zindex"`
	SelectedRelationships         []string             `json:"selectedrelationships"`
	BackPressureObjectThreshold   int64                `json:"backpressureobjectthreshold"`
	BackPressureDataSizeThreshold string               `json:"backpressuredatasizethreshold"`
	FlowFileExpiration            string               `json:"flowfileexpiration"`
	Prioritizers                  []string             `json:"prioritizers"`
	Bends                         []Position           `json:"bends"`
	LoadBalanceStrategy           string               `json:"loadbalancestrategy"`
	PartitioningAttribute         string               `json:"partitioningattribute"`
	LoadBalanceCompression        string               `json:"loadbalancecompression"`
	ComponentType                 string               `json:"componenttype"`
	GroupIdentifier               string               `json:"groupidentifier"`
}

type ConnectableComponent struct {
	// TODO: fields
}

type VersionedLabel struct {
	Identifier      string   `json:"identifier"`
	Name            string   `json:"name"`
	Comments        string   `json:"comments"`
	Position        Position `json:"position"`
	Label           string   `json:"label"`
	Width           float32  `json:"width"`
	Height          float32  `json:"height"`
	Style           object   `json:"style"`
	ComponentType   string   `json:"componenttype"`
	GroupIdentifier string   `json:"groupidentifier"`
}

type VersionedFunnel struct {
	Identifier      string   `json:"identifier"`
	Name            string   `json:"name"`
	Comments        string   `json:"comments"`
	Position        Position `json:"position"`
	ComponentType   string   `json:"componenttype"`
	GroupIdentifier string   `json:"groupidentifier"`
}

type ControllerServiceAPI struct {
	Type   string `json:"type"`
	Bundle Bundle `json:"bundle"`
}

type VersionedControllerService struct {
	Identifier            string                 `json:"identifier"`
	Name                  string                 `json:"name"`
	Comments              string                 `json:"comments"`
	Position              string                 `json:"position"`
	Type                  string                 `json:"type"`
	Bundle                Bundle                 `json:"bundle"`
	ControllerServiceApis []ControllerServiceAPI `json:"controllerserviceapis"`
	Properties            object                 `json:"properties"`
	PropertyDescriptors   object                 `json:"propertydescriptors"`
	AnnotationData        string                 `json:"annotationdata"`
	ComponentType         string                 `json:"componenttype"`
	GroupIdentifier       string                 `json:"groupidentifier"`
}

type Bundle struct {
	Group    string `json:"group"`
	Artifact string `json:"artifact"`
	Version  string `json:"version"`
}

type VersionedFlowCoordinates struct {
	RegistryUrl string `json:"registryurl"`
	BucketId    string `json:"bucketid"`
	FlowId      string `json:"flowid"`
	Version     int32  `json:"version"`
	Latest      bool   `json:"latest"`
}

type VersionedProcessGroup struct {
	Identifier               string                        `json:"identifier"`
	Name                     string                        `json:"name"`
	Comments                 string                        `json:"comments"`
	Position                 Position                      `json:"position"`
	ProcessGroups            []VersionedProcessGroup       `json:"processgroups"`
	RemoteProcessGroups      []VersionedRemoteProcessGroup `json:"remoteprocessgroups"`
	Processors               []VersionedProcessor          `json:"processors"`
	InputPorts               []VersionedPort               `json:"inputports"`
	OutputPorts              []VersionedPort               `json:"outputports"`
	Connections              []VersionedConnection         `json:"connections"`
	Labels                   []VersionedLabel              `json:"labels"`
	Funnels                  []VersionedFunnel             `json:"funnels"`
	ControllerServices       []VersionedControllerService  `json:"controllerservices"`
	VersionedFlowCoordinates VersionedFlowCoordinates      `json:"versionedflowcoordinates"`
	Variables                object                        `json:"variables"`
	ComponentType            string                        `json:"componenttype"`
	GroupIdentifier          string                        `json:"groupidentifier"`
}

type Permissions struct {
	CanRead   bool `json:"canread"`
	CanWrite  bool `json:"canwrite"`
	CanDelete bool `json:"candelete"`
}

type VersionedFlow struct {
	Link              Link        `json:"link"`
	Identifier        string      `json:"identifier"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	BucketIdentifier  string      `json:"bucketidentifier"`
	BucketName        string      `json:"bucketname"`
	CreatedTimestamp  int64       `json:"createdtimestamp"`
	ModifiedTimestamp int64       `json:"modifiedtimestamp"`
	Type              string      `json:"type"`
	Permissions       Permissions `json:"permissions"`
	VersionCount      int64       `json:"versioncount"`
}

type Bucket struct {
	Link             Link        `json:"link"`
	Identifier       string      `json:"identifier"`
	Name             string      `json:"name"`
	CreatedTimestamp int64       `json:"createdtimestamp"`
	Description      string      `json:"description"`
	Permissions      Permissions `json:"permissions"`
}

type NodeProcessGroupStatusSnapshotDTO struct {
	// TODO: NodeProcessGroupStatusSnapshotDTO
}

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

type ConnectionStatusSnapshotEntity struct {
	Id                       string                      `json:"id"`
	ConnectionStatusSnapshot ConnectionStatusSnapshotDTO `json:"connectionstatussnapshot"`
	CanRead                  bool                        `json:"canread"`
}

type ProcessorStatusSnapshotEntity struct {
	Id                      string                     `json:"id"`
	ProcessorStatusSnapshot ProcessorStatusSnapshotDTO `json:"processorstatussnapshot"`
	CanRead                 bool                       `json:"canread"`
}

type ProcessGroupStatusSnapshotEntity struct {
	Id                         string                        `json:"id"`
	ProcessGroupStatusSnapshot ProcessGroupStatusSnapshotDTO `json:"processgroupstatussnapshot"`
	CanRead                    bool                          `json:"canread"`
}

type RemoteProcessGroupStatusSnapshotEntity struct {
	Id                               string                              `json:"id"`
	RemoteProcessGroupStatusSnapshot RemoteProcessGroupStatusSnapshotDTO `json:"remoteprocessgroupstatussnapshot"`
	CanRead                          bool                                `json:"canread"`
}

type PortStatusSnapshotEntity struct {
	Id                 string                `json:"id"`
	PortStatusSnapshot PortStatusSnapshotDTO `json:"portStatusSnapshot"`
	CanRead            bool                  `json:"canread"`
}

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

type versionedFlowSnapshot struct {
	SnapshotMetadata VersionedFlowSnapshotMetadata `json:"snapshotmetadata"`
	FlowContents     VersionedProcessGroup         `json:"flowcontents"`
	Flow             VersionedFlow                 `json:"flow"`
	Bucket           Bucket                        `json:"bucket"`
	Latest           bool                          `json:"latest"`
}

type OutputPortsEntity struct {
	OutputPorts []PortEntity `json:"outputports"`
}

type PortEntity struct {
	Revision                     RevisionDTO      `json:"revision"`
	Id                           string           `json:"id"`
	Uri                          string           `json:"uri"`
	Position                     PositionDTO      `json:"position"`
	Permissions                  PermissionsDTO   `json:"permissions"`
	Bulletins                    []BulletinEntity `json:"bulletins"`
	DisconnectedNodeAcknowledged bool             `json:"disconnected_node_acknowledged"`
	Component                    PortDTO          `json:"component"`
	Status                       PortStatusDTO    `json:"status"`
	PortType                     string           `json:"port_type"`
	OperatePermissions           PermissionsDTO   `json:"operate_permissions"`
}

type PortStatusDTO struct {
	Id                 string                      `json:"id"`
	GroupId            string                      `json:"group_id"`
	Name               string                      `json:"name"`
	Transmitting       bool                        `json:"transmitting"`
	RunStatus          string                      `json:"run_status"`
	StatsLastRefreshed string                      `json:"stats_last_refreshed"`
	AggregateSnapshot  PortStatusSnapshotDTO       `json:"aggregate_snapshot"`
	NodeSnapshots      []NodePortStatusSnapshotDTO `json:"node_snapshots"`
}

type NodePortStatusSnapshotDTO struct {
	NodeId         string                `json:"node_id"`
	Address        string                `json:"address"`
	ApiPort        int32                 `json:"api_port"`
	StatusSnapshot PortStatusSnapshotDTO `json:"status_snapshot"`
}

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
