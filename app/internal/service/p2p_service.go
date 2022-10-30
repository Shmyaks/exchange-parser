// Package service ...
package service

import (
	"sync"

	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
	"github.com/Shmyaks/exchange-parser-server/app/internal/repository"
)

// P2PService structs
type P2PService struct {
	P2PRepository repository.P2PRepository
}

// NewP2PService fabric for CurrencyService
func NewP2PService(P2PRepository repository.P2PRepository) *P2PService {
	return &P2PService{P2PRepository: P2PRepository}
}

// Parse orders
func (s *P2PService) Parse(market markets.P2PMarket, filter filters.P2PFilter) *models.P2PPair {
	pairsBuys, _ := s.P2PRepository.GetAllFromData(market, *filter.SetTradeType(filters.Buy))
	pairsSells, _ := s.P2PRepository.GetAllFromData(market, *filter.SetTradeType(filters.Sell))
	var buyPrice, sellPrice string
	if len(pairsBuys) != 0 {
		buyPrice = pairsBuys[0].Price
	}
	if len(pairsSells) != 0 {
		sellPrice = pairsSells[0].Price
	}
	return models.NewP2PPair(filter.CryptoCurrency,
		filter.Fiat, buyPrice, sellPrice, market, filter.PayType)
}

// ParseAll orders
func (s *P2PService) ParseAll(market markets.P2PMarket) error {
	mp, err := s.P2PRepository.GetPayMethods(market)
	if err != nil {
		return err
	}

	pairsAllCurrencies := [len(filters.P2PCurrencies)][]models.P2PPair{}

	parseFunc := func(
		pairs *[]models.P2PPair,
		wg *sync.WaitGroup,
		mu *sync.Mutex,
		payMethod models.PayMethod,
		fiat models.Fiat,
		currency models.CryptoCurrency) {

		defer wg.Done()
		println(mu, string(currency), fiat, payMethod, wg)
		pair := s.Parse(market, *filters.NewP2PFilter(currency, fiat, payMethod))
		mu.Lock()
		*pairs = append(*pairs, *pair)
		mu.Unlock()
	}

	var wg sync.WaitGroup
	for i, currency := range filters.P2PCurrencies {
		var mu sync.Mutex
		for fiat, payMethods := range mp {
			for _, payMethod := range payMethods {
				wg.Add(1)
				go parseFunc(&pairsAllCurrencies[i], &wg, &mu, payMethod, fiat, currency)
			}
		}
	}
	wg.Wait()
	for _, pairs := range pairsAllCurrencies {
		s.P2PRepository.Set(market, pairs)
	}
	return nil
}

// Get orders from Cache
func (s *P2PService) Get(market markets.P2PMarket, pair models.CurencyPairName) []models.P2PPair {
	return s.P2PRepository.GetFromCache(market, pair)
}

// GetAll orders from Cache
func (s *P2PService) GetAll(market markets.P2PMarket) []models.P2PPair {
	return s.P2PRepository.GetAllFromCache(market)
}
