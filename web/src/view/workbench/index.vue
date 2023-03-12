<template>
  <div class="zong">
    <!-- 上方栏目 -->
    <div class="top">
      <el-row :gutter="20">
        <!-- 待处理 -->
        <el-col :span="6">
          <el-card shadow="hover">
            <template #header>
              <div>
                <el-row>
                  <el-col :span="20">
                    <span class="title1">待处理</span>
                  </el-col>
                  <el-col :span="4">
                    <el-icon class="icon">
                      <Edit />
                    </el-icon>
                  </el-col>
                </el-row>
              </div>
            </template>
            <el-row>
              <el-col :span="6">
                <span class="title2">0</span>
              </el-col>
              <el-col :span="18">
              </el-col>
            </el-row>


          </el-card>
        </el-col>

        <!-- 已发起 -->
        <el-col :span="6">
          <el-card shadow="hover">
            <template #header>
              <div>
                <el-row>
                  <el-col :span="20">
                    <span class="title1">已发起</span>
                  </el-col>
                  <el-col :span="4">
                    <el-icon class="icon">
                      <Tickets />
                    </el-icon>
                  </el-col>
                </el-row>
              </div>
            </template>
            <el-row>
              <el-col :span="20">
                <span class="title2">0</span>
              </el-col>

            </el-row>


          </el-card>
        </el-col>

        <el-col :span="8">
          <el-card shadow="hover">
            <template #header>
              <div>
                <el-row>
                  <el-col :span="20">
                    <span class="title1">已发起</span>
                  </el-col>
                  <el-col :span="4">
                    <el-icon class="icon">
                      <Tickets />
                    </el-icon>
                  </el-col>
                </el-row>
              </div>
            </template>
            <el-row>
              <el-col :span="20">
                <span class="title2">0</span>
              </el-col>

            </el-row>


          </el-card>
        </el-col>

        <el-col :span="3" class="hidden-sm-and-down">
          <img src="@/assets/bg3.png" class="right-img" alt>
        </el-col>
      </el-row>
    </div>
    
    <div class="bottom">
      <el-card class="gva-card quick-entrance">
        <template #header>
          <div class="card-header">
            <span>请假事务</span>
          </div>
        </template>
        <el-row :gutter="50">
          <el-col
            v-for="(card, key) in toolCards" :key="key" :span="4" :xs="8"
            class="quick-entrance-items"
            @click="toTarget(card.name)"
          >
            <div class="quick-entrance-item">
              <div class="quick-entrance-item-icon" :style="{ backgroundColor: card.bg }">
                <el-icon>
                  <component :is="card.icon" :style="{ color: card.color }" />
                </el-icon>
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col>
        </el-row>
      </el-card>
    </div>


  </div>
</template>


<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import 'element-plus/theme-chalk/display.css'

const toolCards = ref([
{
    label: '发起请假',
    icon: 'clock',
    name: 'enteringLeave',
    color: 'rgba(50, 150, 250,1)',
    bg: 'rgba(50, 150, 250,0.3)'
  },
  {
    label: '处理请假',
    icon: 'monitor',
    name: 'myleave',
    color: '#ff9c6e',
    
    bg: 'rgba(255, 156, 110,.3)'
  },
  {
    label: '角色管理',
    icon: 'setting',
    name: 'authority',
    color: '#69c0ff',
    bg: 'rgba(105, 192, 255,.3)'
  },
])

const router = useRouter()

const toTarget = (name) => {
  if(name == "qingjia"){
    return
  }
  router.push({ name })
}


</script>

<script>
export default {
  name: 'WorkBench'
}
</script>


<style lang="scss" scoped>

::v-deep(.el-card__header) {
  padding-bottom: 4px;
  padding-top: 4px;
}
@mixin flex-center {
  display: flex;
  align-items: center;
}


.top {
  padding: 24px;
  padding-bottom: 10px;
  padding-top: 16px;
  background-color: #fff;
  border-radius: 2px;

  .title1 {
    line-height: 34px;
    font-size: 20px;
  }

  .icon {
    line-height: 34px;
    font-size: 34px;
  }

  .title2 {
    line-height: 34px;
    font-size: 34px;
    margin-left: 4px;
  }

  .right-img {
    position: absolute;
    top: -18px;
    right: 20px;
    height: 190px;
  }
}
.bottom {
  margin-top: 20px;


  .gva-card {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 2px;
    height: auto;
    padding: 12px 16px;
    overflow: hidden;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
  }
  .card-header{
    padding-bottom: 10px;
    font-size: 18px;
  }
  .quick-entrance-items {
    @include flex-center;
    justify-content: center;
    text-align: center;
    color: #333;
    .quick-entrance-item {
      padding: 16px 10px;
      margin-top: -16px;
      margin-bottom: -16px;
      border-radius: 4px;
      transition: all 0.2s;
      &:hover{
        box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
        background-color: rgba(217, 217, 217, 0.2);
      }
        cursor: pointer;
        height: auto;
        text-align: center;
        // align-items: center;
        &-icon {
            width: 50px;
            height: 50px !important;
            border-radius: 8px;
            @include flex-center;
            justify-content: center;
            margin: 0 auto;
            i {
                font-size: 24px;
            }
        }
        p {
          margin-top: 16px;
          font-size: 16px;
        }
    }
  }
}



</style>
