package utils

import (
	"fmt"
	"nginx-web/global"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type SSHClient struct {
	IP        string   `json:"IP"`
	Port      string   `json:"Port"`
	UserName  string   `json:"UserName"`
	Passwd    string   `json:"Passwd"`
	TimeOut   int
	Client    *ssh.Client
}

//添加公钥
func (c *SSHClient) AddPublicKey() error {
	PublicKey, _ := ioutil.ReadFile(global.GVA_CONFIG.SSH.Public_key)
	Command := fmt.Sprintf("mkdir -p -m 700 ~/.ssh && echo %s%v >> ~/.ssh/authorized_keys && chmod 600 ~/.ssh/authorized_keys", strings.Replace(string(PublicKey),"\n","",-1), "\r")
	return c.ExecCommand(Command)
}

//SSH可用性测试
func (c *SSHClient) Ping() error {
	return c.ExecCommand("ls")
}

//获取SSH客户端
func (c *SSHClient) SetClient(client *ssh.Client)  {
	c.Client = client
}

//上传文件
func (c *SSHClient) PutFile(Path, Local, FileName string) error {
    sftpClient, err := sftp.NewClient(c.Client)
    if err != nil{
    	return err
	}
    defer sftpClient.Close()
	// create destination file
    _, err = sftpClient.Stat(Path + FileName)
    if err == nil{
		if err = sftpClient.Remove(Path + FileName); err != nil{
			return err
		}
	}
	dstFile, err := sftpClient.Create(Path + FileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := os.Open(Local + FileName)
	if err != nil {
		return err
	}

	// copy source file to destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return err
}

//执行命令
func (c *SSHClient) ExecCommand(Command string) error {
	session, err := c.Client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	err = session.Run(Command)
	return err
}

//重载Nginx配置
func (c *SSHClient) Reload() error {
	return c.ExecCommand("nginx -s reload")
}


//SFTP删除文件
func (c *SSHClient) RemoveFile(FileName string) error {
	sftpClient, err := sftp.NewClient(c.Client)
	if err != nil{
		return err
	}
	defer sftpClient.Close()

	return sftpClient.Remove(FileName)
}

//关闭SSH客户端
func (c *SSHClient) Close()  error {
	return c.Client.Close()
}

//SSH 私钥认证
func (c *SSHClient) VerifySshByKey() error {
	key, err := ioutil.ReadFile(global.GVA_CONFIG.SSH.Private_key)
	if err != nil {
		return err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: c.UserName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addrPort := c.IP +":" + c.Port
	client, err := ssh.Dial("tcp", addrPort, config)
	if err != nil {
		return err
	}
	c.SetClient(client)
	return err
}

//SSH账号密码认证
func (c *SSHClient) VerifySshByPwd() error {
	addrPort := c.IP + ":" + c.Port
	client, err := ssh.Dial("tcp", addrPort, &ssh.ClientConfig{
		User:            c.UserName,
		Auth:            []ssh.AuthMethod{ssh.Password(c.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return err
	}
	c.SetClient(client)
	return nil
}

