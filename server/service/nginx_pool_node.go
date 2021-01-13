package service

import (
	"nginx-web/global"
	"nginx-web/model"
	"nginx-web/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePoolNode
//@description: 创建PoolNode记录
//@param: PNode model.PoolNode
//@return: err error

func CreatePoolNode(PNode model.PoolNode) (err error) {
	err = global.GVA_DB.Create(&PNode).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePoolNode
//@description: 删除PoolNode记录
//@param: PNode model.PoolNode
//@return: err error

func DeletePoolNode(PNode model.PoolNode) (err error) {
	err = global.GVA_DB.Delete(PNode).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePoolNodeByIds
//@description: 批量删除PoolNode记录
//@param: ids request.IdsReq
//@return: err error

func DeletePoolNodeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PoolNode{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePoolNode
//@description: 更新PoolNode记录
//@param: PNode *model.PoolNode
//@return: err error

func UpdatePoolNode(PNode *model.PoolNode) (err error) {
	err = global.GVA_DB.Save(PNode).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPoolNode
//@description: 根据id获取PoolNode记录
//@param: id uint
//@return: err error, PNode model.PoolNode

func GetPoolNode(id uint) (err error, PNode model.PoolNode) {
	err = global.GVA_DB.Where("id = ?", id).First(&PNode).Error
	return
}

func GetPoolNodeByPoolId(id uint) (err error, PNode []model.PoolNode) {
	err = global.GVA_DB.Where("cluster_id = ? and status = ?", id, "Enable").Find(&PNode).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPoolNodeInfoList
//@description: 分页获取PoolNode记录
//@param: info request.PoolNodeSearch
//@return: err error, list interface{}, total int64

func GetPoolNodeInfoList(info request.PoolNodeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PoolNode{})
    var PNodes []model.PoolNode
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Where("cluster_id = ?", info.ClusterId).Find(&PNodes).Error
	return err, PNodes, total
}