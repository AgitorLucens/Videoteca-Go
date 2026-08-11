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
