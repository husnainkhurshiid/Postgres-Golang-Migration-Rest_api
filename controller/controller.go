package controller

import (
	"fmt"
	"net/http"
	"products/database"
	"products/models"
	"products/shared"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostProduct(c *gin.Context) {
	var items []models.Product

	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn, err := database.Connection()
	shared.CheckError(err)

	Query := `INSERT INTO "Products" (id,title,price,info,rating) VALUES ($1,$2,$3,$4,$5)`

	for _, item := range items {
		item.Id = uuid.NewString()
		_, err = conn.Exec(Query, item.Id, item.Title, item.Price, item.Info, item.Rating)
		shared.CheckError(err)
	}
	c.JSON(http.StatusOK, items)
}

func GetProducts(c *gin.Context) {
	var items []models.Product
	conn, err := database.Connection()
	shared.CheckError(err)

	rows, err := conn.Query(`SELECT * From "Products"`)
	shared.CheckError(err)

	for rows.Next() {
		var item models.Product
		err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.Info, &item.Rating)
		shared.CheckError(err)
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

func GetProductById(c *gin.Context) {
	var items []models.Product
	conn, err := database.Connection()
	shared.CheckError(err)

	id := c.Param("id")

	sqlQuery := fmt.Sprintf(`SELECT * From "Products" WHERE id = '%v'`, id)

	rows, err := conn.Query(sqlQuery)
	shared.CheckError(err)

	for rows.Next() {
		var item models.Product
		err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.Info, &item.Rating)
		shared.CheckError(err)
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

func UpdateProduct(c *gin.Context) {
	var items models.Product

	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn, err := database.Connection()
	shared.CheckError(err)

	id := c.Param("id")

	sqlQuery := fmt.Sprintf(`UPDATE "Products" SET title = '%v', price = '%v', info = '%v', rating = '%v' WHERE id = '%v'`, items.Title, items.Price, items.Info, items.Rating, id)

	_, err = conn.Query(sqlQuery)
	shared.CheckError(err)

	c.JSON(http.StatusOK, gin.H{"Record Updated": items})
}

func DeleteProduct(c *gin.Context) {
	conn, err := database.Connection()
	shared.CheckError(err)

	id := c.Param("id")

	sqlQuery := fmt.Sprintf(`DELETE From "Products" WHERE id = '%v'`, id)

	_, err = conn.Query(sqlQuery)
	shared.CheckError(err)

	c.JSON(http.StatusOK, gin.H{"Record Deleted": id})
}
