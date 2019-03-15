package main

import (
	"go/types"
)

type RevisionDTO struct {
	ClientId     string `json:"client_id"`
	Version      int    `json:"version"`
	LastModifier string `json:"last_modifier"`
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
	SupportsLogin bool `json:"supports_login"`
}

type AccessStatusDTO struct {
	Identity string `json:"identity"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

type AccessStatusEntity struct {
	AccessStatus AccessStatusDTO `json:"access_status"`
}

type PermissionsDTO struct {
	CanRead  bool `json:"can_read"`
	CanWrite bool `json:"can_write"`
}

type BulletinDTO struct {
}

type BulletinEntity struct {
	Id          int64       `json:"id"`
	GroupId     string      `json:"group_id"`
	SourceId    string      `json:"source_id"`
	Timestamp   string      `json:"timestamp"`
	NodeAddress string      `json:"node_address"`
	CanRead     bool        `json:"can_read"`
	Bulletin    BulletinDTO `json:"bulletin"`
}

type ProcessGroupDTO struct {
	Id                           string                       `json:"id"`
	VersionedComponentId         string                       `json:"versioned_component_id"`
	ParentGroupId                string                       `json:"parent_group_id"`
	Position                     PositionDTO                  `json:"position"`
	Name                         string                       `json:"name"`
	Comments                     string                       `json:"comments"`
	Variables                    types.Object                 `json:"variables"`
	VersionControlInformation    VersionControlInformationDTO `json:"version_control_information"`
	RunningCount                 int32                        `json:"running_count"`
	StoppedCount                 int32                        `json:"stopped_count"`
	InvalidCount                 int32                        `json:"invalid_count"`
	DisabledCount                int32                        `json:"disabled_count"`
	InactiveRemotePortCount      int32                        `json:"inactive_remote_port_count"`
	UpToDateCount                int32                        `json:"up_to_date_count"`
	LocallyModifiedCount         int32                        `json:"locally_modified_count"`
	StaleCount                   int32                        `json:"stale_count"`
	LocallyModifiedAndStaleCount int32                        `json:"locally_modified_and_stale_count"`
	SyncFailureCount             int32                        `json:"sync_failure_count"`
	InputPortCount               int32                        `json:"input_port_count"`
	OutputPortCount              int32                        `json:"output_port_count"`
	Contents                     FlowSnippetDTO               `json:"contents"`
}

type FlowSnippetDTO struct {
}

type VersionControlInformationDTO struct {
}

type ProcessGroupStatusDTO struct {
	Id                 string                              `json:"id"`
	Name               string                              `json:"name"`
	StatsLastRefreshed string                              `json:"stats_last_refreshed"`
	AggregateSnapshot  ProcessGroupStatusSnapshotDTO       `json:"aggregate_snapshot"`
	NodeSnapshots      []NodeProcessGroupStatusSnapshotDTO `json:"node_snapshots"`
}

type Link struct {
	Type       string       `json:"type"`
	Rels       []string     `json:"rels"`
	Rel        string       `json:"rel"`
	Uri        string       `json:"uri"`
	Params     types.Object `json:"params"`
	Title      string       `json:"title"`
	UriBuilder UriBuilder   `json:"uri_builder"`
}

type UriBuilder struct{}

type VersionedFlowSnapshotMetadata struct {
	Link             Link   `json:"link"`
	BucketIdentifier string `json:"bucket_identifier"`
	FlowIdentifier   string `json:"flow_identifier"`
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
	TargetUri             string        `json:"target_uri"`
	TargetUris            string        `json:"target_uris"`
	CommunicationsTimeout string        `json:"communications_timeout"`
	YieldDuration         string        `json:"yield_duration"`
	TransportProtocol     string        `json:"transport_protocol"`
	LocalNetworkInterface string        `json:"local_network_interface"`
	ProxyHost             string        `json:"proxy_host"`
	ProxyPort             int32         `json:"proxy_port"`
	ProxyUser             string        `json:"proxy_user"`
	InputPorts            []InputPorts  `json:"input_ports"`
	OutputPorts           []OutputPorts `json:"output_ports"`
	ComponentType         string        `json:"component_type"`
	GroupIdentifier       string        `json:"group_identifier"`
}

type InputPorts struct {
	// TODO: InputPorts
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
	DisconnectedNodeAcknowledged bool             `json:"disconnected_node_acknowledged"`
	Component                    UserDTO          `json:"component"`
}

type UserDTO struct {
	Id                   string                      `json:"id"`
	VersionedComponentId string                      `json:"versioned_component_id"`
	ParentGroupId        string                      `json:"permissions_dto"`
	Position             PositionDTO                 `json:"position"`
	Identity             string                      `json:"identity"`
	Configurable         bool                        `json:"configurable"`
	UserGroups           []TenantEntity              `json:"user_groups"`
	AccessPolicies       []AccessPolicySummaryEntity `json:"access_policies"`
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
	DisconnectedNodeAcknowledged bool             `json:"disconnected_node_acknowledged"`
	Component                    UserGroupDTO     `json:"component"`
}

type UserGroupDTO struct {
	Id                   string               `json:"id"`
	VersionedComponentId string               `json:"versioned_component_id"`
	ParentGroupId        string               `json:"parent_group_id"`
	Position             PositionDTO          `json:"position"`
	Identity             string               `json:"identity"`
	Configurable         bool                 `json:"configurable"`
	Users                []TenantEntity       `json:"users"`
	AccessPolicies       []AccessPolicyEntity `json:"access_policies"`
}

type TenantEntity struct {
	// TODO: TenantEntity
}

type AccessPolicyEntity struct {
	// TODO: AccessPolicyEntity
}

type UserGroupsEntity struct {
	UserGroups []UserGroupEntity `json:"user_groups"`
}

type VersionedProcessor struct {
	Identifier                       string       `json:"identifier"`
	Name                             string       `json:"name"`
	Comments                         string       `json:"comments"`
	Position                         Position     `json:"position"`
	Bundle                           Bundle       `json:"bundle"`
	Style                            types.Object `json:"style"`
	Type                             string       `json:"type"`
	Properties                       types.Object `json:"properties"`
	PropertyDescriptors              types.Object `json:"property_descriptors"`
	AnnotationData                   string       `json:"annotation_data"`
	SchedulingPeriod                 string       `json:"scheduling_period"`
	SchedulingStrategy               string       `json:"scheduling_strategy"`
	ExecutionNode                    string       `json:"execution_node"`
	PenaltyDuration                  string       `json:"penalty_duration"`
	YieldDuration                    string       `json:"yield_duration"`
	BulletinLevel                    string       `json:"bulletin_level"`
	RunDurationMillis                int64        `json:"run_duration_millis"`
	ConcurrentlySchedulableTaskCount int64        `json:"concurrently_schedulable_task_count"`
	AutoTerminatedRelationships      []string     `json:"auto_terminated_relationships"`
	ComponentType                    string       `json:"component_type"`
	GroupIdentifier                  string       `json:"group_identifier"`
}

type VersionedPort struct {
	Identifier                       string   `json:"identifier"`
	Name                             string   `json:"name"`
	Comments                         string   `json:"comments"`
	Position                         Position `json:"position"`
	Type                             string   `json:"type"`
	ConcurrentlySchedulableTaskCount int32    `json:"concurrently_schedulable_task_count"`
	ComponentType                    string   `json:"component_type"`
	GroupIdentifier                  string   `json:"group_identifier"`
}

type VersionedConnection struct {
	Identifier                    string               `json:"identifier"`
	Name                          string               `json:"name"`
	Comments                      string               `json:"comments"`
	Position                      Position             `json:"position"`
	Source                        ConnectableComponent `json:"source"`
	Destination                   ConnectableComponent `json:"destination"`
	LabelIndex                    int32                `json:"label_index"`
	ZIndex                        int64                `json:"z_index"`
	SelectedRelationships         []string             `json:"selected_relationships"`
	BackPressureObjectThreshold   int64                `json:"back_pressure_object_threshold"`
	BackPressureDataSizeThreshold string               `json:"back_pressure_data_size_threshold"`
	FlowFileExpiration            string               `json:"flow_file_expiration"`
	Prioritizers                  []string             `json:"prioritizers"`
	Bends                         []Position           `json:"bends"`
	LoadBalanceStrategy           string               `json:"load_balance_strategy"`
	PartitioningAttribute         string               `json:"partitioning_attribute"`
	LoadBalanceCompression        string               `json:"load_balance_compression"`
	ComponentType                 string               `json:"component_type"`
	GroupIdentifier               string               `json:"group_identifier"`
}

type ConnectableComponent struct {
	// TODO: fields
}

type VersionedLabel struct {
	Identifier      string       `json:"identifier"`
	Name            string       `json:"name"`
	Comments        string       `json:"comments"`
	Position        Position     `json:"position"`
	Label           string       `json:"label"`
	Width           float32      `json:"width"`
	Height          float32      `json:"height"`
	Style           types.Object `json:"style"`
	ComponentType   string       `json:"component_type"`
	GroupIdentifier string       `json:"group_identifier"`
}

type VersionedFunnel struct {
	Identifier      string   `json:"identifier"`
	Name            string   `json:"name"`
	Comments        string   `json:"comments"`
	Position        Position `json:"position"`
	ComponentType   string   `json:"component_type"`
	GroupIdentifier string   `json:"group_identifier"`
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
	ControllerServiceApis []ControllerServiceAPI `json:"controller_service_apis"`
	Properties            types.Object           `json:"properties"`
	PropertyDescriptors   types.Object           `json:"property_descriptors"`
	AnnotationData        string                 `json:"annotation_data"`
	ComponentType         string                 `json:"component_type"`
	GroupIdentifier       string                 `json:"group_identifier"`
}

type Bundle struct {
	Group    string `json:"group"`
	Artifact string `json:"artifact"`
	Version  string `json:"version"`
}

type VersionedFlowCoordinates struct {
	RegistryUrl string `json:"registry_url"`
	BucketId    string `json:"bucket_id"`
	FlowId      string `json:"flow_id"`
	Version     int32  `json:"version"`
	Latest      bool   `json:"latest"`
}

type VersionedProcessGroup struct {
	Identifier               string                        `json:"identifier"`
	Name                     string                        `json:"name"`
	Comments                 string                        `json:"comments"`
	Position                 Position                      `json:"position"`
	ProcessGroups            []VersionedProcessGroup       `json:"process_groups"`
	RemoteProcessGroups      []VersionedRemoteProcessGroup `json:"remote_process_groups"`
	Processors               []VersionedProcessor          `json:"processors"`
	InputPorts               []VersionedPort               `json:"input_ports"`
	OutputPorts              []VersionedPort               `json:"output_ports"`
	Connections              []VersionedConnection         `json:"connections"`
	Labels                   []VersionedLabel              `json:"labels"`
	Funnels                  []VersionedFunnel             `json:"funnels"`
	ControllerServices       []VersionedControllerService  `json:"controller_services"`
	VersionedFlowCoordinates VersionedFlowCoordinates      `json:"versioned_flow_coordinates"`
	Variables                types.Object                  `json:"variables"`
	ComponentType            string                        `json:"component_type"`
	GroupIdentifier          string                        `json:"group_identifier"`
}

type Permissions struct {
	CanRead   bool `json:"can_read"`
	CanWrite  bool `json:"can_write"`
	CanDelete bool `json:"can_delete"`
}

type VersionedFlow struct {
	Link              Link        `json:"link"`
	Identifier        string      `json:"identifier"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	BucketIdentifier  string      `json:"bucket_identifier"`
	BucketName        string      `json:"bucket_name"`
	CreatedTimestamp  int64       `json:"created_timestamp"`
	ModifiedTimestamp int64       `json:"modified_timestamp"`
	Type              string      `json:"type"`
	Permissions       Permissions `json:"permissions"`
	VersionCount      int64       `json:"version_count"`
}

