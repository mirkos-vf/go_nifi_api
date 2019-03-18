package types

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
