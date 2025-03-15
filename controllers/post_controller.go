package controllers

import (
	"context"
	"gellyzxc-template-golang-gin/config"
	"gellyzxc-template-golang-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post

	err := config.DB.NewSelect().Model(&posts).Relation("User").Scan(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.NewInsert().Model(&post).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	var post models.Post

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post id is required"})
		return
	}

	_, err := config.DB.NewUpdate().Model(&post).WherePK().Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	var post models.Post

	id, _ := strconv.ParseInt(c.Param("post_id"), 10, 64)

	_, err := config.DB.NewDelete().Model(&post).Where("id = ?", id).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{})
}
