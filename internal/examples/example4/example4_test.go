// Package example4 shows gormmom with Korean fields: the struct passes the check, then works with GORM through its generated columns
//
// example4 은 한국어 필드로 gormmom 을 쓰는 방법을 보여 준다: 구조체가 검증을 통과하고, 생성된 열로 GORM 을 그대로 쓴다
package example4

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/internal/examples/example4/internal/models"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestValidate(t *testing.T) {
	// A plain unit test, no database needed
	// 데이터베이스가 필요 없는 보통의 단위 테스트
	require.NoError(t, gormmom.NewConfigs([]any{&models.T학생{}}, gormmom.NewOptions()).Validate())
}

func TestUse(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// The table, the columns and the indexes get the ASCII names
	// 테이블·열·인덱스는 모두 ASCII 이름이 된다
	done.Done(db.AutoMigrate(&models.T학생{}))
	require.NoError(t, db.Create(&models.T학생{I이름: "김민준", H학년: "grade-1", D동아리: "tennis"}).Error)

	repo := gormrepo.NewRepo(gormclass.Use(&models.T학생{}))
	student, err := repo.With(t.Context(), db).First(func(db *gorm.DB, cls *models.T학생Columns) *gorm.DB {
		return db.Where(cls.I이름.Eq("김민준"))
	})
	require.NoError(t, err)
	require.Equal(t, "tennis", student.D동아리)
}
