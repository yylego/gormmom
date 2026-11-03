package models

// Schüler has German fields; German letters have a case, so the fields are exported with no ASCII prefix
// A name with an umlaut such as Schüler gives a non-ASCII table name, so each column, index and table gets an explicit ASCII name
//
// Schüler hat deutsche Felder; die Buchstaben haben Groß- und Kleinschreibung, daher werden die Felder ohne ASCII-Präfix exportiert
// Ein Name mit Umlaut wie Schüler ergibt einen Tabellennamen, der kein ASCII ist, daher bekommen Spalten, Indizes und Tabelle explizite ASCII-Namen
type Schüler struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"column:name;uniqueIndex:udx_students_name"`
	Klasse string `gorm:"column:grade;index:idx_students_grade"`
	Verein string `gorm:"column:club"`
}

func (*Schüler) TableName() string {
	return "students"
}
