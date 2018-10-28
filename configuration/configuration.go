// TODO:
// Load the configuration from the environement.
// Set the default value and then look for the Env variable.
// Remove the configuration.json file.
// The Env variables will be set in the dockerfile.yml.

package configuration

import (
	"encoding/json"
	"html/template"
	"os"
)

const (
	EndpointDefault       = "127.0.0.1:80"
	URLDatabaseAPIDefault = "127.0.0.1:9900"
	StaticFilePathDefault = "./frontend/static"
	ViewFilePathDefault   = "./frontend/views/*.html"
)

// ServiceConfig contains all the configuration of the service
type ServiceConfig struct {
	Endpoint       string `json:"endpoint"`
	URLDatabaseAPI string `json:"urlDatabaseAPI"`
	StaticFilePath string `json:"staticFilePath"`
	ViewFilePath   string `json:"viewFilePath"`
}

// ExtractConfiguration return the configuration found in the given file.
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	config := ServiceConfig{
		EndpointDefault,
		URLDatabaseAPIDefault,
		StaticFilePathDefault,
		ViewFilePathDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}

// InitTemplate initialize the ServiceConfig's template.
func (config *ServiceConfig) InitTemplate() *template.Template {
	return template.Must(template.ParseGlob(config.ViewFilePath))
}
