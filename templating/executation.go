package templating

import (
	"bytes"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/gkar68/GoAngularBrowserifyBoilerplate/settings"
)

const (
	dataFormat = "Mon Jan 2 2006"
	timeFormat = "15:04:05"
)

func getCurrentDate() string {
	currentDateTime := time.Now()

	var currentDate = currentDateTime.Format(dataFormat)
	return currentDate
	//presc.CurrentTime = currentDateTime.Format("15:04:05")
}

func getCurrentTime() string {
	currentDateTime := time.Now()

	var CurrentTime = currentDateTime.Format(timeFormat)
	return CurrentTime
}

// GetTemplate ...
func GetTemplate(templateFile string) (*template.Template, error) {
	log.Printf("generating template: %s\n", templateFile)
	templateFiles := settings.RunFolder + settings.TemplateFolder + "/" + templateFile
	log.Printf("%s", templateFiles)
	funcMap := template.FuncMap{
		"upper":       strings.ToUpper,
		"currentDate": getCurrentDate,
		"currentTime": getCurrentTime,
	}
	tmplt, err := template.New("anyTemplate").Funcs(funcMap).ParseFiles(templateFiles)
	return tmplt, err
}

//ExecuteTemplateToString ...
func ExecuteTemplateToString(templateFile string, data interface{}) (string, error) {
	
	tmplt, err := GetTemplate(templateFile)
	if err != nil {
		log.Println("an error in reading folder")
		log.Println(err.Error())
		return "", err
	}
	log.Println("execute template: ", templateFile)
	var doc bytes.Buffer
	ex := tmplt.ExecuteTemplate(&doc, templateFile, data)
	if ex != nil {
		return "", ex
	}
	result := doc.String()
	return result, nil
}
