package model

import "errors"


type AutoConfigStruct struct {
    Domain   NginxDomain    `json:"Domain"`
    Rules    []RuleConfig   `json:"Rules"`
    Pools    map[string]*PoolConfig  `json:"Pools"`
}

type RuleConfig struct {
	Rule  DomainRule       `json:"Rule"`
	Ins   []DomainRuleIns  `json:"Ins"`
	Pool  NginxPool        `json:"Pool"`
}

type PoolConfig struct {
	Pool       NginxPool  `json:"Pool"`
	PoolNodes  []PoolNode `json:"PoolNodes"`
}

var AutoConfigErr error = errors.New("创建代码成功并移动文件成功")