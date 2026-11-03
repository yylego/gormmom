// Package example1 shows gormmom with zh-Hans fields: the struct passes the check, then works with GORM through its generated columns
//
// example1 演示简体中文字段的 gormmom 用法：结构体通过校验，之后用生成的列照常配合 GORM 使用
package example1

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/internal/examples/example1/internal/models"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestValidate(t *testing.T) {
	// A plain unit test, no database needed
	// 普通的单元测试，不需要数据库
	require.NoError(t, gormmom.NewConfigs([]any{&models.T学生{}}, gormmom.NewOptions()).Validate())
}

func TestUse(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// The table, the columns and the indexes get the ASCII names
	// 表、列、索引用的都是 ASCII 名字
	done.Done(db.AutoMigrate(&models.T学生{}))
	require.NoError(t, db.Create(&models.T学生{X姓名: "张三", B班级: "grade-1", S社团: "tennis"}).Error)

	repo := gormrepo.NewRepo(gormclass.Use(&models.T学生{}))
	student, err := repo.With(t.Context(), db).First(func(db *gorm.DB, cls *models.T学生Columns) *gorm.DB {
		return db.Where(cls.X姓名.Eq("张三"))
	})
	require.NoError(t, err)
	require.Equal(t, "tennis", student.S社团)
}
