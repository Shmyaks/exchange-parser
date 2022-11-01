package markets

// P2PMarket type of enum
type P2PMarket uint16

// BinanceP2P, ... -> ENUM
const (
	BinanceP2P P2PMarket = iota + 1
	BybitP2P
	HuobiP2P
	OkxP2P
)

// AllP2PMarket array of P2PMarket ENUM
var AllP2PMarket = [4]P2PMarket{BinanceP2P, BybitP2P, HuobiP2P, OkxP2P}

// GetName get name of P2PMarket ENUM
func (m *P2PMarket) GetName() string {
	var name string
	switch *m {
	case BinanceP2P:
		name = "BinanceP2P"
	case BybitP2P:
		name = "BybitP2P"
	case HuobiP2P:
		name = "HuobiP2P"
	case OkxP2P:
		name = "OkxP2P"
	}

	return name
}

// GetMarket Get P2PMarket ENUM string
func GetMarket(name string) (*P2PMarket, error) {
	var market P2PMarket
	switch name {
	case "BinanceP2P":
		market = BinanceP2P
	case "BybitP2P":
		market = BybitP2P
	case "HuobiP2P":
		market = HuobiP2P
	case "OkxP2P":
		market = OkxP2P
	}

	return &market, nil
}
