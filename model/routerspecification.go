package model

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RouterSpecification struct {
	Api        *gin.Engine
	Logger     Logger
	DB         PgxPool
	Redis      *redis.Client
	Threshold  int
	StreamName string
}
