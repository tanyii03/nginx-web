import service from '@/utils/request'

// @Tags PublishHistory
// @Summary 创建PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublishHistory true "创建PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /History/createPublishHistory [post]
export const createPublishHistory = (data) => {
     return service({
         url: "/History/createPublishHistory",
         method: 'post',
         data
     })
 }


// @Tags PublishHistory
// @Summary 删除PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublishHistory true "删除PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /History/deletePublishHistory [delete]
 export const deletePublishHistory = (data) => {
     return service({
         url: "/History/deletePublishHistory",
         method: 'delete',
         data
     })
 }

// @Tags PublishHistory
// @Summary 删除PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /History/deletePublishHistory [delete]
 export const deletePublishHistoryByIds = (data) => {
     return service({
         url: "/History/deletePublishHistoryByIds",
         method: 'delete',
         data
     })
 }

// @Tags PublishHistory
// @Summary 更新PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublishHistory true "更新PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /History/updatePublishHistory [put]
 export const viewPublishHistory = (data) => {
     return service({
         url: "/History/viewPublishHistory",
         method: 'post',
         data
     })
 }


// @Tags PublishHistory
// @Summary 用id查询PublishHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PublishHistory true "用id查询PublishHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /History/findPublishHistory [get]
 export const findPublishHistory = (params) => {
     return service({
         url: "/History/findPublishHistory",
         method: 'get',
         params
     })
 }


// @Tags PublishHistory
// @Summary 分页获取PublishHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PublishHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /History/getPublishHistoryList [get]
 export const getPublishHistoryList = (params) => {
     return service({
         url: "/History/getPublishHistoryList",
         method: 'get',
         params
     })
 }