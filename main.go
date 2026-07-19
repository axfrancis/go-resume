package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go")
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
				Title:   "Application Developer",
				Start:   "Jul 2018",
				End:     "Oct 2021",
				Achievements: []string{
					"Defeated all enemies in single combat",
				},
			},
			{
				Company: "Officeworks",
				Title:   "Junior Developer",
				Start:   "Sep 2017",
				End:     "Jul 2018",
				Achievements: []string{
					"Supported & developed HTML/CSS/JavaScript Cordova iOS apps",
					"Developed & maintained Nightwatch automated tests",
				},
			},
		},
	}

	pdf, err := generatePDF(person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdf)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/resume", resume)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
