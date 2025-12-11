package delivery

import (
	"github.com/gin-gonic/gin"
	"hse-football/internal/domain"
	"net/http"
	"strconv"
	"time"
)

type StadiumHandler struct {
	svc domain.StadiumUsecase
}

func NewStadiumHandler(s domain.StadiumUsecase) *StadiumHandler {
	return &StadiumHandler{svc: s}
}

type createStadiumDTO struct {
	Capacity int64  `json:"capacity" binding:"required"`
	Location string `json:"location" binding:"required"`
}

func (h *StadiumHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/stadiums")
	r.POST("", h.create)
	r.GET("", h.list)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

// create godoc
// @Summary      Создать Stadium
// @Tags         stadium
// @Accept       json
// @Produce      json
// @Param        input  body      createStadiumDTO  true  "Данные"
// @Success      201   {object}  domain.Stadium
// @Failure      400   {object}  domain.ErrorResponse
// @Router       /stadiums [post]
func (h *StadiumHandler) create(c *gin.Context) {
	var dto createStadiumDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e := &domain.Stadium{
		Capacity:  dto.Capacity,
		Location:  dto.Location,
		BuildDate: time.Now(),
	}

	id, err := h.svc.Create(c.Request.Context(), e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// getByID godoc
// @Summary      Получить Stadium
// @Tags         stadium
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  domain.Stadium
// @Failure      404  {object}  domain.ErrorResponse
// @Router       /stadiums/{id} [get]
func (h *StadiumHandler) getByID(c *gin.Context) {
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
// @Summary      Список Stadium
// @Tags         stadium
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200  {array}   domain.Stadium
// @Router       /stadiums [get]
func (h *StadiumHandler) list(c *gin.Context) {
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
// @Summary      Обновить Stadium
// @Tags         stadium
// @Param        id   path      int            true  "ID"
// @Param        input body      createStadiumDTO  true  "Данные"
// @Success      200  {object}  domain.Stadium
// @Router       /stadiums/{id} [put]
func (h *StadiumHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var dto createStadiumDTO
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
	e.Capacity = dto.Capacity
	e.Location = dto.Location
	e.BuildDate = time.Now()

	if err := h.svc.Update(c.Request.Context(), e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

// delete godoc
// @Summary     Удалить Stadium
// @Tags       stadium
// @Param      id path int true "ID"
// @Success     204 "Deleted"
// @Router     /stadiums/{id} [delete]
func (h *StadiumHandler) delete(c *gin.Context) {
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
