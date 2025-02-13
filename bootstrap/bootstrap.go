package bootstrap

import (
	"context"

	"github.com/joho/godotenv"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"

	"github.com/MikelSot/amaris-beer/insfrastructure/handler"
	"github.com/MikelSot/amaris-beer/model"
)

func Run(boot []byte) {
	_ = godotenv.Load()

	ctx := context.Background()

	applicationName := getApplicationName()
	threshold := getHighDemandThreshold()
	steamName := getStreamName()

	dbConfig := model.NewDatabaseConfig()
	db := newDatabase(ctx, dbConfig, applicationName)

	redisConfig := model.NewRedisConfig()
	redis := NewRedisClient(ctx, redisConfig)

	ginEntry := newGinEntry(boot)
	ginEntry.Bootstrap(ctx)

	logger := newLogger()

	api := ginEntry.Router

	handler.InitRoutes(model.RouterSpecification{
		Api:        api,
		Logger:     logger,
		DB:         db,
		Redis:      redis,
		Threshold:  threshold,
		StreamName: steamName,
	})

	rkentry.GlobalAppCtx.WaitForShutdownSig()
	ginEntry.Interrupt(ctx)
}
