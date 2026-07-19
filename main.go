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
		Phone:     "+61 475 305 593",
		Email:     "gday@axf.id.au",
		LinkedIn:  "axfrancis",
		GitHub:    "axfrancis",
		Url:       "axf.id.au/case-studies",
		QrUrl:     "github.com/axfrancis/GO-GETAJOB",
		Summary:   "Senior Software Developer with 8+ years of experience building high-availability critical retail infrastructure, software, and apps with a holistic philosophy and keen eye for customer experience. Strong background in Node, TypeScript/JavaScript, React/React Native, C#/.NET, and deployment through AWS CloudFormation. Proven track record modernising legacy systems, migrating software architectures, and delivering critical software used nationwide.",
		TechnicalSkills: []Category{
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
