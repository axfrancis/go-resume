package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type rgbColor struct {
	red   int
	green int
	blue  int
}

type pdfSettings struct {
	color  rgbColor
	accent rgbColor
	short  bool
}

func parseHex(r *http.Request, name string) (color rgbColor) {
	s := strings.TrimPrefix(r.URL.Query().Get(name), "#")
	color = rgbColor{0, 0, 0}

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
		Summary:   "Senior Software Developer with 8+ years of experience building high-availability critical retail infrastructure, software, and apps with a holistic philosophy, a keen eye for design, and a passion for customer experience. Proven track record modernising legacy systems, migrating software architectures, and delivering software to handle thousands of daily transactions, providing faster and more reliable results, leading to higher customer satisfaction.",
		TechnicalSkills: []TechSkill{
			{
				Icon:   "\uf188",
				Name:   "Languages",
				Values: "TypeScript/JavaScript, Go, C#, C, Swift, Kotlin, SQL, HTML, CSS, Docker, Bash, PowerShell",
			},
			{
				Icon:   "\uf1c9",
				Name:   "Frameworks",
				Values: "Webpack, Node, Express, React, React Native, .NET Core",
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
		JobHistory: []Job{
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
					"Designed and implemented supporting AWS architecture for those apps, automating several manual processes",
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
					"Investigated integrating electronic shelf-edge label technology into those apps",
				},
			},
		},
	}

	settings := pdfSettings{
		color:  rgbColor{0, 0, 0},
		accent: rgbColor{100, 100, 100},
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
