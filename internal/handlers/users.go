package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserReq struct {
	Name  string `json:"name" binding:"required,min=1"`
	Email string `json:"email" binding:"required,email"`
}

func ListUsers(c *gin.Context) {
	// クエリ: /users?page=2
	page := c.DefaultQuery("page", "1")
	c.JSON(http.StatusOK, gin.H{"page": page, "data": []any{}})
}

func CreateUser(c *gin.Context) {
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ここでDB作成など
	c.JSON(http.StatusCreated, gin.H{"id": 1, "name": req.Name, "email": req.Email})
}

func GetUser(c *gin.Context) {
	id := c.Param("id") // パスパラメータ
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func UpdateUser(c *gin.Context) { c.Status(http.StatusNoContent) }
func DeleteUser(c *gin.Context) { c.Status(http.StatusNoContent) }
