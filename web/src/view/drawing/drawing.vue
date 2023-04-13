<template>
	<el-row>
		<el-col :span="6">
			<div class="elementy">
				<div class="row_top">
					<div class="gva-btn-list" style="line-height: 30px;font-size: 16px;">
						参数列表
					</div>
					<div class="gva-btn-list">
						<el-button type="primary" icon="plus" @click="dialogFormVisible2 = true">新增参数</el-button>
					</div>
				</div>
				<div class="line"></div>
				<div>
					<el-table :data="params_list" height="570" style="width: 100%;"
						:header-cell-style="{ 'background': '#e0e6f2', 'color': '#1e55ba', 'line-hight': '30px', 'text-align': 'center' }"
						:header-row-style="{ 'line-height': '30px', }" @row-click="paramClick"
						>
						<!-- :row-style="rowStyle" -->
						<template #empty>
							<span class="tableEmpty" style="width: 202px;height:198px;"></span>
						</template>
						<el-table-column align="center" prop="ID" label="ID" width="44" />
						<el-table-column align="center" prop="title" label="名称" />

						<el-table-column align="center" label="自动填充" width="100">
							<template #default="scope">
								<el-icon v-if="scope.row.autoIncrement == 1" color="#00933c">
									<SuccessFilled />
								</el-icon>
								<el-icon v-if="scope.row.autoIncrement == 0" color="#e22018">
									<CircleCloseFilled />
								</el-icon>
							</template>
						</el-table-column>

					</el-table>
				</div>

			</div>
		</el-col>

		<el-col :span="13">
			<div class="elementy">
				<div class="gva-btn-list" style="line-height: 30px;text-align: center;">
					<div style="text-align: center;width: 100%;font-size: 18px;">模板名：{{ project.title }}</div>
				</div>
				<div class="line"></div>
				<div style="background-color: #e0e6f2;">
					<el-image v-if="project.imgSrc" :src="`${path}/` + project.imgSrc" fit="scale-down"
						style="height: 570px;width: 100%;" @click="" />
				</div>
			</div>
		</el-col>

		<el-col :span="5">
			<div class="elementy">
				<div class="gva-btn-list" style="line-height: 30px;text-align: center;">
					<span v-if="winParam.title">参数名：<span style="font-weight: 600;">{{ winParam.title }}</span></span>
					<span v-else>参数详情</span>
				</div>
				<div class="line"></div>

				<div>
					<el-form :model="winParam" label-width="50px" label-position="top">

						<el-form-item label="" prop="x">
							<el-input type="number" v-model.number="winParam.x" :min="0">
								<template #prepend>X坐标</template>
							</el-input>
							<!-- <el-input-number v-model="winParam.x" :min="0" /> -->
						</el-form-item>
						<el-form-item label="" prop="y">
							<el-input type="number" v-model.number="winParam.y" :min="0">
								<template #prepend>Y坐标</template>
							</el-input>
							<!-- <el-input-number v-model="winParam.y" :min="0" /> -->
						</el-form-item>
						<el-form-item required label="画笔id" prop="brushId">
							<el-input type="number" v-model="winParam.brushId" @click="changeBranch" :min="0">
								<template #prepend>
									<el-button @click="changeBranch">更换画笔</el-button>
								</template>
							</el-input>
						</el-form-item>
						<el-form-item label="填充内容" prop="text" v-if="!winParam.autoIncrementBool">
							<el-input type="textarea" v-model="winParam.text" :disabled="winParam.autoIncrementBool" />
						</el-form-item>

						<el-form-item required label="自动填充(一般不用手动控制!)" prop="increment.winNum" v-if="winParam.autoIncrementBool">
							<el-input type="number" v-model="winParam.increment.winNum" disabled :min="0">
								<template #prepend>
									{{winParam.text}}
								</template>
							</el-input>
						</el-form-item>

					</el-form>
				</div>

				<div class="gva-btn-list">
					<el-button type="primary" icon="plus" @click="saveParam()">保存</el-button>
					<el-button type="danger" icon="delete" @click="deleteParam()">删除该参数</el-button>
				</div>

				<div class="btn-end">
					<el-button type="primary" icon="plus" @click="createImageconfim()">生成图片</el-button>
				</div>
			</div>
		</el-col>

		<!-- 全部弹出层 -->
		<div>
			<!-- 新增画笔弹窗 -->
			<el-dialog v-model="dialogFormVisible" title="新增画笔">
				<el-form ref="authorityForm" :model="branchForm" label-width="100px">
					<el-form-item label="字体名" prop="name">
						<el-input v-model="branchForm.name" />
					</el-form-item>
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

			<!-- 新增参数弹窗 -->
			<el-dialog v-model="dialogFormVisible2" title="新增参数">
				<el-form :model="paramsForm" label-width="100px">
					<el-form-item required label="参数名" prop="title">
						<el-input v-model="paramsForm.title" />
					</el-form-item>
					<el-form-item label="X坐标" prop="x">
						<el-input-number v-model="paramsForm.x" :min="0" />
					</el-form-item>
					<el-form-item label="Y坐标" prop="y">
						<el-input-number v-model="paramsForm.y" :min="0" />
					</el-form-item>
					<el-form-item label="填充内容" prop="text">
						<el-input v-model="paramsForm.text" :disabled="paramsForm.autoIncrementBool" />
					</el-form-item>
					<el-form-item label="是否自增" prop="autoIncrementBool">
						<el-switch v-model="paramsForm.autoIncrementBool" />
					</el-form-item>
					<el-form-item required label="初始值" prop="winNum" v-if="paramsForm.autoIncrementBool">
						<el-input-number v-model="paramsForm.winNum" :min="0" />
					</el-form-item>
				</el-form>

				<template #footer>
					<div class="dialog-footer">
						<el-button @click="dialogFormVisible2 = false">取 消</el-button>
						<el-button type="primary" @click="dialogFormConfirm">确 定</el-button>
					</div>
				</template>
			</el-dialog>


			<!-- 画笔管理弹窗 -->
			<el-dialog v-model="dialogFormVisible3" title="画笔管理">
				<el-table :data="list_branch" style="width: 100%" border height="400">
					<el-table-column align="center" prop="name" label="画笔名" width="130" />
					<!-- <el-table-column align="center" prop="fontSize" label="字体大小" width="100" /> -->

					<el-table-column align="center" label="字体大小" width="100">
						<template #default="scope">
							<el-input type="number" v-model.number="scope.row.fontSize"  :min="0" />
						</template>
					</el-table-column>

					<el-table-column align="center" label="字体颜色" width="100">
						<template #default="scope">
							<el-color-picker v-model="scope.row.fontColor" show-alpha />
						</template>
					</el-table-column>
					<el-table-column fixed="right" label="操作">
						<template #default="scope">
							<el-button link type="danger" size="small" @click="deleteBranchConfirm(scope.row.ID)">删除</el-button>
							<el-button link type="success" size="small" @click="saveBranchConfirm(scope.row)">保存</el-button>
							<el-button link type="primary" size="small" @click="selectBranchConfirm(scope.row.ID)">选用画笔</el-button>
						</template>
					</el-table-column>
				</el-table>
				<div class="gva-btn-list" style="margin-top: 20px;">
					<el-button type="primary" icon="plus" @click="dialogFormVisible = true">创建画笔</el-button>
				</div>
			</el-dialog>

		</div>
	</el-row>
