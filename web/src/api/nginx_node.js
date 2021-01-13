import service from '@/utils/request'

// @Tags NginxNode
// @Summary 创建NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxNode true "创建NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Node/createNginxNode [post]
export const createNginxNode = (data) => {
     return service({
         url: "/Node/createNginxNode",
         method: 'post',
         data
     })
 }


// @Tags NginxNode
// @Summary 删除NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxNode true "删除NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Node/deleteNginxNode [delete]
 export const deleteNginxNode = (data) => {
     return service({
         url: "/Node/deleteNginxNode",
         method: 'delete',
         data
     })
 }

// @Tags NginxNode
// @Summary 删除NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Node/deleteNginxNode [delete]
 export const deleteNginxNodeByIds = (data) => {
     return service({
         url: "/Node/deleteNginxNodeByIds",
         method: 'delete',
         data
     })
 }

// @Tags NginxNode
// @Summary 更新NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxNode true "更新NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Node/updateNginxNode [put]
 export const updateNginxNode = (data) => {
     return service({
         url: "/Node/updateNginxNode",
         method: 'put',
         data
     })
 }


// @Tags NginxNode
// @Summary 用id查询NginxNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxNode true "用id查询NginxNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Node/findNginxNode [get]
 export const findNginxNode = (params) => {
     return service({
         url: "/Node/findNginxNode",
         method: 'get',
         params
     })
 }


// @Tags NginxNode
// @Summary 分页获取NginxNode列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取NginxNode列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Node/getNginxNodeList [get]
 export const getNginxNodeList = (params) => {
     return service({
         url: "/Node/getNginxNodeList",
         method: 'get',
         params
     })
 }