package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const SEARCH_PATH = "catalog"

type DBHandler struct {
	pool   *pgxpool.Pool
	schema string
}

func NewDBHandler(pool *pgxpool.Pool, schema string) *DBHandler {
	return &DBHandler{
		pool:   pool,
		schema: schema,
	}
}

func (h *DBHandler) getDbConnection() (*pgxpool.Conn, error) {
	dbConn, err := h.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}

	_, err = dbConn.Exec(context.Background(), fmt.Sprintf("SET search_path TO %s", h.schema))
	if err != nil {
		dbConn.Release()
		return nil, err
	}

	return dbConn, nil

}
