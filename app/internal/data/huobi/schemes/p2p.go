package schemes

type p2pDetailOrderInfo struct {
	Currency int64  `json:"currency"`
	Price    string `json:"price"`
	CoinID   int64  `json:"coinId"`
	UserName string `json:"userName"`
}

type P2PJSONScheme struct {
	Data []p2pDetailOrderInfo `json:"data"`
}
