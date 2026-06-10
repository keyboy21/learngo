package learngo

import (
	"fmt"
	"io"
	"log"
)

type NumsReader struct {
	nums string
}

func (r NumsReader) Read(p []byte) (n int, err error) {
	var counter int

	for i := range r.nums {
		if string(r.nums[i]) >= "0" && string(r.nums[i]) <= "9" {
			p[counter] = r.nums[i]
			counter++
		}
	}

	return counter, io.EOF
}

func SimpleReader() {
	buf := make([]byte, 10)

	numbers := NumsReader{nums: "1,2,4,5,5,2,6,3,6,7,s,gf,th,jt"}

	count, err := numbers.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Println(count, err)

}
