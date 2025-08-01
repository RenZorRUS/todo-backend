package driver

import (
	"context"
	"errors"

	"github.com/RenZorRUS/todo-backend/src/internal/core/domains/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgxPool(config *configs.DBConfig) (*pgxpool.Pool, error) {
	var errs error

	poolConfig, err := BuildPoolConfig(config)
	errs = errors.Join(errs, err)

	ctx, cancel := context.WithTimeout(context.Background(), config.PoolBuildTimeout)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)

	errs = errors.Join(errs, err)
	if errs != nil {
		return nil, errs
	}

	return pool, nil
}

func BuildPoolConfig(config *configs.DBConfig) (*pgxpool.Config, error) {
	poolConfig, err := pgxpool.ParseConfig(config.ConnUrl)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConnLifetime = config.MaxConnLifeTime
	poolConfig.MaxConnLifetimeJitter = config.MaxConnLifeTimeJitter
	poolConfig.MaxConnIdleTime = config.MaxConnIdleTime
	poolConfig.MaxConns = config.MaxConns
	poolConfig.MinConns = config.MinConns
	poolConfig.MinIdleConns = config.MinIdleConns
	poolConfig.HealthCheckPeriod = config.HealthCheckPeriod

	return poolConfig, nil
}
