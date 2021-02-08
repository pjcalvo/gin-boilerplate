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

		var articleForm forms.ArticleForm

		if c.ShouldBindJSON(&articleForm) != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
			return
		}

		articleID, err := articleModel.Create(userID, articleForm)

		if articleID == 0 && err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Article created", "id": articleID})
	}
}

//All ...
func (ctrl PostController) All(c *gin.Context) {
	if userID := getUserID(c); userID != 0 {

		results, err := articleModel.All(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get articles", "error": err.Error()})
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

			data, err := articleModel.One(userID, id)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Article not found", "error": err.Error()})
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

			var articleForm forms.ArticleForm

			if c.ShouldBindJSON(&articleForm) != nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
				return
			}

			err := articleModel.Update(userID, id, articleForm)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be updated", "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Article updated"})
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

			err := articleModel.Delete(userID, id)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be deleted", "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		}
	}
}
