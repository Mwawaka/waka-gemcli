package main

import (
	"fmt"
	"log"

	"github.com/Mwawaka/go-crazy/internal/utils"
)

func main() {
	result, err := utils.Divide(15, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
