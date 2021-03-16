package model

type Price struct {
	Last            float32 `json:"last"`
	PriceBeforeLast float32 `json:"price_before_last"`
	Open            float32 `json:"open"`
	High            float32 `json:"high"`
	Low             float32 `json:"low"`
	Vwap            float32 `json:"vwap"`
	Volume          float32 `json:"volume"`
	Bid             float32 `json:"bid"`
	Ask             float32 `json:"ask"`
	Text 			string `json:"text"`
}


