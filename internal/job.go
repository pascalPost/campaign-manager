package internal

import (
	"bytes"
	"fmt"
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

type Job struct {
	input            string
	parameter        map[string]string
	outputFolderTmpl string
	dependencies     []Job

	outputFolder string
	exec         Execution
}

func NewJob(input string, parameter map[string]string, outputFolderTmpl string) *Job {
	job := &Job{
		input:            input,
		parameter:        parameter,
		outputFolderTmpl: outputFolderTmpl,
	}

	// parse output folder template to set output folder
	tmplOutput, err := template.New("output").Parse(job.outputFolderTmpl)
	if err != nil {
		slog.Debug("Error in parsing output folder template", "job", fmt.Sprintf("%#v", job), "error", err)
		log.Fatal(err)
	}

	var outputFolder bytes.Buffer
	if err := tmplOutput.Execute(&outputFolder, job.parameter); err != nil {
		slog.Debug("Error in executing output template", "job", fmt.Sprintf("%#v", job), "error", err)
		log.Fatal(err)
	}
	job.outputFolder = outputFolder.String()

	return job
}

// Generate generates the output files from the input templates
func Generate(jobs ...Job) {
	for i, job := range jobs {
		slog.Debug("Processing", "job", fmt.Sprintf("%#v", job), "index", i)

		// create output folder and file
		err := os.MkdirAll(job.outputFolder, os.ModePerm)
		if err != nil {
			slog.Debug("Error in creating output folder", "job", fmt.Sprintf("%#v", job), "index", i)
			log.Fatal(err)
		}

		f, err := os.Create(filepath.Join(job.outputFolder, job.input))
		if err != nil {
			slog.Debug("Error in creating output file", "job", fmt.Sprintf("%#v", job), "index", i)
			log.Fatal(err)
		}

		// generate template
		tmpl, err := template.New(job.input).ParseFiles(job.input)
		if err != nil {
			slog.Debug("Error in parsing input file", "job", fmt.Sprintf("%#v", job), "index", i)
			log.Fatal(err)
		}

		err = tmpl.Execute(f, job.parameter)
		if err != nil {
			slog.Debug("Error in executing input template", "job", fmt.Sprintf("%#v", job), "index", i)
			log.Fatal(err)
		}
	}
}
