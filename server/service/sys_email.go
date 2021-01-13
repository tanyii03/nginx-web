package service

import (
	"nginx-web/utils"
)

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error

func EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}
