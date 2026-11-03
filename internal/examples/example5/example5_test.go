// Package example5 shows gormmom with Spanish fields: the struct passes the check, then works with GORM through its generated columns
//
// example5 muestra gormmom con campos en español: el struct pasa la validación y luego funciona con GORM a través de las columnas generadas
package example5

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/internal/examples/example5/internal/models"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestValidate(t *testing.T) {
	// A plain unit test, no database needed
	// Una prueba unitaria normal, sin base de datos
	require.NoError(t, gormmom.NewConfigs([]any{&models.Estudiante{}}, gormmom.NewOptions()).Validate())
}

func TestUse(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// The table, the columns and the indexes get the ASCII names
	// La tabla, las columnas y los índices reciben los nombres ASCII
	done.Done(db.AutoMigrate(&models.Estudiante{}))
	require.NoError(t, db.Create(&models.Estudiante{Nombre: "Lucía", Año: "grade-1", Club: "tennis"}).Error)

	repo := gormrepo.NewRepo(gormclass.Use(&models.Estudiante{}))
	student, err := repo.With(t.Context(), db).First(func(db *gorm.DB, cls *models.EstudianteColumns) *gorm.DB {
		return db.Where(cls.Nombre.Eq("Lucía"))
	})
	require.NoError(t, err)
	require.Equal(t, "tennis", student.Club)
}
