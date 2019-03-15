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

type UsersEntity struct {
	Generated string `json:"generated"`
	Users []UserEntity `json:"users"`
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
