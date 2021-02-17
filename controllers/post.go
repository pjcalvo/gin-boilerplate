package controllers

import (
	"strconv"

	"github.com/pjcalvo/gin-boilerplate/forms"
	"github.com/pjcalvo/gin-boilerplate/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

//PostController ...
type PostController struct{}

var postModel = new(models.PostModel)

//Create ...
func (ctrl PostController) Create(c *gin.Context) {
	if userID := getUserID(c); userID != 0 {

		var postForm forms.PostForm

		if c.ShouldBindJSON(&postForm) != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			return
		}

		postID, err := postModel.Create(userID, postForm)

		if postID == 0 && err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Post could not be created", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post created", "id": postID})
	}
}

//All ...
func (ctrl PostController) All(c *gin.Context) {
	if userID := getUserID(c); userID != 0 {

		results, err := postModel.All(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get posts", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"results": results})
	}
}

//One ...
func (ctrl PostController) One(c *gin.Context) {
	if userID := getUserID(c); userID != 0 {

		id := c.Param("id")
		if id, err := strconv.ParseInt(id, 10, 64); err == nil {

			data, err := postModel.One(userID, id)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Post not found", "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": data})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		}
	}
}

//Update ...
func (ctrl PostController) Update(c *gin.Context) {
	if userID := getUserID(c); userID != 0 {

		id := c.Param("id")
		if id, err := strconv.ParseInt(id, 10, 64); err == nil {

			var postForm forms.PostForm

			if c.ShouldBindJSON(&postForm) != nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
				return
			}

			err := postModel.Update(userID, id, postForm)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Post could not be updated", "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Post updated"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid parameter", "error": err.Error()})
		}
	}
}

//Delete ...
func (ctrl PostController) Delete(c *gin.Context) {
	if userID := getUserID(c); userID != 0 {

		id := c.Param("id")
		if id, err := strconv.ParseInt(id, 10, 64); err == nil {

			err := postModel.Delete(userID, id)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Post could not be deleted", "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		}
	}
}
