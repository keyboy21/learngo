package learngo

import (
	"fmt"
	"log"
	"os"
	"time"
)

func WriteNumbersToFile() {
	startTime := time.Now()

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	var int int

	for i := range 100 {
		_, err := file.WriteString(fmt.Sprintf("%d\n", i+1))
		if err != nil {
			log.Fatal(err)
		}

		int++
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(startTime))
}

func WriteToIO() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	n, err := file.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Printed:", n)

}

func ReadFromIO() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fileToWrite, err := os.Create("test2.txt")

	defer func() {
		err := fileToWrite.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	n, err := fileToWrite.ReadFrom(file)
	// n, err := os.Stdout.ReadFrom(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Writed:%v to %v", n, fileToWrite.Name())

}
