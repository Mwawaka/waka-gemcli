package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

func Server() {
	errChan := make(chan error)
	// A goroutine which starts a HTTP server
	go func(channel chan error) {
		mux := http.NewServeMux()

		mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s\n", r.Method)
		})

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", 3000),
			Handler: mux,
		}
		fmt.Println("starting the server on port: 3000 ")
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
				channel <- err
			} else {
				fmt.Println("shutting down the server")
				channel <- nil
			}
		}
	}(errChan)

	err := <-errChan

	if err != nil {
		fmt.Println(err)
	}
}
