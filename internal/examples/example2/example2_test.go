// Package example2 shows gormmom with zh-Hant fields: the struct passes the check, then works with GORM through its generated columns
//
// example2 示範用繁體中文欄位的 gormmom 用法：結構體通過校驗，之後用產生的欄位照常搭配 GORM 使用
package example2

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/internal/examples/example2/internal/models"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestValidate(t *testing.T) {
	// A plain unit test, no database needed
	// 普通的單元測試，不需要資料庫
	require.NoError(t, gormmom.NewConfigs([]any{&models.T學生{}}, gormmom.NewOptions()).Validate())
}

func TestUse(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// The table, the columns and the indexes get the ASCII names
	// 表、欄、索引用的都是 ASCII 名字
	done.Done(db.AutoMigrate(&models.T學生{}))
	require.NoError(t, db.Create(&models.T學生{X姓名: "陳小明", N年級: "grade-1", S社團: "tennis"}).Error)

	repo := gormrepo.NewRepo(gormclass.Use(&models.T學生{}))
	student, err := repo.With(t.Context(), db).First(func(db *gorm.DB, cls *models.T學生Columns) *gorm.DB {
		return db.Where(cls.X姓名.Eq("陳小明"))
	})
	require.NoError(t, err)
	require.Equal(t, "tennis", student.S社團)
}
