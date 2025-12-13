package delivery

import (
	"github.com/gin-gonic/gin"
	"hse-football/internal/domain"
	"net/http"
	"strconv"
)

type StaffHandler struct {
	svc domain.StaffUsecase
}

func NewStaffHandler(s domain.StaffUsecase) *StaffHandler {
	return &StaffHandler{svc: s}
}

type createStaffDTO struct {
	Name            string  `json:"staff_name" binding:"required"`
	Surname         string  `json:"staff_surname" binding:"required"`
	Salary          float64 `json:"salary" binding:"required"`
	SpecificationID int64   `json:"specification_id" binding:"required"`
	ClubID          int64   `json:"club_id" binding:"required"`
}

func (h *StaffHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/staffs")
	r.POST("", h.create)
	r.GET("", h.list)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

// create godoc
// @Summary      Создать Staff
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param        input  body      createStaffDTO  true  "Данные"
// @Success      201   {object}  domain.Staff
// @Failure      400   {object}  domain.ErrorResponse
// @Router       /staffs [post]
func (h *StaffHandler) create(c *gin.Context) {
	var dto createStaffDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e := &domain.Staff{
		Name:            dto.Name,
		Surname:         dto.Surname,
		Salary:          dto.Salary,
		SpecificationID: dto.SpecificationID,
		ClubID:          dto.ClubID,
	}

	id, err := h.svc.Create(c.Request.Context(), e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// getByID godoc
// @Summary      Получить Staff
// @Tags         staff
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  domain.Staff
// @Failure      404  {object}  domain.ErrorResponse
// @Router       /staffs/{id} [get]
func (h *StaffHandler) getByID(c *gin.Context) {
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
// @Summary      Список Staff
// @Tags         staff
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200  {array}   domain.Staff
// @Router       /staffs [get]
func (h *StaffHandler) list(c *gin.Context) {
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
// @Summary      Обновить Staff
// @Tags         staff
// @Param        id   path      int            true  "ID"
// @Param        input body      createStaffDTO  true  "Данные"
// @Success      200  {object}  domain.Staff
// @Router       /staffs/{id} [put]
func (h *StaffHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var dto createStaffDTO
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
	e.SpecificationID = dto.SpecificationID
	e.ClubID = dto.ClubID

	if err := h.svc.Update(c.Request.Context(), e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

// delete godoc
// @Summary     Удалить Staff
// @Tags       staff
// @Param      id path int true "ID"
// @Success     204 "Deleted"
// @Router     /staffs/{id} [delete]
func (h *StaffHandler) delete(c *gin.Context) {
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
