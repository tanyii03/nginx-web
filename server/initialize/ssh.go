package initialize

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"nginx-web/global"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
)

func SSH()  {
	//检查是否配置私钥
	exist, _ := PathExists(global.GVA_CONFIG.SSH.Private_key)
	if !exist{
		GenRsaKey(2048)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GenRsaKey(bits int) error {

	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

    err = ioutil.WriteFile(global.GVA_CONFIG.SSH.Private_key, pem.EncodeToMemory(priBlock), 0755)
	if err != nil{
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	key, err :=  ssh.NewPublicKey(publicKey)
	if err != nil{
		return err
	}

	err = ioutil.WriteFile(global.GVA_CONFIG.SSH.Public_key, ssh.MarshalAuthorizedKey(key), 0755)
	if err != nil {
		return err
	}
	return nil
}