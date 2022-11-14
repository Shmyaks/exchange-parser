package filters

import "github.com/Shmyaks/exchange-parser-server/app/internal/models"

// TradeType type
type TradeType string

// BUY, SELL TradeType Filter for P2P method
const (
	Buy  TradeType = "BUY"
	Sell TradeType = "SELL"
)

// P2POption for P2PFilter
type P2POption func(*P2PFilter)

// P2PFilter struct for Filter P2P methods
type P2PFilter struct {
	CryptoCurrency models.CryptoCurrency
	Fiat           models.Fiat
	PayType        models.PayMethod
	TradeType      TradeType
	MinAmount      int
	MinOrders      int
}

// NewP2PFilter fabric for P2PFilter
func NewP2PFilter(cryptoCurrency models.CryptoCurrency, fiat models.Fiat, payTypes models.PayMethod, minAmount int, minOrders int) *P2PFilter {
	return &P2PFilter{CryptoCurrency: cryptoCurrency, Fiat: fiat, PayType: payTypes, TradeType: Buy, MinAmount: minAmount, MinOrders: minOrders}
}

// SetTradeType set TradeType field for P2PFilter
func (p *P2PFilter) SetTradeType(tradeType TradeType) *P2PFilter {
	p.TradeType = tradeType
	return p
}

// P2PCurrencies array for P2P currencies: use only this crypto currencies
var P2PCurrencies = [3]models.CryptoCurrency{"USDT", "BTC", "ETH"}
