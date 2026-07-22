package main

import (
	"bytes"
	"embed"
	"fmt"
	"image/color"

	"github.com/axfrancis/go-qrcode"
	"github.com/phpdave11/gofpdf"
)

//go:embed fonts/*
var assets embed.FS

type PageStyle struct {
	width, margin float64
}

type Font struct {
	name, style, location string
}

type SectionStyle struct {
	iconWidth, cellHeight float64
}

func loadFonts(pdf *gofpdf.Fpdf) {
	fonts := []Font{
		{"Inter", "", "fonts/Inter/static/Inter_18pt-Regular.ttf"},
		{"Inter", "B", "fonts/Inter/static/Inter_18pt-Bold.ttf"},
		{"Caveat", "B", "fonts/Caveat/static/Caveat-Bold.ttf"},
		{"FASolid", "", "fonts/FontAwesome/fa-solid-900.ttf"},
		{"FABrands", "", "fonts/FontAwesome/fa-brands-400.ttf"},
	}

	for _, f := range fonts {
		fontBytes, fontErr := assets.ReadFile(f.location)

		if fontErr != nil {
			continue
		}

		pdf.AddUTF8FontFromBytes(f.name, f.style, fontBytes)
	}
}

func drawLine(pdf *gofpdf.Fpdf, padding float64, pageStyle PageStyle) {
	y := pdf.GetY() + padding
	pdf.Line(pageStyle.margin, y, pageStyle.width-pageStyle.margin, y)
	y += padding
	pdf.SetY(y)
}

func generateHeader(pdf *gofpdf.Fpdf, person Person, settings PDFSettings, pageStyle PageStyle) {
	var storedY float64
	nameWidth := 60.0
	contactWidth := 33.0
	linkWidth := 48.0
	qrSize := 20.0
	paddingWidth := 3.0
	sectionStyle := SectionStyle{5, 5.75}
	x := pageStyle.margin
	y := pageStyle.margin

	qrPng, qrErr := qrcode.EncodeWithColor("https://"+person.QrUrl, qrcode.Medium, int(qrSize*4), color.RGBA{uint8(settings.accent.red), uint8(settings.accent.green), uint8(settings.accent.blue), 255}, color.White)

	if qrErr != nil {
		panic(qrErr)
	}

	pdf.RegisterImageOptionsReader("qr", gofpdf.ImageOptions{ImageType: "PNG"}, bytes.NewReader(qrPng))

	pdf.SetXY(x, y)
	pdf.SetFont("Inter", "B", 28)
	pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
	pdf.MultiCell(nameWidth, 9, fmt.Sprintf("%s\n%s", person.FirstName, person.LastName), "", "L", false)

	x = pageStyle.width - pageStyle.margin - (linkWidth + sectionStyle.iconWidth) - (contactWidth + sectionStyle.iconWidth) - qrSize - (paddingWidth * 4)
	y += 0.5
	pdf.SetXY(x, y)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(sectionStyle.iconWidth, sectionStyle.cellHeight, "\uf57e")
	pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(contactWidth, sectionStyle.cellHeight, person.Location, "", 0, "R", false, 0, "")

	storedY = y
	y += sectionStyle.cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(sectionStyle.iconWidth, sectionStyle.cellHeight, "\uf879")
	pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(contactWidth, sectionStyle.cellHeight, person.Phone, "", 0, "R", false, 0, "")

	y += sectionStyle.cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(sectionStyle.iconWidth, sectionStyle.cellHeight, "\uf0e0")
	pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(contactWidth, sectionStyle.cellHeight, person.Email, "", 0, "R", false, 0, "")

	y = storedY
	x = pageStyle.width - pageStyle.margin - (linkWidth + sectionStyle.iconWidth) - qrSize - (paddingWidth * 2)
	pdf.SetXY(x, y)
	pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
	pdf.SetFont("FABrands", "", 10)
	pdf.Cell(sectionStyle.iconWidth, sectionStyle.cellHeight, "\uf08c")
	pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(linkWidth, sectionStyle.cellHeight, "linkedin.com/in/"+person.LinkedIn, "", 0, "R", false, 0, "https://linkedin.com/"+person.LinkedIn)

	y += sectionStyle.cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
	pdf.SetFont("FABrands", "", 10)
	pdf.Cell(sectionStyle.iconWidth, sectionStyle.cellHeight, "\uf09b")
	pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(linkWidth, sectionStyle.cellHeight, "github.com/"+person.GitHub, "", 0, "R", false, 0, "https://github.com/"+person.GitHub)

	y += sectionStyle.cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(sectionStyle.iconWidth, sectionStyle.cellHeight, "\uf5bc")
	pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(linkWidth, sectionStyle.cellHeight, person.Url, "", 0, "R", false, 0, "https://"+person.Url)

	y = storedY - 1.5
	x = pageStyle.width - pageStyle.margin - qrSize - 1
	pdf.ImageOptions("qr", x, y, qrSize, qrSize, false, gofpdf.ImageOptions{}, 0, "https://"+person.QrUrl)

	y = storedY - 5
	x = pageStyle.width - pageStyle.margin - qrSize - 1 - 31.5
	pdf.SetXY(x, y)
	pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
	pdf.SetFont("Caveat", "B", 10.5)
	pdf.Cell(45, sectionStyle.iconWidth, "See how this PDF was generated!")
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(sectionStyle.iconWidth, sectionStyle.iconWidth, "\uf0ab")
	pdf.SetFont("Inter", "", 11)
	pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)

	y = storedY - 1.5 + (sectionStyle.cellHeight * 3)
	x = pageStyle.width - pageStyle.margin - qrSize - paddingWidth
	pdf.SetDrawColor(225, 225, 225)
	pdf.Line(x, storedY+2.5, x, y-2.5)
	x -= linkWidth + sectionStyle.iconWidth + (paddingWidth * 2)
	pdf.Line(x, storedY+2.5, x, y-2.5)
	pdf.SetY(y)
}

