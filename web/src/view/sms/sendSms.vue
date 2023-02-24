<template>
  <div>
    <warning-bar title="在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示" />
    <div class="gva-table-box bottom-mg-lg">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新建项目</el-button>
        <el-button type="primary" icon="plus" @click="delSmsProject">删除选中项目</el-button>
      </div>

      <el-tabs type="card" @tab-click="handleClick">
        <el-tab-pane :label="obj.project_name" :name="obj.id" v-for="obj in project_list">
          <div>
            <el-descriptions title="" border :column="4">
              <el-descriptions-item label="模板ID" >
                <el-tag size="small">{{obj.template_id}}</el-tag>
              </el-descriptions-item>
              <el-descriptions-item label="备注">{{obj.comment}}</el-descriptions-item>
              <el-descriptions-item label="条数">{{obj.row_num}}</el-descriptions-item>
              <el-descriptions-item label="参数个数">{{obj.param_num}}</el-descriptions-item>
              <el-descriptions-item label="创建时间">{{formatDate(obj.CreatedAt)}}</el-descriptions-item>
              <el-descriptions-item label="更新时间">{{formatDate(obj.UpdatedAt)}}</el-descriptions-item>
              <el-descriptions-item label="模板内容">{{win.project.sms_template.template_content}}</el-descriptions-item>

            </el-descriptions>
            
            <el-table ref="multipleTable" :data="tableData" style="width: 100%;margin-top: 20px;" tooltip-effect="dark" row-key="ID">
              <el-table-column type="selection" width="55" />
              <el-table-column align="left" label="添加日期" width="180">
                <template #default="scope">
                  <span>{{ formatDate(scope.row.CreatedAt) }}</span>
                </template>
              </el-table-column>
              <el-table-column align="left" label="手机号" prop="customerName" width="120" />
              <el-table-column align="left" :label="'参数'+num" prop="customerPhoneData" width="120" v-for="num in obj.param_num" />

              <el-table-column align="left" label="按钮组" min-width="160" fixed="right">
                <template #default="scope">
                  <el-button type="primary" link icon="edit" @click="">变更</el-button>
                  <el-popover v-model="scope.row.visible" placement="top" width="160">
                    <p>确定要删除吗？</p>
                    <div style="text-align: right; margin-top: 8px;">
                      <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                      <el-button type="primary" @click="deleteCustomer(scope.row)">确定</el-button>
                    </div>
                    <template #reference>
                      <el-button type="primary" link icon="delete" @click="scope.row.visible = true">删除</el-button>
                    </template>
                  </el-popover>
                </template>
              </el-table-column>
            </el-table>


            <div class="gva-pagination">

              <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
                @size-change="handleSizeChange" />
            </div>
          </div>
        </el-tab-pane>

      </el-tabs>

    </div>



    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="客户">
      <el-form :inline="true" :model="form" >
        <el-form-item label="项目名">
          <el-input v-model="form.project_name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="form.comment" autocomplete="off" />
        </el-form-item>
        <el-form-item label="模板ID">
          <el-input v-model="form.template_id" autocomplete="off" />
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
  
<script setup>
import {AddSmsProject, GetAllSmsProject, DelSmsProject} from '@/api/txyun/sms.js'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'

const form = ref({
  project_name: '',
  comment: '',
  template_id: '',
})
const win = ref({
  templateId:0,
  selectTab:0,
  temeplate:null,
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

const project_list = ref([]) // 项目列表



// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
}

const handleCurrentChange = (val) => {
  page.value = val
}

const handleClick = async (tab) => {
  win.value.selectTab = tab.index;
  win.value.project = project_list.value[tab.index];
  win.value.temeplate =  win.value.project.sms_template;
  console.log(win.value);
}



// 查询SmsProject
const getTableData = async () => {
  const table = await GetAllSmsProject()
  if (table.code === 0) {
    project_list.value = table.data.list
    if(table.data.list.length!=0){
      handleClick({index:0})
    }
  }
}
getTableData()

const delSmsProject = async()=>{
  const res = await DelSmsProject({"id":win.value.project.ID})
  if(res.code === 0){
    getTableData()
  }
}


const dialogFormVisible = ref(false)
const type = ref('')

const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    project_name: '',
    comment: '',
    template_id: '',
  }
}

const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      form.value.template_id = parseInt(form.value.template_id)
      res = await AddSmsProject(form.value)
      break
    case 'update':
      // res = await updateExaCustomer(form.value)
      break
    default:
      res = await AddSmsProject(form.value)
      break
  }

  if(res.code === 0){
    closeDialog()
    getTableData()
    ElMessage({type: 'success',message: '新增短信项目成功'})
  }

}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

</script>

<script>
export default {
  name: 'Customer'
}
</script>
  
<style>
.project_tag{
  height: 27px;
  font-size: 15px;
  margin-left: 2px;
  margin-right: 10px;
}
</style>
  