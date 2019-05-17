package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/ledongthuc/pdf"
)

func main() {
	path := "600570_2018_n.pdf"
	readPdf(path) // Read local pdf file
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(content)
	// return
}

func readPdf(path string) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		panic(err)
	}
	pages := r.NumPage()
	var buf bytes.Buffer
	fonts := make(map[string]*pdf.Font)
	for i := 1; i <= pages && i < 5; i++ {
		p := r.Page(i)
		for _, name := range p.Fonts() { // cache fonts so we don't continually parse charmap
			if _, ok := fonts[name]; !ok {
				f := p.Font(name)
				fonts[name] = &f
			}
		}
		text, err := p.GetPlainText(fonts)
		if err != nil {
			panic(err)
		}
		buf.WriteString(text)
		log.Println(text)
	}
}

func readPdf2(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage && pageIndex < 2; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		var lastTextStyle pdf.Text
		texts := p.Content().Text
		for _, text := range texts {
			if text == lastTextStyle {
				lastTextStyle.S = lastTextStyle.S + text.S
			} else {
				fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
				lastTextStyle = text
			}
		}
	}
	return "", nil
}
