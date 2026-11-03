// Package example9 shows gormmom with Vietnamese fields: the struct passes the check, then works with GORM through its generated columns
//
// example9 minh họa gormmom với các trường tiếng Việt: struct vượt qua kiểm tra, sau đó dùng với GORM qua các cột được sinh ra
package example9

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/internal/examples/example9/internal/models"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestValidate(t *testing.T) {
	// A plain unit test, no database needed
	// Một bài kiểm thử đơn vị bình thường, không cần cơ sở dữ liệu
	require.NoError(t, gormmom.NewConfigs([]any{&models.HọcSinh{}}, gormmom.NewOptions()).Validate())
}

func TestUse(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// The table, the columns and the indexes get the ASCII names
	// Bảng, các cột và các chỉ mục đều dùng tên ASCII
	done.Done(db.AutoMigrate(&models.HọcSinh{}))
	require.NoError(t, db.Create(&models.HọcSinh{Tên: "Minh", Lớp: "grade-1", CâuLạcBộ: "tennis"}).Error)

	repo := gormrepo.NewRepo(gormclass.Use(&models.HọcSinh{}))
	student, err := repo.With(t.Context(), db).First(func(db *gorm.DB, cls *models.HọcSinhColumns) *gorm.DB {
		return db.Where(cls.Tên.Eq("Minh"))
	})
	require.NoError(t, err)
	require.Equal(t, "tennis", student.CâuLạcBộ)
}
