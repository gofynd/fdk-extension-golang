package mocks

import "fdk-extension-golang/pkg/models"

//GetExtCallback returns extension callback mock
func GetExtCallback() models.ExtCallback {
	return models.ExtCallback{
		Install:   Install,
		Auth:      Auth,
		Uninstall: Uninstall,
	}
}

//Auth ...
func Auth(contextKeys map[string]interface{}) string {
	return "https://light-hound-71.loca.lt"
}

//Install ...
func Install(contextKeys map[string]interface{}) string {
	return ""
}

//Uninstall ...
func Uninstall(contextKeys map[string]interface{}) string {
	return ""
}
