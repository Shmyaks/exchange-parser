package controllers

import "github.com/Shmyaks/exchange-parser-server/app/internal/service"

type ArbitrageController struct {
	service service.ArbitrageService
}

func NewArbitrageController(service service.ArbitrageService) *ArbitrageController {
	return &ArbitrageController{service: service}
}

func (c *ArbitrageController) Start() {
	for {
		if err := c.service.GetRowsFromSub(); err != nil {
			print(err.Error())
		}
	}
}
