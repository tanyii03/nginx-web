<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="集群名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.ClusterName"></el-input>
        </el-form-item>      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增nginx集群</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"

      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="集群名称" prop="ClusterName" width="480"></el-table-column>
    <el-table-column label="集群VIP" prop="Vip" width="360"></el-table-column>
    <el-table-column label="所属部门" prop="Department" width="360"></el-table-column>
<!--    <el-table-column label="用户角色" min-width="150" prop="authorityId">-->
<!--    </el-table-column>-->
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateNginxCluster(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteNginxCluster(scope.row)">确定</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增/编辑集群">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="集群名称:">
            <el-input v-model="formData.ClusterName" clearable placeholder="请输入" ></el-input>
         </el-form-item>
        <el-form-item label="集群VIP:">
          <el-input v-model="formData.Vip" clearable placeholder="请输入" ></el-input>
        </el-form-item>
        <el-form-item label="所属部门" label-width="80px" prop="authorityId">
          <el-cascader
                  v-model="formData.authorityId"
                  :options="authOptions"
                  :show-all-levels="false"
                  :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
                  filterable
          ></el-cascader>
        </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createNginxCluster,
    deleteNginxCluster,
    deleteNginxClusterByIds,
    updateNginxCluster,
    findNginxCluster,
    getNginxClusterList,
} from "@/api/nginx_cluster";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { getAuthorityList } from "@/api/authority";
export default {
  name: "NginxCluster",
  mixins: [infoList],
  data() {
    return {
      listApi: getNginxClusterList,
      dialogFormVisible: false,
      visible: false,
      authOptions: [],
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            ClusterName:"",
            Vip:"",
            Department:"",
            AuthorityId: "",
            
      }
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
    setOptions(authData) {
      this.authOptions = [];
      this.setAuthorityOptions(authData, this.authOptions);
    },
    setAuthorityOptions(AuthorityData, optionsData) {
      AuthorityData &&
      AuthorityData.map(item => {
        if (item.children && item.children.length) {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName,
            children: []
          };
          this.setAuthorityOptions(item.children, option.children);
          optionsData.push(option);
        } else {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName
          };
          optionsData.push(option);
        }
      });
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
        const res = await deleteNginxClusterByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateNginxCluster(row) {
      const res = await findNginxCluster({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reCluster;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          ClusterName:"",
          Vip:"",
          Department:"",
      };
    },
    async deleteNginxCluster(row) {
      this.visible = false;
      const res = await deleteNginxCluster({ ID: row.ID });
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
          res = await createNginxCluster(this.formData);
          break;
        case "update":
          res = await updateNginxCluster(this.formData);
          break;
        default:
          res = await createNginxCluster(this.formData);
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
    const res = await getAuthorityList({ page: 1, pageSize: 999 });
    this.setOptions(res.data.list);
  },

};
</script>

<style>
</style>