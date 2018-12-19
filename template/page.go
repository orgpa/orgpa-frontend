package template

import (
	"orgpa-frontend/configuration"
)

// PageData represent the data transmited to a HTML page
type PageData struct {
	PageName string
	Data     interface{}
	Config   configuration.ServiceConfig
}
