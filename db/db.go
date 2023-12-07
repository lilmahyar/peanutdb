package db

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

type Record struct {
	key string
	value string
	mu sync.Mutex
}

func DbSet(r Record , file *os.File) int64 {
	r.mu.Lock()

	b, err := file.WriteString(r.key + "," + r.value + "\r\n")

	if err != nil {
		log.Fatal("error occurred in DbSet when performing write action, %v", err)
	}

	err = file.Sync()

	if err != nil {
		log.Fatal("error occurred in DbSet when performing write action, %v", err)
	}

	defer r.mu.Unlock()
	
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
