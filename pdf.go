package main

import (
	"bytes"
	"fmt"
	"image/color"

	"github.com/axfrancis/go-qrcode"
	"github.com/phpdave11/gofpdf"
)

type TechSkill struct {
	Name      string
	Values    string
	Icon      string
	BrandIcon bool
}

type Job struct {
	Company      string
	Title        string
	Start        string
	End          string
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
	JobHistory      []Job
}

func generatePDF(person Person) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	margin := 12.0
	x := margin
	y := margin
	pageWidth := 210.0
	iconWidth := 5.0
	nameWidth := 60.0
	contactWidth := 33.0
	linkWidth := 48.0
	cellHeight := 5.75
	qrSize := 20.0
	paddingWidth := 3.0
	paddingHeight := 2.5
	// var storedX float64
	var storedY float64

	pdf.SetMargins(margin, margin, margin)

	pdf.AddUTF8Font("Inter", "", "fonts/Inter/static/Inter_18pt-Regular.ttf")
	pdf.AddUTF8Font("Inter", "B", "fonts/Inter/static/Inter_18pt-Bold.ttf")
	pdf.AddUTF8Font("FASolid", "", "fonts/FontAwesome/fa-solid-900.ttf")
	pdf.AddUTF8Font("FABrands", "", "fonts/FontAwesome/fa-brands-400.ttf")

	// header
	qrPng, qrErr := qrcode.EncodeWithColor("https://"+person.QrUrl, qrcode.Medium, int(qrSize*4), color.RGBA{100, 100, 100, 255}, color.White)

	if qrErr != nil {
		panic(qrErr)
	}

	pdf.RegisterImageOptionsReader("qr", gofpdf.ImageOptions{ImageType: "PNG"}, bytes.NewReader(qrPng))

	pdf.AddPage()

	pdf.SetXY(x, y)
	pdf.SetFont("Inter", "B", 28)
	pdf.SetTextColor(100, 100, 100)
	pdf.MultiCell(nameWidth, 9, fmt.Sprintf("%s\n%s", person.FirstName, person.LastName), "", "L", false)

	x = pageWidth - margin - (linkWidth + iconWidth) - (contactWidth + iconWidth) - qrSize - (paddingWidth * 4)
	y += 0.5
	pdf.SetXY(x, y)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf57e")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(contactWidth, cellHeight, person.Location, "", 0, "R", false, 0, "")

	storedY = y
	y += cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf879")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(contactWidth, cellHeight, person.Phone, "", 0, "R", false, 0, "")

	y += cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf0e0")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(contactWidth, cellHeight, person.Email, "", 0, "R", false, 0, "")

	y = storedY
	x = pageWidth - margin - (linkWidth + iconWidth) - qrSize - (paddingWidth * 2)
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FABrands", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf08c")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(linkWidth, cellHeight, "linkedin.com/in/"+person.LinkedIn, "", 0, "R", false, 0, "https://linkedin.com/"+person.LinkedIn)

	y += cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FABrands", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf09b")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(linkWidth, cellHeight, "github.com/"+person.GitHub, "", 0, "R", false, 0, "https://github.com/"+person.GitHub)

	y += cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf5bc")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(linkWidth, cellHeight, person.Url, "", 0, "R", false, 0, "https://"+person.Url)

	y = storedY - 1.5
	x = pageWidth - margin - qrSize - 1
	pdf.ImageOptions("qr", x, y, qrSize, qrSize, false, gofpdf.ImageOptions{}, 0, "https://"+person.QrUrl)

	y += cellHeight * 3
	x = pageWidth - margin - qrSize - paddingWidth
	pdf.SetDrawColor(225, 225, 225)
	pdf.Line(x, storedY+2.5, x, y-2.5)
	x -= linkWidth + iconWidth + (paddingWidth * 2)
	pdf.Line(x, storedY+2.5, x, y-2.5)

	y += paddingHeight + 1.5
	pdf.Line(margin, y, pageWidth-margin, y)

	// body
	x = margin
	y += paddingHeight
	cellHeight = 5
	pdf.SetXY(x, y)
	pdf.SetFontSize(12)
	pdf.MultiCell(0, cellHeight, person.Summary, "", "J", false)

	y = pdf.GetY() + paddingHeight
	pdf.Line(margin, y, pageWidth-margin, y)

	y += paddingHeight
	pdf.SetXY(x, y)
	cellHeight++

	for _, c := range person.TechnicalSkills {
		var icon string
		if c.Icon == "" {
			icon = "\uf111"
		} else {
			icon = c.Icon
		}
		pdf.SetX(margin)
		pdf.SetFont("Inter", "B", 11)
		pdf.MultiCell(0, cellHeight, c.Name, "", "L", false)
		if c.BrandIcon {
			pdf.SetFont("FABrands", "", 10)
		} else {
			pdf.SetFont("FASolid", "", 10)
		}
		pdf.SetX(pdf.GetX() + 3.5)
		pdf.SetTextColor(100, 100, 100)
		pdf.CellFormat(iconWidth, iconWidth, icon, "", 0, "C", false, 0, "")
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Inter", "", 11)
		pdf.SetX(pdf.GetX() + 2)
		pdf.MultiCell(0, cellHeight, c.Values, "", "L", false)
	}

	y = pdf.GetY() + paddingHeight
	pdf.Line(margin, y, pageWidth-margin, y)
	y += paddingHeight
	pdf.SetY(y)

	for _, j := range person.JobHistory {
		pdf.SetFontSize(12)
		pdf.SetFontStyle("B")
		pdf.Cell(30, cellHeight, j.Company)
		pdf.SetFontSize(11)
		pdf.SetFontStyle("")
		pdf.Cell(30, cellHeight, j.Start+" - "+j.End)
		pdf.SetXY(margin, pdf.GetY()+cellHeight)
		pdf.MultiCell(0, cellHeight, j.Title, "", "L", false)

		for _, a := range j.Achievements {
			pdf.SetX(margin + 2)
			pdf.SetFont("FASolid", "", 4)
			pdf.CellFormat(iconWidth, cellHeight, "\uf111", "", 0, "C", false, 0, "")
			pdf.SetFont("Inter", "", 11)
			pdf.MultiCell(0, cellHeight, a, "", "L", false)
		}

		pdf.SetXY(margin, pdf.GetY()+3.0)
	}

	var buffer bytes.Buffer

	err := pdf.Output(&buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
