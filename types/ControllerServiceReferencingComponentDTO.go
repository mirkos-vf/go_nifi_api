package main

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
