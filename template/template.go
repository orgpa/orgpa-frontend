package template

import (
	"html/template"
	"net/http"
	"orgpa-frontend/configuration"
)

// TemplateEngine is the main engine of the template system
type TemplateEngine struct {
	ViewsPath  string
	HeaderPath string
	FooterPath string
	HeadPath   string
	Config     configuration.ServiceConfig
}

// NewTemplateEngine return a new template engine with the given
// views path and configuration.
func NewTemplateEngine(config configuration.ServiceConfig) TemplateEngine {
	return TemplateEngine{
		ViewsPath:  config.ViewFilePath,
		HeaderPath: config.ViewFilePath + "utils/Header.html",
		FooterPath: config.ViewFilePath + "utils/Footer.html",
		HeadPath:   config.ViewFilePath + "utils/Head.html",
		Config:     config,
	}
}

// GenerateTemplate return a whole template corresponding on given viewName.
func (tmplEng TemplateEngine) GenerateTemplate(viewName string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(tmplEng.ViewsPath+viewName+".html", tmplEng.HeaderPath, tmplEng.FooterPath, tmplEng.HeadPath)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// GenerateAndExecuteTemplate will make a call to GenerateTemplate and then
// it will create a PageData struct and pass it
func (tmplEng TemplateEngine) GenerateAndExecuteTemplate(w http.ResponseWriter, viewName, pageName string, data interface{}) error {
	tmpl, err := tmplEng.GenerateTemplate(viewName)
	if err != nil {
		return err
	}
	pageData := PageData{pageName, data, tmplEng.Config}
	return tmpl.Execute(w, pageData)
}
