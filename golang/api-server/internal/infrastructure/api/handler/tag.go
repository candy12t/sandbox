package handler

import (
	"net/http"
	"strconv"

	"github.com/candy12t/api-server/internal/usecase/form"
	"github.com/candy12t/api-server/internal/usecase/usecase"
	"github.com/gin-gonic/gin"
)

type Tag struct {
	tagUsecase usecase.Tag
}

func NewTag(tagUsecase usecase.Tag) *Tag {
	return &Tag{
		tagUsecase: tagUsecase,
	}
}

func (t *Tag) CreateTag(c *gin.Context) {
	f := new(form.CreateTagParams)
	if err := c.ShouldBindJSON(f); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tag, err := t.tagUsecase.CreateTag(f)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create tag is success!!", "result": tag})
}

func (t *Tag) GetTag(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tag, err := t.tagUsecase.GetTag(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tag)
}

func (t *Tag) GetTags(c *gin.Context) {
	tags, err := t.tagUsecase.GetTags()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tags)
}

func (t *Tag) UpdateTag(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	f := &form.UpdateTagParams{ID: id}
	if err := c.ShouldBindJSON(f); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tag, err := t.tagUsecase.UpdateTag(f)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update tag is success!!", "result": tag})
}

func (t *Tag) DeleteTag(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := t.tagUsecase.DeleteTag(id); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete tag is success!!"})
}
