package main

import (
	"time"
)

// 模型定义、约定、标签、自动迁移和迁移接口

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

/**
约定：默认情况下，GORM 约定使用 ID 作为主键，使用结构体名的复数作为表名，字段名作为列名，使用 CreatedAt、UpdatedAt、DeletedAt时间追踪。
当然，你可以按照自己的需求自定义约定项。比如时间追踪默认是以当前时间填充为零值得 CreatedAt 字段，以当前时间戳秒数填充 UpdatedAt 字段，如果你喜欢使用时间戳记录创建时间，你可以改变 CreatedAt 字段的类型为 int，默认时间戳是 Unix 秒，你还可以使用标签将时间戳的单位改为纳秒或毫秒。

GORM 定义了一个 gorm.Model 结构体，字段包括 ID、CreatedAt、UpdatedAt、DeletedAt，我们可以将它嵌入到我们自定义的结构体中。

在 GO 语言中，根据名称的首字母大小写来定义是否可被导出，GORM 使用可导出的字段进行 CRUD 时拥有全部权限，另外，GORM 可以使用标签控制字段的权限，可以让一个字段的权限是只读、只写、只创建、只更新和忽略该字段。

标签是模型定义时的可选项，GORM 的标签不区分大小写，推荐使用驼峰式命名。GORM 通过标签实现外键定义、表关联和约束。

迁移

Migrator 接口方法
*/

func main() {

}
