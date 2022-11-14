package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"

	"github.com/Shmyaks/exchange-parser-server/app/internal/models"
	"github.com/Shmyaks/exchange-parser-server/app/internal/models/filters"
	pkg "github.com/Shmyaks/exchange-parser-server/app/pkg/redis"
)

type ArbitrateRepository struct {
	db              *sqlx.DB
	redisConnection pkg.Connection
}

func NewArbitrateRepository(db *sqlx.DB, redisConection pkg.Connection) *ArbitrateRepository {
	return &ArbitrateRepository{db: db, redisConnection: redisConection}
}

func (r *ArbitrateRepository) Get(filter filters.ArbitrageFilter) ([]models.ArbitrageRow, error) {
	arbitrageRow := []models.ArbitrageRow{}
	if err := r.db.Select(
		&arbitrageRow, "SELECT * FROM arbitrage ORDER BY percent DESC LIMIT $1 OFFSET $2", filter.Limit,
		filter.Offset,
	); err != nil {
		return arbitrageRow, err
	}
	return arbitrageRow, nil
}

func (r *ArbitrateRepository) Set(arbitrageRow models.ArbitrageRow) error {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(
		context.Background(),
		"INSERT INTO arbitrage (buy, first_pay_type, sell, second_pay_type, percent) VALUES ($1, $2, $3, $4 ,$5)",
		arbitrageRow.Buy, arbitrageRow.FirstPayType, arbitrageRow.Sell, arbitrageRow.SecondPayType,
		arbitrageRow.Percent,
	)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	return tx.Commit()
}

func (r *ArbitrateRepository) Replace(arbitrageRows []models.ArbitrageRow) error {
	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(context.Background(), "TRUNCATE arbitrage; DELETE FROM arbitrage;")
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	_, err = tx.NamedExec(
		`INSERT INTO arbitrage (fiat, asset, buy, first_pay_type, sell, second_pay_type, percent, market_id) VALUES (:fiat, :asset, :buy, :first_pay_type, :sell, :second_pay_type, :percent, :market_id)`,
		arbitrageRows,
	)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *ArbitrateRepository) Sub(name string) *redis.PubSub {
	return r.redisConnection.Pool.Subscribe(context.Background(), name)
}

func (r *ArbitrateRepository) GetFromSub(sub *redis.PubSub) ([]models.ArbitrageRow, error) {
	arbitrageRows := make([]models.ArbitrageRow, 0, 128)
	mes, err := sub.ReceiveMessage(context.Background())
	if err != nil {
		return arbitrageRows, err
	}
	if err = jsoniter.UnmarshalFromString(mes.Payload, &arbitrageRows); err != nil {
		return arbitrageRows, err
	}
	return arbitrageRows, nil
}
