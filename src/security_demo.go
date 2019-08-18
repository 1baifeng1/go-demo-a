package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	testString := "Oh, you're pandaman!"

	md5Inst := md5.New()
	md5Inst.Write([]byte(testString))
	result := md5Inst.Sum([]byte(""))
	fmt.Println("md5Inst = ", result)
	fmt.Printf("md5Inst = %x\n\n", result)

	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(testString))
	result = sha1Inst.Sum([]byte(""))
	fmt.Println("sha1Inst = ", result)
	fmt.Printf("sha1Inst = %x\n\n", result)
}
