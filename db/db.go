package db

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func DbSet(key, value string, file *os.File) int64 {
	b, err := file.WriteString(key + "," + value + "\r\n")

	if err != nil {
		log.Fatal("error occurred in DbSet when performing write action, %v", err)
	}

	err = file.Sync()

	if err != nil {
		log.Fatal("error occurred in DbSet when performing write action, %v", err)
	}

	return int64(b)
}

func DbGet(byteOffset int, file *os.File) string {

	reader := bufio.NewReader(file)

	_, err := reader.Discard(byteOffset)

	//bytes, err := reader.Peek(byteOffset)
	fmt.Println(byteOffset)

	value, _, _ := reader.ReadLine()

	if err != nil {
		log.Fatal("error occurred in DbSet when performing read action, %v", err)
	}

	return string(value)
}
