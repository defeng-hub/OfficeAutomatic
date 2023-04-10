<template>
	<el-row :gutter="20" style="margin-top: 0;padding-top: 0;">

		<!-- 左边 -->
		<el-col :span="16" class="left">
			<div>
				<div class="row_top">
					<div class="gva-btn-list" style="line-height: 30px;font-size: 16px;">
						模板管理
					</div>
					<div class="gva-btn-list">
						<el-button type="primary" icon="plus" @click="dialogFormVisible2 = true">创建模板</el-button>
						<el-button type="warning" icon="delete" size="default" @click="">删除该模板</el-button>
					</div>
				</div>

				<div class="line"></div>
				<el-carousel :interval="10000" type="card" height="570px">
					<el-carousel-item v-for="item,idx in list_template" :key="item.id">
						<el-image :src="`${path}/`+item.imgSrc" 
						fit="cover" style="width: 100%; height: 570px"
						@click="openProject(item)"/>
					</el-carousel-item>
				</el-carousel>
			</div>
		</el-col>

		<!-- 右边画笔 -->
		<el-col :span="8">
			<div class="right">
				<div class="row_top">
					<div class="gva-btn-list" style="line-height: 30px;font-size: 16px;">
						画笔管理
					</div>
					<div class="gva-btn-list">
						<el-button type="primary" icon="plus" @click="dialogFormVisible = true">创建画笔</el-button>
					</div>
				</div>
				<div class="line"></div>
				<el-table :data="list_branch" style="width: 100%">
					<el-table-column align="center" prop="ID" label="画笔ID" width="100" />
					<el-table-column align="center" prop="fontSize" label="字体大小" width="100" />
					<el-table-column align="center" label="字体颜色" width="100">
						<template #default="scope">
							<el-color-picker v-model="scope.row.fontColor" show-alpha/>
						</template>
          </el-table-column>
					<el-table-column fixed="right" label="操作">
						<template #default="scope">
							<el-button link type="primary" size="small" @click="deleteBranchConfirm(scope.row.ID)">删除</el-button>
						</template>
					</el-table-column>
				</el-table>
			</div>
		</el-col>

		<!-- 全部弹出层 -->
		<div>
			<!-- 新增画笔弹窗 -->
			<el-dialog v-model="dialogFormVisible" title="新增画笔">
				<el-form ref="authorityForm" :model="branchForm" label-width="100px">
					<el-form-item label="字体路径" prop="path">
						<el-input v-model="branchForm.path" />
					</el-form-item>
					<el-form-item label="字体大小" prop="authorityName">
						<el-input-number v-model="branchForm.fontSize" />
					</el-form-item>
					<el-form-item label="字体颜色" prop="authorityName">
						<el-color-picker v-model="branchForm.fontColor" show-alpha />
					</el-form-item>
				</el-form>
				<template #footer>
					<div class="dialog-footer">
						<el-button @click="dialogFormVisible = false">取 消</el-button>
						<el-button type="primary" @click="branchFormConfirm">确 定</el-button>
					</div>
				</template>
			</el-dialog>

			<!-- 新增画笔弹窗 -->
			<el-dialog v-model="dialogFormVisible2" title="新增画笔">
				<el-form :model="prjectForm" label-width="100px">
					<el-form-item label="标题" prop="title">
						<el-input v-model="prjectForm.title" />
					</el-form-item>
					<el-form-item label="附件">
						<el-upload :action="`${path}/fileUploadAndDownload/upload?noSave=1`" :before-upload="checkFile"
							:headers="{ 'x-token': userStore.token }" :on-error="uploadError" :on-success="uploadSuccess"
							:limit="1"
							:show-file-list="true" class="upload-btn">
							<el-button type="primary">普通上传</el-button>
						</el-upload>
					</el-form-item>
				</el-form>
				<template #footer>
					<div class="dialog-footer">
						<el-button @click="dialogFormVisible2 = false">取 消</el-button>
						<el-button type="primary" @click="prjectFormConfirm">确 定</el-button>
					</div>
				</template>
			</el-dialog>
		</div>
	</el-row>



