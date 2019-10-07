
package main
 
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "time"
	"fmt"
)
 
type Person struct {
	User_id   int `gorm:"primary_key"` //指定主键并自增
	Name      string
	Pwd       string
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// Car       []Car `gorm:"ForeignKey:User_id;AssociationForeignKey:User_id"` //(一对多)指定外键和关联外键
}
 
// type Car struct {
// 	Car_id     int32 `gorm:"primary_key"`
// 	User_id    int
// 	Car_name   string
// 	Car_gongli string
// }
 
func main() {
	//创建连接
	db, err := gorm.Open("mysql", "fish:77@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	// 自动迁移模式
	db.AutoMigrate(&Person{})
 
	/**
	 * 查询单条数据
	 * &是为了获取结构体的内存地址,这样指针传递速度较快
	 */
	p := Person{}

	    //创建
    db.Create(&Person{User_id:24,Name:"Tom",Pwd:"123"})

    fmt.Println("Create successfully!")

    //查询

    // db.First(&p)

    db.First(&p,8)
    fmt.Println(&p)

	// ////SELECT * FROM users WHERE user_id = 10 LIMIT 1;
    
    p2 := Person{}
	db.Where("user_id = ?", 2).First(&p2)
	// ////SELECT * FROM users WHERE user_id = 9 LIMIT 1;;
	fmt.Println(&p2)
     

    // db.Limit(3).Find(&p2)
    if err := db.Where("name = ?", "jinzhu").First(&p).Error; err != nil {
    // 错误处理...

    	fmt.Println("there is no jinzhu here")
    }

	// /**
	//  * 查询多条数据
	//  *
	//  */
	// persons := []Person{}
	// db.Find(&persons)
	// ////SELECT * FROM users
 
	// db.Where("user_id = ? and name = ?", 10, "jinzhu").Find(&persons)
	// //// SELECT * FROM users WHERE name = "jinzhu" AND user_id = 20;
 
	// db.Where("created_at > ?", "2018-09-22 08:50:09").Or("updated_at < ?", "2018-09-22 08:50:09").Order("created_at asc").Find(&persons)
	// //// SELECT * FROM users WHERE created_at = '2018-09-22 08:50:09' OR updated_at < '2018-09-22 08:50:09';
	// fmt.Println(persons)
 
	// /**
	//  * 关联查询
	//  * 一对多 多对一 多对多
	//  */
	// db.Where("user_id = ?", 10).Preload("Car").Find(&p)
 
 
	// fmt.Println(&p)
 
}