type Bucket struct {
	Link             Link        `json:"link"`
	Identifier       string      `json:"identifier"`
	Name             string      `json:"name"`
	CreatedTimestamp int64       `json:"created_timestamp"`
	Description      string      `json:"description"`
	Permissions      Permissions `json:"permissions"`
}

type NodeProcessGroupStatusSnapshotDTO struct {
}

type ProcessGroupStatusSnapshotDTO struct {
	Id                                string                                   `json:"id"`
	Name                              string                                   `json:"name"`
	ConnectionStatusSnapshots         []ConnectionStatusSnapshotEntity         `json:"connection_status_snapshots"`
	ProcessorStatusSnapshots          []ProcessorStatusSnapshotEntity          `json:"processor_status_snapshots"`
	ProcessGroupStatusSnapshots       []ProcessGroupStatusSnapshotEntity       `json:"process_group_status_snapshots"`
	RemoteProcessGroupStatusSnapshots []RemoteProcessGroupStatusSnapshotEntity `json:"remote_process_group_status_snapshots"`
	InputPortStatusSnapshots          []PortStatusSnapshotEntity               `json:"input_port_status_snapshots"`
	OutputPortStatusSnapshots         []PortStatusSnapshotEntity               `json:"output_port_status_snapshots"`
	VersionedFlowState                string                                   `json:"versioned_flow_state"`
	FlowFilesIn                       string                                   `json:"flow_files_in"`
	BytesIn                           int64                                    `json:"bytes_in"`
	Input                             string                                   `json:"input"`
	FlowFilesQueued                   int32                                    `json:"flow_files_queued"`
	BytesQueued                       int64                                    `json:"bytes_queued"`
	Queued                            string                                   `json:"queued"`
	QueuedCount                       string                                   `json:"queued_count"`
	QueuedSize                        string                                   `json:"queued_size"`
	BytesRead                         int64                                    `json:"bytes_read"`
	Read                              string                                   `json:"read"`
	BytesWritten                      int64                                    `json:"bytes_written"`
	Written                           string                                   `json:"written"`
	FlowFilesOut                      int32                                    `json:"flow_files_out"`
	BytesOut                          int64                                    `json:"bytes_out"`
	Output                            string                                   `json:"output"`
	FlowFilesTransferred              int32                                    `json:"flow_files_transferred"`
	BytesTransferred                  int64                                    `json:"bytes_transferred"`
	Transferred                       string                                   `json:"transferred"`
	BytesReceived                     int64                                    `json:"bytes_received"`
	FlowFilesReceived                 int32                                    `json:"flow_files_received"`
	Received                          string                                   `json:"received"`
	BytesSent                         int64                                    `json:"bytes_sent"`
	FlowFilesSent                     int32                                    `json:"flow_files_sent"`
	Sent                              string                                   `json:"sent"`
	ActiveThreadCount                 int32                                    `json:"active_thread_count"`
	TerminatedThreadCount             int32                                    `json:"terminated_thread_count"`
}

