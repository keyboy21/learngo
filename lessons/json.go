package learngo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int8     `json:"age"`
	Married bool     `json:"married"`
	Hobbies []string `json:"hobbies,omitempty"`
	Address Address  `json:"-"`
}

type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func JsonMarshaling() {
	person := Person{
		Name:    "John",
		Age:     36,
		Married: true,
		Hobbies: []string{"football", "basketball"},
		Address: Address{Country: "USA", City: "Atlanta"},
	}

	bytes, err := json.Marshal(person)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))

	var resultPerson Person

	if err := json.Unmarshal(bytes, &resultPerson); err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultPerson)
}

func UnknownJson() {
	request := `{"first":1,"second":[1,2]}`

	result := make(map[string]any)

	if err := json.Unmarshal([]byte(request), &result); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	for k, v := range result {
		fmt.Printf("Key: %v, Value: %v \n", k, v)
	}
}

func JsonEncoderDecoder() {
	file, err := os.Open("example.json")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var person Person

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&person); err != nil {
		log.Fatal(err)
	}

	fmt.Println(person)

	person.Name = "John2"

	newFile, err := os.Create("example2.json")

	defer func() {
		if err := newFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := json.NewEncoder(newFile).Encode(person); err != nil {
		log.Fatal(err)
	}

}
