package calculator

type Request struct {
	AddresFrom string `json:"addr_from"`
	AddresTo   string `json:"addr_to"`
	Weight     int    `json:"weight"`
	Length     int    `json:"length"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

type GetToken struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type RecieveToken struct {
	Access_token string `json:"access_token"`
}

type PriceSending struct {
	TariffCode        int     `json:"tariff_code"`
	TariffName        string  `json:"tariff_name"`
	TariffDescription string  `json:"tariff_description"`
	DeliveryMode      int     `json:"delivery_mode"`
	DeliverySum       float32 `json:"delivery_sum"`
	PeriodMin         int     `json:"period_min"`
	PeriodMax         int     `json:"period_max"`
}

type ResPrice struct {
	Prices []PriceSending `json:"tariff_codes"`
}

type PostDelivery struct {
	AddresFrom Location  `json:"from_location"`
	AddresTo   Location  `json:"to_location"`
	Packages   []Package `json:"packages"`
}

type Location struct {
	Address string `json:"address"`
}

type Package struct {
	Weight int `json:"weight"`
	Length int `json:"length"`
	Width  int `json:"width"`
	Height int `json:"height"`
}
