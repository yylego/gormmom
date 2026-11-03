package models

// Aluno has Portuguese fields; Portuguese letters have a case, so a field such as Série is exported as it is, with no ASCII prefix
// A field with an accent such as Série gives a non-ASCII column name, so each column, index and table gets an explicit ASCII name
//
// Aluno tem campos em português; as letras têm maiúsculas, então um campo como Série é exportado como está, sem prefixo ASCII
// Um campo com acento como Série gera um nome de coluna não ASCII, então cada coluna, índice e tabela recebe um nome ASCII explícito
type Aluno struct {
	ID    uint   `gorm:"primaryKey"`
	Nome  string `gorm:"column:name;uniqueIndex:udx_students_name"`
	Série string `gorm:"column:grade;index:idx_students_grade"`
	Clube string `gorm:"column:club"`
}

func (*Aluno) TableName() string {
	return "students"
}
