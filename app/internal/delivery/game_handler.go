package delivery

import (
	"github.com/gin-gonic/gin"
	"hse-football/internal/domain"
	"net/http"
	"strconv"
	"time"
)

type GameHandler struct {
	svc domain.GameUsecase
}

func NewGameHandler(s domain.GameUsecase) *GameHandler {
	return &GameHandler{svc: s}
}

type createGameDTO struct {
	StadiumID int64     `json:"stadium_id" binding:"required"`
	Team1ID   int64     `json:"team_1_id" binding:"required"`
	Team2ID   int64     `json:"team_2_id" binding:"required"`
	MatchDate time.Time `json:"match_date" binding:"required"`
}

func (h *GameHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/games")
	r.POST("", h.create)
	r.GET("", h.list)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

// create godoc
// @Summary      Создать Game
// @Tags         game
// @Accept       json
// @Produce      json
// @Param        input  body      createGameDTO  true  "Данные"
// @Success      201   {object}  domain.Game
// @Failure      400   {object}  domain.ErrorResponse
// @Router       /games [post]
func (h *GameHandler) create(c *gin.Context) {
	var dto createGameDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e := &domain.Game{
		StadiumID: dto.StadiumID,
		Team1ID:   dto.Team1ID,
		Team2ID:   dto.Team2ID,
		MatchDate: dto.MatchDate,
	}

	id, err := h.svc.Create(c.Request.Context(), e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// getByID godoc
// @Summary      Получить Game
// @Tags         game
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  domain.Game
// @Failure      404  {object}  domain.ErrorResponse
// @Router       /games/{id} [get]
func (h *GameHandler) getByID(c *gin.Context) {
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
// @Summary      Список Game
// @Tags         game
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200  {array}   domain.Game
// @Router       /games [get]
func (h *GameHandler) list(c *gin.Context) {
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
// @Summary      Обновить Game
// @Tags         game
// @Param        id   path      int            true  "ID"
// @Param        input body      createGameDTO  true  "Данные"
// @Success      200  {object}  domain.Game
// @Router       /games/{id} [put]
func (h *GameHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var dto createGameDTO
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
	e.StadiumID = dto.StadiumID
	e.Team1ID = dto.Team1ID
	e.Team2ID = dto.Team2ID
	e.MatchDate = dto.MatchDate

	if err := h.svc.Update(c.Request.Context(), e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

// delete godoc
// @Summary     Удалить Game
// @Tags       game
// @Param      id path int true "ID"
// @Success     204 "Deleted"
// @Router     /games/{id} [delete]
func (h *GameHandler) delete(c *gin.Context) {
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
