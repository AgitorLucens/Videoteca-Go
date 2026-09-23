package actorcommand

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateActorHandler struct {
	r storage.ActorRepository
}

func NewUpdateActorHandler(r storage.ActorRepository) *UpdateActorHandler {
	return &UpdateActorHandler{r: r}
}

// UpdateActor godoc
// @Summary		Update an actor (admin)
// @Description	Updates an existing actor. Requires role: admin.
// @Tags			Admin - Actors
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			id		path		int							true	"Actor ID"
// @Param			request	body		storage.UpdateActorRequest	true	"Actor data"
// @Success		200		{object}	map[string]interface{}		"data: updated actor"
// @Failure		400		{object}	map[string]interface{}		"Invalid ID or payload"
// @Failure		401		{object}	map[string]interface{}		"Missing or invalid JWT"
// @Failure		403		{object}	map[string]interface{}		"Admin role required"
// @Failure		500		{object}	map[string]interface{}		"Failed to update actor"
// @Router			/admin/actors/{id} [put]
func (h *UpdateActorHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req storage.UpdateActorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actor, err := h.r.UpdateActor(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update actor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": actor})
}
