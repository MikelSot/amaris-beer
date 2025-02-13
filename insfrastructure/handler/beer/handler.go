package beer

import (
	"github.com/gin-gonic/gin"

	"github.com/MikelSot/amaris-beer/domain/beer"
	"github.com/MikelSot/amaris-beer/insfrastructure/handler/request"
	"github.com/MikelSot/amaris-beer/insfrastructure/handler/response"
	"github.com/MikelSot/amaris-beer/model"
)

type handler struct {
	useCase  beer.UseCase
	response response.ApiResponse
}

func newHandler(useCase beer.UseCase, response response.ApiResponse) handler {
	return handler{useCase: useCase, response: response}
}

func (h handler) Create(c *gin.Context) {
	m := model.Beer{}

	if err := c.Bind(&m); err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}

	if err := h.useCase.Create(c.Request.Context(), &m); err != nil {
		c.JSON(h.response.Error(c, "useCase.Create()", err))

		return
	}

	c.JSON(h.response.Created(m))
}

func (h handler) Update(c *gin.Context) {
	m := model.Beer{}

	if err := c.Bind(&m); err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}

	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}
	m.ID = uint(ID)

	if err := h.useCase.Update(c.Request.Context(), m); err != nil {
		c.JSON(h.response.Error(c, "useCase.Update()", err))

		return
	}

	c.JSON(h.response.Updated())
}

// Delete is the handler to delete a beer
func (h handler) Delete(c *gin.Context) {
	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}

	if err := h.useCase.Delete(c.Request.Context(), uint(ID)); err != nil {
		c.JSON(h.response.Error(c, "useCase.Delete()", err))

		return
	}

	c.JSON(h.response.Deleted())
}

func (h handler) GetByID(c *gin.Context) {
	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}

	m, err := h.useCase.GetByID(c.Request.Context(), uint(ID))
	if err != nil {
		c.JSON(h.response.Error(c, "useCase.GetByID()", err))

		return
	}

	c.JSON(h.response.OK(m))
}

func (h handler) GetAll(c *gin.Context) {
	ms, err := h.useCase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(h.response.Error(c, "useCase.GetAll()", err))
		return
	}

	c.JSON(h.response.OK(ms))
}
