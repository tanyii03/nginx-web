package service

import (
	"nginx-web/global"
	"nginx-web/model"
	"nginx-web/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDomainRuleIns
//@description: 创建DomainRuleIns记录
//@param: Ins model.DomainRuleIns
//@return: err error

func CreateDomainRuleIns(Ins model.DomainRuleIns) (err error) {
	err = global.GVA_DB.Create(&Ins).Error
	return err
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDomainRuleIns
//@description: 删除DomainRuleIns记录
//@param: Ins model.DomainRuleIns
//@return: err error

func DeleteDomainRuleIns(Ins model.DomainRuleIns) (err error) {
	err = global.GVA_DB.Delete(Ins).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDomainRuleInsByIds
//@description: 批量删除DomainRuleIns记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDomainRuleInsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DomainRuleIns{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDomainRuleIns
//@description: 更新DomainRuleIns记录
//@param: Ins *model.DomainRuleIns
//@return: err error

func UpdateDomainRuleIns(Ins *model.DomainRuleIns) (err error) {
	err = global.GVA_DB.Save(Ins).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDomainRuleIns
//@description: 根据id获取DomainRuleIns记录
//@param: id uint
//@return: err error, Ins model.DomainRuleIns

func GetDomainRuleIns(id uint) (err error, Ins model.DomainRuleIns) {
	err = global.GVA_DB.Where("id = ?", id).First(&Ins).Error
	return
}

func GetDomainRuleInsByRuleId(id uint) (err error, Ins []model.DomainRuleIns) {
	err = global.GVA_DB.Where("rule_id = ?", id).Find(&Ins).Error
	return
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDomainRuleInsInfoList
//@description: 分页获取DomainRuleIns记录
//@param: info request.DomainRuleInsSearch
//@return: err error, list interface{}, total int64

func GetDomainRuleInsInfoList(info request.DomainRuleInsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.DomainRuleIns{})
    var Inss []model.DomainRuleIns
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Where("rule_id = ? ", info.RuleId).Find(&Inss).Error
	return err, Inss, total
}