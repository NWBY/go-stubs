package main

import (
	"log"
	"os"
	"text/template"
)

type StubDetails struct {
	Name, FileName, Destination string
	Values                      map[string]string
}

func (s *StubDetails) CreateStub() {
	contentsBuff, err := os.ReadFile(s.Name)
	if err != nil {
		log.Fatalf("Unable to read file: %s", s.Name)
	}

	f, err := os.OpenFile(s.Destination+s.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("Unable to open file: %s", s.FileName)
	}
	defer f.Close()

	template, err := template.New(s.FileName).Parse(string(contentsBuff))
	if err != nil {
		log.Fatalf("Unable to parse template: %s", s.Name)
	}
	template.Execute(f, s.Values)
}
