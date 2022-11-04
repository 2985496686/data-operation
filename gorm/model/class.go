package model

// Class 被student关联的表
type Class struct {
	ClassId   int       `gorm:"colum:class_id"`
	ClassName string    `gorm:"colum:class_name"`
	Students  []Student `gorm:"references:ClassId;foreignKey:ClassNo;"`
}

func (class Class) TableName() string {
	return "u_class"
}
