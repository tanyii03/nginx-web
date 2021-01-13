package service

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"nginx-web/global"
	"nginx-web/model"
	"nginx-web/model/request"
	"nginx-web/utils"
	"os"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNginxDomainCert
//@description: 创建NginxDomainCert记录
//@param: Cert model.NginxDomainCert
//@return: err error

func CreateNginxDomainCert(Cert model.NginxDomainCert) (err error) {
	err = global.GVA_DB.Create(&Cert).Error
	return err
}

func CreateNginxDomainCertByFile(CertFile model.NginxDomainCertFile) (err error) {
	Cert := model.NginxDomainCert{CertName:CertFile.CertName}
    if CertFile.Pem == ""{
        err = errors.New("Pem  Null")
    	return
	}

	err = saveCertFile(CertFile)
	if err != nil{
		return
	}

    err, Cert.Issued, Cert.Dns, Cert.Validity, Cert.Deadline = parsePemFile([]byte(CertFile.Pem))
    if err != nil{
    	return err
	}
	err = global.GVA_DB.Create(&Cert).Error
	return err
}

//解析证书信息
func parsePemFile(certPEMBlock []byte) (err error, Origan, Dns, Before, After string) {
	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		err = errors.New("Parse Failed")
		return
	}
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return
	}

	if len(x509Cert.Subject.Organization) > 0 {
		Origan = x509Cert.Subject.Organization[0]
	}

	if len(x509Cert.DNSNames) > 0 {
		Dns = x509Cert.DNSNames[0]
	}

	Before, After = x509Cert.NotBefore.Format("2006-01-02 15:04"), x509Cert.NotAfter.Format("2006-01-02 15:04")
    return
}

//保存证书信息到文件
func saveCertFile(CertFile model.NginxDomainCertFile) (err error) {
	// 写入文件前，先创建文件夹
	CertPath := "CertManager/" + CertFile.CertName + "/"
	if err = utils.CreateDir(CertPath); err != nil {
		return err
	}

	// 生成文件
	PemFile, err := os.OpenFile(CertPath + CertFile.CertName + ".pem", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer PemFile.Close()

	_, err = PemFile.WriteString(CertFile.Pem)
	if err != nil {
		return err
	}
	KeyFile, err := os.OpenFile(CertPath + CertFile.CertName + ".key", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer KeyFile.Close()

	_, err = KeyFile.WriteString(CertFile.Key)
	if err != nil {
		return err
	}

	return
}



//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxDomainCert
//@description: 删除NginxDomainCert记录
//@param: Cert model.NginxDomainCert
//@return: err error

func DeleteNginxDomainCert(Cert model.NginxDomainCert) (err error) {
	err = global.GVA_DB.Delete(Cert).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNginxDomainCertByIds
//@description: 批量删除NginxDomainCert记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNginxDomainCertByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.NginxDomainCert{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNginxDomainCert
//@description: 更新NginxDomainCert记录
//@param: Cert *model.NginxDomainCert
//@return: err error

func UpdateNginxDomainCert(Cert *model.NginxDomainCert) (err error) {
	err = global.GVA_DB.Save(Cert).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxDomainCert
//@description: 根据id获取NginxDomainCert记录
//@param: id uint
//@return: err error, Cert model.NginxDomainCert

func GetNginxDomainCert(id uint) (err error, Cert model.NginxDomainCert) {
	err = global.GVA_DB.Where("id = ?", id).First(&Cert).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNginxDomainCertInfoList
//@description: 分页获取NginxDomainCert记录
//@param: info request.NginxDomainCertSearch
//@return: err error, list interface{}, total int64

func GetNginxDomainCertInfoList(info request.NginxDomainCertSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.NginxDomainCert{})
    var Certs []model.NginxDomainCert
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&Certs).Error
	return err, Certs, total
}