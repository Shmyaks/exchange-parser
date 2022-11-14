package service

import (
	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	"github.com/Shmyaks/exchange-parser-server/app/internal/repository"
)

type ArbitrageService struct {
	repository repository.ArbitrateRepository
}

func NewArbitrageService(repository repository.ArbitrateRepository) *ArbitrageService {
	return &ArbitrageService{repository: repository}
}
func (s *ArbitrageService) Get(filter filters.ArbitrageFilter) ([]models.ArbitrageRow, error) {
	return s.repository.Get(filter)
}

func (s *ArbitrageService) Set(arbitrageRow models.ArbitrageRow) error {
	return s.repository.Set(arbitrageRow)
}

func (s *ArbitrageService) GetRowsFromSub() error {
	sub := s.repository.Sub("P2P_ARBITRAGE")
	for {
		arbitrageRows, err := s.repository.GetFromSub(sub)
		if err != nil {
			return err
		}
		if err = s.repository.Replace(arbitrageRows); err != nil {
			return err
		}
	}
}
