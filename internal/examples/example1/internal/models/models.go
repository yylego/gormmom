package models

// T学生 has zh-Hans fields, and gives each column, index and table an explicit database-safe ASCII name
// Chinese characters have no case, so each field starts with one of A-Z to be exported
//
// T学生 的字段名用简体中文，列名、索引名、表名都显式给了数据库安全的 ASCII 名字
// 汉字没有大小写，所以每个字段前面加一个 A-Z 的字母才能导出
type T学生 struct {
	ID  uint   `gorm:"primaryKey"`
	X姓名 string `gorm:"column:name;uniqueIndex:udx_students_name"`
	B班级 string `gorm:"column:grade;index:idx_students_grade"`
	S社团 string `gorm:"column:club"`
}

func (*T学生) TableName() string {
	return "students"
}
