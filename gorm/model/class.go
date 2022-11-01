package model

type Class struct {
	ClassId   int    `gorm:"colum:class_id"`
	ClassName string `gorm:"colum:class_name"`
}

func (class Class) TableName() string {
	return "u_class"
}
