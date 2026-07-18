package main

import (
	"bytes"

	"github.com/phpdave11/gofpdf"
	"github.com/skip2/go-qrcode"
)

func generatePDF() ([]byte, error) {
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
	// var storedX float64
	var storedY float64

	pdf.AddUTF8Font("Inter", "", "fonts/Inter/static/Inter_18pt-Regular.ttf")
	pdf.AddUTF8Font("Inter", "B", "fonts/Inter/static/Inter_18pt-Bold.ttf")
	pdf.AddUTF8Font("FASolid", "", "fonts/FontAwesome/fa-solid-900.ttf")
	pdf.AddUTF8Font("FABrands", "", "fonts/FontAwesome/fa-brands-400.ttf")

	qrPng, qrErr := qrcode.Encode("https://github.com/axfrancis/GO-GETAJOB", qrcode.Medium, int(qrSize))

	if qrErr != nil {
		panic(qrErr)
	}

	pdf.RegisterImageOptionsReader("qr", gofpdf.ImageOptions{ImageType: "PNG"}, bytes.NewReader(qrPng))

	pdf.AddPage()

	pdf.SetXY(x, y)
	pdf.SetFont("Inter", "B", 28)
	pdf.SetTextColor(100, 100, 100)
	pdf.MultiCell(nameWidth, 9, "Anthony\nFrancis", "", "L", false)

	x = pageWidth - margin - (linkWidth + iconWidth) - (contactWidth + iconWidth) - qrSize - 12 - 12
	y += 0.5
	pdf.SetXY(x, y)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf57e")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(contactWidth, cellHeight, "Melbourne, VIC", "", 0, "R", false, 0, "")

	storedY = y
	y += cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf879")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(contactWidth, cellHeight, "+61 475 305 593", "", 0, "R", false, 0, "")

	y += cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf0e0")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(contactWidth, cellHeight, "gday@axf.id.au", "", 0, "R", false, 0, "")

	y = storedY
	x = pageWidth - margin - (linkWidth + iconWidth) - (qrSize + 2) - 10
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FABrands", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf08c")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(linkWidth, cellHeight, "linkedin.com/in/axfrancis", "", 0, "R", false, 0, "https://linkedin.com/axfrancis")

	y += cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FABrands", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf09b")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(linkWidth, cellHeight, "github.com/axfrancis", "", 0, "R", false, 0, "https://github.com/axfrancis")

	y += cellHeight
	pdf.SetXY(x, y)
	pdf.SetTextColor(100, 100, 100)
	pdf.SetFont("FASolid", "", 10)
	pdf.Cell(iconWidth, cellHeight, "\uf5bc")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Inter", "", 11)
	pdf.CellFormat(linkWidth, cellHeight, "axf.id.au/case-studies", "", 0, "R", false, 0, "https://axf.id.au/case-studies")

	y = storedY - 1.25
	x = pageWidth - margin - qrSize
	pdf.ImageOptions("qr", x, y, qrSize, qrSize, false, gofpdf.ImageOptions{}, 0, "https://github.com/axfrancis/GO-GETAJOB")

	y += cellHeight * 3
	x = pageWidth - margin - qrSize - 6
	pdf.SetDrawColor(225, 225, 225)
	pdf.Line(x, storedY+2.5, x, y-2.5)
	x -= linkWidth + iconWidth + 12
	pdf.Line(x, storedY+2.5, x, y-2.5)

	y += 5
	pdf.Line(margin, y, pageWidth-margin, y)

	var buffer bytes.Buffer

	err := pdf.Output(&buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
