package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	var (
		inputFile string
		tmplFile  string
	)

	flag.StringVar(&inputFile, "in", "", "input json file")
	flag.StringVar(&tmplFile, "f", "", "template output file")

	flag.Parse()

	// get input
	var input io.Reader = os.Stdin
	if inputFile != "" {
		f, err := os.Open(inputFile)
		if err != nil {
			log.Fatalf("cannot open input file %s: %s", inputFile, err)
		}
		defer f.Close()
		input = f
	}

	// get template
	var tmpl string = strings.Join(flag.Args(), " ")
	if tmplFile != "" {
		data, err := ioutil.ReadFile(tmplFile)
		if err != nil {
			log.Fatalf("cannot read template file %s: %s", tmplFile, err)
		}
		tmpl = string(data)
	}

	// decode JSON input
	var v interface{}
	dec := json.NewDecoder(input)
	dec.UseNumber()
	if err := dec.Decode(&v); err != nil {
		log.Fatalf("cannot decode JSON input: %s", err)
	}

	// parse template
	tpl, err := template.New("jstpl").Parse(tmpl)
	if err != nil {
		log.Fatalf("cannot parse template: %s", err)
	}

	// finally, execute template with json input as data
	if err := tpl.Execute(os.Stdout, v); err != nil {
		log.Fatalf("cannot execute template: %s", err)
	}
}
