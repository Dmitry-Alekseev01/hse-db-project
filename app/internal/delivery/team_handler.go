package delivery

import (
	"github.com/gin-gonic/gin"
	"hse-football/internal/domain"
	"net/http"
	"strconv"
)

type TeamHandler struct {
	svc domain.TeamUsecase
}

func NewTeamHandler(s domain.TeamUsecase) *TeamHandler {
	return &TeamHandler{svc: s}
}

type createTeamDTO struct {
	Name   string  `json:"team_name" binding:"required"`
	Budget float64 `json:"budget" binding:"required"`
	ClubID int64   `json:"club_id" binding:"required"`
}

func (h *TeamHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/teams")
	r.POST("", h.create)
	r.GET("", h.list)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

// create godoc
// @Summary      Создать Team
// @Tags         team
// @Accept       json
// @Produce      json
// @Param        input  body      createTeamDTO  true  "Данные"
// @Success      201   {object}  domain.Team
// @Failure      400   {object}  domain.ErrorResponse
// @Router       /teams [post]
func (h *TeamHandler) create(c *gin.Context) {
	var dto createTeamDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e := &domain.Team{
		Name:   dto.Name,
		Budget: dto.Budget,
		ClubID: dto.ClubID,
	}

	id, err := h.svc.Create(c.Request.Context(), e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// getByID godoc
// @Summary      Получить Team
// @Tags         team
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  domain.Team
// @Failure      404  {object}  domain.ErrorResponse
// @Router       /teams/{id} [get]
func (h *TeamHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	res, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// list godoc
// @Summary      Список Team
// @Tags         team
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200  {array}   domain.Team
// @Router       /teams [get]
func (h *TeamHandler) list(c *gin.Context) {
	limit := 100
	offset := 0
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil {
			limit = v
		}
	}
	if o := c.Query("offset"); o != "" {
		if v, err := strconv.Atoi(o); err == nil {
			offset = v
		}
	}
	res, err := h.svc.List(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// update godoc
// @Summary      Обновить Team
// @Tags         team
// @Param        id   path      int            true  "ID"
// @Param        input body      createTeamDTO  true  "Данные"
// @Success      200  {object}  domain.Team
// @Router       /teams/{id} [put]
func (h *TeamHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var dto createTeamDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	// Update fields
	e.Name = dto.Name
	e.Budget = dto.Budget
	e.ClubID = dto.ClubID

	if err := h.svc.Update(c.Request.Context(), e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

// delete godoc
// @Summary     Удалить Team
// @Tags       team
// @Param      id path int true "ID"
// @Success     204 "Deleted"
// @Router     /teams/{id} [delete]
func (h *TeamHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
