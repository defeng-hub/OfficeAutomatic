<template>
  <div>
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="搜索">
          <el-input v-model="searchInfo.keyword" placeholder="名称,手机号,邮箱" />
        </el-form-item>
        <el-form-item label="部门" prop="status">
          <el-cascader
              v-model="searchInfo.authorityIds"
              :options="authOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
            />
        </el-form-item>
        <el-form-item>
          <el-select v-model="searchInfo.userTeachingGradeID" clearable  class="m-2" placeholder="Select">
              <el-option
                v-for="item in teachingoptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >查询</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addUser">新增用户</el-button>
      </div>
      <el-table
        :data="tableData"
        row-key="ID"
      >
        <el-table-column align="center" label="头像" min-width="75">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <!-- <el-table-column align="left" label="ID" min-width="50" prop="ID" /> -->
        <el-table-column align="center" label="用户名" min-width="150" prop="userName" />
        <el-table-column align="center" label="姓名" min-width="100" prop="nickName" />
        <el-table-column align="center" label="手机号" min-width="140" prop="phone" />
        <el-table-column align="center" label="邮箱" min-width="180" prop="email" />
        <el-table-column align="center" label="本职工作单位/职务" min-width="180" prop="desc0" />
        <el-table-column align="center" label="加入公司时间" min-width="140" prop="joinCompanyTime" />
        <el-table-column align="center" label="用户角色" min-width="200">
          <template #default="scope">
            <el-cascader
              v-model="scope.row.authorityIds"
              :options="authOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
              @visible-change="(flag)=>{changeAuthority(scope.row,flag,0)}"
              @remove-tag="(removeAuth)=>{changeAuthority(scope.row,false,removeAuth)}"
            />
          </template>
        </el-table-column>
        <!-- <el-table-column align="left" label="启用" min-width="150">
          <template #default="scope">
            <el-switch
              v-model="scope.row.enable"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column> -->

        <el-table-column label="操作" min-width="200" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="openEdit(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="edit" @click="openCatUserDialog(scope.row)">查看</el-button>

            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除此用户吗</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="deleteUserFunc(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete">删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>

      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- (新增/修改)用户信息弹框 -->
    <el-dialog
      v-model="addUserDialog"
      class="user-dialog"
      title="用户"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
      width="50%"
      draggable
    >
      <div style="height:60vh;overflow:auto;padding:0 0px;">
        <el-form ref="userForm" :rules="rules" :model="userInfo" 
        label-width="160px" :label-position="'top'">
          <el-form-item v-if="dialogFlag === 'add'" label="用户名" prop="userName">
            <el-input v-model="userInfo.userName" />
          </el-form-item>
          <el-form-item v-if="dialogFlag === 'add'" label="密码" prop="password">
            <el-input v-model="userInfo.password" />
          </el-form-item>

          <el-form-item label="姓名" prop="nickName">
            <el-input v-model="userInfo.nickName" />
          </el-form-item>
          <el-form-item label="手机号" prop="phone">
            <el-input v-model="userInfo.phone" />
          </el-form-item>

          <el-form-item label="邮箱" prop="email">
            <el-input v-model="userInfo.email" />
          </el-form-item>
          <el-form-item label="性别" prop="sex">
            <!-- <el-input v-model="userInfo.sex" /> -->
            <el-select v-model="userInfo.sex" class="m-2" placeholder="Select">
              <el-option
                v-for="item in sexoptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="出生日期" prop="birthdate">
            <el-input v-model="userInfo.birthdate" />
          </el-form-item>
          <el-form-item label="通信地址" prop="address">
            <el-input v-model="userInfo.address" />
          </el-form-item>
          <el-form-item label="职工号" prop="wno">
            <el-input v-model="userInfo.wno" />
          </el-form-item>
          <el-form-item label="职务" prop="zhiwu">
            <el-input v-model="userInfo.zhiwu" />
          </el-form-item>
          <el-form-item label="教学等级" prop="userTeachingGradeID">
            <el-select v-model="userInfo.userTeachingGradeID" class="m-2" placeholder="Select">
              <el-option
                v-for="item in teachingoptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="加入公司时间" prop="joinCompanyTime">
            <el-input v-model="userInfo.joinCompanyTime" />
          </el-form-item>
          <el-form-item label="参加工作时间" prop="joinWorkTime">
            <el-input v-model="userInfo.joinWorkTime" />
          </el-form-item>
          <el-form-item label="本职工作单位/职务" prop="desc0">
            <el-input v-model="userInfo.desc0" />
          </el-form-item>
          <el-form-item label="个人简历" prop="resume">
            <el-input v-model="userInfo.resume" :autosize="{ minRows: 5, maxRows: 15 }"
            type="textarea" placeholder="无" />
          </el-form-item>
          <el-form-item label="教师技能等级/职务变动情况记录" prop="desc1">
            <el-input v-model="userInfo.desc1" :autosize="{ minRows: 5, maxRows: 15 }"
            type="textarea" placeholder="无" />
          </el-form-item>
          <el-form-item label="本职工作变动情况记录" prop="desc2">
            <el-input v-model="userInfo.desc2" :autosize="{ minRows: 5, maxRows: 15 }"
            type="textarea" placeholder="无" />
          </el-form-item>

          <el-form-item label="角色" prop="authorityId">
            <el-cascader
              v-model="userInfo.authorityIds"
              style="width:100%"
              :options="authOptions"
              :show-all-levels="false"
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
            />
          </el-form-item>
          <el-form-item label="启用" prop="disabled">
            <el-switch
              v-model="userInfo.enable"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
            />
          </el-form-item>
          <el-form-item label="头像" label-width="80px" >
            <div style="display:inline-block" @click="openHeaderChange">
              <img v-if="userInfo.headerImg" alt="头像" class="header-img-box" :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg">
              <div v-else class="header-img-box">从媒体库选择</div>
            </div>
          </el-form-item>

        </el-form>

      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" v-if="dialogFlag === 'edit'" link icon="magic-stick" @click="resetPasswordFunc(userInfo)">重置密码</el-button>

          <el-button @click="closeAddUserDialog">取 消</el-button>
          <el-button type="primary" @click="enterAddUserDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />

    <el-dialog v-model="catUserDialog" class="user-dialog"
      :close-on-press-escape="false"
      width="70%" draggable>
      <el-descriptions border :column="3">
        <el-descriptions-item label="姓名" width="120">
          {{ win.winuser.nickName }}
        </el-descriptions-item>

        <el-descriptions-item label="性别" width="120">
          {{ win.winuser.sex }}
        </el-descriptions-item>
        
        <el-descriptions-item label="出生日期" width="120">
          {{ win.winuser.birthdate }}
        </el-descriptions-item>

        <el-descriptions-item label="联系电话">
          {{ win.winuser.phone }}
        </el-descriptions-item>
        
        <el-descriptions-item label="通信地址">
          {{ win.winuser.address }}
        </el-descriptions-item>

        <el-descriptions-item label="邮箱">
          {{ win.winuser.email }}
        </el-descriptions-item>

        <el-descriptions-item label="部门">

        </el-descriptions-item>

        <el-descriptions-item label="职务">{{ win.winuser.zhiwu }}</el-descriptions-item>

        <el-descriptions-item label="工号">{{ win.winuser.wno }}</el-descriptions-item>

        <el-descriptions-item label="教学技能等级">{{ win.winuser.userTeachingGrade.title }}</el-descriptions-item>
        <el-descriptions-item label="进入公司时间">{{ win.winuser.joinCompanyTime }}</el-descriptions-item>
        <el-descriptions-item label="参加工作时间">{{ win.winuser.joinWorkTime }}</el-descriptions-item>
        <el-descriptions-item label="个人简历"  span="3">
          <el-input v-model="win.winuser.resume" :autosize="{ minRows: 3, maxRows: 5 }"
            type="textarea" readonly placeholder="无" />
        </el-descriptions-item>

        <el-descriptions-item label="教师技能等级 / 职务变动情况记录"  span="3">
          <el-input v-model="win.winuser.desc1" :autosize="{ minRows: 3, maxRows: 5 }"
            type="textarea" readonly placeholder="无" />
        </el-descriptions-item>
        
        <el-descriptions-item label="本职工作变动情况记录"  span="3">
          <el-input  readonly true v-model="win.winuser.desc2" :autosize="{ minRows: 3, maxRows: 5 }"
            type="textarea" placeholder="无" />
        </el-descriptions-item>

      </el-descriptions>
      <template #footer>
          <div class="dialog-footer">
            <el-button size="large" type="primary" @click="print(1)">打 印</el-button>
            <el-button size="large" type="primary" @click="print(2)">导出Excel</el-button>
          </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'User',
}
</script>

