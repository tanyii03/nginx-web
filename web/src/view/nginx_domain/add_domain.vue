<template>
    <div class="big">
        <el-row>
            <el-col :span="16">
                <el-card  class="card_item box-card">
                    <div  slot="header" class="clearfix label">
                        <span>新增站点</span>
                    </div>
                    <div>
                        <el-form :model="formData" label-position="right" label-width="80px">
                        <el-form-item label="站点名称:">
                            <el-input v-model="formData.DomainName" clearable placeholder="请输入" ></el-input>
                        </el-form-item>
                        <el-form-item label="域名:">
                            <el-input v-model="formData.Domain" clearable placeholder="请输入" ></el-input>
                        </el-form-item>
                        <el-form-item label="端口:">
                            <el-input v-model="formData.Port" clearable placeholder="请输入" ></el-input>
                        </el-form-item>
                        <el-form-item label="Pool集群:">
                        <el-select  name="Pool集群:" style="width: 480px" v-model="formData.Pool" filterable placeholder="请选择">
                            <el-option
                                    v-for="item in poolOption"
                                    :key="item.ID"
                                    :label="item.PoolName"
                                    :value="item.ID">
                            </el-option>
                        </el-select>
                        </el-form-item>
                        <el-form-item label="Nginx集群:">
                        <el-select  label="Nginx集群:" style="width: 480px" v-model="formData.Nginx" filterable placeholder="请选择">
                            <el-option
                                    v-for="item in nginxOption"
                                    :key="item.ID"
                                    :label="item.ClusterName"
                                    :value="item.ID">
                            </el-option>
                        </el-select>
                        </el-form-item>
                        </el-form>
                    </div>
                </el-card>
            </el-col>
            <el-col :span="8">
                <el-card  class="card_item box-card">
                    <div  slot="header" class="clearfix label">
                        <span>Https配置</span>
                    </div>
                    <div>
                        <el-form :model="formData" label-position="right" label-width="120px">
                        <el-form-item label="配置Https:">
                            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.UseHttps" clearable ></el-switch>
                        </el-form-item>
                        <el-form-item label="Https路径类型:" v-show="formData.UseHttps">
                            <el-input v-model="formData.HttpsRewrite" clearable placeholder="请输入" ></el-input>
                        </el-form-item>
                        <el-form-item label="Https端口:"    v-show="formData.UseHttps">
                            <el-input v-model="formData.HttpsPort" clearable placeholder="请输入" ></el-input>
                        </el-form-item>
                        <el-form-item label="Https证书:"    v-show="formData.UseHttps">
                            <el-select   style="width: 240px" v-model="formData.HttpsCert" filterable placeholder="请选择">
                                <el-option
                                        v-for="item in certOption"
                                        :key="item.ID"
                                        :label="item.CertName"
                                        :value="item.CertName">
                                </el-option>
                            </el-select>
                        </el-form-item>
                        </el-form>
                    </div>
                </el-card>
            </el-col>
        </el-row>
        <el-button type="primary"  style="display:block;margin:0 auto" @click="save">提交</el-button>
    </div>

</template>

<script>
import {
    createNginxDomain,
} from "@/api/nginx_domain";  //  此处请自行替换地址
import {getAllNginxClusterList} from "@/api/nginx_cluster";
import {getAllNginxPoolList} from "@/api/nginx_pool";
import { getNginxDomainCertList } from "@/api/nginx_domain_cert";
export default {
    name:"addDomain",
    data(){
        return {
            type: "",
            nginxOption:[],
            poolOption:[],
            certOption:[],
            formData: {
                DomainName:"",
                Domain:"",
                Port:"80",
                Pool:"",
                NginxCluster:"",
                UseHttps:false,
                HttpsRewrite:"http,https",
                HttpsPort:"443",
                HttpsCert:"",

            }
        };
    },
    mounted(){
        this.getAllNginxCluster()
        this.getAllNginxPool()
        this.getAllNginxDomainCert()
    },
    methods:{
        async save() {
            let res;
            res = await createNginxDomain(this.formData);
            if (res.code == 0) {
                this.$message({
                    type:"success",
                    message:"创建/更改成功"
                });
                this.formData = {
                    DomainName:"",
                        Domain:"",
                        Port:"80",
                        Pool:"",
                        NginxCluster:"",
                        UseHttps:false,
                        HttpsRewrite:"",
                        HttpsPort:"443",
                        HttpsCert:"",
                }
            }
        },
        async getAllNginxPool(){
            const res =  await getAllNginxPoolList()
            if(res.code == 0){
                this.poolOption = res.data.nginxPools
            }
        },
        async getAllNginxCluster(){
            const res =  await getAllNginxClusterList()
            if(res.code == 0){
                this.nginxOption = res.data.nginxClusters;
            }
        },
        async getAllNginxDomainCert(){
            let res;
            res = await getNginxDomainCertList({Page: 1, PageSize:100});
            if (res.code == 0){
                this.certOption = res.data.list
            }
        },
    }
};
</script>


<style>
    .box-card {
        height: 720px;
    }
    .label {
        font-size: 18px;
    }

    .clearfix:before,

    .clearfix:after {
        display: table;
        content: "";
    }
    .clearfix:after {
        clear: both
    }

</style>

