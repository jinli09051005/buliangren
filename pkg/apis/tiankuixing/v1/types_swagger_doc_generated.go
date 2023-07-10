package v1

var map_UpdateConfig = map[string]string{
	"":         "UpdateConfig",
	"metadata": "Standard object's metadata.",
}

func (UpdateConfig) SwaggerDoc() map[string]string {
	return map_UpdateConfig
}

var map_UpdateConfigList = map[string]string{
	"": "UpdateConfigList",
}

func (UpdateConfigList) SwaggerDoc() map[string]string {
	return map_UpdateConfigList
}

var map_UpdateConfigSpec = map[string]string{
	"":               "UpdateConfigSpec defines the desired state of UpdateConfig",
	"imageName":      "Image Name",
	"configMapName":  "ConfigMap Name",
	"deploymentName": "Deployment Name",
	"counts":         "Numbers",
}

func (UpdateConfigSpec) SwaggerDoc() map[string]string {
	return map_UpdateConfigSpec
}

var map_UpdateConfigStatus = map[string]string{
	"": "UpdateConfigStatus defines the observed state of UpdateConfig",
}

func (UpdateConfigStatus) SwaggerDoc() map[string]string {
	return map_UpdateConfigStatus
}
