package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sdek/calculator"
)

func main() {
	req := calculator.Request{
		AddresFrom: "Россия, г. Москва, Cлавянский бульвар д.1",
		AddresTo:   "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43",
		Weight:     10,
		Length:     10,
		Height:     10,
		Width:      10,
	}

	cl := calculator.NewClient(&http.Client{})
	tk, err := cl.GetAuth()
	if err != nil {
		fmt.Printf("error recieve token %v", err)
	}
	rp, err := cl.Calculate(req, tk)
	if err != nil {
		fmt.Printf("calcualtion error %v", err)
	}
	res, _ := json.Marshal(rp)
	fmt.Println(string(res))

}
