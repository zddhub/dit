package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("index")
	if err != nil {
		fmt.Println("open index file error")
	}
	defer file.Close()

	sha1 := make([]byte, 20)
	flags := make([]byte, 2)
	b := make([]byte, 4)

	// header
	file.Read(b)
	fmt.Printf("signature: %s\n", b)
	file.Read(b)
	fmt.Printf("version: %d\n", b[3])

	file.Read(b)
	n := binary.BigEndian.Uint32(b)
	fmt.Println("entries no.:", n)

	//entries, only parse version 2
	fmt.Println("============================")
	for i := uint32(0); i < n; i++ {
		fmt.Println("----------------------")
		fmt.Println("index:", i)

		file.Read(b)
		fmt.Println("ctime seconds:", binary.BigEndian.Uint32(b))
		file.Read(b)
		fmt.Println("ctime nanosecond:", binary.BigEndian.Uint32(b))
		file.Read(b)
		fmt.Println("mtime seconds:", binary.BigEndian.Uint32(b))
		file.Read(b)
		fmt.Println("mtime nanosecond:", binary.BigEndian.Uint32(b))

		file.Read(b)
		fmt.Println("dev:", binary.BigEndian.Uint32(b))
		file.Read(b)
		fmt.Println("ino:", binary.BigEndian.Uint32(b))

		file.Read(b)
		fmt.Printf("mode: %o\n", binary.BigEndian.Uint32(b))

		file.Read(b)
		fmt.Printf("uid: %d\n", binary.BigEndian.Uint32(b))
		file.Read(b)
		fmt.Printf("gid: %d\n", binary.BigEndian.Uint32(b))
		file.Read(b)
		fmt.Println("file size:", binary.BigEndian.Uint32(b))

		file.Read(sha1)
		fmt.Printf("sha1: %x\n", sha1)

		file.Read(flags)
		nameLength := binary.BigEndian.Uint16(flags)
		nameLength = nameLength & 0x0fff
		fmt.Printf("flags: %x, name length: %d\n", flags, nameLength)

		name := make([]byte, nameLength)
		file.Read(name)
		fmt.Printf("path name: %s\n", name)

		// 1-8 nul bytes as necessary to pad the entry to a multiple of eight bytes
		// while keeping the name NUL-terminated.
		skip := make([]byte, 8-(62+nameLength)%8)
		file.Read(skip)
		fmt.Println("----------------------")
	}
	fmt.Println("============================")

	// Extensions
	file.Read(b)
	fmt.Printf("extension signature: %s\n", b)

	file.Read(b)
	n = binary.BigEndian.Uint32(b)
	fmt.Printf("size of the extension: %d\n", n)

	reader := bufio.NewReader(file)
	size := uint32(0)
	for size < n {
		str, err := reader.ReadString(0)
		if err != nil {
			fmt.Println("str err", err)
			break
		}
		size += uint32(len(str) + 1)
		fmt.Println("str:", str)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("err", err)
			break
		}
		size += uint32(len(line))
		fmt.Printf("line: %s\n", line)
		// if entry_count == -1
		if int(line[0]) == 45 && int(line[1]) == 49 {
			continue
		}
		num, err := reader.Read(sha1)
		fmt.Printf("sha1: %d %x\n", num, sha1)
		size += uint32(20)
	}
	fmt.Println("============================")
	num, err := reader.Read(sha1)
	fmt.Printf("checksum: %d %x\n", num, sha1)
}
