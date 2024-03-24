package cm

import (
	"bytes"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"text/template"
)

type Status struct {
}

type Execution struct {
	command string
	status  Status
}

type Unit struct {
	input            string
	parameter        map[string]string
	outputFolderTmpl string
	dependencies     []Unit

	outputFolder string
	exec         Execution
}

func NewUnit(input string, parameter map[string]string, outputFolderTmpl string) *Unit {
	u := &Unit{
		input:            input,
		parameter:        parameter,
		outputFolderTmpl: outputFolderTmpl,
	}

	// parse output folder template to set output folder
	tmplOutput, err := template.New("output").Parse(u.outputFolderTmpl)
	if err != nil {
		slog.Debug("Error in parsing output folder template for unit %v: %s", u, err)
		log.Fatal(err)
	}

	var outputFolder bytes.Buffer
	if err := tmplOutput.Execute(&outputFolder, u.parameter); err != nil {
		slog.Debug("Error in executing output template for unit %v: %s", u, err)
		log.Fatal(err)
	}
	u.outputFolder = outputFolder.String()

	return u
}

// Generate generates the output files from the input templates
func Generate(units ...Unit) {
	for i, u := range units {
		slog.Debug("Processing unit %i: %v", i, u)

		// create output folder and file
		err := os.MkdirAll(u.outputFolder, os.ModePerm)
		if err != nil {
			slog.Debug("Error in creating output folder for unit %i: %s", i, err)
			log.Fatal(err)
		}

		f, err := os.Create(filepath.Join(u.outputFolder, u.input))
		if err != nil {
			slog.Debug("Error in creating output file for unit %i: %s", i, err)
			log.Fatal(err)
		}

		// generate template
		tmpl, err := template.New(u.input).ParseFiles(u.input)
		if err != nil {
			slog.Debug("Error in parsing input file for unit %i: %s", i, err)
			log.Fatal(err)
		}

		err = tmpl.Execute(f, u.parameter)
		if err != nil {
			slog.Debug("Error in executing input template for unit %i: %s", i, err)
			log.Fatal(err)
		}
	}
}
