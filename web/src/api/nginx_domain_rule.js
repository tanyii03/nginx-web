import service from '@/utils/request'

// @Tags DomainRule
// @Summary 创建DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRule true "创建DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Rule/createDomainRule [post]
export const createDomainRule = (data) => {
     return service({
         url: "/Rule/createDomainRule",
         method: 'post',
         data
     })
 }


// @Tags DomainRule
// @Summary 删除DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRule true "删除DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Rule/deleteDomainRule [delete]
 export const deleteDomainRule = (data) => {
     return service({
         url: "/Rule/deleteDomainRule",
         method: 'delete',
         data
     })
 }

// @Tags DomainRule
// @Summary 删除DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Rule/deleteDomainRule [delete]
 export const deleteDomainRuleByIds = (data) => {
     return service({
         url: "/Rule/deleteDomainRuleByIds",
         method: 'delete',
         data
     })
 }

// @Tags DomainRule
// @Summary 更新DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRule true "更新DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Rule/updateDomainRule [put]
 export const updateDomainRule = (data) => {
     return service({
         url: "/Rule/updateDomainRule",
         method: 'put',
         data
     })
 }


// @Tags DomainRule
// @Summary 用id查询DomainRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainRule true "用id查询DomainRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Rule/findDomainRule [get]
 export const findDomainRule = (params) => {
     return service({
         url: "/Rule/findDomainRule",
         method: 'get',
         params
     })
 }


// @Tags DomainRule
// @Summary 分页获取DomainRule列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DomainRule列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Rule/getDomainRuleList [get]
 export const getDomainRuleList = (params) => {
     return service({
         url: "/Rule/getDomainRuleList",
         method: 'get',
         params
     })
 }