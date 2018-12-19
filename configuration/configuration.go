package configuration

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// ServiceConfig contains all the configuration of the service
type ServiceConfig struct {
	Endpoint       string `required:"true"`
	URLDatabaseAPI string `required:"true" envconfig:"URL_DATABASE_API"`
	StaticFilePath string `required:"false" split_words:"true" default:"./frontend/static/"`
	ViewFilePath   string `required:"false" split_words:"true" default:"./frontend/views/"`
}

// ExtractConfiguration will extract the configuration from
// the environement and return a ServiceConfig struct containing
// the whole service configuration.
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	var config ServiceConfig
	err := envconfig.Process("orgpa", &config)
	if err != nil {
		return ServiceConfig{}, err
	}

	if strings.HasSuffix(config.StaticFilePath, "/") == false {
		config.StaticFilePath += "/"
	}
	if strings.HasSuffix(config.ViewFilePath, "/") == false {
		config.ViewFilePath += "/"
	}
	return config, nil
}
