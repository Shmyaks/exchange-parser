package schemes

type paymentJSONScheme struct {
	PaymentName string `json:"paymentName"`
	PaymentType int    `json:"paymentType"`
}

type AliasJSONScheme struct {
	Result []paymentJSONScheme `json:"result"`
}
