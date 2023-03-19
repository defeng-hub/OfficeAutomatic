<template>
	<div>
		<div style="display: flex;">
			<el-card shadow="never" style="width: 48%;">
				<template #header>
					<div class="card-header">
						<span>一、请假表单填写</span>
					</div>
				</template>

				<el-form label-position="right" label-width="80px" :model="Leaveform">
					<el-form-item label="请假类型" required>
						<el-select v-model="Leaveform.leaveType" placeholder="请选择活动区域">
							<el-option label="事假" :value="0"></el-option>
							<el-option label="调休" :value="1"></el-option>
							<el-option label="病假" :value="2"></el-option>
							<el-option label="年假" :value="3"></el-option>
							<el-option label="产假" :value="4"></el-option>
							<el-option label="陪产假" :value="5"></el-option>
							<el-option label="婚假" :value="6"></el-option>
							<el-option label="例假" :value="7"></el-option>
							<el-option label="丧假" :value="8"></el-option>
							<el-option label="哺乳假" :value="9"></el-option>
						</el-select>
					</el-form-item>

					<!-- 时间勾选 -->
					<el-form-item label="请假时间" required>
						<el-date-picker v-model="Leaveform.selectTime" type="datetimerange" start-placeholder="请假开始日期"
							end-placeholder="请假结束日期">
						</el-date-picker>
					</el-form-item>

					<el-form-item label="请假天数" required>
						<el-input-number v-model="Leaveform.daynum" :min="0" placeholder="自动计算请假时长" :precision="1" :step="0.5"></el-input-number>
					</el-form-item>

					<el-form-item label="请假事由" required>
						<el-input v-model="Leaveform.leaveContent" placeholder="至少输入5个字" :autosize="{ minRows: 5, maxRows: 15 }"
							type="textarea"></el-input>
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
			</el-card>
			
			<div style="width: 2%;"></div>

			<el-card shadow="never" style="width: 48%;">
				<template #header>
					<div class="card-header">
						<span>二、选择审核人</span>
					</div>
				</template>
				<el-timeline>
					<el-timeline-item>
						<el-row style="padding-top:3px;">
							<el-col :span="11">
								<span style="line-height:26px;">部门领导审批</span>
							</el-col>
							<el-col :span="11">
								<div class="selectuser">
									<el-select v-model="Leaveform.shenpiUserID" placeholder="请选择活动区域">
										<el-option label="人事部门--bytest" :value="3"></el-option>
										<el-option label="教研部门--雯雯" :value="4"></el-option>
										<el-option label="IT部门--王德丰" :value="1"></el-option>

										<el-option label="总管--孙老师" :value="2"></el-option>
									</el-select>
								</div>
							</el-col>
						</el-row>
					</el-timeline-item>

					<el-timeline-item>
						<el-row style="padding-top:0px;">
							<el-col :span="11">
								<span style="line-height:26px;">若超过三天主管审批</span>
							</el-col>
							<el-col :span="11">
								<div class="selectuser">
									<div class="selectuser">
									<el-select disabled v-model="Leaveform.shenpiUser2ID" placeholder="请选择活动区域">
										<el-option label="无需二次审批" :value="0"></el-option>
										<el-option label="总管--孙老师" :value="2"></el-option>
									</el-select>
								</div>
								</div>
							</el-col>
						</el-row>
					</el-timeline-item>
				</el-timeline>

				<div>
					<el-button type="primary" class="button" @click="submit()">提交请假申请</el-button>
				</div>
			</el-card>
		</div>


		<el-dialog v-model="visible" title="请假流程">
			<!-- 时间线 -->
			<el-timeline>
				<el-timeline-item v-for="(activity, index) in leaveline" :key="index" :icon="activity.icon" :type="activity.type"
					:color="activity.color" :size="activity.size" :hollow="activity.hollow" :timestamp="activity.timestamp">
					{{ activity.content }}
				</el-timeline-item>
			</el-timeline>

			<!-- 按钮区域 -->
			<template #footer>
				<span class="dialog-footer">
					<el-button @click="visible = false">我已知晓</el-button>
				</span>
			</template>
		</el-dialog>
	</div>
</template>

<script setup>
import { ref } from 'vue';
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage } from 'element-plus'
import {CreateLeaveForm} from '@/api/renshiguanli/leave.js'
const userStore = useUserStore()


const visible = ref(true)
const Leaveform = ref({
	leaveType: 0,//请假类型
	userId: 0,// 当前用户ID
	leaveContent: "",//请假事由
	selectTime: '', // 时间选择器
	beginTime: null,//请假时间
	endTime: null,//结束时间
	image: "",//附加图片
	approval: 0,//审核状态
	shenpiUserID: 1, //审批人ID
	shenpiUser2ID: 0, //审批人ID
	daynum:0,

})

// upload 组件-start
const fullscreenLoading = ref(false)
const path = ref(import.meta.env.VITE_BASE_API)
const emit = defineEmits(['on-success'])

const checkFile = (file) => {
  fullscreenLoading.value = true
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 3
  if (!isJPG && !isPng) {
    ElMessage.error('上传图片只能是 jpg或png 格式!')
    fullscreenLoading.value = false
  }
  if (!isLt2M) {
    ElMessage.error('未压缩的上传图片大小不能超过 3MB')
    fullscreenLoading.value = false
  }
  return (isPng || isJPG) && isLt2M
}

const uploadSuccess = (res) => {
  const { data } = res
	if(res.code == 0){
		Leaveform.value.image = data.file.url
		ElMessage({type: 'success',message: '上传成功'})
	}
  if (data.file) {
    emit('on-success', data.file.url)
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}


// upload 组件-end

// leave 请假时间线
const leaveline = [
	{ content: '填写请假表单', timestamp: '2018-04-12 20:46', size: 'large', type: 'primary' },
	{ content: '部门主管领导审批', timestamp: '2018-04-03 20:46', size: 'large', hollow: true, },
	{ content: '若请假天数超过3天, 再转孙老师审批', timestamp: '2018-04-03 20:46', size: 'large', hollow: true, },
	{ content: '查看审核结果', timestamp: '2018-04-03 20:46', hollow: true, size: 'large', }
]
const submit = async () => {
	Leaveform.value.beginTime = Leaveform.value.selectTime[0]
	Leaveform.value.endTime = Leaveform.value.selectTime[1]
	console.log(Leaveform.value)
	let res = await CreateLeaveForm(Leaveform.value)
	console.log(res)
}

</script>

<script>
export default {
	name: 'EnteringLeave'
}
</script>
<style lang="scss" scoped>
.top {
	padding: 24px;
	padding-bottom: 10px;
	padding-top: 16px;
	background-color: #fff;
	border-radius: 2px;
}

.card-header {
	font-weight: 600;
	color: #191a23;
}

.selectuser {
	border-radius: 4px;
	opacity: 0.68;
	width: 170px;
	height: 30px;
	display: flex;
	align-items: center;
	text-align: center;
	justify-content: center;
}

.timecenter {
	font-size: 24px;
	transform: rotate(90deg);
	margin: 0 18px;
	// margin-left: 14px;
	// margin-top: 8px;
	color: #9a9a9a;
}

.juzhong {
	display: flex;
	justify-content: center;
	align-items: center;
}
</style>