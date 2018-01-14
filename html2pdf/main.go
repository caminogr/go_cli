package main

import (
	"fmt"
	"log"

	"github.com/kaminora/go_cli/html2pdf/scrape"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	urls := scrape.GetURLs()
	var page *wkhtmltopdf.Page
	for _, url := range urls {
		page = wkhtmltopdf.NewPage(url)

		// This seconds make it possible to convert about 100 page
		page.RunScript.Set("setInterval(function(){if(document.readyState=='complete') window.status='done';},25000)")
		page.WindowStatus.Set("done")

		pdfg.AddPage(page)
	}

	fmt.Println("start to create pdf")
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./output.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("complete")
}
