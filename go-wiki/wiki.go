package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Page records the title and the body
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// remeber the os.FileMode is 0600, not 600
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "test1", Body: []byte("test12345")}
	p1.save()
	p2, err := loadPage("test1")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(p2.Body))
	}

}
