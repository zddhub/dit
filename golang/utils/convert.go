package utils

import "fmt"

func BytesToSha1(bytes [20]byte) (sha1 string) {
	return fmt.Sprintf("%x", bytes)
}

func Sha1ToBytes(sha1 string) (bytes [20]byte) {
	for i := range bytes {
		if 2*i+1 >= len(sha1) {
			break
		}
		bytes[i] = htob(sha1[2*i])*16 + htob(sha1[2*i+1])
	}
	return bytes
}

func htob(hex byte) (b byte) {
	if hex <= '9' {
		b = hex - '0'
	} else {
		b = hex - 'a' + 10
	}
	return b
}
