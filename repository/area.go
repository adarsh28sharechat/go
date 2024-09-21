package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ObjectRepo struct {
	DB *sql.DB
}

type Object struct {
	ObjectType string `json:"objectType"`
	Side       int    `json:"side"`
}

// Insert - Insert objects to db
func AddObject(c *gin.Context) {
	var inputVar Object

	if err := c.ShouldBindJSON(&inputVar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	side := inputVar.Side

	obj := Object{
		ObjectType: inputVar.ObjectType,
		Side:       side,
	}

	log.Println("input--->>", obj)

	query := "INSERT INTO object (object_type, side) VALUES (?, ?)"
	result, err := c.MustGet("db").(*sql.DB).Exec(query, obj.ObjectType, obj.Side)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Update - Update object in db
func UpdateObject(c *gin.Context) {
	var inputVar Object

	if err := c.ShouldBindJSON(&inputVar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is missing"})
		return
	}
	side := inputVar.Side

	obj := Object{
		ObjectType: inputVar.ObjectType,
		Side:       side,
	}

	query := "UPDATE object SET object_type = ?, side = ? WHERE object_id = ?"
	_, err := c.MustGet("db").(*sql.DB).Exec(query, obj.ObjectType, obj.Side, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Object not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Object updated successfully"})
}

func RemoveObject(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is missing"})
		return
	}

	query := "DELETE FROM object  WHERE object_id = ?"
	_, err := c.MustGet("db").(*sql.DB).Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Object deleted successfully"})
}

func CalculateArea(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is missing"})
		return
	}

	query := "SELECT * FROM object WHERE object_id = ?"
	row := c.MustGet("db").(*sql.DB).QueryRow(query, id)

	var objID int
	var objType string
	var objSide int

	err := row.Scan(&objID, &objType, &objSide)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Object not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	const pi = 3.14
	log.Println("Object retrieved successfully:", objID)

	var result float64

	switch objType {
	case "square":
		result = float64(objSide * objSide)
	default:
		result = pi * float64(objSide*objSide)

	}

	c.JSON(http.StatusOK, gin.H{"object_area": result})
}
