<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="集群名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.PoolName"></el-input>
        </el-form-item>              
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增Pool集群</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"

      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="集群名称" prop="PoolName" width="400"></el-table-column>
    
    <el-table-column label="负载均衡策略" prop="Policy" width="150"></el-table-column>
    
    <el-table-column label="长连接数" prop="Keepalive" width="120"></el-table-column> 
    
    <el-table-column label="健康检查类型" prop="CheckType" width="120"></el-table-column> 
    
    <el-table-column label="检查时间间隔" prop="CheckInterval" width="120"></el-table-column> 
    
    <el-table-column label="检查超时时间" prop="CheckTimeout" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateNginxPool(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteNginxPool(scope.row)">确定</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增集群">
      <el-form :model="formData" label-position="right" label-width="180px">
         <el-form-item label="集群名称:">
            <el-input v-model="formData.PoolName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="负载均衡策略:">
            <el-select v-model="formData.Policy" clearable placeholder="请选择">
              <el-option v-for="(item, index) in PolicyOptions" :key="index" :label="item.label"
                         :value="item.value" :disabled="item.disabled"></el-option>
            </el-select>
      </el-form-item>
       
         <el-form-item label="长连接数:"><el-input v-model.number="formData.Keepalive" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="健康检查类型:">
            <el-input v-model="formData.CheckType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="检查时间间隔:"><el-input v-model.number="formData.CheckInterval" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="检查超时时间:"><el-input v-model.number="formData.CheckTimeout" clearable placeholder="请输入"></el-input>
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
    createNginxPool,
    deleteNginxPool,
    deleteNginxPoolByIds,
    updateNginxPool,
    findNginxPool,
    getNginxPoolList
} from "@/api/nginx_pool";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "NginxPool",
  mixins: [infoList],
  data() {
    return {
      listApi: getNginxPoolList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      PolicyOptions:[{
        "label":"round-robin",
        "value":"round-robin"
      },{
        "label":"ip_hash",
        "value":"ip_hash"
      }],
      formData: {
            PoolName:"",
            Policy:"round-robin",
            Keepalive:20,
            CheckType:"tcp",
            CheckInterval:3000,
            CheckTimeout:3000,

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
        const res = await deleteNginxPoolByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateNginxPool(row) {
      const res = await findNginxPool({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rePool;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          PoolName:"",
          Policy:"round-robin",
          Keepalive:20,
          CheckType:"tcp",
          CheckInterval:3000,
          CheckTimeout:3000,
      };
    },
    async deleteNginxPool(row) {
      this.visible = false;
      const res = await deleteNginxPool({ ID: row.ID, PoolName: row.PoolName });
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
          res = await createNginxPool(this.formData);
          break;
        case "update":
          res = await updateNginxPool(this.formData);
          break;
        default:
          res = await createNginxPool(this.formData);
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

<style>
</style>