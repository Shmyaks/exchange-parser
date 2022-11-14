// Package internal ...
package internal

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/Shmyaks/exchange-parser-server/app/config"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/binance"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/bybit"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/huobi"
	"github.com/Shmyaks/exchange-parser-server/app/internal/data/okx"
	"github.com/Shmyaks/exchange-parser-server/app/internal/delivery/controllers"
	"github.com/Shmyaks/exchange-parser-server/app/internal/repository"
	"github.com/Shmyaks/exchange-parser-server/app/internal/service"
	pkg "github.com/Shmyaks/exchange-parser-server/app/pkg/redis"

	"github.com/imroc/req/v3"

	jsoniter "github.com/json-iterator/go"
)

type packageContainer struct {
	redisConnection pkg.Connection
	db              *sqlx.DB
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
	P2PController       controllers.P2PController
	SpotController      controllers.SPOTController
	ArbitrageController controllers.ArbitrageController
}
type serviceContainer struct {
	P2PService       service.P2PService
	SPOTService      service.SPOTService
	ArbitrageService service.ArbitrageService
}

type repositoryContainer struct {
	p2pRepository       repository.P2PRepository
	spotRepository      repository.SPOTRepository
	arbitrageRepository repository.ArbitrateRepository
}

// ApplicationContainer app Container
type ApplicationContainer struct {
	PackageContainer    packageContainer
	repositoryContainer repositoryContainer
	ServiceContainer    serviceContainer
	Controllers         controllersContainer
}

func initPackages() *packageContainer {
	db, err := sqlx.Connect(
		"pgx", fmt.Sprintf(
			"postgres://%s:%s@localhost:5432/%s", config.Env.PostgresUser, config.Env.PostgresPassword,
			config.Env.PostgresDB,
		),
	)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(15)
	db.SetMaxOpenConns(15)
	return &packageContainer{
		redisConnection: *pkg.NewConnection(config.Env.RedisConnection),
		db:              db,
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
		p2pRepository: *repository.NewP2PRepository(
			[]data.P2P{
				&dataContainer.dataP2PBinance,
				&dataContainer.dataP2PBybit,
				&dataContainer.dataP2PHuobi,
				&dataContainer.dataP2POkx}, packages.redisConnection,
		),
		spotRepository: *repository.NewSPOTRepository(
			[]data.SPOT{
				&dataContainer.dataBinance,
				&dataContainer.dataBybit,
				&dataContainer.dataHuobi,
				&dataContainer.dataOkx,
			}, packages.redisConnection,
		),
		arbitrageRepository: *repository.NewArbitrateRepository(packages.db, packages.redisConnection),
	}
}

func initServices(repositories repositoryContainer, packages packageContainer) *serviceContainer {
	SPOTservice := service.NewSPOTService(repositories.spotRepository)
	P2PService := service.NewP2PService(repositories.p2pRepository)
	ArbitrageService := service.NewArbitrageService(repositories.arbitrageRepository)
	return &serviceContainer{
		P2PService:       *P2PService,
		SPOTService:      *SPOTservice,
		ArbitrageService: *ArbitrageService,
	}
}

func initControllers(services serviceContainer) *controllersContainer {
	return &controllersContainer{
		P2PController:       *controllers.NewP2PController(services.P2PService),
		SpotController:      *controllers.NewSPOTController(services.SPOTService),
		ArbitrageController: *controllers.NewArbitrageController(services.ArbitrageService),
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
