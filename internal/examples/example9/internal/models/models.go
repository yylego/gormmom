package models

// HọcSinh has Vietnamese fields; Vietnamese letters have a case, so a field such as Tên is exported as it is, with no ASCII prefix
// Most Vietnamese words have a tone mark, such as Tên, which gives a non-ASCII column name, so each column, index and table gets an explicit ASCII name
//
// HọcSinh có các trường bằng tiếng Việt; chữ cái có chữ hoa, nên trường như Tên được xuất ra trực tiếp, không cần tiền tố ASCII
// Hầu hết từ tiếng Việt đều có dấu, như Tên, sẽ tạo ra tên cột không phải ASCII, nên mỗi cột, chỉ mục và bảng đều được đặt tên ASCII rõ ràng
type HọcSinh struct {
	ID       uint   `gorm:"primaryKey"`
	Tên      string `gorm:"column:name;uniqueIndex:udx_students_name"`
	Lớp      string `gorm:"column:grade;index:idx_students_grade"`
	CâuLạcBộ string `gorm:"column:club"`
}

func (*HọcSinh) TableName() string {
	return "students"
}
