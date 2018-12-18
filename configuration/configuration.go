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
	"strings"

	"github.com/kelseyhightower/envconfig"
)

const (
	EndpointDefault       = "127.0.0.1:80"
	URLDatabaseAPIDefault = "127.0.0.1:9900"
	StaticFilePathDefault = "./frontend/static"
	ViewFilePathDefault   = "./frontend/views/*.html"
)

// ExtractConfiguration will extract the configuration from
// the environement and return a ServiceConfig struct containing
// the whole service configuration.
//
// If an environment variable is missing a non nil error will be
// returned.

// ServiceConfig contains all the configuration of the service
type ServiceConfig struct {
	Endpoint       string `required:"true"`
	URLDatabaseAPI string `required:"true" envconfig:"URL_DATABASE_API"`
	StaticFilePath string `required:"true" split_words:"true"`
	ViewFilePath   string `required:"true" split_words:"true"`
}

// ExtractConfiguration return the configuration found in the given file.
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	var config ServiceConfig
	err := envconfig.Process("orgpa", &config)
	if err != nil {
		return ServiceConfig{}, err
	}

	if strings.HasSuffix(config.StaticFilePath, "/") == false {
		config.StaticFilePath += "/"
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
