package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	templatePath := filepath.Join("./templates/inventory.tmpl")

	sweaters := NewInventory("wool", 17)
	tmpl, _ := template.ParseFiles(templatePath)

	if err := tmpl.Execute(os.Stdout, sweaters); err != nil {
		return err
	}
	return nil
}

type Inventory struct {
	Material string
	Count    int
}

func NewInventory(material string, count int) Inventory {
	return Inventory{
		Material: material,
		Count:    count,
	}
}
