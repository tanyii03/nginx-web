<template>
    <div class="big">
        <el-row :gutter="1"  >
        <el-col :span="4" >
            <el-card  class="card_item box-card">
            <div class="item">
                <el-select v-model="ClusterId" filterable placeholder="请选择" @change="getAllNginxDomain">
                    <el-option
                            v-for="item in nginxOption"
                            :key="item.ID"
                            :label="item.ClusterName"
                            :value="item.ID">
                    </el-option>
                </el-select>
            </div>
            <div calss="item">
                <el-select v-model="DomainId" filterable placeholder="请选择" @change="getNginxDomain">
                    <el-option
                            v-for="item in domainOption"
                            :key="item.ID"
                            :label="item.DomainName"
                            :value="item.ID">
                    </el-option>
                </el-select>
            </div>

            </el-card>
        </el-col>
        <el-col :span="20">
            <el-card  class="card_item box-card">
            <el-tabs v-model="activeName"  class="navTab">
                <el-tab-pane label="基本配置" name="first">
                    <el-row v-show="selectDomain">
                        <el-col :span="16">
                            <el-card  class="card_item box-card-2">
                                <div  slot="header" class="clearfix label">
                                    <span>编辑站点</span>
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
                            <el-card  class="card_item box-card-2">
                                <div  slot="header" class="clearfix label">
                                    <span>Https配置</span>
                                </div>
                                <div>
                                    <el-form :model="formData" label-position="right" label-width="120px">
                                        <el-form-item label="配置Https:">
                                            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.UseHttps" clearable ></el-switch>
                                        </el-form-item>
                                        <el-form-item label="Https路径类型:" v-show="formData.UseHttps" >
                                            <el-input v-model="formData.HttpsRewrite" clearable placeholder="请输入" ></el-input>
                                        </el-form-item>
                                        <el-form-item label="Https端口:" v-show="formData.UseHttps" >
                                            <el-input v-model="formData.HttpsPort" clearable placeholder="请输入" ></el-input>
                                        </el-form-item>
                                        <el-form-item label="Https证书:" v-show="formData.UseHttps" >
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
                </el-tab-pane>
                <el-tab-pane label="转发规则" name="second">
                <el-row :gutter="1"  v-show="selectDomain">
                    <el-col :span="24" >
                    <el-card  class="card_item box-card-2">
                    <div  slot="header" class="clearfix label">
                        <span>转发规则列表</span>
                        <el-button  style="float: right; padding: 4px 6px; line-height: 0;" type="text" class="label" @click="openDialog">新增规则</el-button>
                    </div>
                        <el-table :data="tableData" border stripe style="width: 100%" tooltip-effect="dark">
                            <el-table-column label="匹配类型" prop="MatchType"></el-table-column>

                            <el-table-column label="转发路径" prop="Path"></el-table-column>

                            <el-table-column label="Https类型" prop="HttpsRewrite"></el-table-column>

                            <el-table-column label="操作">
                                <template slot-scope="scope">
                                    <el-button class="table-button" @click="updateDomainRule(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
                                    <el-popover placement="top" width="160" v-model="scope.row.visible">
                                        <p>确定要删除吗？</p>
                                        <div style="text-align: right; margin: 0">
                                            <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                                            <el-button type="primary" size="mini" @click="deleteDomainRule(scope.row)">确定</el-button>
                                        </div>
                                        <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
                                    </el-popover>
                                </template>
                            </el-table-column>
                        </el-table>

                        <el-pagination
                                :current-page="page"
                                :page-size="pageSize"
                                :page-sizes="[10, 30, 50, 100]"
                                :style="{float:'right',padding:'20px'}"
                                :total="total"
                                @current-change="handleCurrentChange"
                                @size-change="handleSizeChange"
                                layout="total, sizes, prev, pager, next, jumper"
                        ></el-pagination>

                    </el-card>
                    </el-col>
                </el-row>
                </el-tab-pane>
                <el-tab-pane label="发布记录" name="third" >
                    <el-row   v-show="selectDomain">
                    <div class="block ">
                        <el-timeline >
                            <el-timeline-item
                                    v-for="(activity, index) in publishHistory"
                                    :key="index"
                                    :timestamp="activity.CreatedAt">
                                {{activity.User}} {{activity.Comment}}
                            </el-timeline-item>
                        </el-timeline>
                    </div>
                    </el-row>
                </el-tab-pane>
            </el-tabs>
                <div v-show="selectDomain">
                <div class="domainSave">
                    <el-button type="primary" plain @click="updateNginxDomain">保存</el-button>
                </div>
                <div class="domainPreview">
                    <el-button type="success" plain @click="previewNginxDomain">预览</el-button>
                </div>
                <div class="domainPublish">
                    <el-button type="warning" plain @click="publishNginxDomain">发布</el-button>
                </div>
                <div class="domainDelete">
                    <el-button type="danger" plain @click="deleteNginxDomain">删除</el-button>
                </div>
                </div>
            </el-card>
        </el-col>
        </el-row>

        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增转发规则">
            <el-form :model="ruleData" label-position="right" label-width="150px">
                <el-form-item label="匹配类型:">
                    <el-input v-model="ruleData.MatchType" clearable placeholder="请输入" ></el-input>
                </el-form-item>

                <el-form-item label="路径/正则表达式:">
                    <el-input v-model="ruleData.Path" clearable placeholder="请输入" ></el-input>
                </el-form-item>
            </el-form>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button @click="enterDialog" type="primary">确 定</el-button>
            </div>
        </el-dialog>

        <el-dialog :before-close="closeRuleDialog" :visible.sync="dialogRuleFormVisible" title="修改转发规则">
            <el-form :model="ruleData" label-position="right" label-width="150px">
                <el-form-item label="匹配类型:">
                    <el-select style="width: 480px" v-model="ruleData.MatchType" placeholder="请选择">
                        <el-option
                                v-for="item in matchTypeOptions"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="路径/正则表达式:">
                    <el-input v-model="ruleData.Path" clearable placeholder="请输入" ></el-input>
                </el-form-item>
            </el-form>

            <div  class="clearfix label">
                <span>指令列表</span>
                <el-button style="float: right; padding: 4px 6px; line-height: 2;" type="text" @click="openInsDialog"  icon="el-icon-circle-plus">新增指令</el-button>
            </div>

            <el-table :data="insTableData" border stripe style="width: 100%" tooltip-effect="dark">
                <el-table-column label="指令类型" prop="InsType"></el-table-column>
                <el-table-column label="指令参数" prop="InsValue"></el-table-column>
                <el-table-column label="操作">
                    <template slot-scope="scope">
                        <el-popover placement="top" width="160" v-model="scope.row.visible">
                            <p>确定要删除吗？</p>
                            <div style="text-align: right; margin: 0">
                                <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                                <el-button type="primary" size="mini" @click="deleteDomainRuleIns(scope.row)">确定</el-button>
                            </div>
                            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
                        </el-popover>
                    </template>
                </el-table-column>
            </el-table>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeRuleDialog">取 消</el-button>
                <el-button @click="enterRuleDialog" type="primary">确 定</el-button>
            </div>
        </el-dialog>

        <el-dialog :before-close="closeInsDialog" :visible.sync="dialogInsFormVisible" title="新增指令">
            <el-form :model="insData" label-position="right" label-width="150px">
                <el-form-item label="指令类型:">
                    <el-select    v-model="insData.InsType" placeholder="请选择">
                        <el-option
                                v-for="item in insTypeOptions"
                                :key="item"
                                :label="item"
                                :value="item">
                        </el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="Pool选择:" v-if="insData.InsType=='proxy_pass'">
                    <el-select  style="width: 480px" v-model="insData.InsValue" filterable placeholder="请选择">
                        <el-option
                                v-for="item in poolOption"
                                :key="item.ID"
                                :label="item.PoolName"
                                :value="item.PoolName">
                        </el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="指令参数:" v-else>
                    <el-input v-model="insData.InsValue" clearable placeholder="请输入" ></el-input>
                </el-form-item>
            </el-form>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeInsDialog">取 消</el-button>
                <el-button @click="enterInsDialog" type="primary">确 定</el-button>
            </div>
        </el-dialog>

        <el-dialog :before-close="closePublishDialog" :visible.sync="dialogPublishFormVisible" title="发布版本">
            <el-form :model="publishData" label-position="right" label-width="150px">
                <el-form-item label="版本信息:">
                    <el-input v-model="publishData.Comment" clearable placeholder="请输入" ></el-input>
                </el-form-item>
            </el-form>
            <template>
                <el-alert
                        title="错误终止发布"
                        type="info"
                        center
                        show-icon>
                </el-alert>
            </template>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closePublishDialog">取 消</el-button>
                <el-button @click="enterPublishDialog" type="primary">确 定</el-button>
            </div>
        </el-dialog>

        <el-dialog :before-close="closePreviewDialog" :visible.sync="dialogPreviewFormVisible" title="预览配置">
            <el-input
                    type="textarea"
                    placeholder="请输入内容"
                    v-model="previewData"
                    autosize
                    readonly
                    show-word-limit
            >
            </el-input>
        </el-dialog>

    </div>
