<template>
  <div>
    <div  class= "label">
      <span>发布列表</span>
      <el-divider></el-divider>
    </div>
    <div>
    <el-table
      :data="tableData"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column label="开始时间" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>

    <el-table-column label="结束时间" width="180">
        <template slot-scope="scope">{{scope.row.UpdatedAt|formatDate}}</template>
    </el-table-column>

    <el-table-column label="任务类型" prop="Type" width="300"></el-table-column>

    <el-table-column label="发布内容" prop="Comment" width="240"></el-table-column>
    
    <el-table-column label="用户" prop="User" width="180" align="center"></el-table-column>

    <el-table-column label="操作类型" prop="Operate" width="180" align="center"></el-table-column>

    <el-table-column label="状态" prop="Status" width="120" align="center">
        <template slot-scope="scope">
            <el-tag :type="scope.row.Status === '成功' ? 'success' : 'warning'"  disable-transitions>{{scope.row.Status}}
            </el-tag>
        </template>
    </el-table-column>

    <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="viewPublishHistory(scope.row)" size="small" type="primary" icon="el-icon-edit">查看</el-button>
        </template>
      </el-table-column>
    </el-table>

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

      <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="发布记录">
          <el-input
                  type="textarea"
                  placeholder="请输入内容"
                  v-model="viewData"
                  autosize
                  readonly
                  show-word-limit
          >
          </el-input>
      </el-dialog>


  </div>
</template>

<script>
import {
    createPublishHistory,
    deletePublishHistory,
    deletePublishHistoryByIds,
    viewPublishHistory,
    findPublishHistory,
    getPublishHistoryList
} from "@/api/nginx_domain_publish_history";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "PublishHistory",
  mixins: [infoList],
  data() {
    return {
      listApi: getPublishHistoryList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      viewData:"",
      multipleSelection: [],formData: {
            Comment:"",
            User:"",
            Type:"",
            Operate:"发布",
            Status:"成功",
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
        const res = await deletePublishHistoryByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async viewPublishHistory(row) {
      const res = await viewPublishHistory({ Uuid: row.Uuid });
      if (res.code == 0) {
        this.viewData = res.data;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        Comment:"",
        User:"",
        Type:"",
        Operate:"",
        Status:"",
      };
    },
    async deletePublishHistory(row) {
      this.visible = false;
      const res = await deletePublishHistory({ ID: row.ID });
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
          res = await createPublishHistory(this.formData);
          break;
        default:
          res = await createPublishHistory(this.formData);
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

  .label {
    font-size: 18px;
  }

</style>