package models

// PayMethod Pay methods for P2P
type PayMethod string

// Tinkoff, Sberbank, Raif, Rosbank, QIWI ENUM for PayMethod
const (
	Tinkoff  PayMethod = "Tinkoff"
	Sberbank PayMethod = "Sberbank"
	Raif     PayMethod = "RaiffeisenBank"
	Rosbank  PayMethod = "Rosbank"
	QIWI     PayMethod = "QIWI"
)

// P2PCurrencies ENUM for P2P CryptoCurrency
var P2PCurrencies = [3]PayMethod{"USDT", "BTC", "ETH"}
