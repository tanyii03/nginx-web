package service

import (
	"nginx-web/global"
	"nginx-web/model"
	"nginx-web/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNginxDomain
//@description: 创建NginxDomain记录
//@param: Domain model.NginxDomain
//@return: err error

func CreateNginxDomain(Domain model.NginxDomain) (err error) {
	err = global.GVA_DB.Create(&Domain).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxDomain
//@description: 删除NginxDomain记录
//@param: Domain model.NginxDomain
//@return: err error

func DeleteNginxDomain(Domain model.NginxDomain, User string) (err error) {
	History := model.PublishHistory{Operate:"删除",
		                            DomainId: Domain.ID,
		                            //Type : "删除站点: " + Domain.DomainName,
		                            //ClusterId: Domain.Nginx,
		                            //Config : "Domain."+ Domain.DomainName + ".conf",
		                            User: User,
	}
	err = CreatePublishHistory(History)
	if err != nil{
		return err
	}
	err = global.GVA_DB.Delete(Domain).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxDomainByIds
//@description: 批量删除NginxDomain记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNginxDomainByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.NginxDomain{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNginxDomain
//@description: 更新NginxDomain记录
//@param: Domain *model.NginxDomain
//@return: err error

func UpdateNginxDomain(Domain *model.NginxDomain) (err error) {
	err = global.GVA_DB.Save(Domain).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxDomain
//@description: 根据id获取NginxDomain记录
//@param: id uint
//@return: err error, Domain model.NginxDomain

func GetNginxDomain(id uint) (err error, Domain model.NginxDomain) {
	err = global.GVA_DB.Where("id = ?", id).First(&Domain).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxDomainInfoByClusterId
//@description: 根据Nginx Cluster id获取NginxDomain记录
//@param: id uint
//@return: err error, reDomains []model.NginxDomain

func GetNginxDomainInfoByClusterId(info model.NginxCluster) (err error, reDomains []model.NginxDomain) {
	db := global.GVA_DB.Model(&model.NginxDomain{})
	var Domains []model.NginxDomain
	err = db.Where("nginx = ?", info.ID).Find(&Domains).Error
	return err, Domains
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxDomainInfoByPoolId
//@description: 根据Nginx Pool id获取NginxDomain数量
//@param: info model.NginxPool
//@return: err error, total int64
func GetNginxDomainInfoByPoolId(info model.NginxPool) (err error, total int64) {
	db := global.GVA_DB.Model(&model.NginxDomain{})
	err = db.Where("pool = ?", info.ID).Count(&total).Error
	return err, total
}

