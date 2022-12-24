package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    entity "todo-app/internal/entities"
)

func (h *Handler) IndexList(c *gin.Context) {
    lists, err := h.services.List.GetAll()

    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, lists)
}

func (h *Handler) ShowList(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    list, err := h.services.List.GetById(id)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, list)
}

func (h *Handler) CreateList(c *gin.Context) {
    var input entity.List
    if err := c.BindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    list, err := h.services.List.Create(&input)

    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateList(c *gin.Context) {
    var input entity.List
    if err := c.BindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    list, err := h.services.List.Update(&input)

    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, list)
}

func (h *Handler) DeleteList(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = h.services.List.Delete(id)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.AbortWithStatus(204)
}