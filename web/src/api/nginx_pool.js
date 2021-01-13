import service from '@/utils/request'

// @Tags NginxPool
// @Summary 创建NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxPool true "创建NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Pool/createNginxPool [post]
export const createNginxPool = (data) => {
     return service({
         url: "/Pool/createNginxPool",
         method: 'post',
         data
     })
 }


// @Tags NginxPool
// @Summary 删除NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxPool true "删除NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Pool/deleteNginxPool [delete]
 export const deleteNginxPool = (data) => {
     return service({
         url: "/Pool/deleteNginxPool",
         method: 'delete',
         data
     })
 }

// @Tags NginxPool
// @Summary 删除NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Pool/deleteNginxPool [delete]
 export const deleteNginxPoolByIds = (data) => {
     return service({
         url: "/Pool/deleteNginxPoolByIds",
         method: 'delete',
         data
     })
 }

// @Tags NginxPool
// @Summary 更新NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxPool true "更新NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Pool/updateNginxPool [put]
 export const updateNginxPool = (data) => {
     return service({
         url: "/Pool/updateNginxPool",
         method: 'put',
         data
     })
 }


// @Tags NginxPool
// @Summary 用id查询NginxPool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxPool true "用id查询NginxPool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Pool/findNginxPool [get]
 export const findNginxPool = (params) => {
     return service({
         url: "/Pool/findNginxPool",
         method: 'get',
         params
     })
 }


// @Tags NginxPool
// @Summary 分页获取NginxPool列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取NginxPool列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Pool/getNginxPoolList [get]
 export const getNginxPoolList = (params) => {
     return service({
         url: "/Pool/getNginxPoolList",
         method: 'get',
         params
     })
 }


// @Tags NginxPool
// @Summary 获取NginxPool列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "获取NginxPool列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Pool/getAllNginxPoolList [get]
export const getAllNginxPoolList = (params) => {
    return service({
        url: "/Pool/getAllNginxPoolList",
        method: 'get',
        params
    })
}