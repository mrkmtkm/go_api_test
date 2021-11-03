package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// データベース
	Dialect = "mysql"

	// ユーザー名
	DBUser = "test"

	// パスワード
	DBPass = "test"

	// プロトコル
	DBProtocol = "tcp(127.0.0.1:3306)"

	// DB名
	DBName = "testdb"
)

func connectGorm() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s?parseTime=true"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect)

	if err != nil {
		log.Println(err.Error())
	}
	return db
}

func main() {
	db := connectGorm()
	defer db.Close()
	db.AutoMigrate(&User{})

	//ユーザーの作成
	//db.Create(&User{Name: "test", Email: "test@test.com"})

	//ユーザーの宣言
	//var users []User
	var user User

	//全ての取得
	//db.Find(&users)

	//条件に合ったものの取得
	//db.Where("ID = ?", "2").Find(&user)

	//更新（１つのみ）
	/*
		db.First(&user)
		user.Name = "uptesttt"
		db.Save(&user)
	*/

	//複数の条件付き更新
	//db.Model(&User{}).Where("ID = ?", 1).Update("Name", "test")
	//db.Model(&user).Update("name", "hello")

	//削除
	//db.Where("ID = ?", 1).Delete(&user)

	//fmt.Println(user)

	//API関連
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		db.Where("ID = ?", "2").Find(&user)
		c.JSON(http.StatusOK, gin.H{
			"ID":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		})
	})

	engine.POST("/post", func(c *gin.Context) {
		name := c.Query("name")
		email := c.Query("email")
		db.Create(&User{Name: name, Email: email})
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	engine.PUT("/put", func(c *gin.Context) {
		postName := c.Query("name")
		id := c.Query("id")
		db.Model(&User{}).Where("ID = ?", id).Update("Name", postName)
		c.JSON(200, gin.H{
			"message": "更新メソッドです",
		})
	})

	engine.DELETE("/delete", func(c *gin.Context) {
		id := c.Query("id")
		db.Where("ID = ?", id).Delete(&user)
		c.JSON(200, gin.H{
			"message": "削除メソッドです",
		})
	})
	engine.Run(":8080")

}

type User struct {
	gorm.Model //ID,CreatedAt,UpdatedAt,DeletedAtが入ってる
	Name       string
	Email      string
}
