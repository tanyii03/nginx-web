<template>
<div>
    <el-form :model="formData" label-position="right" label-width="200px">
             <el-form-item label="集群名称:">
                <el-input v-model="formData.PoolName" clearable placeholder="请输入" ></el-input>
          </el-form-item>

          <el-form-item label="负载均衡策略:">
            <el-select v-model="formData.Policy" clearable placeholder="请选择">
                <el-option v-for="(item, index) in formData.PolicyOptions" :key="index" :label="item.label"
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
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createNginxPool,
    updateNginxPool,
    findNginxPool
} from "@/api/nginx_pool";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "NginxPool",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            PoolName:"",
            Policy:"",
            Keepalive:20,
            CheckType:"tcp",
            CheckInterval:3000,
            CheckTimeout:3000,
            PolicyOptions:[{
                "label":"轮询",
                "value":"round-robin"
            },{
                "label":"权重",
                "value":"weight"
            },{
                "label":"哈希",
                "value":"ip_hash"
            }],
            
      }
    };
  },
  methods: {
    async save() {
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findNginxPool({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rePool
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
}
};
</script>

<style>
</style>