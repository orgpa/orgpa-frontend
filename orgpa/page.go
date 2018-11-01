package orgpa

import "orgpa-frontend/configuration"

// PageInfo is used to render information on HTML pages
type PageInfo struct {
	Config configuration.ServiceConfig
	Title  string
	Data   interface{}
}

// Return a new PageInfo object with the given title,
// data and configuration.
func newPageInfo(title string, data interface{}, config configuration.ServiceConfig) PageInfo {
	return PageInfo{
		Config: config,
		Title:  title,
		Data:   data,
	}
}
