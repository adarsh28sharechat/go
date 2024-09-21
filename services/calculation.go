package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdditionInput struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type Output struct {
	Result int `json:"result"`
}

func Addition(c *gin.Context) {
	var input AdditionInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := input.Num1 + input.Num2

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func Subtraction(c *gin.Context) {
	var input AdditionInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := input.Num1 - input.Num2

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func Multiplication(c *gin.Context) {
	var input AdditionInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := input.Num1 * input.Num2

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func Division(c *gin.Context) {
	var input AdditionInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Num1 == 0 || input.Num2 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "num1 or num2 must be positive"})
		return
	}
	result := input.Num1 / input.Num2

	c.JSON(http.StatusOK, gin.H{"result": result})
}
