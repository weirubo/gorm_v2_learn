package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 查询
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
	// 查询
	// 检索单个对象
	// 主键升序 LIMIT 1
	/*student := Student{}
	gormDB.First(&student)
	fmt.Println(student)*/

	// 没有指定排序字段 LIMIT 1
	/*student := Student{}
	gormDB.Take(&student)
	fmt.Println(student)*/

	// 主键降序 LIMIT 1
	/*student := Student{}
	result := gormDB.Last(&student)
	fmt.Println(student)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))*/

	// 根据主键检索
	/*student := Student{}
	// gormDB.First(&student, 16)
	gormDB.First(&student, "16")
	fmt.Println(student)*/

	/*var students  []Student
	gormDB.Find(&students, []int{15, 16, 17})
	fmt.Println(students)*/
	// 检索对象
	/*var students  []Student
	result := gormDB.Find(&students)
	fmt.Println(students)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)*/
	// 条件查询
	// String
	/*student := Student{}
	gormDB.Where("name = ?", "bear").First(&student)
	fmt.Println(student)*/

	/*var students []Student
	gormDB.Where("name <> ?", "bear").Find(&students)
	fmt.Println(students)*/

	/*var students []Student
	gormDB.Where("name IN ?", []string{"bear", "panda"}).Find(&students)
	fmt.Println(students)*/

	/*var students []Student
	gormDB.Where("name LIKE ?", "%a%").Find(&students)
	fmt.Println(students)*/

	/*var students []Student
	gormDB.Where("name = ? AND age > ?", "coco", 18).Find(&students)
	fmt.Println(students)*/

	/*var students []Student
	lastWeek := time.Now().Add(time.Hour*24*-7)
	// fmt.Println(lastWeek)
	gormDB.Where("updated_at > ?", lastWeek).Find(&students)
	fmt.Println(students)*/

	/*var students []Student
	lastWeek := time.Now().Add(time.Hour*24*-7)
	today := time.Now()
	gormDB.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&students)
	fmt.Println(students)*/
	// Struct
	/*var student Student
	gormDB.Where(&Student{Name: "coco", Age: 19}).Find(&student)
	fmt.Println(student)*/
	// Map
	/*var student Student
	gormDB.Where(map[string]interface{}{"name": "coco", "age": 19}).Find(&student)
	fmt.Println(student)*/

	// 主键切片条件
	/*var students []Student
	gormDB.Where([]int{15,16,17}).Find(&students)
	fmt.Println(students)*/

	// 内联条件
	// 根据非整型主键获取记录
	/*student := Student{}
	gormDB.Find(&student, "id = ?", 17)
	fmt.Println(student)*/

	// 普通条件查询
	/*var student Student
	gormDB.Find(&student, "name = ?", "cat")
	fmt.Println(student)*/

	// 多条件查询
	/*student := Student{}
	gormDB.Find(&student, "name <> ? AND age >= ?", "cat", 0)
	fmt.Println(student)*/

	// struct 条件查询
	/*var student Student
	gormDB.Find(&student, &Student{Name: "cat2", Age: 0})
	fmt.Println(student)*/

	// map 条件查询
	/*var student Student
	gormDB.Find(&student, map[string]interface{}{"name":"cat2", "age":0})
	fmt.Println(student)*/

	// Not 条件
	// 普通 Not 条件
	/*var student Student
	gormDB.Not("name = ?", "cat").First(&student)
	fmt.Println(student)*/

	// Not In
	/*var students []Student
	gormDB.Not(map[string]interface{}{"id":[]int{15,17}}).Find(&students)
	fmt.Println(students)*/

	// Struct
	/*var student Student
	gormDB.Not(Student{Name: "cat", Age: 18}).Find(&student)
	fmt.Println(student)*/

	// 主键
	/*var students []Student
	gormDB.Not([]int{1,2,3,15,16}).Find(&students)
	fmt.Println(students)*/

	// Or
	/*var students []Student
	gormDB.Where("name = ?", "cat").Or("name = ?", "cat2").Find(&students)
	fmt.Println(students)*/

	// struct
	/*var students []Student
	gormDB.Where("name = ?", "cat2").Or(Student{Name: "cat", Age: 19}).Find(&students)
	fmt.Println(students)*/

	// map
	/*var students []Student
	gormDB.Where("name = ?", "cat2").Or(map[string]interface{}{"name":"cat", "age":19}).Find(&students)
	fmt.Println(students)*/

	// 查询特定字段
	/*var students []Student
	gormDB.Select("name", "age").Find(&students)
	fmt.Println(students)*/

	// slice
	/*var students []Student
	gormDB.Select([]string{"name", "age"}).Find(&students)
	fmt.Println(students)*/
	// order排序
	// Limit & Offset
	// Group & Having
	// Distinct
	// Joins
	// Scan
	// 高级查询

	// 原生 SQL

}
