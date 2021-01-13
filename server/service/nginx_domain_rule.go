package service

import (
	"nginx-web/global"
	"nginx-web/model"
	"nginx-web/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDomainRule
//@description: 创建DomainRule记录
//@param: Rule model.DomainRule
//@return: err error

func CreateDomainRule(Rule model.DomainRule) (err error) {
	err = global.GVA_DB.Create(&Rule).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDomainRule
//@description: 删除DomainRule记录
//@param: Rule model.DomainRule
//@return: err error

func DeleteDomainRule(Rule model.DomainRule) (err error) {
	err = global.GVA_DB.Delete(Rule).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDomainRuleByIds
//@description: 批量删除DomainRule记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDomainRuleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DomainRule{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDomainRule
//@description: 更新DomainRule记录
//@param: Rule *model.DomainRule
//@return: err error

func UpdateDomainRule(Rule *model.DomainRule) (err error) {
	err = global.GVA_DB.Save(Rule).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDomainRule
//@description: 根据id获取DomainRule记录
//@param: id uint
//@return: err error, Rule model.DomainRule

func GetDomainRule(id uint) (err error, Rule model.DomainRule) {
	err = global.GVA_DB.Where("id = ?", id).First(&Rule).Error
	return
}


func GetDomainRuleByDomainId(id uint) (err error, Rules []model.DomainRule) {
	err = global.GVA_DB.Where("domain_id = ?", id).Find(&Rules).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDomainRuleInfoList
//@description: 分页获取DomainRule记录
//@param: info request.DomainRuleSearch
//@return: err error, list interface{}, total int64

func GetDomainRuleInfoList(info request.DomainRuleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.DomainRule{})
    var Rules []model.DomainRule
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Where("domain_id = ? ", info.DomainId).Find(&Rules).Error
	return err, Rules, total
}


