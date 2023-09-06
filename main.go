package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]
	srcfd, err := os.OpenFile(src, os.O_RDONLY, os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}
	infilestats, _ := srcfd.Stat()
	infilestats.Size()
	dstfd, err := syscall.Open(dst, syscall.O_WRONLY|syscall.O_CREAT, 0777)
	if err != nil {
		log.Fatal(err)
	}
	offset := int64(0)
	writeComplete := int64(0)
	for writeComplete < infilestats.Size() {
		writtenBytes, err := syscall.Sendfile(dstfd, int(srcfd.Fd()), &offset, int(infilestats.Size()))
		if err != nil {
			log.Fatal(err)
		}
		writeComplete += int64(writtenBytes)
	}
}
