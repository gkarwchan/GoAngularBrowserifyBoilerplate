package settings

import (
	"log"

	"github.com/namsral/flag"
)

var (
	config string
	/* ********** Web App settings *************** */
	StaticAssets string
	WebHost      string
	WebPort      int
	UseSSL       bool
	CertFile     string
	KeyFile      string

	TemplateFolder string
	RunFolder      string
	EndPointURL    string
	//Smtp Settings
	//SmtpHost
	SmtpHost string
	//SmtpPort
	SmtpPort int
	//SmtpUsername
	SmtpUsername string
	//SmtpPassword
	SmtpPassword string

	DatabaseStore string

	JwtFile        string
	GplusKey       string
	GplusSecret    string
	FacebookKey    string
	FacebookSecret string
	TwitterKey     string
	TwitterSecret  string
)

func init() {
	// flag.StringVar(&config, "config", "", "config file to use")
	flag.IntVar(&WebPort, "web-port", 8081, "Port to bind to")
	flag.StringVar(&EndPointURL, "endpoint", "http://localhost:8081/", "end point")
	flag.BoolVar(&UseSSL, "use-ssl", false, "Use SSL")
	flag.StringVar(&CertFile, "cert-file", "certs/mycert1.cer", "SSL certificate")
	flag.StringVar(&KeyFile, "key-file", "certs/mycert1.key", "SSL key")

	flag.StringVar(&StaticAssets, "static-assets", "output", "Static asset folder")
	flag.StringVar(&RunFolder, "run-folder", "", "Run Folder")
	flag.StringVar(&TemplateFolder, "template-folder", "serverFiles/htmlTemplates", "Template Folder")

	flag.StringVar(&SmtpHost, "smtp-host", "smtp.gmail.com", "SMTP host name")
	flag.IntVar(&SmtpPort, "smtp-port", 587, "SMTP port")
	flag.StringVar(&SmtpUsername, "smtp-username", "gkarwchan@getfishtank.ca", "SMTP username")
	flag.StringVar(&SmtpPassword, "smtp-password", "", "SMTP password")

	flag.StringVar(&DatabaseStore, "database", "mongodb://127.0.0.1:27017/model", "database")

	flag.StringVar(&GplusKey, "gplus-key", "348931298414-nv1k83aifd844hut6d1lsebc14dpt4im.apps.googleusercontent.com", "GPlus Key")
	flag.StringVar(&GplusSecret, "gplus-secret", "DdrXObGcvf5wr8yLByuNRs40", "GPlus secret")
	flag.StringVar(&FacebookKey, "facebookKey", "", "facebook key")
	flag.StringVar(&FacebookSecret, "facebookSecret", "", "facebook secret")
	flag.StringVar(&TwitterKey, "twitterKey", "", "twitter key")
	flag.StringVar(&TwitterSecret, "twitterSecret", "", "twitter secret")
	flag.Parse()

	log.Printf("Config file						%s", config)
	log.Printf("Web folder						%s", StaticAssets)
	log.Printf("Web host						%s", WebHost)
	log.Printf("Web port						%s", WebPort)
	log.Printf("End point						%s", EndPointURL)
	log.Printf("Use SSL							%s", UseSSL)
	log.Printf("certification file				%s", CertFile)
	log.Printf("key file						%s", KeyFile)
	log.Printf("run folder						%s", RunFolder)
	log.Printf("template folder					%s", TemplateFolder)
	log.Printf("end point						%s", EndPointURL)
	log.Printf("SMTP host						%s", SmtpHost)
	log.Printf("SMTP port 						%s", SmtpPort)
	log.Printf("SMTP username					%s", SmtpUsername)
	log.Printf("SMTP password					%s", SmtpPassword)
	log.Printf("Database						%s", DatabaseStore)
	log.Printf("JWT File						%s", JwtFile)
	log.Printf("Google plus						%s", GplusKey)
	log.Printf("Google Secret					%s", GplusSecret)
	log.Printf("facebook						%s", FacebookKey)
	log.Printf("facebook secret					%s", FacebookSecret)
	log.Printf("twitter							%s", TwitterKey)
	log.Printf("twitter secret					%s", TwitterSecret)

}
