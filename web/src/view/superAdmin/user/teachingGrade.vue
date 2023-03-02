<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <el-table ref="multipleTable" :data="tableData" style="width: 100%" tooltip-effect="dark" row-key="ID">
        <el-table-column type="selection" width="55" />


        <el-table-column align="left" label="标题" prop="title" width="120" />
        <el-table-column align="left" label="时薪" prop="wage" width="120" />
        <el-table-column align="left" label="顺序" prop="shunxv" width="120" />
        <el-table-column align="left" label="修改日期" width="180">
          <template #default="scope">
            <span>{{ formatDate(scope.row.updated_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组" min-width="160">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="update(scope.row)">变更</el-button>
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="deleteRow(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete" @click="scope.row.visible = true">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>


    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="客户">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="教学等级">
          <el-input v-model="form.title" autocomplete="off" />
        </el-form-item>
        <el-form-item label="时薪" type="number">
          <el-input v-model.number="form.wage" type="" autocomplete="off" />
        </el-form-item>
        <el-form-item label="顺序" type="number">
          <el-input v-model.number="form.shunxv" autocomplete="off" />
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
import {
  GetAllUserTeachingGrade,
  DeleteUserTeachingGrade,
  ChangeUserTeachingGrade
} from '@/api/user'

import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'

const form = ref({
  ID:0,
  title: '',
  shunxv: 0,
  wage: 0,
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await GetAllUserTeachingGrade()
  console.log("ccc",table)
  if (table.code === 0) {
    tableData.value = table.data
  }
}

getTableData()

const dialogFormVisible = ref(false)
const type = ref('')
const update = async (row) => {
  dialogFormVisible.value = true;
  type.value = 'update'
  form.value = {
    ID: row.ID,
    title: row.title,
    wage: row.wage,
    shunxv: row.shunxv,
  }
}
const deleteRow = async (row) => {
  // console.log(row.ID)
  let res = await DeleteUserTeachingGrade({ id: row.ID })
  if (res.code === 0) {
    getTableData()
    ElMessage({
      type: 'success',
      message: res.msg,
    })
  }
}

const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    ID:0,
    title: '',
    shunxv: 0,
    wage: 0,
}
}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
  form.value = {
    ID:0,
    title: '',
    shunxv: 0,
    wage: 0,
}
}
const enterDialog = async () => {
  let res
  // console.log(form.value)
  switch (type.value) {
    case 'create':
      res = await ChangeUserTeachingGrade(form.value);
      break
    case 'update':
      res = await ChangeUserTeachingGrade(form.value);
      break
    default:
      res = await ChangeUserTeachingGrade(form.value);
      break
  }

  if (res.code === 0) {
    closeDialog()
    getTableData()
  }
}


</script>
  
<script>

export default {
  name: 'TeachingGrade'
}
</script>
  
<style></style>
  