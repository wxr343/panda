package models

type Object struct {
	ID
	Name       string  `json:"name" gorm:"not null;comment:物体名称"`
	Length     float32 `json:"length" gorm:"not null;index;comment:物体长度"`
	LengthUnit string  `json:"lengthUnit" gorm:"not null;index;comment:物体长度单位"`
	Mass       float32 `json:"mass" gorm:"not null;index;comment:物体质量"`
	MassUnit   string  `json:"massUnit" gorm:"not null;index;comment:物体质量单位"`
	Timestamps
	SoftDeletes
}
