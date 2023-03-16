<template>
	<div>
		<div class="gva-table-box">
			<el-tabs  v-model="activeName" tab-position="top">
        		<el-tab-pane label="我发起的请假" name="faqi" >
					<myLeaveYifaqi/>
				</el-tab-pane>

				<el-tab-pane label="待处理的请假" name="daichuli" >
					<myLeaveDaichuli/>
				</el-tab-pane>

				<el-tab-pane label="已处理的请假" name="yichuli" >
					<myLeaveYichuli/>
				</el-tab-pane>
			</el-tabs>
		</div>
	</div>
</template>

<script setup>
import myLeaveYifaqi from './myLeave-Yifaqi.vue'
import myLeaveDaichuli from './myLeave-Daichuli.vue'
import myLeaveYichuli from './myLeave-Yichuli.vue'

import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const activeName = ref("faqi")

</script>

<script>

export default {
	name: 'TeachingGrade'
}
</script>

<style>
</style>
    