package orgpa

import (
	"orgpa-frontend/configuration"
	"orgpa-frontend/template"
)

type ServerHandler struct {
	TmplEngine template.TemplateEngine
	Config     configuration.ServiceConfig
}

func newServerHandler(config configuration.ServiceConfig, templateEngine template.TemplateEngine) ServerHandler {
	return ServerHandler{
		TmplEngine: templateEngine,
		Config:     config,
	}
}
