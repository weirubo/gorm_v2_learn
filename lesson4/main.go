package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 新增、删除、修改、查询
var (
	sqlDB  *sql.DB
	gormDB *gorm.DB
)

type Student struct {
	gorm.Model
	Name  string
	Age   uint
	Email string
}

func InitDB() {
	// driverName
	driverName := "mysql"
	// DSN
	dbUser := "root"
	dbPassword := "root"
	protocol := "tcp"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "blog"
	parseTime := true
	loc := "Local"
	charset := "utf8mb4"

	dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s", dbUser, dbPassword, protocol, dbHost, dbPort, dbName, charset, parseTime, loc)

	// 数据库连接
	if sqlDB == nil {
		sqlDB, _ = sql.Open(driverName, dataSourceName)
	}

	err := sqlDB.Ping()
	if err != nil {
		fmt.Printf("sqlDB.Ping() err:%s\n", err)
		return
	}
	// gorm 是用的 sql 包的连接池
	sqlDB.SetMaxOpenConns(10)           // 设置连接池最大打开连接数
	sqlDB.SetMaxIdleConns(5)            // 设置连接池最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间

	// 使用现有数据库连接初始化 *gorm.DB
	// gorm 配置
	gormDB, err = gorm.Open(
		mysql.New(
			mysql.Config{
				Conn: sqlDB,
			},
		),
		&gorm.Config{
			SkipDefaultTransaction: true, // 关闭写入操作默认启用事务
			DisableAutomaticPing:   true, // 关闭自动 Ping 数据库
		},
	)
	if err != nil {
		fmt.Printf("gorm.Open() err:%s\n", err)
		return
	}
	// gorm 还支持使用 mysql 驱动的高级配置和使用自定义驱动。
}

func main() {
	InitDB()

	// 新增
	// 普通创建
	/*stu := Student{
		Name:  "rose",
		Age:   28,
		Email: "rose@88.com",
	}*/
	/*tx := gormDB.Create(&stu)
	insertID := stu.ID
	fmt.Printf("主键 ID：%d\n", insertID)

	insertErr := tx.Error
	fmt.Printf("插入错误：%v\n", insertErr)
	num := tx.RowsAffected
	fmt.Printf("影响行数:%d\n", num)*/

	// 选定字段创建
	// gormDB.Select("Name", "Age").Create(&stu)

	// 排除字段创建
	// gormDB.Omit("Age", "Email").Create(&stu)

	// 批量创建
	/*stus := []Student{
		{
			Name: "coco",
			Age: 19,
			Email: "coco@88.com",
		},
		{
			Name: "bear",
			Age: 12,
			Email: "bear@88.com",
		},
	}
	gormDB.Create(&stus)
	for _, stu := range stus {
		fmt.Printf("ID:%d\n", stu.ID)
	}*/

	// 根据 map 创建，需要注意的是不会自动填充 gorm.Model 结构体定义的字段
	// map 单条
	/*stuMap := map[string]interface{}{
		"Name": "name1",
		"Age": 22,
		"Email": "name1@8.com",
	}*/
	// map 切片
	/*stusMap := []map[string]interface{}{
		{
			"Name": "apple",
			"Age": 20,
			"Email": "apple@88.com",
		},
		{
			"Name": "pear",
			"Age": 21,
			"Email": "pear@88.com",
		},
	}*/
	// gormDB.Model(&Student{}).Create(stuMap)
	// gormDB.Model(&Student{}).Create(stusMap)

	// 删除
	// 删除单条
	// stu := Student{}
	// gormDB.Where("name = ?", "apple").Delete(&stu)

	// 根据主键删除
	// gormDB.Delete(&Student{}, 2)
	// gormDB.Delete(&Student{}, "2")
	// gormDB.Delete(&Student{},[]int{8,9})

	// 批量删除
	// gormDB.Where("email LIKE ?", "%88%").Delete(Student{})
	// gormDB.Delete(Student{}, "email LIKE ?", "%88%")

	// 阻止全部删除

	// 软删除
	// 查找被软删除的记录

	// 永久删除
	// gormDB.Unscoped().Where("id = ?", 5).Delete(&Student{})
	// gormDB.Unscoped().Where("id IN ?", []int{8,9}).Delete(&Student{})

	// 修改
	// 保存所有字段
	/*student := Student{}
	gormDB.Find(&student)
	student.Name="cat"
	student.Age=18
	student.Email="cat@88.com"
	student.ID=15
	gormDB.Save(&student)*/
	// 更新单个列
	/*student := Student{}
	student.ID = 15
	gormDB.Model(&student).Update("email", "cat@gmail.com")*/

	// gormDB.Model(&Student{}).Where("age = ?", 19).Update("age", 21)

	/*student := Student{}
	student.ID = 15
	gormDB.Model(&student).Where("email = ?", "cat@gmail.com").Update("name", "bigFace")*/
	// 更新多列
	// struct 只更新非零值的字段
	/*student := Student{}
	student.ID = 1
	gormDB.Model(&student).Updates(Student{Name: "book", Age: 20})*/

	// map
	/*student := Student{}
	student.ID = 15
	gormDB.Model(&student).Updates(map[string]interface{}{"name" : "panda", "age": 30})*/

	// 更新选定字段
	// Select
	/*student := Student{}
	student.ID = 2
	gormDB.Model(&student).Select("name").Updates(map[string]interface{}{"name":"lucy", "email":"lucy@88.com"})*/
	// Omit
	/*student := Student{}
	student.ID = 3
	gormDB.Model(&student).Omit("name").Updates(map[string]interface{}{"name": "dog", "age": 29, "email":"dog@gmail.com"})*/
	// 批量更新
	// struct
	// gormDB.Model(&Student{}).Where("name=?", "milk").Updates(Student{Name: "tom", Age: 18})
	// map
	// gormDB.Model(&Student{}).Where("name = ?", "frank").Updates(map[string]interface{}{"name":"milk", "email": "milk@88.com"})
	// gormDB.Table("students").Where("name = ?", "tom").Updates(map[string]interface{}{"name":"honey", "email" :"honey@88.com"})

	// 阻止全局更新

	// 更新的记录数
	/*result := gormDB.Table("students").Where("name = ?", "honey").Updates(map[string]interface{}{"name":"life", "email" :"life@88.com"})
	rows := result.RowsAffected
	fmt.Printf("更新的记录数：%d\n", rows)
	fmt.Println(result.Error)*/

	// 查询
	// 检索单个对象
	// 根据主键检索
	// 检索对象
	// 条件查询
	// 查询特定字段
	// order排序
	// Limit & Offset
	// Group & Having
	// Distinct
	// Joins
	// Scan
	// 高级查询

	// 原生 SQL

}
