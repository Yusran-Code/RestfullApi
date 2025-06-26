package main

import (
	"log"
	"strconv"

	"RestfullApi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=LuckyCode dbname=data sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection DB failed", err)
	}

	DB = db
	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("migration failed")
	}
	log.Println("migration succes")

	//router
	r := gin.Default()
	r.GET("/product", getProducts)
	r.GET("/product/:id", getProduct)
	r.POST("/product", createProduct)
	r.PUT("/product/:id", updateProduct)
	r.DELETE("/product/:id", deleteProduct)

	r.Run(":8080")
}

func getProducts(c *gin.Context) {
	var products []models.Product
	if err := DB.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed get products "})
		return
	}

	c.JSON(200, products)

}

func getProduct(c *gin.Context) {
	var product models.Product
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(404, gin.H{"error": "Get product Not Found"})
		return
	}

	if err := DB.First(&product, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed Get Id Product"})
		return
	}
	c.JSON(200, product)
}

func createProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "failed to save"})
		return
	}

	if err := DB.Create(&product); err != nil {
		c.JSON(500, gin.H{"error": "failed to Create product"})
		return
	}

	c.JSON(201, product)
}

func updateProduct(c *gin.Context) {
	var product models.Product
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "failed update id"})
		return
	}

	if err := DB.First(&product, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "unknown id update "})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(404, gin.H{"error": "not found input update"})
		return
	}

	if err := DB.Model(&product).Updates(&input).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed update product"})
		return
	}

	c.JSON(200, product)
}

func deleteProduct(c *gin.Context) {
	var product models.Product
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found id delete"})
		return
	}

	if err := DB.First(&product, id).Error; err != nil {
		c.JSON(401, gin.H{"error": "asasasas"})
		return
	}

	if err := DB.Delete(&product, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed delete product"})
		return
	}
	c.JSON(200, product)

}
