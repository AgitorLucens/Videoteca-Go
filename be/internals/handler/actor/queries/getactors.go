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

// GetActors godoc
// @Summary		List all actors (admin)
// @Description	Returns all actors. Requires role: admin.
// @Tags			Admin - Actors
// @Produce		json
// @Security		BearerAuth
// @Success		200	{object}	handler.Response{data=[]storage.Actor}
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403	{object}	map[string]interface{}	"Admin role required"
// @Failure		500	{object}	handler.Response	"Failed to get actors"
// @Router			/admin/actors [get]
func (h *GetActorsHandler) Handle(c *gin.Context) {
	actors, err := h.r.GetAllActors()
	if err != nil {
		handler.WriteInternalServerError(c, "failed to get actors")
		return
	}
	handler.WriteSuccess(c, actors)
}