</template>

<script>
export default {
	name: 'Drawing'
}
</script>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import {
	CreateBranch, DeleteBranch, GetAllBranch, ChangeParamBranch,ChangeParam,ChangeBrush,
	GetProjectById, GetParamsByProjectId, CreateParam,DeleteParam,CreateImage
} from '@/api/drawing/project'
import { ElMessage } from 'element-plus'
import { useRoute, useRouter } from 'vue-router'
import { closeThisPage } from '@/utils/closeThisPage'

const path = ref(import.meta.env.VITE_BASE_API)
const router = useRouter()
const route = useRoute()
const params_list = ref([])

const project = ref({})

const initData = async () => {
	if (route.query.projectid) { } else {
		router.replace({ name: 'drawingProject' })
		return
	}
	const res = await GetProjectById({ ID: parseInt(route.query.projectid) })
	console.log(res)
	if (res.code != 0) {
		// 没拿到项目, 直接退出
		closeThisPage()
		return
	} else {
		project.value = res.data;
		const res2 = await GetParamsByProjectId({ ID: res.data.ID });
		console.log(res2)
		if (res2.code === 0) {
			params_list.value = res2.data;
		}
		// 获取画笔
		const res3 = await GetAllBranch()
		if (res3.code === 0) {
			list_branch.value = res3.data;
		}

	}
}
initData()

