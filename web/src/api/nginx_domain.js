import service from '@/utils/request'

// @Tags NginxDomain
// @Summary 创建NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomain true "创建NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Domain/createNginxDomain [post]
export const createNginxDomain = (data) => {
     return service({
         url: "/Domain/createNginxDomain",
         method: 'post',
         data
     })
 }


// @Tags NginxDomain
// @Summary 删除NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomain true "删除NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Domain/deleteNginxDomain [delete]
 export const deleteNginxDomain = (data) => {
     return service({
         url: "/Domain/deleteNginxDomain",
         method: 'delete',
         data
     })
 }

// @Tags NginxDomain
// @Summary 删除NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Domain/deleteNginxDomain [delete]
 export const deleteNginxDomainByIds = (data) => {
     return service({
         url: "/Domain/deleteNginxDomainByIds",
         method: 'delete',
         data
     })
 }

// @Tags NginxDomain
// @Summary 更新NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomain true "更新NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Domain/updateNginxDomain [put]
 export const updateNginxDomain = (data) => {
     return service({
         url: "/Domain/updateNginxDomain",
         method: 'put',
         data
     })
 }


// @Tags NginxDomain
// @Summary 用id查询NginxDomain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomain true "用id查询NginxDomain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Domain/findNginxDomain [get]
 export const findNginxDomain = (params) => {
     return service({
         url: "/Domain/findNginxDomain",
         method: 'get',
         params
     })
 }


// @Tags NginxDomain
// @Summary 分页获取NginxDomain列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取NginxDomain列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Domain/getNginxDomainList [get]
 export const getNginxDomainList = (params) => {
     return service({
         url: "/Domain/getNginxDomainList",
         method: 'get',
         params
     })
 }




export const previewNginxDomain = (data) => {
    return service({
        url: "/Domain/previewNginxDomain",
        method: 'put',
        data
    })
}
