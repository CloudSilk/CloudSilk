package model

type SystemConfig struct {
	ModelID
	Code   string `gorm:"size:50;comment:代号"`
	Key    string `gorm:"size:50;comment:项"`
	Value  string `gorm:"size:4000;comment:值"`
	Remark string `gorm:"size:500;comment:备注"`
}
