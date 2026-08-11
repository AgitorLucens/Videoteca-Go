package actorqueries

import (
	"be/internals/storage"
	"github.com/gin-gonic/gin"

	"be/internals/handler"
)

type GetActorsHandler struct {
	r storage.ActorRepository
}

func NewGetActorsHandler(r storage.ActorRepository) *GetActorsHandler {
	return &GetActorsHandler{r: r}
}

func (h *GetActorsHandler) Handle(c *gin.Context) {
	actors, err := h.r.GetAllActors()
	if err != nil {
		handler.WriteInternalServerError(c, "failed to get actors")
		return
	}
	handler.WriteSuccess(c, actors)
}