func generateSummary(pdf *gofpdf.Fpdf, person Person) {
	sectionStyle := SectionStyle{
		cellHeight: 5,
	}
	pdf.SetFontSize(12)
	pdf.MultiCell(0, sectionStyle.cellHeight, person.Summary, "", "J", false)
}

func generateTechSkills(pdf *gofpdf.Fpdf, person Person, settings PDFSettings, pageStyle PageStyle) {
	sectionStyle := SectionStyle{
		iconWidth:  5,
		cellHeight: 5.5,
	}
	for _, c := range person.TechnicalSkills {
		var icon string
		if c.Icon == "" {
			icon = "\uf111"
		} else {
			icon = c.Icon
		}
		pdf.SetX(pageStyle.margin)
		pdf.SetFont("Inter", "B", 11)
		pdf.MultiCell(0, sectionStyle.cellHeight, c.Name, "", "L", false)
		if c.BrandIcon {
			pdf.SetFont("FABrands", "", 10)
		} else {
			pdf.SetFont("FASolid", "", 10)
		}
		pdf.SetX(pdf.GetX() + 3.5)
		pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
		pdf.CellFormat(sectionStyle.iconWidth, sectionStyle.iconWidth, icon, "", 0, "C", false, 0, "")
		pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)
		pdf.SetFont("Inter", "", 11)
		pdf.SetX(pdf.GetX() + 2)
		pdf.MultiCell(0, sectionStyle.cellHeight, c.Values, "", "L", false)
	}
}

