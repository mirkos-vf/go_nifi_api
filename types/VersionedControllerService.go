package main

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
