<template>
    <div class="big">
        <el-row >
            <el-col :span="4" >
                <el-card  class="box-card" >
                    <div  slot="header" class="clearfix label">
                        <span>Nginx集群列表</span>
                    </div>
                    <div>
                        <el-table :data="listData"   style="width: 100%" :show-header="false" :highlight-current-row="true">
                            <el-table-column    align="left">
                                <template slot-scope="scope">
                                    <el-link  @click="getNginxNode(scope.row)">{{ scope.row.ClusterName }}</el-link>
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                </el-card>
            </el-col>
            <el-col :span="20">
                <el-card class="box-card">
                    <div slot="header" class="clearfix label">
                        <span>Nginx节点</span>
                        <el-button style="float: right; padding: 4px 0; line-height: 0;" type="text" class="label" @click="openDialog">新增节点</el-button>
                    </div>
                    <div >
                        <el-table :data="tableData" style="width: 100%"  >
                            <el-table-column label="Node" prop="NodeName" width="480" align="center">
                            </el-table-column>
                            <el-table-column label="IP" prop="IP" width="240" align="center">
                            </el-table-column>
                            <el-table-column label="状态" prop="Status" width="240" align="center">
                            </el-table-column>
                            <el-table-column label="按钮组">
                                <template slot-scope="scope">
                                    <el-button class="table-button" @click="updateNginxNode(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
                                    <el-popover placement="top" width="160" v-model="scope.row.visible">
                                        <p>确定要删除吗？</p>
                                        <div style="text-align: right; margin: 0">
                                            <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                                            <el-button type="primary" size="mini" @click="deleteNginxNode(scope.row)">确定</el-button>
                                        </div>
                                        <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
                                    </el-popover>
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                </el-card>
            </el-col>
        </el-row>

        <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增Node">
            <el-form :model="formData" label-position="right" label-width="100px">
                <el-form-item label="Node 名称:">
                    <el-input v-model="formData.NodeName" clearable placeholder="请输入" ></el-input>
                </el-form-item>
                <el-form-item label="IP:">
                    <el-input v-model="formData.IP" clearable placeholder="请输入" ></el-input>
                </el-form-item>
                <el-form-item label="SSH端口:">
                    <el-input v-model="formData.Port" clearable placeholder="请输入" ></el-input>
                </el-form-item>
                <el-form-item label="用户名称:">
                    <el-input v-model="formData.UserName" clearable placeholder="请输入" ></el-input>
                </el-form-item>
                <el-form-item label="密码:">
                    <el-input v-model="formData.Passwd" type="password" clearable placeholder="密码不保存，只做认证" ></el-input>
                </el-form-item>
            </el-form>
            <div class="dialog-footer" slot="footer">
                <el-button @click="closeDialog">取 消</el-button>
                <el-button @click="enterDialog" type="primary">验 证</el-button>
            </div>
        </el-dialog>
    </div>

</template>

<style>
    .box-card {
        height: 720px;
    }
    .item {
        padding: 10px 0;
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

<script>
import {
    createNginxNode,
    deleteNginxNode,
    deleteNginxNodeByIds,
    updateNginxNode,
    findNginxNode,
    getNginxNodeList
} from "@/api/nginx_node";  //  此处请自行替换地址
import {getAllNginxClusterList} from "@/api/nginx_cluster";
import { formatTimeToStr } from "@/utils/date";
// import infoList from "@/mixins/infoList";
export default {
        name: "NginxNode",
        // mixins: [infoList],
        data() {
            return {
                listApi: getNginxNodeList,
                dialogFormVisible: false,
                visible: false,
                type: "",
                deleteVisible: false,
                multipleSelection: [],formData: {
                    NodeName:"",
                    IP:"",
                    Port:"22",
                    UserName:"root",
                    Passwd:"",
                    Status:"Create",
                    ClusterId:0,
                },
                ClusterId:0,
                listData:[],
                tableData:[],
            };
        },

        filters: {
            formatDate: function(time) {
                if (time != null && time != "") {
                    var date = new Date(time);
                    return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
                } else {
                    return "";
                }
            },
            formatBoolean: function(bool) {
                if (bool != null) {
                    return bool ? "是" :"否";
                } else {
                    return "";
                }
            }
        },
        mounted(){
            this.getAllNginxCluster()
        },
        methods: {
            //条件搜索前端看此方法
            onSubmit() {
                this.page = 1
                this.pageSize = 10
                this.getTableData()
            },

            handleSelectionChange(val) {
                this.multipleSelection = val
            },
            async onDelete() {
                const ids = []
                if(this.multipleSelection.length == 0){
                    this.$message({
                        type: 'warning',
                        message: '请选择要删除的数据'
                    })
                    return
                }
                this.multipleSelection &&
                this.multipleSelection.map(item => {
                    ids.push(item.ID)
                })
                const res = await deleteNginxNodeByIds({ ids })
                if (res.code == 0) {
                    this.$message({
                        type: 'success',
                        message: '删除成功'
                    })
                    this.deleteVisible = false
                    this.getTableData()
                }
            },
            async getNginxNode(row) {
                this.ClusterId = this.formData.ClusterId = row.ID
                this.getTableData()
            },
            async getTableData() {
                const res = await getNginxNodeList({ ClusterId: this.ClusterId })
                if (res.code == 0){
                    this.tableData = res.data.Nodes;
                }
            },
            async updateNginxNode(row) {
                const res = await findNginxNode({ ID: row.ID });
                this.type = "update";
                if (res.code == 0) {
                    this.formData = res.data.reNode;
                    this.dialogFormVisible = true;
                }
            },
            closeDialog() {
                this.dialogFormVisible = false;
                this.formData = {
                    NodeName:"",
                    IP:"",
                    Port:"22",
                    UserName:"root",
                    Passwd:"",
                    Status:"Create",
                    ClusterId: this.formData.ClusterId,
                };
            },
            async getAllNginxCluster(){
                const res =  await getAllNginxClusterList()
                if(res.code == 0){
                    this.listData = res.data.nginxClusters;
                }
            },
            async deleteNginxNode(row) {
                this.visible = false;
                const res = await deleteNginxNode({ ID: row.ID });
                if (res.code == 0) {
                    this.$message({
                        type: "success",
                        message: "删除成功"
                    });
                    this.getTableData();
                }
            },
            async enterDialog() {
                let res;
                switch (this.type) {
                    case "create":
                        res = await createNginxNode(this.formData);
                        break;
                    case "update":
                        res = await updateNginxNode(this.formData);
                        break;
                    default:
                        res = await createNginxNode(this.formData);
                        break;
                }
                if (res.code == 0) {
                    this.$message({
                        type:"success",
                        message:"创建/更改成功"
                    })
                    this.closeDialog();
                    this.getTableData();
                }
            },
            openDialog() {
                this.type = "create";
                this.dialogFormVisible = true;
            }
        },
        async created() {
            await this.getTableData();
        }
    };
</script>
