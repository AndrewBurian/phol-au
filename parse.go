package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"go.burian.dev/private-health-au/scheme"
)

func main() {
	var err error

	//err = printSingle()
	err = decodeSingle()

	if err != nil {
		panic(err)
	}

}

func decodeSingle() error {
	file, err := os.Open("oneplan.xml")
	if err != nil {
		return err
	}

	productList := new(scheme.ProductList)

	decoder := xml.NewDecoder(file)
	err = decoder.Decode(productList)
	if err != nil {
		return err
	}

	out, _ := xml.MarshalIndent(productList, "", "  ")
	fmt.Println(string(out))

	return nil
}
