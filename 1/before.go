package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Invoices []struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {

	playsFile, err := ioutil.ReadFile("plays.json")
	if err != nil {
		fmt.Println(err)
	}

	var plays map[string]Play
	if err := json.Unmarshal(playsFile, &plays); err != nil {
		fmt.Println(err)
	}

	invoiceFile, err := ioutil.ReadFile("invoices.json")
	if err != nil {
		fmt.Println(err)
	}

	var invoices Invoices
	if err := json.Unmarshal(invoiceFile, &invoices); err != nil {
		fmt.Println(err)
	}

	fmt.Println(plays)
	fmt.Println(invoices)

}

func statement(invoice string, plays map[string]Play) {

}
