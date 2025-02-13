package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type RouterSpecification struct {
	Api        *gin.Engine
	Logger     Logger
	DB         PgxPool
	TxDB       *pgxpool.Pool
	Redis      *redis.Client
	Stream     *redis.Client
	Threshold  int
	StreamName string
}
