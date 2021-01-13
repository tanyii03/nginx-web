package service

import (
	"errors"
	"nginx-web/global"
	"nginx-web/model"
	"nginx-web/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNginxPool
//@description: 创建NginxPool记录
//@param: Pool model.NginxPool
//@return: err error

func CreateNginxPool(Pool model.NginxPool) (err error) {
	err = global.GVA_DB.Create(&Pool).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxPool
//@description: 删除NginxPool记录
//@param: Pool model.NginxPool
//@return: err error

func DeleteNginxPool(Pool model.NginxPool) (err error) {
	_, Pool = GetNginxPool(Pool.ID)
    err, total := GetNginxDomainInfoByPoolId(Pool)
    if total != 0{
    	return errors.New("Domain使用集群")
	}

	err = global.GVA_DB.Delete(Pool).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxPoolByIds
//@description: 批量删除NginxPool记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNginxPoolByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.NginxPool{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNginxPool
//@description: 更新NginxPool记录
//@param: Pool *model.NginxPool
//@return: err error

func UpdateNginxPool(Pool *model.NginxPool) (err error) {
	err = global.GVA_DB.Save(Pool).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxPool
//@description: 根据id获取NginxPool记录
//@param: id uint
//@return: err error, Pool model.NginxPool

func GetNginxPool(id uint) (err error, Pool model.NginxPool) {
	err = global.GVA_DB.Where("id = ?", id).First(&Pool).Error
	return
}

func GetNginxPoolByName(PoolName string) (err error, Pool model.NginxPool) {
	err = global.GVA_DB.Where("pool_name = ?", PoolName).First(&Pool).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxPoolInfoList
//@description: 分页获取NginxPool记录
//@param: info request.NginxPoolSearch
//@return: err error, list interface{}, total int64

func GetNginxPoolInfoList(info request.NginxPoolSearch, authId string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
   // 创建db
	db := global.GVA_DB.Model(&model.NginxPool{})
   var Pools []model.NginxPool
   // 如果有条件搜索 下方会自动创建搜索语句
   if info.PoolName != "" {
      db = db.Where("`pool_name` LIKE ?","%"+ info.PoolName+"%")
   }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Joins("JOIN sys_data_authority_id ON sys_data_authority_id.sys_authority_authority_id = ? ",
		authId).Where("authority_id = sys_data_authority_id.data_authority_id_authority_id").Find(&Pools).Error
	return err, Pools, total
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllNginxPoolInfoList
//@description: 获取Pool记录
//@param: info request.AuthId
//@return: err error, list []model.NginxPool
func GetAllNginxPoolInfoList(authId string) (err error, list []model.NginxPool) {
	//err = global.GVA_DB.Find(&list).Error
	err = global.GVA_DB.Joins("JOIN sys_data_authority_id ON sys_data_authority_id.sys_authority_authority_id = ? ",
		authId).Where("authority_id = sys_data_authority_id.data_authority_id_authority_id").Find(&list).Error
	return
}