type ConnectionStatusSnapshotEntity struct {
	Id                       string                      `json:"id"`
	ConnectionStatusSnapshot ConnectionStatusSnapshotDTO `json:"connection_status_snapshot"`
	CanRead                  bool                        `json:"can_read"`
}

type ProcessorStatusSnapshotEntity struct {
	Id                      string                     `json:"id"`
	ProcessorStatusSnapshot ProcessorStatusSnapshotDTO `json:"processor_status_snapshot"`
	CanRead                 bool                       `json:"can_read"`
}

type ProcessGroupStatusSnapshotEntity struct {
	Id                         string                        `json:"id"`
	ProcessGroupStatusSnapshot ProcessGroupStatusSnapshotDTO `json:"process_group_status_snapshot"`
	CanRead                    bool                          `json:"can_read"`
}

type RemoteProcessGroupStatusSnapshotEntity struct {
	Id                               string                              `json:"id"`
	RemoteProcessGroupStatusSnapshot RemoteProcessGroupStatusSnapshotDTO `json:"remote_process_group_status_snapshot"`
	CanRead                          bool                                `json:"can_read"`
}

type PortStatusSnapshotEntity struct {
	Id                 string                `json:"id"`
	PortStatusSnapshot PortStatusSnapshotDTO `json:"portStatusSnapshot"`
	CanRead            bool                  `json:"can_read"`
}

