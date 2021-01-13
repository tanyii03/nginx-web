import service from '@/utils/request'

// @Tags NginxCluster
// @Summary 创建NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxCluster true "创建NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cluster/createNginxCluster [post]
export const createNginxCluster = (data) => {
     return service({
         url: "/Cluster/createNginxCluster",
         method: 'post',
         data
     })
 }


// @Tags NginxCluster
// @Summary 删除NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxCluster true "删除NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Cluster/deleteNginxCluster [delete]
 export const deleteNginxCluster = (data) => {
     return service({
         url: "/Cluster/deleteNginxCluster",
         method: 'delete',
         data
     })
 }

// @Tags NginxCluster
// @Summary 删除NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Cluster/deleteNginxCluster [delete]
 export const deleteNginxClusterByIds = (data) => {
     return service({
         url: "/Cluster/deleteNginxClusterByIds",
         method: 'delete',
         data
     })
 }

// @Tags NginxCluster
// @Summary 更新NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxCluster true "更新NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Cluster/updateNginxCluster [put]
 export const updateNginxCluster = (data) => {
     return service({
         url: "/Cluster/updateNginxCluster",
         method: 'put',
         data
     })
 }


// @Tags NginxCluster
// @Summary 用id查询NginxCluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxCluster true "用id查询NginxCluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Cluster/findNginxCluster [get]
 export const findNginxCluster = (params) => {
     return service({
         url: "/Cluster/findNginxCluster",
         method: 'get',
         params
     })
 }


// @Tags NginxCluster
// @Summary 分页获取NginxCluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取NginxCluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cluster/getNginxClusterList [get]
 export const getNginxClusterList = (params) => {
     return service({
         url: "/Cluster/getNginxClusterList",
         method: 'get',
         params
     })
 }


// @Tags NginxCluster
// @Summary 分页获取NginxCluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取NginxCluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cluster/getNginxClusterList [get]
export const getAllNginxClusterList = () => {
    return service({
        url: "/Cluster/getAllNginxClusterList",
        method: 'get',
    })
}


