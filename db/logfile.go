package db

import (
	"log"
	"os"
)

func GetFileSize(f *os.File) int64 {
	fileInfo, err := os.Stat(f.Name())

	if err != nil {
		log.Fatal("error occured when getting file info:")
	}

	return fileInfo.Size()
}
