package schemes

import "github.com/Shmyaks/exchange-parser-server/app/internal/models"

type p2pOrderJSONScheme struct {
	Price          string                `json:"price"`
	Fiat           models.Fiat           `json:"quoteCurrency"`
	CryptoCurrency models.CryptoCurrency `json:"baseCurrency"`
	NickName       string                `json:"nickName"`
}

type trateTypeJSONScheme struct {
	Buy  []p2pOrderJSONScheme `json:"buy"`
	Sell []p2pOrderJSONScheme `json:"sell"`
}

// P2PJSONScheme general json scheme for okx
type P2PJSONScheme struct {
	Data trateTypeJSONScheme `json:"data"`
}
