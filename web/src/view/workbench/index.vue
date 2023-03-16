<template>
  <div class="zong">
    <!-- 上方栏目 -->
    <div class="top">
      <el-row :gutter="40">
        <!-- 待处理 -->
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="top-element">
              <div class="left">
                <div class="title">待处理</div>
                <div class="titlenum">355</div>
              </div>
              <div class="right">
                <img style="width: 94px; height: 94px" src="@/assets/workbench/1.jpg"/>
              </div>
            </div>
          </el-card>
        </el-col>

        <!-- 已发起 -->
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="top-element">
              <div class="left">
                <div class="title">已发起</div>
                <div class="titlenum">355</div>
              </div>
              <div class="right">
                <img style="width: 94px; height: 94px" src="@/assets/workbench/2.jpg"/>
              </div>
            </div>
          </el-card>
        </el-col>

        <!-- 已处理 -->
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="top-element">
              <div class="left">
                <div class="title">已处理</div>
                <div class="titlenum">355</div>
              </div>
              <div class="right">
                <img style="width: 94px; height: 94px" src="@/assets/workbench/3.jpg"/>
              </div>
            </div>
          </el-card>
        </el-col>

        <!-- 我收到的 -->
        <el-col :span="6">
          <el-card shadow="hover">
            <div class="top-element">
              <div class="left">
                <div class="title">我收到的</div>
                <div class="titlenum">355</div>
              </div>
              <div class="right">
                <img style="width: 94px; height: 94px" src="@/assets/bg3.png"/>
              </div>
            </div>
          </el-card>
        </el-col>

        <el-col :span="3" class="hidden-sm-and-down">
          <img src="" class="right-img" alt>
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

.top-element{
  display: flex;
  justify-content: left;
  height: 105px;
  padding: 0 10px;
  .left{
    width: 50%;
    margin-top: 8px;
    .title{
      height: 25px;
      font-size: 18px;
      font-weight: 500;
      color: #999999;
      line-height: 25px;
    }
    .titlenum{
      margin-top: 14px;
      width: 58px;
      height: 45px;
      font-size: 32px;
      font-weight: 600;
      color: #333333;
      line-height: 45px;
    }
  }
  .right{
    img{
      margin-top: 4px;
      margin-left: 10px;
    }
    width: 50%;
  }
}

.top {
  padding: 30px;
  padding-bottom: 10px;
  padding-top: 16px;
  background-color: #fff;
  border-radius: 2px;
  box-sizing: border-box;
  height: auto;
  box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);

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
