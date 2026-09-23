package actorcommand

import (
	"be/internals/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteActorHandler struct {
	r storage.ActorRepository
}

func NewDeleteActorHandler(r storage.ActorRepository) *DeleteActorHandler {
	return &DeleteActorHandler{r: r}
}

// DeleteActor godoc
// @Summary		Delete an actor (admin)
// @Description	Deletes an actor by ID. Requires role: admin.
// @Tags			Admin - Actors
// @Produce		json
// @Security		BearerAuth
// @Param			id	path		int	true	"Actor ID"
// @Success		200	{object}	map[string]interface{}	"actor deleted successfully"
// @Failure		400	{object}	map[string]interface{}	"Invalid ID"
// @Failure		401	{object}	map[string]interface{}	"Missing or invalid JWT"
// @Failure		403	{object}	map[string]interface{}	"Admin role required"
// @Failure		500	{object}	map[string]interface{}	"Failed to delete actor"
// @Router			/admin/actors/{id} [delete]
func (h *DeleteActorHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.r.DeleteActor(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete actor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "actor deleted successfully"})
}
