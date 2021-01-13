import service from '@/utils/request'

// @Tags DomainRuleIns
// @Summary 创建DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRuleIns true "创建DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Ins/createDomainRuleIns [post]
export const createDomainRuleIns = (data) => {
     return service({
         url: "/Ins/createDomainRuleIns",
         method: 'post',
         data
     })
 }


// @Tags DomainRuleIns
// @Summary 删除DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRuleIns true "删除DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Ins/deleteDomainRuleIns [delete]
 export const deleteDomainRuleIns = (data) => {
     return service({
         url: "/Ins/deleteDomainRuleIns",
         method: 'delete',
         data
     })
 }

// @Tags DomainRuleIns
// @Summary 删除DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Ins/deleteDomainRuleIns [delete]
 export const deleteDomainRuleInsByIds = (data) => {
     return service({
         url: "/Ins/deleteDomainRuleInsByIds",
         method: 'delete',
         data
     })
 }

// @Tags DomainRuleIns
// @Summary 更新DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRuleIns true "更新DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Ins/updateDomainRuleIns [put]
 export const updateDomainRuleIns = (data) => {
     return service({
         url: "/Ins/updateDomainRuleIns",
         method: 'put',
         data
     })
 }


// @Tags DomainRuleIns
// @Summary 用id查询DomainRuleIns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRuleIns true "用id查询DomainRuleIns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Ins/findDomainRuleIns [get]
 export const findDomainRuleIns = (params) => {
     return service({
         url: "/Ins/findDomainRuleIns",
         method: 'get',
         params
     })
 }


// @Tags DomainRuleIns
// @Summary 分页获取DomainRuleIns列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DomainRuleIns列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Ins/getDomainRuleInsList [get]
 export const getDomainRuleInsList = (params) => {
     return service({
         url: "/Ins/getDomainRuleInsList",
         method: 'get',
         params
     })
 }