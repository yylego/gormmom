package models

// T生徒 has Japanese fields, and gives each column, index and table an explicit database-safe ASCII name
// Japanese letters have no case, so each field starts with one of A-Z to be exported
//
// T生徒 はフィールド名を日本語で書き、列名・インデックス名・テーブル名にはデータベースで安全な ASCII 名を明示する
// 日本語の文字には大文字・小文字がないため、各フィールドの先頭に ASCII の大文字を付けてエクスポートする
type T生徒 struct {
	ID   uint   `gorm:"primaryKey"`
	N名前  string `gorm:"column:name;uniqueIndex:udx_students_name"`
	G学年  string `gorm:"column:grade;index:idx_students_grade"`
	Tクラブ string `gorm:"column:club"`
}

func (*T生徒) TableName() string {
	return "students"
}
