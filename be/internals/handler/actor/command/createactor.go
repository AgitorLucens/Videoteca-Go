package actorcommand

import (
	"be/internals/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateActorHandler struct {
	r storage.ActorRepository
}

func NewCreateActorHandler(r storage.ActorRepository) *CreateActorHandler {
	return &CreateActorHandler{r: r}
}

func (h *CreateActorHandler) Handle(c *gin.Context) {
	var req storage.CreateActorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actor, err := h.r.CreateActor(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create actor"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": actor})
}
