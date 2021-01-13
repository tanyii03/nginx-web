import service from '@/utils/request'

// @Tags PoolNode
// @Summary 创建PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PoolNode true "创建PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PNode/createPoolNode [post]
export const createPoolNode = (data) => {
     return service({
         url: "/PNode/createPoolNode",
         method: 'post',
         data
     })
 }


// @Tags PoolNode
// @Summary 删除PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PoolNode true "删除PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PNode/deletePoolNode [delete]
 export const deletePoolNode = (data) => {
     return service({
         url: "/PNode/deletePoolNode",
         method: 'delete',
         data
     })
 }

// @Tags PoolNode
// @Summary 删除PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PNode/deletePoolNode [delete]
 export const deletePoolNodeByIds = (data) => {
     return service({
         url: "/PNode/deletePoolNodeByIds",
         method: 'delete',
         data
     })
 }

// @Tags PoolNode
// @Summary 更新PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PoolNode true "更新PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PNode/updatePoolNode [put]
 export const updatePoolNode = (data) => {
     return service({
         url: "/PNode/updatePoolNode",
         method: 'put',
         data
     })
 }


// @Tags PoolNode
// @Summary 用id查询PoolNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PoolNode true "用id查询PoolNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PNode/findPoolNode [get]
 export const findPoolNode = (params) => {
     return service({
         url: "/PNode/findPoolNode",
         method: 'get',
         params
     })
 }


// @Tags PoolNode
// @Summary 分页获取PoolNode列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PoolNode列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PNode/getPoolNodeList [get]
 export const getPoolNodeList = (params) => {
     return service({
         url: "/PNode/getPoolNodeList",
         method: 'get',
         params
     })
 }