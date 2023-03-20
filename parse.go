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

	//out, _ := xml.MarshalIndent(productList, "", "  ")
	//fmt.Println(string(out))

	var coveredServices, uncoveredService []*scheme.GeneralHealthServiceCategory
	for _, service := range productList.Product[0].GeneralHealthCover.GeneralHealthServices.GeneralHealthService {
		if service.Covered {
			coveredServices = append(coveredServices, service)
		} else {
			uncoveredService = append(uncoveredService, service)
		}
	}

	fmt.Println("This plan covers:")
	for _, service := range coveredServices {
		fmt.Printf(" - %s\n", service.Title)
	}
	fmt.Println("It does NOT cover:")
	for _, service := range uncoveredService {
		fmt.Printf(" - %s\n", service.Title)
	}

	return nil
}
