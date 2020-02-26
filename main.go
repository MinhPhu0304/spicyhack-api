package main

import (
	"net/http"

	"spicyhack/utils"

	"github.com/gin-gonic/gin"
)

var frontends = []string{"React", "Vanila"}
var backends = []string{"Ruby on Rails", ".NET Core", "ExpressJS", "Django", "Flask", "PHP", "Laravel", "Spring Boot"}
var databases = []string{"PostgreSQL", "MySQL", "MongoDB", "DynamoDB", "SQLite", "Maria DB", "CouchDB", "ArangoDB"}

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/10x", extraHardHandler)
	router.Run(":3000")
}

type Architecture struct {
	Frontend string `json:"frontend"`
	Backend  string `json:"backend"`
	Database string `json:"database"`
}

func rootHandler(context *gin.Context) {
	result := Architecture{
		Frontend: frontends[utils.GetRandomIndex(len(frontends))],
		Backend:  backends[utils.GetRandomIndex(len(backends))],
		Database: databases[utils.GetRandomIndex(len(databases))],
	}

	context.JSON(http.StatusOK, gin.H{
		"message":      "OK",
		"architecture": result,
	})
}

var frontendsHard = []string{"Apache Cordova",
	"React",
	"AngularJS",
	"Angular",
	"VueJS",
	"Wix",
	"Squarespace",
	"Shopify",
	"Vanilla JS",
	"Ember",
	"jQuery",
	"Backbone JS",
	"Svelte",
	"Flutter",
	"HTML Canvas",
	"three.js",
	"deck.gl",
	"Wordpress",
	"Drupal",
	"Magento",
	"Silverstripe",
	"Joomla",
	"Google ARCore",
	"Apple ARKit",
	"Android App",
	"Chatbot",
	"iOS App",
	"Any VR Platform",
	"Unreal Engine",
	"GameMaker",
	"Google Home",
	"Alexa Skill",
	"Messenger Bot",
	"Unity",
	"Visual Basic",
}
var backendsHard = []string{
	"Ruby on Rails",
	"Apache Tomcat",
	"kubernetes",
	"GraphQL",
	"Relay",
	"Docker Swarm",
	"openshift",
	"Apache Mesos",
	"Apache Aurora",
	".NET Framework",
	".NET Core",
	"ExpressJS",
	"Django",
	"Flask",
	"PHP",
	"Rust",
	"Laravel",
	"Spring Boot",
	"Cobol",
	"Fortran",
	"Forth",
	"ASM.js",
	"Prolog",
	"brainfuck",
	"Erlang",
	"Chicken esolang",
	"Whitespace esolang",
	"TrumpScript",
	"LOLCODE",
	"GNU Assembler (GAS)",
	"Microsoft Macro Assembler (MASM)",
	"Perl",
	"F#",
	"Dlang",
	"Magento",
}

var databasesHard = []string{"PostgreSQL", "MySQL", "MongoDB", "DynamoDB", "SQLite", "Maria DB", "CouchDB", "ArangoDB"}

var projectManagement = []string{
	"Trello",
	"Asana",
	"Jira",
	"Notion",
	"monday.com",
	"Wrike",
	"Basecamp",
	"Whiteboard",
	"Miro",
	"Airtable",
	"Evernote",
	"Microsoft Word",
}

var textEditor = []string{
	"Notepad/TextEdit",
	"Notepad++",
	"Microsoft Word",
	"LibreOffice Writer",
	"Sublime Text",
	"Visual Studio Code",
	"Brackets",
	"Atom",
	"Vim",
	"Emacs",
	"Nano",
	"WebStorm",
	"Bluefish",
	"Spacemacs",
	"UltraEdit",
}

type SmartAssArchitecture struct {
	Frontend          string `json:"frontend"`
	Backend           string `json:"backend"`
	Database          string `json:"database"`
	Editor            string `json:"textEditor"`
	ProjectManagement string `json:"projectManagement"`
}

func extraHardHandler(context *gin.Context) {
	result := SmartAssArchitecture{
		Frontend:          frontendsHard[utils.GetRandomIndex(len(frontendsHard))],
		Backend:           backendsHard[utils.GetRandomIndex(len(backendsHard))],
		Database:          databasesHard[utils.GetRandomIndex(len(databasesHard))],
		Editor:            textEditor[utils.GetRandomIndex(len(textEditor))],
		ProjectManagement: projectManagement[utils.GetRandomIndex(len(projectManagement))],
	}

	context.JSON(http.StatusOK, gin.H{
		"message":      "OK",
		"architecture": result,
	})
}
