package services

import (
	"errors"
	"panda/app/common/request"
	"panda/app/models"
	"panda/global"
)

type objectService struct {
}

var ObjectService = new(objectService)

// ObjectPush 插入数据库
func (objectService *objectService) ObjectPush(params request.Object) (err error, object models.Object) {
	var result = global.App.DB.Where("name = ?", params.Name).Select("id").First(&models.Object{})
	if result.RowsAffected != 0 {
		err = errors.New("物体已经存在")
		return
	}

	lengthUnits := []request.LengthUnit{request.Meter, request.Kilometer, request.Astronomical, request.LunarDistance, request.LightYear}
	massUnits := []request.MassUnit{request.Kilogram, request.MetricTon, request.LunarMass, request.EarthMass, request.SolarMass, request.SolarSystemMass, request.MilkyWayMass}

	lengthUnitMatch := false
	for _, units := range lengthUnits {
		unit := string(units)
		if params.LengthUnit == unit {
			lengthUnitMatch = true
			break
		}
	}

	massUnitMatch := false
	for _, units := range massUnits {
		unit := string(units)
		if params.MassUnit == unit {
			massUnitMatch = true
			break
		}
	}
	if !massUnitMatch || !lengthUnitMatch {
		err = errors.New("长度或质量单位错误")
		return
	}
	object = models.Object{
		Name:       params.Name,
		Length:     params.Length,
		LengthUnit: params.LengthUnit,
		Mass:       params.Mass,
		MassUnit:   params.MassUnit,
	}
	err = global.App.DB.Create(&object).Error
	return
}

// ObjectQuery 查询数据库
func (objectService *objectService) ObjectQuery(params request.ObjectQuery) (err error, object []models.Object) {

	var result = global.App.DB.Where("length >= ? and lengthUnit = ?", params.Length, params.LengthUnit).Find(&object)
	if result.RowsAffected != 0 {
		//err = errors.New("不存在对象")
		return
	}
	err = errors.New("不存在对象")
	return
}

// ObjectDelete 删除数据库对象
func (objectService *objectService) ObjectDelete(params request.ObjectDelete) (err error) {
	err = global.App.DB.Where("name = ?", params.Name).Delete(&models.Object{}).Error
	if err != nil {
		global.App.Log.Sugar().Error(params.Name, ",删除失败err：", err)
		return
	}

	return nil
}
