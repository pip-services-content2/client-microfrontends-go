package version1

type MicrofrontendV1 struct {
	Id            string            `json:"id"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	PathPrefix    string            `json:"path_prefix"`
	Icon          string            `json:"icon"`
	Type          string            `json:"type"`
	RemoteEntry   string            `json:"remote_entry"`
	ExposedModule string            `json:"exposed_module"`
	ElementName   string            `json:"element_name"`
	Params        map[string]string `json:"params"`
}

func NewEmptyMicrofrontendV1() *MicrofrontendV1 {
	return &MicrofrontendV1{}
}
