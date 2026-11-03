package models

// Élève has French fields; French letters have a case, so a field such as Prénom is exported as it is, with no ASCII prefix
// A field with an accent such as Prénom gives a non-ASCII column name, so each column, index and table gets an explicit ASCII name
//
// Élève a des champs en français ; les lettres ont des majuscules, donc un champ comme Prénom est exporté tel quel, sans préfixe ASCII
// Un champ accentué comme Prénom donne un nom de colonne non ASCII, donc chaque colonne, index et table reçoit un nom ASCII explicite
type Élève struct {
	ID     uint   `gorm:"primaryKey"`
	Prénom string `gorm:"column:name;uniqueIndex:udx_students_name"`
	Classe string `gorm:"column:grade;index:idx_students_grade"`
	Club   string `gorm:"column:club"`
}

func (*Élève) TableName() string {
	return "students"
}
