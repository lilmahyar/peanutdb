package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"peanutdb/db"
)

var (
	help = `
	invalid command, use help command for more information 
	enter any key to exist :
	`
)

var indexMap = make(map[string]int64)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println(help)

		var keyStroke string
		fmt.Scan(&keyStroke)

		log.Fatal(help)
	}

	f, err := os.OpenFile("peanutdb.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		log.Fatal("error occurred when opening file %s , %v", "peanutdb.log", err)
	}

	command := flag.Args()[0]
	args := flag.Args()[1:]

	executeCommand(command, args, f)
	executeCommand("db_get", []string{"mahyar"}, f)

}

func executeCommand(command string, args []string, dbFile *os.File) {
	switch command {
	case "help":
		log.Fatal(help)
	case "db_set":

		if len(args) != 2 {
			log.Fatal("db_set requires two arguments <key> <value>, see help for more information")
		}

		key, value := args[0], args[1]

		bytesSize := db.DbSet(key, value, dbFile)

		currentFileSize := getFileSize(dbFile)

		offset := currentFileSize - bytesSize

		insertKeyIndex(key, offset)

	case "db_get":
		if len(args) != 1 {
			log.Fatal("db_get requires one argument <key>, see help for more information")
		}

		key := args[0]
		offset := indexMap[key]

		value := db.DbGet(int(offset), dbFile)

		fmt.Println(value)
	default:
		log.Fatal(help)
	}
}

func getFileSize(f *os.File) int64 {
	fileInfo, err := os.Stat(f.Name())

	if err != nil {
		log.Fatal("error occurred when getting file info for %s", f.Name())
	}

	return fileInfo.Size()
}

func insertKeyIndex(key string, offset int64) {
	indexMap[key] = offset
}
