package models

// T학생 has Korean fields, and gives each column, index and table an explicit database-safe ASCII name
// Korean letters have no case, so each field starts with one of A-Z to be exported
//
// T학생 은 필드 이름을 한국어로 쓰고, 열 이름·인덱스 이름·테이블 이름에는 데이터베이스에 안전한 ASCII 이름을 명시한다
// 한국어 글자에는 대소문자가 없으므로, 각 필드 앞에 ASCII 대문자를 붙여 내보낸다
type T학생 struct {
	ID   uint   `gorm:"primaryKey"`
	I이름  string `gorm:"column:name;uniqueIndex:udx_students_name"`
	H학년  string `gorm:"column:grade;index:idx_students_grade"`
	D동아리 string `gorm:"column:club"`
}

func (*T학생) TableName() string {
	return "students"
}
