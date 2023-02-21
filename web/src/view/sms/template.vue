<template>
	<div>
		<warning-bar title="模板状态，其中: 0表示审核通过，1表示审核中，-1表示审核未通过或审核失败。" />

		<div class="gva-table-box">
			<div class="gva-btn-list">
				<el-button type="primary" icon="refresh" @click="refreshData">刷新短信模板列表</el-button>
			</div>
			<el-table ref="multipleTable" stripe  highlight-current-row :data="tableData" style="width: 100%" tooltip-effect="dark" row-key="template_id">
				<el-table-column type="selection" width="50" />
				<el-table-column align="center" label="模板ID"   prop="template_id" width="100" />
				<el-table-column align="center" label="模板内容" prop="template_content" width="400" />
				<el-table-column align="center" label="模板状态" prop="status_code" width="120" />
				<el-table-column align="center" label="模板名称" prop="template_name" width="120" />
				<el-table-column align="center" label="提交审核时间" prop="create_time" width="120" />
				
				<el-table-column align="center" label="按钮组" min-width="100">
					<template #default="scope">
            <el-button type="primary" link icon="edit" >发短信</el-button>

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



	</div>
</template>
  
<script setup>
import {
	deleteExaCustomer,
} from '@/api/customer'

import { UpdateTemplates,GetSmsList } from '@/api/txyun/sms'
import { ref } from 'vue'

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
	const table = await GetSmsList({ page: page.value, pageSize: pageSize.value })
	if (table.code === 0) {
		tableData.value = table.data.list
		total.value = table.data.total
		page.value = table.data.page
		pageSize.value = table.data.pageSize
	}
}

getTableData()


const type = ref('')


const deleteCustomer = async (row) => {
	row.visible = false
	const res = await deleteExaCustomer({ ID: row.ID })
	if (res.code === 0) {
		ElMessage({
			type: 'success',
			message: '删除成功'
		})
		if (tableData.value.length === 1 && page.value > 1) {
			page.value--
		}
		getTableData()
	}
}
</script>
  

<script>
import WarningBar from '@/components/warningBar/warningBar.vue'
import { UpdateTemplates,GetSmsList } from '@/api/txyun/sms'
import { ElMessage } from 'element-plus'
export default {
	name: 'sms',
	data() {
		return {

		}
	},
	methods: {
		async refreshData() {
			let res = await UpdateTemplates();
			if (res.code === 0) {
				ElMessage({
					// 	'success' | 'warning' | 'info' | 'error'
					type: 'success',
					message: '更新成功',
					duration: 3000,
				})
			} else {
				ElMessage({
					type: 'error',
					message: res.msg,
					duration: 3000,
				})
			}
		}
	},
	created() {
		console.log("aaa")
	}
}

</script>
  
<style></style>
  