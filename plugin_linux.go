package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	pluginsFiles := GetPluginsFiles()
	for _, file := range pluginsFiles {
		_, _ = plugin.Open(file)
		log.Infoln("已加载动态插件" + file)
	}

}

func GetPluginsFiles() []string {
	_, err := os.Stat("plugins")
	if os.IsNotExist(err) {
		return []string{}
	}
	var datas []string
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, "so") {
			datas = append(datas, path)
		}
		return nil
	})
	return datas
}
