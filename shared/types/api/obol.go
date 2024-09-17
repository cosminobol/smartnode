package api



type GetValidatorPublicKeysData struct {
	endpoint     string `json:"endpoint"`
}

type ManageCharonServiceData struct {
	action string `json:"action"`
	serviceName string `json:"serviceName"`
}

type GetCharonHealthData struct {
	endpoint     string `json:"endpoint"`
}

type DvExitSignData struct {
	endpoint     string `json:"endpoint"`
}

type DvExitBroadcastData struct {
	endpoint     string `json:"endpoint"`
	validatorPublicKeys     string `json:"validatorPublicKeys"`
	publishTimeout     string `json:"publishTimeout"`
}

type CharonDkgData struct {
	password     string `json:"password"`
}

type CreateENRData struct {
	password     string `json:"password"`
}