package resume

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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

func Resume(w http.ResponseWriter, r *http.Request) {
	person := getPerson()
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
	w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s %s Resume.pdf"`, strings.ToUpper(person.LastName), person.FirstName))
	w.Write(pdf)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", Resume)

	fmt.Println("Listening on :" + port)
	http.ListenAndServe(":"+port, nil)
}
