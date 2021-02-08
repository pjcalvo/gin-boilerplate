package controllers

import (
	"strconv"

	"github.com/pjcalvo/gin-boilerplate/forms"
	"github.com/pjcalvo/gin-boilerplate/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

//CommentController ...
type CommentController struct{}

var commentModel = new(models.CommentModel)

//Create ...
func (ctrl CommentController) Create(c *gin.Context) {

	postID := c.Param("postid")
	if postID, err := strconv.ParseInt(postID, 10, 64); err == nil {
		var commentForm forms.CommentForm

		if c.ShouldBindJSON(&commentForm) != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			return
		}

		commentID, err := commentModel.Create(postID, commentForm)

		if commentID == 0 && err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Comment could not be created", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Comment created", "id": commentID})
	} else {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "PostID is missing"})
	}
}

//All ...
func (ctrl CommentController) All(c *gin.Context) {

	postID := c.Param("postid")
	if postID, err := strconv.ParseInt(postID, 10, 64); err == nil {

		results, err := commentModel.All(postID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get comments", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"results": results})
	}
}

//One ...
func (ctrl CommentController) One(c *gin.Context) {
	postID := c.Param("postid")
	if postID, err := strconv.ParseInt(postID, 10, 64); err == nil {

		id := c.Param("id")
		if id, err := strconv.ParseInt(id, 10, 64); err == nil {

			data, err := commentModel.One(postID, id)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Comment not found", "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": data})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		}
	}
}

//Update ...
func (ctrl CommentController) Update(c *gin.Context) {
	postID := c.Param("postid")
	if postID, err := strconv.ParseInt(postID, 10, 64); err == nil {

		id := c.Param("id")
		if id, err := strconv.ParseInt(id, 10, 64); err == nil {

			var commentForm forms.CommentForm

			if c.ShouldBindJSON(&commentForm) != nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
				return
			}

			err := commentModel.Update(postID, id, commentForm)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Comment could not be updated", "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Comment updated"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid parameter", "error": err.Error()})
		}
	}
}

//Delete ...
func (ctrl CommentController) Delete(c *gin.Context) {
	postID := c.Param("postid")
	if postID, err := strconv.ParseInt(postID, 10, 64); err == nil {

		id := c.Param("id")
		if id, err := strconv.ParseInt(id, 10, 64); err == nil {

			err := commentModel.Delete(postID, id)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Comment could not be deleted", "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		}
	}
}
