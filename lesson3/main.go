package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 自动迁移和迁移接口的方法

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

	// 自动迁移
	err := gormDB.AutoMigrate(&Student{})
	if err != nil {
		fmt.Printf("自动迁移失败，err:%s\n", err)
		return
	}

	// 迁移接口的方法

	// 返回当前操作的数据库名
	// currentDBName := gormDB.Migrator().CurrentDatabase()
	// fmt.Printf("当前操作的数据库名:%s\n", currentDBName)

	// 数据库表操作
	// 判断数据库表是否已存在
	isExist := gormDB.Migrator().HasTable(&Student{})
	// isExist := gormDB.Migrator().HasTable("students")
	// fmt.Printf("数据库表是否存在:%t\n", isExist)
	// 创建数据库表
	if isExist == false { // 数据库表不存在
		// 默认情况下，GORM 会约定使用 ID 作为表的主键，可以通过标签 primaryKey 将其它字段设为主键。通过将多个字段设为主键，以达到创建复合主键，整型字段设为主键，默认为启用 AutoIncrement，如果需要禁用，使用标签autoIncrement:false。
		// GORM 约定使用结构体名的复数形式作为表名，不过也可以根据需求修改，可以实现Tabler 接口来更改默认表名，不过这种方式不支持动态变化，它会被缓存下来以便后续使用，如果想要使用动态表名，可以使用Scopes，关于 Scopes 的使用方法，本文暂不展开。
		// GORM 约定使用结构体的字段名作为数据表的字段名，可以通过标签 column 修改。
		/*err := gormDB.Migrator().CreateTable(&Student{})
		if err != nil {
			fmt.Printf("创建数据库表失败，错误:%s\n", err)
			return
		}
		fmt.Println("创建数据库表成功")*/
	} else { // 数据库表已存在
		// 重命名数据库表
		// newName := "stu_" + time.Now().Format("2006-01-02 15:04:05")
		// gormDB.Migrator().RenameTable("students", newName)
		// gormDB.Migrator().RenameTable("students", "stu")
		// gormDB.Migrator().RenameTable(&Student{}, &Stu{})

		// 删除数据库表
		// gormDB.Migrator().DropTable("students")
		// gormDB.Migrator().DropTable(&Student{})
	}

	// 数据库表字段操作
	// 添加字段
	/*type Student struct {
		Score uint
	}
	err := gormDB.Migrator().AddColumn(&Student{}, "Score")
	if err != nil {
		fmt.Printf("添加字段错误,err:%s\n", err)
		return
	}*/

	// 删除字段
	// gormDB.Migrator().DropColumn(&Student{}, "email")

	// 修改字段
	/*type Student struct{
		Name string
		UserName string
	}
	gormDB.Migrator().RenameColumn(&Student{}, "name", "user_name")*/

	// 检查字段是否存在
	// isExistField := gormDB.Migrator().HasColumn(&Student{}, "name")
	// fmt.Printf("字段是否存在:%t\n", isExistField)

	// 数据库表索引操作
	type Student struct {
		Name     string `gorm:"index:idx_name"`
		UserName string `gorm:"index:idx_user_name"`
	}
	// 创建索引

	/*err = gormDB.Migrator().CreateIndex(&Student{}, "Name")
	if err != nil {
		fmt.Printf("创建索引失败1，err:%s\n", err)
		return
	}*/
	/*err = gormDB.Migrator().CreateIndex(&Student{}, "idx_name")
	if err != nil {
		fmt.Printf("创建索引失败2，err:%s\n", err)
		return
	}*/

	// 删除索引
	// gormDB.Migrator().DropIndex(&Student{}, "idx_name")
	// gormDB.Migrator().DropIndex(&Student{}, "UserName")

	// 修改索引
	err = gormDB.Migrator().RenameIndex(&Student{}, "UserName", "Name")
	if err != nil {
		fmt.Printf("修改索引名称失败，err:%s\n", err)
		return
	}
	// gormDB.Migrator().RenameIndex(&Student{}, "idx_name", "idx_user_name")

	// 查询索引
	// isExistIndex := gormDB.Migrator().HasIndex(&Student{}, "Name")
	// isExistIndex := gormDB.Migrator().HasIndex(&Student{}, "idx_name")
	// fmt.Printf("查询索引是否存在：%t\n", isExistIndex)

	// 数据库表约束操作

}
