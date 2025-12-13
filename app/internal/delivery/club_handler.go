package delivery

import (
	"github.com/gin-gonic/gin"
	"hse-football/internal/domain"
	"net/http"
	"strconv"
	"time"
)

type ClubHandler struct {
	svc domain.ClubUsecase
}

func NewClubHandler(s domain.ClubUsecase) *ClubHandler {
	return &ClubHandler{svc: s}
}

type createClubDTO struct {
	Name    string  `json:"club_name" binding:"required"`
	Website *string `json:"website,omitempty"`
}

func (h *ClubHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/clubs")
	r.POST("", h.create)
	r.GET("", h.list)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

// create godoc
// @Summary      Создать новый клуб
// @Description  Добавляет клуб в базу
// @Tags         club
// @Accept       json
// @Produce      json
// @Param        club  body      domain.Club  true  "Данные клуба"
// @Success      201   {object}  domain.Club
// @Failure      400   {object}  domain.ErrorResponse
// @Router       /clubs [post]
func (h *ClubHandler) create(c *gin.Context) {
	var dto createClubDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	club := &domain.Club{
		Name:         dto.Name,
		CreationDate: time.Now(),
		Website:      dto.Website,
	}
	id, err := h.svc.Create(c.Request.Context(), club)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// getByID godoc
// @Summary      Получить клуб по ID
// @Description  Возвращает данные клуба по его идентификатору
// @Tags         club
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID клуба"
// @Success      200  {object}  domain.Club
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Router       /clubs/{id} [get]
func (h *ClubHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	club, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, club)
}

// list godoc
// @Summary      Получить список клубов
// @Description  Возвращает все клубы из базы данных
// @Tags         club
// @Accept       json
// @Produce      json
// @Success      200  {array}   domain.Club
// @Router       /clubs [get]
func (h *ClubHandler) list(c *gin.Context) {
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
	clubs, err := h.svc.List(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clubs)
}

// update godoc
// @Summary      Обновить клуб
// @Description  Обновляет клуб по ID
// @Tags         club
// @Accept       json
// @Produce      json
// @Param        id   path      int            true  "ID клуба"
// @Param        club body      createClubDTO  true  "Данные клуба для обновления"
// @Success      200  {object}  domain.Club
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /clubs/{id} [put]
func (h *ClubHandler) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var dto createClubDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	club, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "club not found"})
		return
	}

	club.Name = dto.Name
	club.Website = dto.Website

	if err := h.svc.Update(c.Request.Context(), club); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, club)
}

// delete godoc
// @Summary 	Удалить клуб
// @Description Удаляет клуб по идентификатору
// @Tags 		club
// @Param 		id path int true "ID клуба"
// @Success 	204 "Успешно удалено"
// @Failure 	404 {object} domain.ErrorResponse
// @Failure 	500 {object} domain.ErrorResponse
// @Router 		/clubs/{id} [delete]
func (h *ClubHandler) delete(c *gin.Context) {
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
