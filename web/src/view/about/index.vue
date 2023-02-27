<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="12">
        <el-card>
          <template #header>
            博远天合教育
          </template>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                  <img
                    class="org-img dom-center"
                    src="@/assets/bylogo.png"
                    alt="gin-vue-admin"
                  >
              </el-col>
            </el-row>
          </div>
        </el-card>

        <!-- byIT团队 -->
        <el-card style="margin-top: 20px">
          <template #header>
            <div>ByIT团队</div>
          </template>
          <div>
            <el-row>
              <el-col :span="8" :offset="8">
                  <img
                    class="org-img dom-center"
                    src="@/assets/bylogo.png"
                    alt="flipped-aurora"
                  >
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
      
      <!-- 提交记录 -->
      <el-col :span="12">
        <el-card>
          <template #header>
            <div>提交记录</div>
          </template>
          <div>
            <el-timeline>
              <el-timeline-item
                v-for="(item,index) in dataTimeline"
                :key="index"
                :timestamp="item.from"
                placement="top"
              >
                <el-card>
                  <p>{{ item.title }}: {{ item.message }}</p>
                </el-card>
              </el-timeline-item>
            </el-timeline>
          </div>

        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'About',
}
</script>

<script setup>
import { ref } from 'vue'
import { Commits} from '@/api/github'
import { formatTimeToStr } from '@/utils/date'
const page = ref(0)

const dataTimeline = ref([])
const loadCommits = () => {
  Commits(page.value).then(({ data }) => {
    data.forEach((element) => {
      if (element.commit.message) {
        dataTimeline.value.push({
          from: formatTimeToStr(element.commit.author.date, 'yyyy-MM-dd'),
          title: element.commit.author.name,
          showDayAndMonth: true,
          message: element.commit.message,
        })
      }
    })
  })
}
loadCommits()

</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
    "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>
