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