type PortStatusSnapshotDTO struct {
	Id                string `json:"id"`
	GroupId           string `json:"group_id"`
	Name              string `json:"name"`
	ActiveThreadCount int32  `json:"active_thread_count"`
	FlowFilesIn       int32  `json:"flow_files_in"`
	BytesIn           int64  `json:"bytes_in"`
	Input             string `json:"input"`
	FlowFilesOut      int32  `json:"flow_files_out"`
	BytesOut          int64  `json:"bytes_out"`
	Output            string `json:"output"`
	Transmitting      bool   `json:"transmitting"`
	RunStatus         string `json:"run_status"`
}

type RemoteProcessGroupStatusSnapshotDTO struct {
	Id                 string `json:"id"`
	GroupId            string `json:"group_id"`
	Name               string `json:"name"`
	TargetUri          string `json:"target_uri"`
	TransmissionStatus string `json:"transmission_status"`
	ActiveThreadCount  int    `json:"active_thread_count"`
	FlowFilesSent      int    `json:"flow_files_sent"`
	BytesSent          int    `json:"bytes_sent"`
	Sent               string `json:"sent"`
	FlowFilesReceived  int    `json:"flow_files_received"`
	BytesReceived      int    `json:"bytes_received"`
	Received           string `json:"received"`
}

