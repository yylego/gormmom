// Package example6 shows gormmom with Portuguese fields: the struct passes the check, then works with GORM through its generated columns
//
// example6 mostra o gormmom com campos em português: o struct passa na validação e depois funciona com o GORM pelas colunas geradas
package example6

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/gormmom"
	"github.com/yylego/gormmom/internal/examples/example6/internal/models"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestValidate(t *testing.T) {
	// A plain unit test, no database needed
	// Um teste unitário comum, sem banco de dados
	require.NoError(t, gormmom.NewConfigs([]any{&models.Aluno{}}, gormmom.NewOptions()).Validate())
}

func TestUse(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// The table, the columns and the indexes get the ASCII names
	// A tabela, as colunas e os índices recebem os nomes ASCII
	done.Done(db.AutoMigrate(&models.Aluno{}))
	require.NoError(t, db.Create(&models.Aluno{Nome: "João", Série: "grade-1", Clube: "tennis"}).Error)

	repo := gormrepo.NewRepo(gormclass.Use(&models.Aluno{}))
	student, err := repo.With(t.Context(), db).First(func(db *gorm.DB, cls *models.AlunoColumns) *gorm.DB {
		return db.Where(cls.Nome.Eq("João"))
	})
	require.NoError(t, err)
	require.Equal(t, "tennis", student.Clube)
}
