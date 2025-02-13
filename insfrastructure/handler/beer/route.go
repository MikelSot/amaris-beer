package beer

import (
	"github.com/MikelSot/amaris-beer/domain/beer"
	"github.com/MikelSot/amaris-beer/domain/beerview"
	"github.com/MikelSot/amaris-beer/insfrastructure/handler/response"
	beerStorage "github.com/MikelSot/amaris-beer/insfrastructure/postgres/beer"
	"github.com/MikelSot/amaris-beer/insfrastructure/postgres/transaction"
	"github.com/MikelSot/amaris-beer/insfrastructure/redis"
	"github.com/MikelSot/amaris-beer/model"
	"github.com/gin-gonic/gin"
)

const (
	_privateRoutesPrefix = "/beers"
)

func NewRouter(spec model.RouterSpecification) {
	handler := buildHandler(spec)

	privateRoutes(spec.Api, handler)
}

func buildHandler(spec model.RouterSpecification) handler {
	response := response.New(spec.Logger)

	tx := transaction.New(spec.TxDB)
	storage := beerStorage.New(spec.DB)

	cache := redis.NewRedis(spec.Redis)
	stream := redis.NewStream(spec.Stream, spec.StreamName)

	beerViewUseCase := beerview.New(cache, stream, spec.Threshold)

	useCase := beer.New(storage, tx, beerViewUseCase)

	return newHandler(useCase, response)
}

func privateRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group(_privateRoutesPrefix, middlewares...)

	routes.POST("", h.Create)
	routes.PUT("/:id", h.Update)
	routes.DELETE("/:id", h.Delete)

	routes.GET("/:id", h.GetByID)
	routes.GET("", h.GetAll)
}
