package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Salary struct {
	Basic float64
}

type Employee struct {
	Firstname     string
	Lastname      string
	Age           int
	MonthlySalary []Salary
}

func ReadFromStdin() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("What is your name ? >")
		input, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, "error reading input: ", err)
			continue
		}

		fmt.Println(strings.TrimSpace(input))
	}

}

func ReadFromFile() {
	// file, err := os.OpenFile("Makefile", os.O_RDONLY,0644)
	file, err := os.Open("Makefile")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("> %s\n", line)
	}

	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v", scanner.Err())
		return
	}
}

func WriteStructToFile() {

	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := Employee{
		Firstname: "Sammy",
		Lastname:  "Mwawaka",
		Age:       24,
		MonthlySalary: []Salary{
			{
				Basic: 400000.0,
			},
			{
				Basic: 30000.0,
			},
		},
	}

	// formattedData, err := json.MarshalIndent(data, "", " ")
	// os.WriteFile("data.txt", formattedData, 0644)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(data); err != nil {
		log.Fatal(err)
	}

}

func ReadStructFromFile() {

	var emp Employee
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(&emp); err != nil {
		log.Fatal(err)
	}

	// data, err := os.ReadFile("data.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err = json.Unmarshal(data, &emp); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("%+v", emp)
}

func GoDot() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(os.Getenv("API_KEY"))

}
