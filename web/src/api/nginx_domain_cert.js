import service from '@/utils/request'

// @Tags NginxDomainCert
// @Summary 创建NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomainCert true "创建NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cert/createNginxDomainCert [post]
export const createNginxDomainCert = (data) => {
     return service({
         url: "/Cert/createNginxDomainCert",
         method: 'post',
         data
     })
 }


// @Tags NginxDomainCert
// @Summary 删除NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomainCert true "删除NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Cert/deleteNginxDomainCert [delete]
 export const deleteNginxDomainCert = (data) => {
     return service({
         url: "/Cert/deleteNginxDomainCert",
         method: 'delete',
         data
     })
 }

// @Tags NginxDomainCert
// @Summary 删除NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Cert/deleteNginxDomainCert [delete]
 export const deleteNginxDomainCertByIds = (data) => {
     return service({
         url: "/Cert/deleteNginxDomainCertByIds",
         method: 'delete',
         data
     })
 }

// @Tags NginxDomainCert
// @Summary 更新NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomainCert true "更新NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Cert/updateNginxDomainCert [put]
 export const updateNginxDomainCert = (data) => {
     return service({
         url: "/Cert/updateNginxDomainCert",
         method: 'put',
         data
     })
 }


// @Tags NginxDomainCert
// @Summary 用id查询NginxDomainCert
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NginxDomainCert true "用id查询NginxDomainCert"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Cert/findNginxDomainCert [get]
 export const findNginxDomainCert = (params) => {
     return service({
         url: "/Cert/findNginxDomainCert",
         method: 'get',
         params
     })
 }


// @Tags NginxDomainCert
// @Summary 分页获取NginxDomainCert列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取NginxDomainCert列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Cert/getNginxDomainCertList [get]
 export const getNginxDomainCertList = (params) => {
     return service({
         url: "/Cert/getNginxDomainCertList",
         method: 'get',
         params
     })
 }