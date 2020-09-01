package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 模型定义、约定、标签、自动迁移和迁移接口

/**
约定：默认情况下，GORM 约定使用 ID 作为主键，使用结构体名的复数作为表名，字段名作为列名，使用 CreatedAt、UpdatedAt、DeletedAt时间追踪。
当然，你可以按照自己的需求自定义约定项。比如时间追踪默认是以当前时间填充为零值得 CreatedAt 字段，以当前时间戳秒数填充 UpdatedAt 字段，如果你喜欢使用时间戳记录创建时间，你可以改变 CreatedAt 字段的类型为 int，默认时间戳是 Unix 秒，你还可以使用标签将时间戳的单位改为纳秒或毫秒。

GORM 定义了一个 gorm.Model 结构体，字段包括 ID、CreatedAt、UpdatedAt、DeletedAt，我们可以将它嵌入到我们自定义的结构体中。

在 GO 语言中，根据名称的首字母大小写来定义是否可被导出，GORM 使用可导出的字段进行 CRUD 时拥有全部权限，另外，GORM 可以使用标签控制字段的权限，可以让一个字段的权限是只读、只写、只创建、只更新和忽略该字段。

标签是模型定义时的可选项，GORM 的标签不区分大小写，推荐使用驼峰式命名。GORM 通过标签实现外键定义、表关联和约束。

迁移

Migrator 接口方法
*/

// 模型定义
type Student struct {
	ID        uint
	Name      string
	Age       uint
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

var (
	sqlDB  *sql.DB
	gormDB *gorm.DB
)

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
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic() err:%s\n", err)
			return
		}
	}()
	// 初始化数据库
	InitDB()

	// 返回当前操作的数据库名
	currentDBName := gormDB.Migrator().CurrentDatabase()
	fmt.Printf("当前操作的数据库名:%s\n", currentDBName)

	// 数据库表操作
	// 判断数据库表是否已存在
	isExist := gormDB.Migrator().HasTable(&Student{})
	// isExist := gormDB.Migrator().HasTable("students")
	fmt.Printf("数据库表是否存在:%t\n", isExist)
	// 创建数据库表
	if isExist == false { // 数据库表不存在
		err := gormDB.Migrator().CreateTable(&Student{})
		if err != nil {
			fmt.Printf("创建数据库表失败，错误:%s\n", err)
			return
		}
		fmt.Println("创建数据库表成功")
	} else { // 数据库表已存在
		// 重命名数据库表
		newName := "stu_" + time.Now().Format("2006-01-02 15:04:05")
		gormDB.Migrator().RenameTable("students", newName)
		// gormDB.Migrator().RenameTable("students", "stu")
		// gormDB.Migrator().RenameTable(&Student{}, &Stu{})

		// 删除数据库表
		gormDB.Migrator().DropTable("students")
		// gormDB.Migrator().DropTable(&Student{})
	}

	// 数据库表字段操作

	// 数据库表索引操作

	// 数据库表约束操作

}
