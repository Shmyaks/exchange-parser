package markets

// Market scheme of json
type Market struct {
	ID   uint16 `json:"id"`
	Name string `json:"name"`
}

// NewMarket fabric
func NewMarket(id uint16, name string) *Market {
	return &Market{ID: id, Name: name}
}