</template>
  
<script>
export default {
	name: 'DrawingProject'
}
</script>

<script setup>
import { CreateProject,GetAllProject,GetAllBranch,CreateBranch,
	DeleteBranch } from '@/api/drawing/project'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'
import { useUserStore } from '@/pinia/modules/user'
const path = ref(import.meta.env.VITE_BASE_API)

const list_template = ref([]) //项目列表
const list_branch = ref([]) //画笔列表
const dialogFormVisible = ref(false)
const dialogFormVisible2 = ref(false)
const branchForm = ref({
	path:"uploads/file/9c5d46e22ae2db1f0d838611657c14d9_20230410213023.ttf",//画笔默认路径
	fontSize:24,
	fontColor:'rgba(255, 206, 102, 1)'
})
const prjectForm = ref({
	title:"",
	imgSrc:"",
})

const initData = async() => {
  const res = await GetAllProject()
  if (res.code === 0) {
    list_template.value =  res.data;
  }

	const res2 = await GetAllBranch()
  if (res2.code === 0) {
    list_branch.value =  res2.data;
  }
	console.log(list_template.value)
	console.log(list_branch.value)
}
initData()

// 新增画笔
const branchFormConfirm = async()=>{
	if (branchForm.value.path == ""){
		ElMessage({
      type: 'warning',
      message: '路径为空'
    })
		return 
	}
	const res = await CreateBranch(branchForm.value)
	console.log(res)
  if (res.code === 0) {
		dialogFormVisible.value = false
		initData()
  }
}
// 删除画笔
const deleteBranchConfirm = async(id)=>{
	const res = await DeleteBranch({id:id})
	console.log(res)
  if (res.code === 0) {
		initData()
  }
}

// 文件上传
const userStore = useUserStore()
const fullscreenLoading = ref(false)
const emit = defineEmits(['on-success'])
const checkFile = (file) => {
  fullscreenLoading.value = true
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 15
  if (!isJPG && !isPng) {
    ElMessage.error('上传图片只能是 jpg或png 格式!')
    fullscreenLoading.value = false
  }
  if (!isLt2M) {
    ElMessage.error('未压缩的上传图片大小不能超过 15MB')
    fullscreenLoading.value = false
  }
  return (isPng || isJPG) && isLt2M
}
const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}
const uploadSuccess = (res) => {
  const { data } = res
	if(res.code == 0){
		prjectForm.value.imgSrc = data.file.url
		ElMessage({type: 'success',message: '上传成功'})
	}
  if (data.file) {
    emit('on-success', data.file.url)
  }
}

// 新增项目
const prjectFormConfirm = async()=>{
	console.log(prjectForm.value)
	if (prjectForm.value.title == "" || prjectForm.value.imgSrc == ""){
		ElMessage({
      type: 'warning',
      message: '请填写完整'
    })
		return 
	}
	const res = await CreateProject(prjectForm.value)
	console.log(res)
  if (res.code === 0) {
		dialogFormVisible2.value = false
		initData()
  }
}

// 点击图片
const openProject = async(item)=>{
	console.log(item)
}
</script>

<style lang="scss" scoped>
.demo-image__error .image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: var(--el-fill-color-light);
  color: var(--el-text-color-secondary);
  font-size: 30px;
}
.line{
	border-bottom: 1px solid #e8e8e8;
	margin-bottom: 20px;
}
.row_top{
	display: flex;
	justify-content: space-between;
}
.left{
	padding: 30px;
	height:680px;
  padding-bottom: 10px;
  padding-top: 16px;
  background-color: #fff;
  border-radius: 2px;
  box-sizing: border-box;
  box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
}
.right{
	padding: 30px;
	height:680px;
  padding-bottom: 10px;
  padding-top: 16px;
  background-color: #fff;
  border-radius: 2px;
  box-sizing: border-box;
  box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
}
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 700px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>