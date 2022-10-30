// Package internal ...
package internal

import (
	"github.com/Shmyaks/exchange-parser-server/app/config"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/binance"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/bybit"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/huobi"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/okx"
	"github.com/Shmyaks/exchange-parser-server/app/internal/delivery/controllers"
	"github.com/Shmyaks/exchange-parser-server/app/internal/repository"
	"github.com/Shmyaks/exchange-parser-server/app/internal/service"
	"github.com/Shmyaks/exchange-parser-server/app/pkg/redis"

	"github.com/imroc/req/v3"

	jsoniter "github.com/json-iterator/go"
)

type packageContainer struct {
	redisConnectionPackage redis.Connection
}
type dataContainer struct {
	dataP2PBinance binance.P2PData
	dataP2PBybit   bybit.P2PData
	dataP2PHuobi   huobi.P2PData
	dataP2POkx     okx.P2PData

	dataBinance binance.SPOTData
	dataBybit   bybit.SPOTData
	dataHuobi   huobi.SPOTData
	dataOkx     okx.SPOTData
}
type controllersContainer struct {
	P2PController  controllers.P2PController
	SpotController controllers.SPOTController
}
type serviceContainer struct {
	P2PService  service.P2PService
	SPOTService service.SPOTService
}

type repositoryContainer struct {
	p2pRepository  repository.P2PRepository
	spotRepository repository.SPOTRepository
}

// ApplicationContainer app Container
type ApplicationContainer struct {
	PackageContainer    packageContainer
	repositoryContainer repositoryContainer
	ServiceContainer    serviceContainer
	Controllers         controllersContainer
}

func initPackages() *packageContainer {
	conf := config.GetConfig()
	return &packageContainer{
		redisConnectionPackage: *redis.NewConnection(conf.RedisConnection),
	}
}

func initData() *dataContainer {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	client := req.NewClient().DisableAutoDecode().SetJsonMarshal(json.Marshal).
		SetJsonUnmarshal(json.Unmarshal)
	return &dataContainer{
		dataP2PBinance: *binance.NewP2PData(*client),
		dataP2PBybit:   *bybit.NewP2PData(*client),
		dataP2PHuobi:   *huobi.NewP2PData(*client),
		dataP2POkx:     *okx.NewP2PData(*client),
		dataBinance:    *binance.NewSPOTData(*client),
		dataBybit:      *bybit.NewSPOTData(*client),
		dataHuobi:      *huobi.NewSPOTData(*client),
		dataOkx:        *okx.NewSPOTData(*client),
	}
}

func initRepositories(dataContainer dataContainer, packages packageContainer) *repositoryContainer {
	return &repositoryContainer{
		p2pRepository: *repository.NewP2PRepository([]data.P2P{
			&dataContainer.dataP2PBinance,
			&dataContainer.dataP2PBybit,
			&dataContainer.dataP2PHuobi,
			&dataContainer.dataP2POkx}, packages.redisConnectionPackage),
		spotRepository: *repository.NewSPOTRepository([]data.SPOT{
			&dataContainer.dataBinance,
			&dataContainer.dataBybit,
			&dataContainer.dataHuobi,
			&dataContainer.dataOkx,
		}, packages.redisConnectionPackage),
	}
}

func initServices(repositories repositoryContainer, packages packageContainer) *serviceContainer {
	SPOTservice := service.NewSPOTService(repositories.spotRepository)
	P2PService := service.NewP2PService(repositories.p2pRepository)

	return &serviceContainer{
		P2PService:  *P2PService,
		SPOTService: *SPOTservice,
	}
}

func initControllers(services serviceContainer) *controllersContainer {
	return &controllersContainer{
		P2PController:  *controllers.NewP2PController(services.P2PService),
		SpotController: *controllers.NewSPOTController(services.SPOTService),
	}
}

// InitApplication initialization full app
func InitApplication() *ApplicationContainer {

	packageContainer := initPackages()
	dataContainer := initData()
	repositoryContainer := initRepositories(*dataContainer, *packageContainer)
	serviceContainer := initServices(*repositoryContainer, *packageContainer)
	controllerContainer := initControllers(*serviceContainer)

	return &ApplicationContainer{
		PackageContainer:    *packageContainer,
		repositoryContainer: *repositoryContainer,
		ServiceContainer:    *serviceContainer,
		Controllers:         *controllerContainer}
}