func generateJobHistory(pdf *gofpdf.Fpdf, person Person, settings PDFSettings, pageStyle PageStyle) {
	sectionStyle := SectionStyle{
		iconWidth:  5,
		cellHeight: 5,
	}

	for _, j := range person.JobHistory {
		pdf.SetFontSize(12)
		pdf.SetFontStyle("B")
		pdf.Cell(30, sectionStyle.cellHeight, j.Company)
		pdf.SetFontSize(11)
		pdf.SetFontStyle("")
		pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
		pdf.Cell(30, sectionStyle.cellHeight, j.Start+" - "+j.End)
		pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)
		pdf.SetXY(pageStyle.margin, pdf.GetY()+sectionStyle.cellHeight)
		pdf.SetFontStyle("B")
		pdf.MultiCell(0, sectionStyle.cellHeight+1, j.Title, "", "L", false)
		pdf.SetFontStyle("")

		for _, a := range j.Achievements {
			pdf.SetXY(pageStyle.margin+2, pdf.GetY()+0.75)
			pdf.SetFont("FASolid", "", 4)
			pdf.CellFormat(sectionStyle.iconWidth, sectionStyle.cellHeight, "\uf111", "", 0, "C", false, 0, "")
			pdf.SetFont("Inter", "", 11)
			pdf.MultiCell(0, sectionStyle.cellHeight, a, "", "L", false)
		}

		pdf.SetXY(pageStyle.margin, pdf.GetY()+3.0)
	}
}

func generateMajorProjects(pdf *gofpdf.Fpdf, person Person, settings PDFSettings, pageStyle PageStyle) {
	sectionStyle := SectionStyle{
		cellHeight: 5,
		iconWidth:  5,
	}
	pdf.SetFontSize(14)
	pdf.SetFontStyle("B")
	pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
	pdf.MultiCell(0, sectionStyle.cellHeight+5, "Major Projects", "", "L", false)
	pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)

	for _, j := range person.MajorProjects {
		pdf.SetFontSize(12)
		pdf.SetFontStyle("B")
		pdf.Cell(50, sectionStyle.cellHeight, j.Title)
		pdf.SetFontStyle("")
		pdf.SetFontSize(11)
		pdf.SetTextColor(settings.accent.red, settings.accent.green, settings.accent.blue)
		pdf.Cell(35, sectionStyle.cellHeight, j.Company)
		pdf.Cell(30, sectionStyle.cellHeight, j.Start+" - "+j.End)
		pdf.SetTextColor(settings.color.red, settings.color.green, settings.color.blue)
		pdf.SetXY(pageStyle.margin, pdf.GetY()+sectionStyle.cellHeight+1)
		pdf.MultiCell(0, sectionStyle.cellHeight-0.5, j.Description, "", "J", false)

		for _, a := range j.Achievements {
			pdf.SetXY(pageStyle.margin+2, pdf.GetY()+0.75)
			pdf.SetFont("FASolid", "", 4)
			pdf.CellFormat(sectionStyle.iconWidth, sectionStyle.cellHeight, "\uf111", "", 0, "C", false, 0, "")
			pdf.SetFont("Inter", "", 11)
			pdf.MultiCell(0, sectionStyle.cellHeight-0.5, a, "", "L", false)
		}

		pdf.SetXY(pageStyle.margin, pdf.GetY()+3.0)
	}
}

func generatePDF(person Person, settings PDFSettings) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pageStyle := PageStyle{210, 12}
	paddingHeight := 2.5

	loadFonts(pdf)

	pdf.SetMargins(pageStyle.margin, pageStyle.margin, pageStyle.margin)
	pdf.SetAutoPageBreak(true, pageStyle.margin)

	pdf.AddPage()

	generateHeader(pdf, person, settings, pageStyle)
	drawLine(pdf, paddingHeight, pageStyle)

	generateSummary(pdf, person)
	drawLine(pdf, paddingHeight, pageStyle)

	generateTechSkills(pdf, person, settings, pageStyle)
	drawLine(pdf, paddingHeight, pageStyle)

	generateJobHistory(pdf, person, settings, pageStyle)
	drawLine(pdf, paddingHeight, pageStyle)

	if settings.short != true {
		generateMajorProjects(pdf, person, settings, pageStyle)
		drawLine(pdf, paddingHeight, pageStyle)
	}

	var buffer bytes.Buffer

	err := pdf.Output(&buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
