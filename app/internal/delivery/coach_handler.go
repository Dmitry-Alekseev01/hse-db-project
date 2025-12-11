package delivery

import (
	"github.com/gin-gonic/gin"
	"hse-football/internal/domain"
	"net/http"
	"strconv"
)

type CoachHandler struct {
	svc domain.CoachUsecase
}

func NewCoachHandler(s domain.CoachUsecase) *CoachHandler {
	return &CoachHandler{svc: s}
}

type createCoachDTO struct {
	Name    string  `json:"coach_name" binding:"required"`
	Surname string  `json:"coach_surname" binding:"required"`
	Salary  float64 `json:"salary" binding:"required"`
	Phone   string  `json:"phone" binding:"required"`
	TeamID  int64   `json:"team_id" binding:"required"`
}

func (h *CoachHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/coachs")
	r.POST("", h.create)
	r.GET("", h.list)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

// create godoc
// @Summary      Создать Coach
// @Tags         coach
// @Accept       json
// @Produce      json
// @Param        input  body      createCoachDTO  true  "Данные"
// @Success      201   {object}  domain.Coach
// @Failure      400   {object}  domain.ErrorResponse
// @Router       /coachs [post]
func (h *CoachHandler) create(c *gin.Context) {
	var dto createCoachDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e := &domain.Coach{
		Name:    dto.Name,
		Surname: dto.Surname,
		Salary:  dto.Salary,
		Phone:   dto.Phone,
		TeamID:  dto.TeamID,
	}

	id, err := h.svc.Create(c.Request.Context(), e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// getByID godoc
// @Summary      Получить Coach
// @Tags         coach
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  domain.Coach
// @Failure      404  {object}  domain.ErrorResponse
// @Router       /coachs/{id} [get]
func (h *CoachHandler) getByID(c *gin.Context) {
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
// @Summary      Список Coach
// @Tags         coach
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200  {array}   domain.Coach
// @Router       /coachs [get]
func (h *CoachHandler) list(c *gin.Context) {
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
// @Summary      Обновить Coach
// @Tags         coach
// @Param        id   path      int            true  "ID"
// @Param        input body      createCoachDTO  true  "Данные"
// @Success      200  {object}  domain.Coach
// @Router       /coachs/{id} [put]
func (h *CoachHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var dto createCoachDTO
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
	e.Salary = dto.Salary
	e.Phone = dto.Phone
	e.TeamID = dto.TeamID

	if err := h.svc.Update(c.Request.Context(), e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

// delete godoc
// @Summary     Удалить Coach
// @Tags       coach
// @Param      id path int true "ID"
// @Success     204 "Deleted"
// @Router     /coachs/{id} [delete]
func (h *CoachHandler) delete(c *gin.Context) {
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
