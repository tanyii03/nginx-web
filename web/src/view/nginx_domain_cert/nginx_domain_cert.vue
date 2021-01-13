<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增证书</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div>
      <el-row :gutter="330">
      <el-col :span=3 v-for="(data, index) in tableData" :key="index" class="cert-row">
<!--        <el-button  class="cert-delete-button text" @click="deleteNginxDomainCert(data)">删除</el-button>-->
      <el-card class="cert-card">
        <div slot="header"  class="clearfix text">
          <span>{{data.CertName}}</span>
          <el-button  type="text" class="cert-delete-button text" @click="deleteNginxDomainCert(data)">删除</el-button>
        </div>
        <div class="text cert-item">颁发给：{{data.Issued}}</div>
        <div class="text cert-item">DNS：{{data.Dns}}</div>
        <div class="text cert-item">有效期：{{data.Validity}}</div>
        <div class="text cert-item">截止至：{{data.Deadline}}</div>
      </el-card>
    </el-col>
    </el-row>
    </div>

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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增证书">
      <el-form :model="formData" label-position="right" label-width="78px">
         <el-form-item label="证书名称:">
            <el-input v-model="formData.CertName" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="证书文件:">
            <el-input v-model="formData.Pem" clearable placeholder="请输入证书PEM/CRT" type="textarea" :autosize="{ minRows: 5}"></el-input>
      </el-form-item>
       
         <el-form-item label="证书密钥:">
            <el-input v-model="formData.Key" clearable placeholder="请输入密钥KEY" type="textarea" :autosize="{ minRows: 5}"></el-input>
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
    createNginxDomainCert,
    deleteNginxDomainCert,
    deleteNginxDomainCertByIds,
    updateNginxDomainCert,
    findNginxDomainCert,
    getNginxDomainCertList
} from "@/api/nginx_domain_cert";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "NginxDomainCert",
  mixins: [infoList],
  data() {
    return {
      listApi: getNginxDomainCertList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            CertName:"",
            Pem:"",
            Key:"",
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
        const res = await deleteNginxDomainCertByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateNginxDomainCert(row) {
      const res = await findNginxDomainCert({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reCert;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          CertName:"",
          Pem:"",
          Key:"",
      };
    },
    async deleteNginxDomainCert(row) {
      this.$confirm('删除证书文件, 是否继续?', '', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then( async () => {
        this.visible = false;
        const res =  await deleteNginxDomainCert({ ID: row.ID });
        if (res.code == 0) {
          this.$message({
            type: "success",
            message: "删除成功"
          });
          this.getTableData();
        }
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    },





    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createNginxDomainCert(this.formData);
          break;
        case "update":
          res = await updateNginxDomainCert(this.formData);
          break;
        default:
          res = await createNginxDomainCert(this.formData);
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
  .text {
    font-size: 14px;
    font-weight: 700;
  }

  .cert-item {
    margin-bottom: 18px;
  }

  .cert-delete-button{
    font-size: 15px;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }
  .cert-card {
    width: 320px;
    height: 210px;
  }
  .cert-row{
      height: 225px;
  }

</style>