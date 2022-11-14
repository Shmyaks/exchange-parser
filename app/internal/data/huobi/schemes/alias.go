package schemes

type payMethodAliasJSONScheme struct {
	Name        string `json:"name"`
	PayMethodID int    `json:"payMethodId"`
}

type dataAliasJSONScheme struct {
	PayMethod []payMethodAliasJSONScheme `json:"payMethod"`
}

type AliasJSONScheme struct {
	Data dataAliasJSONScheme `json:"data"`
}
