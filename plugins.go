// +build go1.8

package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"plugin"
	"text/template"
)

func init() {
	// get plugins path
	ex, err := os.Executable()
	if err != nil {
		log.Printf("warning: cannot get executable path: %s", err)
	}
	exPath := path.Dir(ex)

	// find all plugins
	os.Chdir(exPath)
	pluginPaths, _ := filepath.Glob("jstpl-plugin-*")

	for _, pluginPath := range pluginPaths {
		p, err := plugin.Open(pluginPath)
		if err != nil {
			log.Printf("warning: cannot load plugin %q: %s", pluginPath, err)
			continue
		}

		getFunctionsSymbol, err := p.Lookup("Functions")
		if err != nil {
			log.Printf("warning: failed to load Functions for plugin %q: %s", pluginPath, err)
			continue
		}

		getFunctions, ok := getFunctionsSymbol.(func() template.FuncMap)
		if !ok {
			log.Println("warning: wrong signature for Functions method")
			continue
		}

		for k, v := range getFunctions() {
			_, exists := templateFuncs[k]
			if exists {
				log.Printf("warning: function named %q declared twice", k)
			}
			templateFuncs[k] = v
		}
	}
}
