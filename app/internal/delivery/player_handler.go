package delivery

import (
	"github.com/gin-gonic/gin"
	"hse-football/internal/domain"
	"net/http"
	"strconv"
	"time"
)

type PlayerHandler struct {
	svc domain.PlayerUsecase
}

func NewPlayerHandler(s domain.PlayerUsecase) *PlayerHandler {
	return &PlayerHandler{svc: s}
}

type createPlayerDTO struct {
	Name      string    `json:"player_name" binding:"required"`
	Surname   string    `json:"player_surname" binding:"required"`
	Number    int64     `json:"player_number" binding:"required"`
	Salary    float64   `json:"salary" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
	TeamID    int64     `json:"team_id" binding:"required"`
	StatusID  int64     `json:"status_id" binding:"required"`
}

func (h *PlayerHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/players")
	r.POST("", h.create)
	r.GET("", h.list)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

// create godoc
// @Summary      Создать Player
// @Tags         player
// @Accept       json
// @Produce      json
// @Param        input  body      createPlayerDTO  true  "Данные"
// @Success      201   {object}  domain.Player
// @Failure      400   {object}  domain.ErrorResponse
// @Router       /players [post]
func (h *PlayerHandler) create(c *gin.Context) {
	var dto createPlayerDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e := &domain.Player{
		Name:      dto.Name,
		Surname:   dto.Surname,
		Number:    dto.Number,
		Salary:    dto.Salary,
		Phone:     dto.Phone,
		BirthDate: dto.BirthDate,
		TeamID:    dto.TeamID,
		StatusID:  dto.StatusID,
	}

	id, err := h.svc.Create(c.Request.Context(), e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// getByID godoc
// @Summary      Получить Player
// @Tags         player
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  domain.Player
// @Failure      404  {object}  domain.ErrorResponse
// @Router       /players/{id} [get]
func (h *PlayerHandler) getByID(c *gin.Context) {
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
// @Summary      Список Player
// @Tags         player
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200  {array}   domain.Player
// @Router       /players [get]
func (h *PlayerHandler) list(c *gin.Context) {
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
// @Summary      Обновить Player
// @Tags         player
// @Param        id   path      int            true  "ID"
// @Param        input body      createPlayerDTO  true  "Данные"
// @Success      200  {object}  domain.Player
// @Router       /players/{id} [put]
func (h *PlayerHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var dto createPlayerDTO
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
	e.Surname = dto.Surname
	e.Number = dto.Number
	e.Salary = dto.Salary
	e.Phone = dto.Phone
	e.BirthDate = dto.BirthDate
	e.TeamID = dto.TeamID
	e.StatusID = dto.StatusID

	if err := h.svc.Update(c.Request.Context(), e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

// delete godoc
// @Summary     Удалить Player
// @Tags       player
// @Param      id path int true "ID"
// @Success     204 "Deleted"
// @Router     /players/{id} [delete]
func (h *PlayerHandler) delete(c *gin.Context) {
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
