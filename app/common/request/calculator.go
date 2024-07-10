package request

type LengthUnit string

const (
	Meter         LengthUnit = "米"    // 米
	Kilometer     LengthUnit = "千米"   // 公里
	Astronomical  LengthUnit = "天文单位" // 天文单位
	LunarDistance LengthUnit = "地月距离" // 地月距离
	LightYear     LengthUnit = "光年"   // 光年
)

type MassUnit string

const (
	Kilogram        MassUnit = "千克"    // 千克
	MetricTon       MassUnit = "吨"     // 公吨
	LunarMass       MassUnit = "月球质量"  // 月球质量
	EarthMass       MassUnit = "地球质量"  // 地球质量
	SolarMass       MassUnit = "太阳质量"  // 太阳质量
	SolarSystemMass MassUnit = "太阳系质量" // 太阳系质量
	MilkyWayMass    MassUnit = "银河系质量" // 银河系质量
)

type Object struct {
	Name       string  `json:"name" gorm:"column:name;not null;comment:物体名称"`
	Length     float32 `json:"length" gorm:"column:length;not null;index;comment:物体长度"`
	LengthUnit string  `json:"lengthUnit" gorm:"column:lengthUnit;not null;index;comment:物体长度单位"`
	Mass       float32 `json:"mass" gorm:"column:mass;not null;index;comment:物体质量"`
	MassUnit   string  `json:"massUnit" gorm:"column:massUnit;not null;index;comment:物体质量单位"`
}

type ObjectQuery struct {
	Length     float32 `json:"length" gorm:"column:length;not null;index;comment:物体长度"`
	LengthUnit string  `json:"lengthUnit" gorm:"column:lengthUnit;not null;index;comment:物体长度单位"`
}

type ObjectDelete struct {
	Name string `json:"name" gorm:"column:name;not null;comment:物体名称"`
}
