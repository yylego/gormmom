// Package example7 shows gormmom with French fields: the struct passes the check, then works with GORM through its generated columns
//
// example7 montre gormmom avec des champs en français : la struct passe la validation, puis fonctionne avec GORM grâce aux colonnes générées
package example7

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/internal/examples/example7/internal/models"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestValidate(t *testing.T) {
	// A plain unit test, no database needed
	// Un test unitaire ordinaire, sans base de données
	require.NoError(t, gormmom.NewConfigs([]any{&models.Élève{}}, gormmom.NewOptions()).Validate())
}

func TestUse(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// The table, the columns and the indexes get the ASCII names
	// La table, les colonnes et les index reçoivent les noms ASCII
	done.Done(db.AutoMigrate(&models.Élève{}))
	require.NoError(t, db.Create(&models.Élève{Prénom: "Chloé", Classe: "grade-1", Club: "tennis"}).Error)

	repo := gormrepo.NewRepo(gormclass.Use(&models.Élève{}))
	student, err := repo.With(t.Context(), db).First(func(db *gorm.DB, cls *models.ÉlèveColumns) *gorm.DB {
		return db.Where(cls.Prénom.Eq("Chloé"))
	})
	require.NoError(t, err)
	require.Equal(t, "tennis", student.Club)
}
