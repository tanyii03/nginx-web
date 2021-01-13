package service

import (
	"nginx-web/global"
	"nginx-web/model"
	"nginx-web/model/request"
	"nginx-web/utils"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNginxNode
//@description: 创建NginxNode记录
//@param: Node model.NginxNode
//@return: err error

func CreateNginxNode(Node model.NginxNode) (err error) {
	err = VerifyNginxNode(Node)
	if err != nil{
		Node.Status = "Stopped"
	}else {
		Node.Status = "Running"
	}
	err = global.GVA_DB.Create(&Node).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxNode
//@description: 删除NginxNode记录
//@param: Node model.NginxNode
//@return: err error

func DeleteNginxNode(Node model.NginxNode) (err error) {
	err = global.GVA_DB.Delete(Node).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxNodeByIds
//@description: 批量删除NginxNode记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNginxNodeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.NginxNode{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNginxNode
//@description: 更新NginxNode记录
//@param: Node *model.NginxNode
//@return: err error

func UpdateNginxNode(Node *model.NginxNode) (err error) {
	Node.Status = "Running"
	err = VerifyNginxNode(*Node)
	if err != nil{
		Node.Status = "Stopped"
	}
	err = global.GVA_DB.Save(Node).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxNode
//@description: 根据id获取NginxNode记录
//@param: id uint
//@return: err error, Node model.NginxNode

func GetNginxNode(id uint) (err error, Node model.NginxNode) {
	err = global.GVA_DB.Where("id = ?", id).First(&Node).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxNodeByClusterId
//@description: 根据id获取NginxNode记录
//@param: id uint
//@return: err error, Nodes []model.NginxNode

func GetNginxNodeByClusterId(ClusterId uint) (err error, Nodes []model.NginxNode) {
	err = global.GVA_DB.Where("cluster_id = ?", ClusterId).Find(&Nodes).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxNodeInfoList
//@description: 分页获取NginxNode记录
//@param: info request.NginxNodeSearch
//@return: err error, list interface{}, total int64

func GetNginxNodeInfoList(info request.NginxNodeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.NginxNode{})
    var Nodes []model.NginxNode
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Nodes).Error
	return err, Nodes, total
}

func VerifyNginxNode(Node model.NginxNode) (err error) {
	Client := utils.SSHClient{IP: Node.IP,
		                      UserName: Node.UserName,
		                      Port: Node.Port,
		                      Passwd: strings.TrimSpace(Node.Passwd),
	}
	if err = Client.VerifySshByKey(); err != nil && Client.Passwd != ""{
		err := Client.VerifySshByPwd()
		if err != nil{
			return err
		}
		err = Client.AddPublicKey()
		if err != nil{
			return err
		}
		Client.Close()
	}else {
		return err
	}

    return nil
}