const dialogFormVisible = ref(false)
const dialogFormVisible2 = ref(false)
const dialogFormVisible3 = ref(false)

// 参数详情 表单
const paramsForm = ref({
	projectId: 0,
	title: "",
	text: "", //填充内容
	x: 20,
	y: 20,
	autoIncrementBool: false,
	autoIncrement: 0,//是否自增 0不自增 1自增填充, 选用1就不会填充本身的Text， 使用winnum
	winNum: 1,
	brushId: 0, //画笔id
})

const dialogFormConfirm = async () => {
	paramsForm.value.projectId = project.value.ID;
	console.log(paramsForm.value)
	const res = await CreateParam(paramsForm.value);
	// console.log(res)
	if (res.code === 0) {
		dialogFormVisible2.value = false
		initData()
	}
}


const winParam = ref({})
// 参数被点击
const paramClick = async (row, column, event) => {
	// console.log(row.ID)
	winParam.value = row;
	if (row.autoIncrement == 0) {
		winParam.value.autoIncrementBool = false;
	} else {
		winParam.value.autoIncrementBool = true;
	}
}
// 更改选择行颜色
const rowStyle = async ({row})=>{
	console.log(winParam.value.ID === row.ID)
	if (winParam.value.ID === row.ID) {
    return {"background-color": "#F7EDED"};
	}
	return { cursor: 'pointer' };

}

// 画笔相关
const list_branch = ref([]) //画笔列表
const branchForm = ref({
	path: "uploads/file/9c5d46e22ae2db1f0d838611657c14d9_20230410213023.ttf",//画笔默认路径
	fontSize: 24,
	name:"微软雅黑",
	fontColor: 'rgba(0, 0, 0, 1)'
})
// 更换画笔
const changeBranch = async () => {
	// console.log("aaa")
	dialogFormVisible3.value = true;
}

// 删除画笔
const deleteBranchConfirm = async (id) => {
	const res = await DeleteBranch({ id: id })
	console.log(res)
	if (res.code === 0) {
		ElMessage({
			type: 'success',
			message: '删除成功'
		})
		initData()
	}
}
// 选择画笔
const selectBranchConfirm = async (id) => {
	console.log( winParam.value.ID,id)
	const res = await ChangeParamBranch({ paramId: winParam.value.ID, branchId: id })
	// console.log(res)
	if (res.code === 0) {
		winParam.value.brushId = id;
		dialogFormVisible3.value = false;
		initData()
		ElMessage({
			type: 'success',
			message: '更换画笔成功'
		})
	}
}
// 更改画笔
const saveBranchConfirm = async (row) => {
	if (row.path == "" || row.fontSize == 0 || row.fontColor=="") {
		ElMessage({
			type: 'warning',
			message: '参数为空'
		})
		return
	}
	const res = await ChangeBrush(row)
	console.log(res)
	if (res.code === 0) {
		initData()
		ElMessage({
			type: 'success',
			message: '更新成功'
		})
	}
}

// 新增画笔
const branchFormConfirm = async () => {
	if (branchForm.value.path == "") {
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
// 保存画笔
const saveParam = async()=>{
	if(winParam.value.ID){}else{
		return
	}
	// console.log(winParam.value)
	const res = await ChangeParam(winParam.value)
	// console.log(res)
	if(res.code ===0){
		ElMessage({
			type: 'success',
			message: '更新成功'
		})
	}
}

const deleteParam = async()=>{
	if(winParam.value.ID){}else{
		return
	}
	// console.log(winParam.value)
	const res = await DeleteParam({id:winParam.value.ID})
	// console.log(res)
	if(res.code ===0){
		ElMessage({
			type: 'success',
			message: '删除成功'
		})
		initData()
	}
}

const createImageconfim = async()=>{
	const res = await CreateImage({id:project.value.ID});
	console.log(res)

}

</script>

<style lang="scss" scoped>
.row_top {
	display: flex;
	justify-content: space-between;
}

.line {
	border-bottom: 1px solid #e8e8e8;
	margin-bottom: 20px;
}


.elementy {
	margin: 6px;
	padding: 30px;
	height: 680px;
	padding-bottom: 10px;
	padding-top: 16px;
	background-color: #fff;
	border-radius: 2px;
	box-sizing: border-box;
	box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
}
.btn-end{
	position: absolute;
	bottom: 50px;
	display: flex;
	justify-content: center;
	align-items:center;
}

</style>