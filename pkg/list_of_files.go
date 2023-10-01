package pkg

import (
	"log"
	"os"
	"strings"
)

func GetAllArticlesFiles() []string {
	entries, err := os.ReadDir("resources/template/articles")
	if err != nil {
		log.Fatal(err)
	}
	files := make([]string, 0)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		strings.HasSuffix(name, ".html")
		files = append(files, name[:len(name)-5])
	}
	return files
}