</template>

<style>
    .box-card {
        height: 720px;
    }
    .box-card-2 {
        height: 620px;
    }
    .item {
        padding: 10px 0;
    }
    .label {
        font-size: 18px;
    }
    .domainSave{
        position: absolute;right:200px;top:30px;
        font-weight: 600;
        font-size: 14px;
    }
    .domainPreview{
        position: absolute;right:150px;top:30px;
        font-weight: 600;
        font-size: 14px;
    }
    .domainPublish{
        position: absolute;right:100px;top:30px;
        font-weight: 600;
        font-size: 14px;
    }
    .domainDelete{
        position: absolute;right:50px;top:30px;
        font-weight: 600;
        font-size: 14px;
    }

</style>

<script>
import {getAllNginxClusterList} from "@/api/nginx_cluster";
import {getAllNginxPoolList} from "@/api/nginx_pool";
import {
    deleteNginxDomain,
    updateNginxDomain,
    findNginxDomain,
    getNginxDomainList,
    previewNginxDomain
} from "@/api/nginx_domain";  //
import {
    createDomainRule,
    deleteDomainRule,
    updateDomainRule,
    findDomainRule,
    getDomainRuleList
} from "@/api/nginx_domain_rule";  //
import {
    createDomainRuleIns,
    deleteDomainRuleIns,
    getDomainRuleInsList
} from "@/api/domain_rule_ins";  //
import {
    createPublishHistory,
    getPublishHistoryList
} from "@/api/nginx_domain_publish_history";  //
import { getNginxDomainCertList } from "@/api/nginx_domain_cert";
import infoList from "@/mixins/infoList";
export default {
        name:"Domain",
        mixins: [infoList],
        data() {
            return {
                listApi: getDomainRuleList,
                activeName: 'first',
                nginxOption: [],
                domainOption: [],
                poolOption:[],
                certOption:[],
                ClusterId: '',
                DomainId:'',
                RuleId:'',
                selectDomain: false,
                dialogFormVisible:false,
                dialogRuleFormVisible:false,
                dialogInsFormVisible:false,
                dialogPublishFormVisible:false,
                dialogPreviewFormVisible:false,
                ruleData:{
                    MatchType:"PREFIX",
                    Path:"",
                    HttpsRewrite:"Http,Https",
                    InsNum:0,
                    DomainId:"",
                },
                publishData: {
                    Comment:"",
                    User:"",
                    DomainId:"",
                    Type:"",
                    Operate:"",
                    Status:"执行",
                    Config:"",
                    ClusterId:"",
                },
                previewData:"",
                formData: {
                    DomainName:"",
                    Domain:"",
                    Port:"80",
                    Pool:"",
                    Nginx:"",
                    UseHttps:false,
                    HttpsRewrite:"Http,Https",
                    HttpsPort:"443",
                    HttpsCert:"",

                },
                matchTypeOptions:[
                    "EXACT",
                    "PREFIX",
                    "REGEX_CASE_INSENSITIVE",
                    "REGEX_CASE_SENSITIVE",
                ],
                insTableData:[],
                insData: {
                    InsType:"",
                    InsValue:"",
                    RuleId:"",
                },
                insTypeOptions:[
                    "proxy_pass",
                    "rewrite",
                    "return",
                    "set",
                    "proxy_set_header",
                    "allow",
                    "deny",
                    "expires",
                    "custom",
                ],
                publishHistory: [],
            };
        },
        mounted(){
            this.getAllNginxCluster()
            this.getAllNginxPool()
            this.getAllNginxDomainCert()
        },
        methods: {
            async getAllNginxCluster(){
                const res =  await getAllNginxClusterList()
                if(res.code == 0){
                    this.nginxOption = res.data.nginxClusters;
                }
            },
            async getAllNginxDomain(){
                const res =  await getNginxDomainList({ ID: this.ClusterId })
                this.DomainId='';
                if(res.code == 0){
                    this.domainOption = res.data.nginxDomains;
                }
            },
            async getNginxDomain(){
                const res =  await findNginxDomain({ ID: this.DomainId })
                this.searchInfo = { DomainId: this.DomainId }
                if(res.code == 0){
                    this.formData = res.data.reDomain;
                    this.selectDomain = true;
                    this.getTableData();
                    this.getDomainPublishHistory();
                }
            },
            async getAllNginxPool(){
                const res =  await getAllNginxPoolList()
                if(res.code == 0){
                    this.poolOption = res.data.nginxPools
                }
            },
            async updateNginxDomain() {
                const res = await updateNginxDomain(this.formData);
                if (res.code == 0) {
                    this.$message({
                        type: "success",
                        message: "更新成功"
                    });
                    this.getNginxDomain();
                }
            },
            async previewNginxDomain() {
                const res = await previewNginxDomain(this.formData);
                if (res.code == 0) {
                    this.previewData = res.data
                    this.dialogPreviewFormVisible = true;
                }
            },
            async closePreviewDialog(){
                this.dialogPreviewFormVisible = false;
                this.previewData = "";
            },
            async deleteNginxDomain() {
                this.selectDomain = false;
                const res = await deleteNginxDomain({ ID: this.formData.ID, Nginx: this.ClusterId, DomainName: this.formData.DomainName });
                if (res.code == 0) {
                    this.$message({
                        type: "success",
                        message: "删除成功"
                    });
                    this.DomainId='';
                    this.getAllNginxDomain()
                }
            },
            openDialog() {
                this.type = "create";
                this.dialogFormVisible = true;
            },
            closeDialog() {
                this.dialogFormVisible = false;
                this.ruleData = {
                };
            },
            async enterDialog() {
                let res;
                this.ruleData.DomainId = this.DomainId
                res = await createDomainRule(this.ruleData);
                if (res.code == 0) {
                    this.$message({
                        type:"success",
                        message:"创建/更改成功"
                    })
                    this.closeDialog();
                    this.getTableData();
                }
            },
            async deleteDomainRule(row) {
                this.visible = false;
                const res = await deleteDomainRule({ ID: row.ID });
                if (res.code == 0) {
                    this.$message({
                        type: "success",
                        message: "删除成功"
                    });
                    this.getTableData();
                }
            },
            async updateDomainRule(row) {
                const res = await findDomainRule({ ID: row.ID });
                this.type = "update";
                if (res.code == 0) {
                    this.ruleData = res.data.reRule;
                    this.dialogRuleFormVisible = true;
                    this.RuleId = this.insData.RuleId = row.ID
                    this.getInsTableData()
                }
            },
            closeRuleDialog() {
                this.dialogRuleFormVisible = false;
                this.ruleData = {
                    MatchType:"",
                    Path:"",
                    HttpsRewrite:"",
                    InsNum:0,
                    DomainId:this.DomainId,
                };
                this.RuleId = "";
            },
            async enterRuleDialog() {
                let res;
                res = await updateDomainRule(this.ruleData);
                if (res.code == 0) {
                    this.$message({
                        type:"success",
                        message:"创建/更改成功"
                    })
                    this.closeRuleDialog();
                    this.getTableData();
                }
            },
            closeInsDialog() {
                this.dialogInsFormVisible = false;
                this.insData = {
                    InsType:"",
                    InsValue:"",
                };
            },
            async deleteDomainRuleIns(row) {
                this.visible = false;
                const res = await deleteDomainRuleIns({ ID: row.ID });
                if (res.code == 0) {
                    this.$message({
                        type: "success",
                        message: "删除成功"
                    });
                    this.getInsTableData();
                }
            },
            openInsDialog() {
                this.dialogInsFormVisible = true;
            },
            async enterInsDialog() {
                let res;
                res = await createDomainRuleIns(this.insData);
                if (res.code == 0) {
                    this.$message({
                        type:"success",
                        message:"创建/更改成功"
                    })
                    this.closeInsDialog();
                    this.getInsTableData();
                }
            },
            async getInsTableData(){
                let res;
                res = await getDomainRuleInsList({Page: 1, PageSize:100, RuleId: this.RuleId});
                if (res.code == 0){
                    this.insTableData = res.data.list
                }
            },
            async getDomainPublishHistory(){
              let res;
              res = await getPublishHistoryList({Page: 1, PageSize:10, DomainId: this.DomainId});
              if (res.code == 0){
                this.publishHistory = res.data.list
              }
            },
            async publishNginxDomain(){
                this.dialogPublishFormVisible = true;
            },
            closePublishDialog() {
                this.dialogPublishFormVisible = false;
                this.publishData = {
                    Comment:"",
                    User:"",
                    DomainId:"",
                    Type:"",
                    Operate:"",
                    Status:"执行",
                    Config:"",
                    ClusterId:"",
                };
            },
            async enterPublishDialog() {
                let res;
                this.publishData.DomainId  = this.DomainId;
                // this.publishData.ClusterId = this.ClusterId
                this.publishData.Operate = "发布";
                // this.publishData.Type = "部署站点: " + this.formData.DomainName;
                // this.publishData.Config = "Domain."+ this.formData.DomainName + ".conf"
                res = await createPublishHistory(this.publishData);
                if (res.code == 0) {
                    this.$message({
                        type:"success",
                        message:"发布成功"
                    })
                    this.closePublishDialog();
                    this.getDomainPublishHistory();
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