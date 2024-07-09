package request

type LengthUnit string

const (
	Meter         LengthUnit = "m"  // 米
	Kilometer     LengthUnit = "km" // 公里
	Astronomical  LengthUnit = "au" // 天文单位
	LunarDistance LengthUnit = "ld" // 地月距离
	LightYear     LengthUnit = "ly" // 光年
)

type MassUnit string

const (
	Kilogram        MassUnit = "kg"              // 千克
	MetricTon       MassUnit = "t"               // 公吨
	LunarMass       MassUnit = "LunarMass"       // 月球质量
	EarthMass       MassUnit = "EarthMass"       // 地球质量
	SolarMass       MassUnit = "SolarMass"       // 太阳质量
	SolarSystemMass MassUnit = "SolarSystemMass" // 太阳系质量
	MilkyWayMass    MassUnit = "MilkyWayMass"    // 银河系质量
)

type Object struct {
	Name       string  `json:"name" gorm:"not null;comment:物体名称"`
	Length     float32 `json:"length" gorm:"not null;index;comment:物体长度"`
	LengthUnit string  `json:"lengthUnit" gorm:"not null;index;comment:物体长度单位"`
	Mass       float32 `json:"mass" gorm:"not null;default:'';comment:物体质量"`
	MassUnit   string  `json:"massUnit" gorm:"not null;default:'';comment:物体质量单位"`
}
