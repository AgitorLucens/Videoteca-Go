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

// CreateActor godoc
// @Summary		Create an actor (admin)
// @Description	Creates a new actor. Requires role: admin.
// @Tags			Admin - Actors
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		storage.CreateActorRequest	true	"Actor data"
// @Success		201		{object}	map[string]interface{}		"data: created actor"
// @Failure		400		{object}	map[string]interface{}		"Invalid request payload"
// @Failure		401		{object}	map[string]interface{}		"Missing or invalid JWT"
// @Failure		403		{object}	map[string]interface{}		"Admin role required"
// @Failure		500		{object}	map[string]interface{}		"Failed to create actor"
// @Router			/admin/actors [post]
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
