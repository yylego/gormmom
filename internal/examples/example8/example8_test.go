// Package example8 shows gormmom with German fields: the struct passes the check, then works with GORM through its generated columns
//
// example8 zeigt gormmom mit deutschen Feldern: die Struct besteht die Prüfung und funktioniert dann über die generierten Spalten mit GORM
package example8

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/internal/examples/example8/internal/models"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestValidate(t *testing.T) {
	// A plain unit test, no database needed
	// Ein gewöhnlicher Unit-Test, ohne Datenbank
	require.NoError(t, gormmom.NewConfigs([]any{&models.Schüler{}}, gormmom.NewOptions()).Validate())
}

func TestUse(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// The table, the columns and the indexes get the ASCII names
	// Tabelle, Spalten und Indizes bekommen die ASCII-Namen
	done.Done(db.AutoMigrate(&models.Schüler{}))
	require.NoError(t, db.Create(&models.Schüler{Name: "Jonas", Klasse: "grade-1", Verein: "tennis"}).Error)

	repo := gormrepo.NewRepo(gormclass.Use(&models.Schüler{}))
	student, err := repo.With(t.Context(), db).First(func(db *gorm.DB, cls *models.SchülerColumns) *gorm.DB {
		return db.Where(cls.Name.Eq("Jonas"))
	})
	require.NoError(t, err)
	require.Equal(t, "tennis", student.Verein)
}