type ProcessorStatusSnapshotDTO struct {
	Id                    string `json:"id"`
	GroupId               string `json:"group_id"`
	Name                  string `json:"name"`
	Type                  string `json:"type"`
	RunStatus             string `json:"run_status"`
	ExecutionNode         string `json:"execution_node"`
	BytesRead             int64  `json:"bytes_read"`
	BytesWritten          int64  `json:"bytes_written"`
	Read                  string `json:"read"`
	Written               string `json:"written"`
	FlowFilesIn           int32  `json:"flow_files_in"`
	BytesIn               int64  `json:"bytes_in"`
	Input                 string `json:"input"`
	FlowFilesOut          int32  `json:"flow_files_out"`
	BytesOut              int64  `json:"bytes_out"`
	Output                string `json:"output"`
	TaskCount             int32  `json:"task_count"`
	TasksDurationNanos    int64  `json:"tasks_duration_nanos"`
	Tasks                 string `json:"tasks"`
	TasksDuration         string `json:"tasks_duration"`
	ActiveThreadCount     int32  `json:"active_thread_count"`
	TerminatedThreadCount int32  `json:"terminated_thread_count"`
}

type ConnectionStatusSnapshotDTO struct {
	Id              string `json:"id"`
	GroupId         string `json:"group_id"`
	Name            string `json:"name"`
	SourceId        string `json:"source_id"`
	SourceName      string `json:"source_name"`
	DestinationId   string `json:"destination_id"`
	DestinationName string `json:"destination_name"`
	FlowFilesIn     int32  `json:"flow_files_in"`
	BytesIn         int64  `json:"bytes_in"`
	Input           string `json:"input"`
	FlowFilesOut    int32  `json:"flow_files_out"`
	BytesOut        int64  `json:"bytes_out"`
	Output          string `json:"output"`
	FlowFilesQueued int32  `json:"flow_files_queued"`
	BytesQueued     int64  `json:"bytes_queued"`
	Queued          string `json:"queued"`
	QueuedSize      string `json:"queued_size"`
	QueuedCount     string `json:"queued_count"`
	PercentUseCount int32  `json:"percent_use_count"`
	PercentUseBytes int32  `json:"percent_use_bytes"`
}

type versionedFlowSnapshot struct {
	SnapshotMetadata VersionedFlowSnapshotMetadata `json:"snapshot_metadata"`
	FlowContents     VersionedProcessGroup         `json:"flow_contents"`
	Flow             VersionedFlow                 `json:"flow"`
	Bucket           Bucket                        `json:"bucket"`
	Latest           bool                          `json:"latest"`
}

type ProcessGroupEntity struct {
	Revision                     RevisionDTO           `json:"revision"`
	Id                           string                `json:"id"`
	Url                          string                `json:"url"`
	Position                     PositionDTO           `json:"position"`
	Permissions                  PermissionsDTO        `json:"permissions"`
	Bulletins                    []BulletinEntity      `json:"bulletins"`
	DisconnectedNodeAcknowledged bool                  `json:"disconnected_node_acknowledged"`
	Component                    ProcessGroupDTO       `json:"component"`
	Status                       ProcessGroupStatusDTO `json:"status"`
	VersionedFlowSnapshot        versionedFlowSnapshot `json:"versioned_flow_snapshot"`
	RunningCount                 int32                 `json:"running_count"`
	StoppedCount                 int32                 `json:"stopped_count"`
	InvalidCount                 int32                 `json:"invalid_count"`
	DisabledCount                int32                 `json:"disabled_count"`
	ActiveRemotePortCount        int32                 `json:"active_remote_port_count"`
	InactiveRemotePortCount      int32                 `json:"inactive_remote_port_count"`
	VersionedFlowState           string                `json:"versioned_flow_state"`
	UpToDateCount                int32                 `json:"up_to_date_count"`
	LocallyModifiedCount         int32                 `json:"locally_modified_count"`
	StaleCount                   int32                 `json:"stale_count"`
	LocallyModifiedAndStaleCount int32                 `json:"locally_modified_and_stale_count"`
	SyncFailureCount             int32                 `json:"sync_failure_count"`
	InputPortCount               int32                 `json:"input_port_count"`
	OutputPortCount              int32                 `json:"output_port_count"`
}
