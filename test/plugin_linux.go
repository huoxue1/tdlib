package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

func init() {
	pluginsFiles := GetPluginsFiles()
	for _, file := range pluginsFiles {

	}
	plugin.Open("")
}

func GetPluginsFiles() []string {
	_, err := os.Stat("plugins")
	if os.IsNotExist(err) {
		return []string{}
	}
	var datas []string
	filepath.Walk("./plugins", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, "so") {
			datas = append(datas, path)
		}
	})
	return datas
}
