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



func main()  {
	dsn := "host=localhost user=postgres password=LuckyCode dbname=data sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal("Connection DB failed",err)
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
	r.POST("/product",createProduct)
	r.PUT("/product/:id",updateProduct)
	r.DELETE("/product/:id",deleteProduct)

	r.Run(":8080")
}

func getProducts(c *gin.Context)  {
	var products []models.Product
	if err:= DB.Find(&products).Error; err != nil {
		c.JSON(500,gin.H{"error": "products not found"})
		return
	}

	c.JSON(200, products)
	

}

func getProduct(c *gin.Context)  {
	var product models.Product
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "shdhsgdhsgdsd"})
		return
	}
	
	if err := DB.First(&product,id).Error; err != nil {
		c.JSON(404,gin.H{"error": "nosdsdsdsdsd"})
		return
	}
  c.JSON(200,product)
}

func createProduct(c *gin.Context)  {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error":"asasasa"})
		return 
	}

	if err := DB.Create(&product); err != nil {
		c.JSON(400,gin.H{"error": "sdsdsdsddsd"})
		return 
	}

	c.JSON(200,product)
}

func updateProduct(c *gin.Context)  {
	var product models.Product
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "sdsdsdssdsdsd"})
		return 
	}

	if err := DB.First(&product,id).Error; err != nil {
		c.JSON(400, gin.H{"error": "erasaresaresraes"})
		return 
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(300,gin.H{"error": "sdsdsdsdsds"})
		return
	}

	if err := DB.Model(&product).Updates(&input).Error; err != nil {
		c.JSON(400,gin.H{"error": "sdsdssdd"})
		return 
	}

	c.JSON(200, product)
}

func deleteProduct(c *gin.Context)  {
	var product models.Product
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "sdsddsd"})
		return 
	}

	if  err:= DB.First(&product,id).Error; err != nil {
		c.JSON(401, gin.H{"error": "asasasas"})
		return 
	}

	if err := DB.Delete(&product,id).Error; err != nil {
		c.JSON(400, gin.H{"error": "sdsdsdsdsd"})
		return 
	}
	c.JSON(200, product)


}