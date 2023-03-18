package calculator

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient(cl *http.Client) *Client {
	return &Client{
		client: cl,
	}
}

func (cl *Client) GetAuth() (RecieveToken, error) {
	req, err := http.NewRequest("POST", "https://api.edu.cdek.ru/v2/oauth/token?grant_type=client_credentials&client_id=EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI&client_secret=PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG", nil)
	if err != nil {
		return RecieveToken{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := cl.client.Do(req)
	if err != nil {
		return RecieveToken{}, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return RecieveToken{}, err
	}
	var rt RecieveToken
	json.Unmarshal(bodyBytes, &rt)
	return rt, nil
}

func (cl *Client) Calculate(rq Request, tk RecieveToken) (ResPrice, error) {
	var responseObject ResPrice
	ps := ConvertToDelivery(rq)
	myBytes, _ := json.Marshal(ps)
	req, err := http.NewRequest("POST", "https://api.edu.cdek.ru/v2/calculator/tarifflist", bytes.NewBuffer(myBytes))
	if err != nil {
		return ResPrice{}, err
	}
	bearer := "Bearer " + tk.Access_token
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := cl.client.Do(req)
	if err != nil {
		return ResPrice{}, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResPrice{}, err
	}
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject, nil
}

func ConvertToDelivery(rq Request) PostDelivery {
	return PostDelivery{
		AddresFrom: Location{Address: rq.AddresFrom},
		AddresTo:   Location{Address: rq.AddresTo},
		Packages: []Package{{
			Weight: rq.Weight,
			Length: rq.Length,
			Width:  rq.Width,
			Height: rq.Height}}}
}
