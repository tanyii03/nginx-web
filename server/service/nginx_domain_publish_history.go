package service

import (
	uuid2 "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"nginx-web/global"
	"nginx-web/model"
	"nginx-web/model/request"
	"nginx-web/utils"
	"os"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePublishHistory
//@description: 创建PublishHistory记录
//@param: History model.PublishHistory
//@return: err error

func CreatePublishHistory(History model.PublishHistory) (err error) {
	History.Uuid = uuid2.NewV4().String()
	err, History.Domain = GetNginxDomain(History.DomainId)
	if err != nil{
		return err
	}

	switch History.Operate {
	case "发布":
        History.Type = "部署站点：" + History.Domain.DomainName
	case "删除":
		History.Type = "删除站点：" + History.Domain.DomainName
	}
	err = PublishNginxClusterConfig(History.Domain, History.Uuid, History.Operate)
	if err != nil{
		History.Status = "失败"
	}else {
		History.Status = "成功"
	}
	err = global.GVA_DB.Create(&History).Error
	return err
}

func PublishNginxClusterConfig(Domain model.NginxDomain, logId string, Oprate string) error {
	log := log.New()
	autoPath := "log/publish/"
	// 写入文件前，先创建文件夹
	if err := utils.CreateDir(autoPath); err != nil {
		return err
	}

	file, err := os.OpenFile(autoPath + logId + ".log", os.O_CREATE|os.O_WRONLY, 0666)
	log.Out = file
	log.Info("Start Deploy Domain Config ", Domain.DomainName)

    err, Nodes := GetNginxNodeByClusterId(Domain.Nginx)
    for _, Node := range Nodes{
    	log.Info("Start Executing In ", Node.IP)
		Client := utils.SSHClient{IP: Node.IP,
			                      UserName: Node.UserName,
			                      Port: Node.Port,
		}
		err = Client.VerifySshByKey()
		if err != nil{
			log.Warning("Verify Key Failed In ", Node.IP)
			return err
		}

		switch Oprate {
		case "发布":
			if Domain.UseHttps{
				_  = Client.PutFile(global.GVA_CONFIG.Nginx.CertPath, "certManager/"+Domain.HttpsCert + "/", Domain.HttpsCert+".pem")
				_  = Client.PutFile(global.GVA_CONFIG.Nginx.CertPath, "certManager/"+Domain.HttpsCert + "/", Domain.HttpsCert+".key")
			}
			err = Client.PutFile(global.GVA_CONFIG.Nginx.ConfigPath, "autoCode/", "Domain."+Domain.DomainName+".conf")

		case "删除":
			err = Client.RemoveFile(global.GVA_CONFIG.Nginx.ConfigPath + "Domain." + Domain.DomainName + ".conf")
		}

		if err != nil{
			log.Warning("Failed Put/Remove Config In ", Node.IP, err)
			return err
		}

		err = Client.Reload()
		if err != nil{
			log.Warning("Failed Executing In ", Node.IP, err)
			return err
		}

		Client.Close()
		log.Info("Success Executing In ", Node.IP)
	}
    return nil
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePublishHistory
//@description: 删除PublishHistory记录
//@param: History model.PublishHistory
//@return: err error

func DeletePublishHistory(History model.PublishHistory) (err error) {
	err = global.GVA_DB.Delete(History).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePublishHistoryByIds
//@description: 批量删除PublishHistory记录
//@param: ids request.IdsReq
//@return: err error

func DeletePublishHistoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PublishHistory{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePublishHistory
//@description: 更新PublishHistory记录
//@param: History *model.PublishHistory
//@return: err error

func UpdatePublishHistory(History *model.PublishHistory) (err error) {
	err = global.GVA_DB.Save(History).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPublishHistory
//@description: 根据id获取PublishHistory记录
//@param: id uint
//@return: err error, History model.PublishHistory

func GetPublishHistory(id uint) (err error, History model.PublishHistory) {
	err = global.GVA_DB.Where("id = ?", id).First(&History).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPublishHistoryInfoList
//@description: 分页获取PublishHistory记录
//@param: info request.PublishHistorySearch
//@return: err error, list interface{}, total int64

func GetPublishHistoryInfoList(info request.PublishHistorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PublishHistory{})
    var Historys []model.PublishHistory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.DomainId != 0{
    	db = db.Where(" domain_id = ? ", info.DomainId)
	}
	err = db.Count(&total).Error
	err = db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&Historys).Error
	return err, Historys, total
}