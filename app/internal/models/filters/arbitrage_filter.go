package filters

type ArbitrageFilter struct {
	Offset int32
	Limit  int32
}

func NewArbitrageFilter(offset int32, limit int32) *ArbitrageFilter {
	return &ArbitrageFilter{Offset: offset, Limit: limit}
}
