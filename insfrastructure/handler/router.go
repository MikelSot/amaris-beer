package handler

import (
	"github.com/MikelSot/amaris-beer/insfrastructure/handler/beer"
	"github.com/MikelSot/amaris-beer/model"
)

func InitRoutes(spec model.RouterSpecification) {
	// B
	beer.NewRouter(spec)
}
