package main   


import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type User struct {
	Id  int `gorm:"primary_key"`
	Name string 
	Pwd string
	CreateAt  time.Time  
	UpdateAt  time.Time
}


func main() {
	db,err:=gorm.Open("mysql","fish:77@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic("database connect failed!")
	}

	defer db.Close()

	 // 自动迁移模式
    db.AutoMigrate(&User{})

    //实例化
    user := User{}

    //创建
    db.Create(&User{Id:1,Name:"Mike",Pwd:"123"})

    //查询
    db.First(&user, 1).Scan(&user)
}

