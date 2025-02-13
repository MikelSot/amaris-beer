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
	db, txDb := newDatabase(ctx, dbConfig, applicationName)

	streamConfig := model.NewStreamConfig()
	stream := newRedisClient(ctx, streamConfig)

	redisConfig := model.NewRedisConfig()
	redis := newRedisClient(ctx, redisConfig)

	ginEntry := newGinEntry(boot)
	ginEntry.Bootstrap(ctx)

	logger := newLogger()

	api := ginEntry.Router

	handler.InitRoutes(model.RouterSpecification{
		Api:        api,
		Logger:     logger,
		DB:         db,
		TxDB:       txDb,
		Redis:      redis,
		Stream:     stream,
		StreamName: steamName,
		Threshold:  threshold,
	})

	rkentry.GlobalAppCtx.WaitForShutdownSig()
	ginEntry.Interrupt(ctx)
}
