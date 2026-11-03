// Package example3 shows gormmom with Japanese fields: the struct passes the check, then works with GORM through its generated columns
//
// example3 は日本語フィールドでの gormmom の使い方を示す：構造体が検証を通り、生成された列で GORM をそのまま使える
package example3

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/internal/examples/example3/internal/models"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestValidate(t *testing.T) {
	// A plain unit test, no database needed
	// データベースのいらない普通の単体テスト
	require.NoError(t, gormmom.NewConfigs([]any{&models.T生徒{}}, gormmom.NewOptions()).Validate())
}

func TestUse(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// The table, the columns and the indexes get the ASCII names
	// テーブル・列・インデックスはすべて ASCII 名になる
	done.Done(db.AutoMigrate(&models.T生徒{}))
	require.NoError(t, db.Create(&models.T生徒{N名前: "佐藤", G学年: "grade-1", Tクラブ: "tennis"}).Error)

	repo := gormrepo.NewRepo(gormclass.Use(&models.T生徒{}))
	student, err := repo.With(t.Context(), db).First(func(db *gorm.DB, cls *models.T生徒Columns) *gorm.DB {
		return db.Where(cls.N名前.Eq("佐藤"))
	})
	require.NoError(t, err)
	require.Equal(t, "tennis", student.Tクラブ)
}
