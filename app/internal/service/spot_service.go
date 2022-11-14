// Package service file for services
package service

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/markets"
	"github.com/Shmyaks/exchange-parser-server/app/internal/repository"
)

// SPOTService structs
type SPOTService struct {
	SPOTRepository repository.SPOTRepository
}

// NewSPOTService favric for CurrencyService
func NewSPOTService(repository repository.SPOTRepository) *SPOTService {
	return &SPOTService{SPOTRepository: repository}
}

// Parse orders
func (s *SPOTService) Parse(market markets.SPOTMarket) error {
	resp, err := s.SPOTRepository.GetAllFromData(market)
	if err != nil {
		return err
	}
	return s.SPOTRepository.SetMany(market, resp)
}

// Get orders from Cache
func (s *SPOTService) Get(market markets.SPOTMarket, currencyPairName models.CurencyPairName) *models.CurrencyPair {
	return s.SPOTRepository.GetFromCache(market, currencyPairName)
}

// GetAll orders from Cache
func (s *SPOTService) GetAll(market markets.SPOTMarket) []models.CurrencyPair {
	return s.SPOTRepository.GetAllFromCache(market)
}
