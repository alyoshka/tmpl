package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"time"
)

func main() {
	el := flag.Float64("el", 0., "Electricity")
	hot := flag.Float64("hot", 0., "Hot water")
	cold := flag.Float64("cold", 0., "Cold water")
	flag.Parse()

	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("failed to parse template: %s", err)
	}
	report := struct {
		Date string
		El   float64
		Hot  float64
		Cold float64
	}{
		Date: time.Now().Format("02.01.2006"),
		El:   *el,
		Hot:  *hot,
		Cold: *cold,
	}
	f, err := os.OpenFile("index.html", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	err = t.Execute(f, report)
	if err != nil {
		log.Fatalf("failed to execute template: %s", err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("faield to close file: %s", err)
	}
}
