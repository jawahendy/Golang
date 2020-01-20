package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `sql:"type:text;"`
}

var DB *gorm.DB

func main() {

	var err error

	DB, err := gorm.Open("mysql", "root:@/learngin?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	defer DB.Close()

	// migrate schema
	//DB.AutoMigrate(&Article{})

	// Create Db
	// DB.Create(&Article{Title: "judul-kedua", Slug: "judul-kedua", Desc: "ini merupakan judul aja"})
	// DB.Create(&Article{Title: "judul-ketiga", Slug: "judul-ketiga", Desc: "ini merupakan judul aja"})

	r := gin.Default()

	v1 := r.Group("/api/v1/")
	{
		art := v1.Group("/article")
		{
			art.GET("/", getHome)
			art.GET("/:id", getArticle)
			art.POST("/post", newpost)
			art.POST("/", newarticel)
		}
	}

	r.Run()
}

func getHome(c *gin.Context) {

	items := []Article{}
	DB.Find(&items)

	c.JSON(200, gin.H{
		"status": "sukses bro",
		"data":   items,
	})
}

func getArticle(c *gin.Context) {
	id := c.Param("id")

	var item Article

	if DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "sukses bro",
		"data":   item,
	})
}

func newarticel(c *gin.Context) {
	item := Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
	}

	DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "sukses bro",
		"data":   item,
	})
}

func newpost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}
