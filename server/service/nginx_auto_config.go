package service

import (
	"fmt"
	"nginx-web/model"
	"nginx-web/utils"
	"os"
	"strings"
	"text/template"
)

type tplConfigData struct {
	template         *template.Template
	locationPath     string
	autoConfigPath   string
}


func CreateConfigTemp(autoConfig model.AutoConfigStruct) (err error) {
	configPath := "resource/nginx/nginx.conf.tpl"
	autoConfig.Pools = make(map[string]*model.PoolConfig)
	data := tplConfigData{
		locationPath: configPath,
		autoConfigPath: "Domain."+ autoConfig.Domain.DomainName + ".conf",
	}

	//获取Pool数据
	_, Pool := GetNginxPool(autoConfig.Domain.Pool)
	autoConfig.Domain.PoolCluster = Pool
	autoConfig.Pools[Pool.PoolName] = &model.PoolConfig{Pool: Pool}

	//获取Rule数据
    _, Rules := GetDomainRuleByDomainId(autoConfig.Domain.ID)
	RuleConfigs := make([]model.RuleConfig,len(Rules))
    for index, v := range Rules{
    	var Pool model.NginxPool
    	_, Ins := GetDomainRuleInsByRuleId(v.ID)
    	for _, ins := range Ins{
    		if ins.InsType == "proxy_pass" && strings.TrimSpace(ins.InsValue) != ""{
    			_, Pool = GetNginxPoolByName(ins.InsValue)
				autoConfig.Pools[Pool.PoolName] = &model.PoolConfig{Pool: Pool}
			}
		}
    	RuleConfigs[index] = model.RuleConfig{ Rule: v, Ins: Ins, Pool: Pool}
	}
    autoConfig.Rules = RuleConfigs

    //获取Pool节点
    for key, PoolConfig := range autoConfig.Pools{
    	_, Nodes := GetPoolNodeByPoolId(PoolConfig.Pool.ID)
		autoConfig.Pools[key].PoolNodes = Nodes
	}

	// 生成 *Template, 填充 template 字段
	data.template, err = template.ParseFiles(data.locationPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	autoPath := "autoCode/"

	// 写入文件前，先创建文件夹
	if err = utils.CreateDir(autoPath); err != nil {
		return err
	}
    data.autoConfigPath = autoPath + data.autoConfigPath
	// 生成文件
	f, err := os.OpenFile(data.autoConfigPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	defer f.Close()
	if err != nil {
		return err
	}

	if err = data.template.Execute(f, autoConfig); err != nil {
		return err
	}
	return nil
}