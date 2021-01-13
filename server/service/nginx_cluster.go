package service

import (
	"nginx-web/global"
	"nginx-web/model"
	"nginx-web/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNginxCluster
//@description: 创建NginxCluster记录
//@param: Cluster model.NginxCluster
//@return: err error

func CreateNginxCluster(Cluster model.NginxCluster) (err error) {
	_, sa := GetAuthorityInfo(model.SysAuthority{AuthorityId:Cluster.AuthorityId})
	Cluster.Department = sa.AuthorityName
	err = global.GVA_DB.Create(&Cluster).Error
	return err
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxCluster
//@description: 删除NginxCluster记录
//@param: Cluster model.NginxCluster
//@return: err error

func DeleteNginxCluster(Cluster model.NginxCluster) (err error) {
	err = global.GVA_DB.Delete(Cluster).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxClusterByIds
//@description: 批量删除NginxCluster记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNginxClusterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.NginxCluster{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNginxCluster
//@description: 更新NginxCluster记录
//@param: Cluster *model.NginxCluster
//@return: err error

func UpdateNginxCluster(Cluster *model.NginxCluster) (err error) {
	_, sa := GetAuthorityInfo(model.SysAuthority{AuthorityId:Cluster.AuthorityId})
	Cluster.Department = sa.AuthorityName
	err = global.GVA_DB.Save(Cluster).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxCluster
//@description: 根据id获取NginxCluster记录
//@param: id uint
//@return: err error, Cluster model.NginxCluster

func GetNginxCluster(id uint) (err error, Cluster model.NginxCluster) {
	err = global.GVA_DB.Where("id = ?", id).First(&Cluster).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxClusterInfoList
//@description: 分页获取NginxCluster记录
//@param: info request.NginxClusterSearch
//@return: err error, list interface{}, total int64

func GetNginxClusterInfoList(info request.NginxClusterSearch, authId string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.NginxCluster{})
    var Clusters []model.NginxCluster
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.ClusterName != "" {
		db = db.Where("`cluster_name` LIKE ? ","%"+ info.ClusterName+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Joins("JOIN sys_data_authority_id ON sys_data_authority_id.sys_authority_authority_id = ? ",
		authId).Where("authority_id = sys_data_authority_id.data_authority_id_authority_id").Find(&Clusters).Error
	return err, Clusters, total
}





//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllNginxClusterInfoList
//@description: 获取NginxCluster记录
//@param: info request.AuthId
//@return: err error, list []model.NginxCluster
func GetAllNginxClusterInfoList(authId string) (err error, list []model.NginxCluster) {
	//err = global.GVA_DB.Find(&list).Error
	err = global.GVA_DB.Joins("JOIN sys_data_authority_id ON sys_data_authority_id.sys_authority_authority_id = ? ",
		authId).Where("authority_id = sys_data_authority_id.data_authority_id_authority_id").Find(&list).Error
	return
}