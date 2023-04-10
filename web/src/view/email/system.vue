<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="320px" label-position="left">
      <!-- 用户相关 -->
      <el-collapse>
        <el-collapse-item>
          <template #title>
            <h3>
              <span>自动化邮件推送设置--用户相关</span>
              <el-icon style="margin-left: 6px;top: 2px;">
                <user />
              </el-icon>
            </h3>

          </template>

          <el-form-item  label="一、新增用户时 发送邮件通知该用户">
            <el-checkbox v-model="config.email_send_setting['user-register']" />
          </el-form-item>

          <el-form-item label="二、重置用户密码时 发送邮件通知该用户">
            <el-checkbox v-model="config.email_send_setting['user-reset']" />
          </el-form-item>

        </el-collapse-item>

      </el-collapse>

      <!-- 请假相关 -->
      <el-collapse>
        <el-collapse-item>
          <template #title>
            <h3>
              <span>自动化邮件推送设置--请假相关</span>
              <el-icon style="margin-left: 6px;top: 2px;">
                <message />
              </el-icon>
            </h3>

          </template>

          <el-form-item label="一、用户申请请假时 发送邮件通知审核人">
            <el-checkbox v-model="config.email_send_setting['leave-inform-shenheren']" />
          </el-form-item>

          <el-form-item label="二、请假审核完毕后 发送邮件通知申请人">
            <el-checkbox v-model="config.email_send_setting['leave-inform-shenqingren']" />
          </el-form-item>

        </el-collapse-item>

      </el-collapse>
    </el-form>
    <div class="gva-btn-list">
      <el-button type="primary" @click="update">立即更新</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'EmailConfig'
}
</script>

<script setup>
import { getSystemConfig, setSystemConfig } from '@/api/system'
import { emailTest } from '@/api/email'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const activeNames = reactive([])
const config = ref({
  email: {},
  email_send_setting:{},
})

const initForm = async() => {
  const res = await getSystemConfig()
  if (res.code === 0) {
    console.log("配置",res.data.config)
    config.value = res.data.config
  }
}
initForm()
const update = async() => {
  const res = await setSystemConfig({ config: config.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '配置文件设置成功'
    })
    await initForm()
  }
}
const email = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '邮件发送成功'
    })
    await initForm()
  } else {
    ElMessage({
      type: 'error',
      message: '邮件发送失败'
    })
  }
}

</script>

<style lang="scss">
.system {
  background: #fff;
  padding:36px;
  border-radius: 2px;
  h2 {
    padding: 10px;
    margin: 10px 0;
    font-size: 16px;
    box-shadow: -4px 0px 0px 0px #e7e8e8;
  }
  ::v-deep(.el-input-number__increase){
    top:5px !important;
  }
  .gva-btn-list{
    margin-top:16px;
  }
}
</style>
