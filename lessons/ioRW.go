package learngo

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type NumsReader struct {
	nums string
}

func (r NumsReader) Read(buf []byte) (n int, err error) {
	var counter int

	for i := range r.nums {
		if string(r.nums[i]) >= "0" && string(r.nums[i]) <= "9" {
			buf = append(buf, r.nums[i])
			counter++
		}
	}

	return counter, io.EOF
}

func SimpleReader() {
	numbers := NumsReader{nums: "1,2,4,5,5,2,6,3,6,7,s,gf,th,jt,0"}
	buf := make([]byte, len(numbers.nums))

	count, err := numbers.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Println(count, err)
	fmt.Println(string(buf))
}

type Rows struct {
	text string
}

func (r *Rows) Read(p []byte) (int, error) {
	if len(r.text) == 0 {
		return 0, io.EOF
	}

	idx := strings.IndexByte(r.text, '\n')

	var line string
	if idx >= 0 {
		line = r.text[:idx]
		r.text = r.text[idx+1:]
	} else {
		line = r.text
		r.text = ""
	}

	n := copy(p, line)

	if len(r.text) == 0 {
		return n, io.EOF
	}

	return n, nil

	// var i int

	// for i := range len(r.text) {
	// 	if string(r.text[i]) == "\n" {
	// 		r.text = r.text[i+1:]
	// 		break
	// 	}

	// 	p[i] = r.text[i]

	// 	if i == len(r.text)-1 {
	// 		r.text = ""
	// 		return i + 1, io.EOF
	// 	}
	// }

	// return i + 1, nil
}

func RowsReader() {
	rows := Rows{text: "first line\nsecond line\nthird line"}

	var (
		count int
		err   error
	)

	buf := make([]byte, 100)

	for err != io.EOF {
		count, err = rows.Read(buf)
		fmt.Println(string(buf), count)
	}

}

type NumsToWrite struct {
	storedNums []byte
}

func (n *NumsToWrite) Write(nums []byte) (c int, err error) {
	if len(nums) == 0 {
		return 0, io.EOF
	}

	count := 0

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] >= '0' && nums[i] <= '9' {
	// 		n.storedNums[count] = nums[i]
	// 		count++
	// 	}
	// }

	for _, val := range nums {
		if val >= '0' && val <= '9' {
			n.storedNums = append(n.storedNums, val)
			count++
		}
	}

	return count, io.EOF
}

func SimpleWriter() {
	nums := []byte{'1', '2', '9', '7', '=', '+', '.'}
	writer := NumsToWrite{storedNums: make([]byte, len(nums))}

	count, err := writer.Write(nums)

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Println(string(writer.storedNums), count)

}

func OsFile() {
	newFile, err := os.Create("new.txt")

	if err != nil {
		log.Fatal(err)
	}

	n, err := newFile.WriteString("hello world")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)

	err = newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("new.txt")

	defer func() {
		file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 12)
	n, err = file.Read(buf)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read count:%v, data: %v", n, string(buf))

}
