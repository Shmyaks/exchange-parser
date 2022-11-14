package models

// PayMethod Pay methods for P2P
type PayMethod string

// Tinkoff, Sberbank, Raif, Rosbank, QIWI ENUM for PayMethod
const (
	Tinkoff   PayMethod = "TINKOFF"
	Sberbank  PayMethod = "SBERBANK"
	Raif      PayMethod = "RAIFFEISENBANK"
	Rosbank   PayMethod = "ROSBANK"
	QIWI      PayMethod = "QIWI"
	YOUMoney  PayMethod = "YOUMONEY"
	MTSBank   PayMethod = "MTSBANK"
	AdvCash   PayMethod = "ADVCASH"
	Payeer    PayMethod = "PAYEER"
	Alfabank  PayMethod = "ALFABANK"
	Wise      PayMethod = "WISE"
	Zelle     PayMethod = "ZELLE"
	Revolut   PayMethod = "REVOLUT"
	Ziraat    PayMethod = "ZIRAAT"
	DenizBank PayMethod = "DENIZBANK"
	VakifBank PayMethod = "VAKIFBANK"
	Akbank    PayMethod = "AKBANK"
)

// P2PCurrencies ENUM for P2P CryptoCurrency
var P2PCurrencies = [3]PayMethod{"USDT", "BTC", "ETH"}
