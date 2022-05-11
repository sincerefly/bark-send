package sqlite3

import (
	"log"
	"os"
	"path/filepath"
)

func InitStore() {

	path := "data"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	file, err := os.Create(filepath.Join(path, "database.db"))
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
