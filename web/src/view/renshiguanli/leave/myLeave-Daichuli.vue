<template>
	<div>
		<el-table ref="multipleTable" :height="500"
			:data="tableData" style="width: 100%" tooltip-effect="dark" row-key="ID">
			<el-table-column type="selection" width="55" />

			<el-table-column align="left" label="请假类型" prop="leaveType" width="120" />
			<el-table-column align="left" label="请假开始日期" width="180">
				<template #default="scope">
					<span>{{ formatDate(scope.row.beginTime) }}</span>
				</template>
			</el-table-column>
			<el-table-column align="left" label="请假结束日期" width="180">
				<template #default="scope">
					<span>{{ formatDate(scope.row.endTime) }}</span>
				</template>
			</el-table-column>
			<el-table-column align="left" label="请假内容" prop="leaveContent" width="120" />
			<el-table-column align="left" label="审核状态" prop="approval" width="120" />

			<el-table-column align="left" label="按钮组" min-width="160">
				<template #default="scope">
					<el-button type="primary" link icon="edit" @click="update(scope.row)">审批</el-button>
				</template>
			</el-table-column>
		</el-table>

		<div class="gva-pagination">
			<el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
				layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
				@size-change="handleSizeChange" />
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

<script>
export default {
	name: 'myLeaveDaichuli'
}
</script>

<script setup>
import {
	ChangeUserTeachingGrade
} from '@/api/user'

import {
	GetMyselfLeaves,
	DeleteLeaveForm,
	GetDaichuliLeaves
} from '@/api/renshiguanli/leave'

import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'

const form = ref({
	ID: 0,
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
	const table = await GetDaichuliLeaves({ page: page.value, pageSize: pageSize.value })
	console.log("ccc", table)
	if (table.code === 0) {
		tableData.value = table.data.list
		total.value = table.data.total
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


const closeDialog = () => {
	dialogFormVisible.value = false
	form.value = {
		ID: 0,
		title: '',
		shunxv: 0,
		wage: 0,
	}
}
const openDialog = () => {
	type.value = 'create'
	dialogFormVisible.value = true
	form.value = {
		ID: 0,
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