<script setup>

import {
  getUserList,
  GetAllUserTeachingGrade,
  setUserAuthorities,
  register,
  deleteUser,
  ExportUserExcel
} from '@/api/user'

import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
const path = ref(import.meta.env.VITE_BASE_API + '/')
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
        AuthorityData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            }
            setAuthorityOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            }
            optionsData.push(option)
          }
        })
}
const print = async (parame)=>{
  if(parame == 1){
    window.print();

  }else if(parame == 2){
    let res = await ExportUserExcel({...win.value.winuser})
    if(res.code === 0){
      window.open(res.data.url)
    }
  }
}
const searchInfo = ref({
  keyword:"",
  authorityIds:null,
  userTeachingGradeID:null,
})
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const catUserDialog = ref(false)
const win = ref({})
const sexoptions = ref([
  {value: 0,label: '未选择',},
  {value: 1,label: '男',},
  {value: 2,label: '女',},
])
const teachingoptions = ref([
  {value: 1,label: '获取失败',},
])

const getTeachingoptions = async () => {
  const table = await GetAllUserTeachingGrade()
  // console.log("ccc",table)
  if (table.code === 0) {
    teachingoptions.value = []
    table.data.forEach(element => {
      teachingoptions.value.push(
        {value: element.ID,label: element.title,}
      )
    });
  }
}
getTeachingoptions()

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
const getTableData = async() => {
  let req = { 
    page: page.value,
    pageSize: pageSize.value,
    keyword:searchInfo.value.keyword,
  }
  if(searchInfo.value.authorityIds){
    req.authorityIds = searchInfo.value.authorityIds
  }
  if(searchInfo.value.userTeachingGradeID){
    req.userTeachingGradeID = searchInfo.value.userTeachingGradeID
  }
  console.log(req)
  const table = await getUserList(req)
  if (table.code === 0) {
    // console.log("userlist",table)
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
const onSubmit = async()=>{
  // console.log(searchInfo.value)
  getTableData()
}
const onReset = async()=>{
  searchInfo.value.keyword = ''
  searchInfo.value.authorityIds = []
  searchInfo.value.userTeachingGradeID = null;
  getTableData()
}


watch(() => tableData.value, () => {
  setAuthorityIds()
})

const initPage = async() => {
  getTableData()
  const res = await getAuthorityList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
    '是否将此用户密码重置为123456?',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(async() => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    user.authorityIds = user.authorities && user.authorities.map(i => {
      return i.authorityId
    })
  })
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  
  setAuthorityOptions(authData, authOptions.value)
  // console.log("权限",authOptions.value)
}

const deleteUserFunc = async(row) => {
  ElMessageBox.confirm(
    '是否将此用户删除！！',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(async() => {
    const res = await deleteUser({ id: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      row.visible = false
      await getTableData()
    }
  })

}

// 弹窗相关
const userInfo = ref({
  userName: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  authorityIds: [],
  enable: 1,
})

const rules = ref({
  userName: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入用户密码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' }
  ],
  nickName: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' }
  ],
  phone: [
    { pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入合法手机号', trigger: 'blur' },
  ],
  email: [
    { pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g, message: '请输入正确的邮箱', trigger: 'blur' },
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async() => {
  
  userForm.value.validate(async valid => {
    if (valid) {
      userInfo.value.authorityId = userInfo.value.authorityIds[0]
      const req = {
        ...userInfo.value
      }
      console.log(req)
      // return 
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
  userInfo.value = {}
}

const tempAuth = {}
const changeAuthority = async(row, flag, removeAuth) => {
  if (flag) {
    if (!removeAuth) {
      tempAuth[row.ID] = [...row.authorityIds]
    }
    return
  }
  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    authorityIds: row.authorityIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '角色设置成功' })
  } else {
    if (!removeAuth) {
      row.authorityIds = [...tempAuth[row.ID]]
      delete tempAuth[row.ID]
    } else {
      row.authorityIds = [removeAuth, ...row.authorityIds]
    }
  }
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}
const openCatUserDialog = (row)=>{
  catUserDialog.value = true;
  // console.log(row)
  win.value.winuser = row;
}

const switchEnable = async(row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: `${req.enable === 2 ? '禁用' : '启用'}成功` })
    await getTableData()
    userInfo.value.headerImg = ''
    userInfo.value.authorityIds = []
  }
}

</script>

<style lang="scss">
.user-dialog {
  .header-img-box {
    width: 150px;
    height: 150px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
.nickName{
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.pointer{
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>
