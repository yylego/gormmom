package models

// Estudiante has Spanish fields; Spanish letters have a case, so a field such as Año is exported as it is, with no ASCII prefix
// A field with an accent such as Año gives a non-ASCII column name, so each column, index and table gets an explicit ASCII name
//
// Estudiante tiene campos en español; las letras tienen mayúsculas, así que un campo como Año se exporta tal cual, sin prefijo ASCII
// Un campo con acento como Año da un nombre de columna no ASCII, así que cada columna, índice y tabla recibe un nombre ASCII explícito
type Estudiante struct {
	ID     uint   `gorm:"primaryKey"`
	Nombre string `gorm:"column:name;uniqueIndex:udx_students_name"`
	Año    string `gorm:"column:grade;index:idx_students_grade"`
	Club   string `gorm:"column:club"`
}

func (*Estudiante) TableName() string {
	return "students"
}
