// Package schemes is json schemes of Bybit
package schemes

import "github.com/Shmyaks/exchange-parser-server/app/internal/models"

type bybitP2PDetailOrderInfo struct {
	TokenID  models.CryptoCurrency `json:"tokenId"`
	Price    string                `json:"price"`
	Fiat     models.Fiat           `json:"currencyId"`
	NickName string                `json:"nickName"`
}

type bybitP2PDetailInfo struct {
	Count int64                     `json:"count"`
	Items []bybitP2PDetailOrderInfo `json:"items"`
}

// P2PJSONScheme general json scheme for bybit
type P2PJSONScheme struct {
	Result bybitP2PDetailInfo `json:"result"`
}
