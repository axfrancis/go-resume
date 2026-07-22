package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type RGBColor struct {
	red   int
	green int
	blue  int
}

type PDFSettings struct {
	color  RGBColor
	accent RGBColor
	short  bool
}

type TechSkill struct {
	Name      string
	Values    string
	Icon      string
	BrandIcon bool
}

type MajorProject struct {
	Company      string
	Title        string
	Start        string
	End          string
	Description  string
	Achievements []string
}

type Person struct {
	FirstName       string
	LastName        string
	Location        string
	Phone           string
	Email           string
	LinkedIn        string
	GitHub          string
	Url             string
	QrUrl           string
	Summary         string
	TechnicalSkills []TechSkill
	JobHistory      []MajorProject
	MajorProjects   []MajorProject
}

func parseHex(r *http.Request, name string) (color RGBColor) {
	s := strings.TrimPrefix(r.URL.Query().Get(name), "#")
	color = RGBColor{0, 0, 0}

	if len(s) != 6 {
		return color
	}

	n, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		fmt.Printf("there is an err from strconv")
		return color
	}

	color.red = int((n >> 16) & 0xff)
	color.green = int((n >> 8) & 0xff)
	color.blue = int(n & 0xff)

	return
}

func resume(w http.ResponseWriter, r *http.Request) {
	person := Person{
		FirstName: "Anthony",
		LastName:  "Francis",
		Location:  "Melbourne, VIC",
		Phone:     "+61 416 769 371",
		Email:     "gday@axf.id.au",
		LinkedIn:  "axfrancis",
		GitHub:    "axfrancis",
		Url:       "axf.id.au/case-studies",
		QrUrl:     "github.com/axfrancis/GO-GETAJOB",
		Summary:   "Senior Software Developer with 8+ years of experience building high-availability critical retail infrastructure, software, and apps with a holistic philosophy, a keen eye for design, and a passion for customer experience. Proven track record modernising legacy systems, migrating software architectures, and delivering software that handles thousands of daily transactions, provides faster and more reliable results, and leads to higher customer satisfaction.",
		TechnicalSkills: []TechSkill{
			{
				Icon:   "\uf188",
				Name:   "Languages",
				Values: "TypeScript/JavaScript, Go, C#, C, Swift, Kotlin, SQL",
			},
			{
				Icon:   "\uf1c9",
				Name:   "Frameworks",
				Values: "Webpack, Node, Express, React, React Native, .NET Core, Docker, Bash, PowerShell",
			},
			{
				Icon:      "\uf375",
				Name:      "AWS Stack",
				Values:    "CloudFormation, DynamoDB, EC2, ECS, RDS, S3, Lambda, Secrets Manager, Route 53",
				BrandIcon: true,
			},
			{
				Icon:   "\uf1c0",
				Name:   "Databases",
				Values: "SQL Server, MySQL/MariaDB, Redis, AWS DynamoDB",
			},
		},
		JobHistory: []MajorProject{
			{
				Company: "Officeworks",
				Title:   "Senior Developer",
				Start:   "Aug 2021",
				End:     "Present",
				Achievements: []string{
					"Led the integration of AI LLMs into the development & code review process, allowing for easy identification and remediation of missed bugs",
					"Implemented smart caching and request routing for overloaded SAP infrastructure to reduce slow loading times at peak periods",
					"Identified and implemented cost-saving measures in cloud infrastructure, including demand-based scaling and consolidation of microservices, saving tens of thousands of dollars each year",
					"Ensured continuity of service during migration of PDT devices from iOS to Android, integrating new hardware with existing apps",
					"Performed Pentest/Audit remediation, added code linting infrastructure, and contibuted to unit testing, improving developer coding experience, and ensuring PCI compliance",
					"Coached junior developers, conducting peer reviews, training, and mentoring",
				},
			},
			{
				Company: "Officeworks",
				Title:   "Developer",
				Start:   "Jul 2018",
				End:     "Aug 2021",
				Achievements: []string{
					"Migrated 5+ internal-facing HTML/CSS/JS Cordova iOS apps to React Native while maintaining 100% uptime and functionality",
					"Designed and implemented supporting AWS architecture for the internal apps using a code-as-architecture CloudFormation approach",
					"Developed several internal-facing dashboards to automate previously manual processes & save stakeholders many work hours",
					"Spearheaded the design and implementation of several Click & Collect features to reduce customer waiting time and increase NPS",
					"Collaborated with external vendors to implement custom wrappers for hardware integrations",
				},
			},
			{
				Company: "Officeworks",
				Title:   "Junior Application Developer",
				Start:   "Sep 2017",
				End:     "Jul 2018",
				Achievements: []string{
					"Supported & developed internal HTML/CSS/JavaScript Cordova iOS apps",
					"Developed & maintained Nightwatch automated tests",
					"Investigated integrating electronic shelf-edge label technology into internal apps",
				},
			},
		},
		MajorProjects: []MajorProject{
			{
				Company:     "Officeworks",
				Title:       `"Pollie" POS`,
				Start:       "2023",
				End:         "Present",
				Description: `Developed in-house, Pollie is a point of sale system with a .NET Core/C# back-end, a JQuery/React JavaScript/TypeScript based front-end, and several supporting applications for interfacing with register hardware i.e. scanners, printers, and customer displays. Pollie handles thousands of daily transactions across hundreds of stores nationwide.`,
				Achievements: []string{
					"Refactored large parts of the JS front-end into TypeScript, which enabled the easy identification and removal of dead code and hidden bugs",
					"Implemented custom tax logic and reporting for GST edge cases when dealing with Government entities",
					"Performed Pentest/Audit remediation on large parts of the front and back ends to secure each endpoint, prevent prompt injection, remove hardcoded secrets from the codebase, enable the secure storage of secrets in AWS, etc.",
				},
			},
			{
				Company:     "Officeworks",
				Title:       "Click & Collect App",
				Start:       "2018",
				End:         "2023",
				Description: "Click & Collect is React Native/TypeScript based internal-facing app that enables team members in-store to pick orders, manage customer collections, and interface with 3rd party APIs e.g. DoorDash. Efficient and timely handling of Click & Collect orders is crucial to overall NPS scores, with a focus on customer time spent waiting in-store.",
				Achievements: []string{
					"Implemented a smart caching service to sit between the app and SAP, in order to reduce load on SAP endpoints and enable faster load times for team members",
					"Integrated with DoorDash APIs to facilitate communication between store team members and dashers from within the app",
					`Spearheaded customer experience initiatives such as "Valet" which enables team members at POS to notify those at back of house of when a customer has arrived and prompt them to bring the order to POS, and "I'm on my way!" which enables customers to let TMs at stores know when to bring larger orders to the front of store to await collection. Both of these initiatives resulted in a massive NPS spike for Click & Collect orders`,
				},
			},
			{
				Company:     "Self Published",
				Title:       "Free Cell Ad-Free",
				Start:       "2020",
				End:         "Present",
				Description: `Started during the COVID pandemic as a way to pass the time, Free Cell Ad-Free is a React Native/TypeScript based game, inspired by Microsoft's FreeCell: free to play card games with no interruptions.`,
				Achievements: []string{
					"Launched on iOS and Android, with hundreds of reviews averaging over four stars",
					`"No ads, no frills" philosophy treats users as people, not wallets`,
				},
			},
		},
	}

	settings := PDFSettings{
		color:  RGBColor{0, 0, 0},
		accent: RGBColor{100, 100, 100},
		short:  false,
	}

	if r.URL.Query().Has("color") {
		settings.color = parseHex(r, "color")
	}

	if r.URL.Query().Has("accent") {
		settings.accent = parseHex(r, "accent")
	}

	if r.URL.Query().Get("short") == "true" {
		settings.short = true
	}

	pdf, err := generatePDF(person, settings)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `inline; filename="FRANCIS Anthony Resume.pdf`)
	w.Write(pdf)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", resume)

	fmt.Println("Listening on :" + port)
	http.ListenAndServe(":"+port, nil)
}
