package models

// T學生 has zh-Hant fields, and gives each column, index and table an explicit database-safe ASCII name
// Chinese characters have no case, so each field starts with one of A-Z to be exported
//
// T學生 的欄位名用繁體中文，欄名、索引名、表名都明確給了資料庫安全的 ASCII 名字
// 漢字沒有大小寫，所以每個欄位前面加一個 A-Z 的字母才能匯出
type T學生 struct {
	ID  uint   `gorm:"primaryKey"`
	X姓名 string `gorm:"column:name;uniqueIndex:udx_students_name"`
	N年級 string `gorm:"column:grade;index:idx_students_grade"`
	S社團 string `gorm:"column:club"`
}

func (*T學生) TableName() string {
	return "students"
}
