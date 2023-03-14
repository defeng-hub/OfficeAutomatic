<template>
	<div>
		<el-table ref="multipleTable" :height="500" :data="tableData" style="width: 100%" tooltip-effect="dark"
			row-key="ID">
			<el-table-column type="selection" width="30" />

			<el-table-column align="center" label="请假类型" width="100">
				<template #default="scope">
					<span>{{ LeaveTypeOptions[scope.row.leaveType] }}</span>
				</template>
			</el-table-column>

			<el-table-column align="center" label="请假开始日期" width="160">
				<template #default="scope">
					<span>{{ formatDate(scope.row.beginTime) }}</span>
				</template>
			</el-table-column>
			<el-table-column align="center" label="请假结束日期" width="160">
				<template #default="scope">
					<span>{{ formatDate(scope.row.endTime) }}</span>
				</template>
			</el-table-column>
			<el-table-column align="center" label="请假内容" prop="leaveContent" width="200" />

			<el-table-column align="center" label="审核状态" width="100">
				<template #default="scope">
					<span :class="scope.row.approval == 1 ? 'shenhetongguo' : 'shenheqita'">{{
						ApprovalOptions[scope.row.approval] }}</span>
				</template>
			</el-table-column>
			<el-table-column align="center" label="按钮组" min-width="160">

				<template #default="scope">
					<el-button type="primary" link icon="files" @click="chakanxiangqing(scope.row)">查看详情</el-button>

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

		<div class="gva-pagination">
			<el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
				layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
				@size-change="handleSizeChange" />
		</div>

		<!-- 右抽屉 -->
		<el-drawer v-if="drawershow" v-model="drawershow" class="auth-drawer" :with-header="false" size="38%" title="请假详情">
			<el-descriptions title="" border :column="2" direction="vertical">
				<el-descriptions-item label="申请人" :span="1">
					{{win.row.User.nickName}}
				</el-descriptions-item>

				<el-descriptions-item label="请假类型" :span="1">
					<el-tag size="small">{{ LeaveTypeOptions[win.row.leaveType] }}</el-tag>
				</el-descriptions-item>

				<el-descriptions-item label="申请时间" :span="2">{{ formatDate(win.row.created_at) }}</el-descriptions-item>

				<el-descriptions-item label="请假内容" :span="2">
					<el-input :value="win.row.leaveContent" :show-word-limit="true"
					 :autosize="{ minRows: 5, maxRows: 15 }" disable 
					 type="textarea" placeholder="请假内容" />
				</el-descriptions-item>

				<el-descriptions-item label="请假开始时间" :span="1">{{ formatDate(win.row.beginTime) }}</el-descriptions-item>
				<el-descriptions-item label="请假结束时间" :span="1">{{ formatDate(win.row.endTime) }}</el-descriptions-item>
				
				<el-descriptions-item label="附加图片" :span="2">
					<el-image v-if="win.row.image" :src="'http://localhost:8888/'+win.row.image"/>
					<span v-else>无</span>
				</el-descriptions-item>

				<el-descriptions-item label="审核状态" :span="1">
					{{ApprovalOptions[win.row.approval]}} 
				</el-descriptions-item>
			</el-descriptions>
			<div style="height: 50px;"></div>
		</el-drawer>
	</div>
</template>

<script>
export default {
	name: 'myLeaveYifaqi'
}
</script>

<script setup>
import {
	GetMyselfLeaves,
	DeleteLeaveForm
} from '@/api/renshiguanli/leave'
import { getDict, showDictLabel, showDict } from '@/utils/dictionary'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'


const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const win = ref({})

// 分页-start
const handleSizeChange = (val) => {
	pageSize.value = val
	getTableData()
}

const handleCurrentChange = (val) => {
	page.value = val
	getTableData()
}
// 分页-end

// 查询
const getTableData = async () => {
	const table = await GetMyselfLeaves({ page: page.value, pageSize: pageSize.value })
	console.log("ccc", table)
	if (table.code === 0) {
		tableData.value = table.data.list
		total.value = table.data.total
	}
}

getTableData()

const deleteRow = async (row) => {
	// console.log(row.ID)
	let res = await DeleteLeaveForm({ id: row.ID })
	if (res.code === 0) {
		getTableData()
		ElMessage({
			type: 'success',
			message: res.msg,
		})
	}
}



// 字典-start
const ApprovalOptions = ref({}) // 审核状态
const LeaveTypeOptions = ref({}) //请假类型

const loadDict = async () => {
	let Options = await getDict('Approval')
	ApprovalOptions.value = await showDict(Options)

	Options = await getDict('LeaveType')
	LeaveTypeOptions.value = await showDict(Options)

	// console.log("字典",ApprovalOptions.value,LeaveTypeOptions.value)
}
loadDict()
// 字典-end

// 右抽屉
const drawershow = ref(false)
const chakanxiangqing = async (row) => {
	drawershow.value = true;
	// console.log(row)
	win.value.row = row;
	console.log(win.value)
}

</script>

<style lang="scss" scoped>
.shenhetongguo {
	color: #5cb87a;
	font-size: 18px;
}

.shenheqita {
	font-size: 15px;
}
</style>