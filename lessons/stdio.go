package learngo

import (
	"fmt"
	"log"
	"os"
)

func Input() {
	var (
		text  string
		text2 string
	)

	count, err := fmt.Scan(&text, &text2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count, text, text2)

	newFile, err := os.Create("new.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	c, err := fmt.Println(newFile, "hello world")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Write:", c)

	file, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("error on open new.txt")
	}
	defer file.Close()

	n, err := file.WriteString("hello world")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Write count", n)

	fileToRead, err := os.Open("new.txt")
	if err != nil {
		log.Fatal(err)
	}

	var (
		text3 string
		text4 string
	)
	count, err = fmt.Fscan(fileToRead, &text3, &text4)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error on scan")
	}

	fmt.Println("Readed:", count, text3, text4)
}
