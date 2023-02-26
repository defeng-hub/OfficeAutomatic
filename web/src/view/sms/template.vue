<template>
	<div>
		<warning-bar title="模板状态，其中: 0表示审核通过，1表示审核中，-1表示审核未通过或审核失败。" />
		<div class="gva-table-box">
			<div class="gva-btn-list">
				<el-button type="primary" icon="refresh" @click="refreshData">刷新短信模板列表</el-button>
			</div>
			<el-table ref="multipleTable" stripepageSize highlight-current-row :data="tableData" style="width: 100%"
				tooltip-effect="dark" row-key="template_id">
				<el-table-column type="selection" width="50" />
				<el-table-column align="center" label="模板ID" prop="template_id" width="100" />
				<el-table-column align="center" label="模板内容" prop="template_content" width="550" />
				<el-table-column align="center" label="模板状态" prop="status_code" width="80" />
				<el-table-column align="center" label="模板名称" prop="template_name" width="180" />
				<el-table-column align="center" label="提交审核时间" prop="create_time" width="120"></el-table-column>


				<el-table-column align="center" label="操作" min-width="100">
					<template #default="scope">
						<el-button type="primary" link icon="edit" @click="openDialog(scope.row)">发短信</el-button>

						<el-popover v-model="scope.row.visible" placement="top" width="160">
							<p>确定要删除吗？</p>
							<div style="text-align: right; margin-top: 8px;">
								<el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
								<el-button type="primary" @click="deleteSms(scope.row)">确定</el-button>
							</div>

							<!-- 原始样式 -->
							<template #reference>
								<el-button type="primary" link icon="delete"
									@click="scope.row.visible = true">删除</el-button>
							</template>
						</el-popover>

					</template>
				</el-table-column>

			</el-table>
		</div>

		<el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="发送短信">
			<el-tag>{{ '模板ID:' + form.template_id }}</el-tag>
			<el-row :gutter="20">
				<el-col v-for="obj, idx in form.tpl_params" :span="8">
					<el-input v-model="form.tpl_params[idx]" :placeholder="`参数${idx + 1}`" />
				</el-col>
			</el-row>

			<el-row>
				<el-input v-model="form.tpl_phones" :autosize="{ minRows: 5, maxRows: 15 }" type="textarea"
					placeholder="手机号列表" />
			</el-row>

			<template #footer>
				<div class="dialog-footer">
					<el-button @click="form.tpl_phones = &quot;&quot;">清空</el-button>
					<el-button type="warning" @click="closeDialog">取 消</el-button>

					<el-button type="primary" @click="enterDialog">确 定</el-button>
				</div>
			</template>
		</el-dialog>

	</div>
</template>
  
<script setup>
import { UpdateTemplates, GetSmsList, SendSms } from '@/api/txyun/sms'
import { ref } from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 分页开始
const handleSizeChange = (val) => {
	pageSize.value = val
	getTableData()
}

const handleCurrentChange = (val) => {
	page.value = val
	getTableData()
}
// 分页结束

// 查
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

// 更新数据文档
const refreshData = async () => {
	const res = await UpdateTemplates()
	if (res.code === 0) {
		ElMessage({
			// 	'success' | 'warning' | 'info' | 'error'
			type: 'success',
			message: '更新成功',
			duration: 3000,
		})
		getTableData()
	} else {
		ElMessage({
			type: 'error',
			message: res.msg,
			duration: 3000,
		})
	}
}

// 删
const deleteSms = async (row) => {
	console.log(row)
}

// dialog弹窗
const form = ref({
	template_id: '', // 模板id
	template_content: '', // 模板内容
	tpl_params: [],
	parms_num: 0,
	tpl_phones: '', // 手机号列表
})
const dialogFormVisible = ref(false)

const closeDialog = async () => {
	dialogFormVisible.value = false
}
const openDialog = (row) => {
	dialogFormVisible.value = true
	form.value.parms_num = findParmsNum(row.template_content)
	form.value.template_content = row.template_content
	form.value.template_id = row.template_id
	form.value.tpl_params = []

	for (let index = 0; index < form.value.parms_num; index++) {
		form.value.tpl_params.push('')
	}
}
const enterDialog = async () => {
	// console.log(form.value)
	const res = await SendSms(form.value)
	// console.log(res)
	if (res.code === 0) {
		dialogFormVisible.value = false
		ElMessage({ type: 'success', message: '发信成功' })
		form.value.tpl_phones = "";
	}
}

// 传入string,获取{1},{2}, 这个的最大数量
const findParmsNum = (content) => {
	if (content.indexOf('{5}') != -1) {
		if (content.indexOf('{6}') != -1) {
			if (content.indexOf('{7}') != -1) {
				if (content.indexOf('{8}') != -1) {
					if (content.indexOf('{9}') != -1) {
						return 9
					} else { return 8 }
				} else { return 7 }
			} else { return 6 }
		} else { return 5 }
	} else {
		if (content.indexOf('{4}') != -1) { return 4 } else {
			if (content.indexOf('{3}') != -1) { return 3 } else {
				if (content.indexOf('{2}') != -1) { return 2 } else {
					if (content.indexOf('{1}') != -1) { return 1 } else { return 0 }
				}
			}
		}
	}
}

</script>
  
<script>
export default {
	name: 'Sms',
	data() {
		return {
		}
	},
	created() {
		console.log('created')
	},
	methods: {
	}
}

</script>
  
<style></style>
  