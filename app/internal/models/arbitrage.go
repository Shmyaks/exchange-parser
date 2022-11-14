package models

type ArbitrageRow struct {
	ID            int     `json:"id" db:"id"`
	Buy           float32 `json:"buy" db:"buy"`
	Asset         string  `json:"asset" db:"asset"`
	Fiat          Fiat    `json:"fiat" db:"fiat"`
	FirstPayType  string  `json:"first_pay_type" db:"first_pay_type"`
	Sell          float32 `json:"sell" db:"sell"`
	SecondPayType string  `json:"second_pay_type" db:"second_pay_type"`
	Percent       float32 `json:"percent" db:"percent"`
	MarketID      int     `json:"market_id" db:"market_id"`
}
