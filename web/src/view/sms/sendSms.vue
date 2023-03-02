<template>
  <div>
    <warning-bar title="在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示" />
    <div class="gva-table-box bottom-mg-lg">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="this.dialogFormVisible = true"
        >新建项目</el-button>
        <el-button
          type="primary"
          icon="delete"
          @click="delSmsProject"
        >删除选中项目</el-button>
      </div>

      <el-tabs
        type="card"
        @tab-click="changeTab"
        v-model="activeName"
      >
        <el-tab-pane
          :label="obj.project_name"
          :name="obj.ID"
          v-for="obj in project_list"
        >
          <div>
            <el-descriptions
              title=""
              border
              :column="4"
            >
              <el-descriptions-item label="模板ID">
                <el-tag size="small">{{obj.template_id}}</el-tag>
              </el-descriptions-item>
              <el-descriptions-item
                label="备注">{{obj.comment}}</el-descriptions-item>
              <el-descriptions-item
                label="条数">{{obj.row_num}}</el-descriptions-item>
              <el-descriptions-item
                label="参数个数">{{obj.param_num}}</el-descriptions-item>
              <el-descriptions-item
                label="创建时间">{{formatDate(obj.created_at)}}</el-descriptions-item>
              <el-descriptions-item
                label="更新时间">{{formatDate(obj.updated_at)}}</el-descriptions-item>
              <el-descriptions-item
                label="模板内容">{{win.project.sms_template.template_content}}</el-descriptions-item>

            </el-descriptions>

            <el-table
              ref="multipleTable"
              :data="rows"
              style="width: 100%;margin-top: 20px;"
              tooltip-effect="dark"
              row-key="ID"
              @selection-change="handleSelectionChange"
            >
              <el-table-column
                type="selection"
                width="55"
                align="center"
              />
              <el-table-column
                align="center"
                label="添加日期"
                width="170"
              >
                <template #default="scope">
                  <span>{{ formatDate(scope.row.created_at) }}</span>
                </template>
              </el-table-column>
              <el-table-column
                align="center"
                label="手机号"
                prop="phone"
                width="120"
              />
              <el-table-column
                align="center"
                :label="'参数'+num"
                :prop="`param${num}`"
                width="120"
                v-for="num in obj.param_num"
              />

              <el-table-column
                align="center"
                label="按钮组"
                min-width="160"
                fixed="right"
              >
                <template #default="scope">
                  <el-button
                    type="primary"
                    link
                    icon="edit"
                    @click=""
                  >变更</el-button>
                  <el-popover
                    v-model="scope.row.visible"
                    placement="top"
                    width="160"
                  >
                    <p>确定要删除吗？</p>
                    <div style="text-align: right; margin-top: 8px;">
                      <el-button
                        type="primary"
                        link
                        @click="scope.row.visible = false"
                      >取消</el-button>
                      <el-button
                        type="primary"
                        @click="DelRow(scope.row.ID)"
                      >确定</el-button>
                    </div>
                    <template #reference>
                      <el-button
                        type="primary"
                        link
                        icon="delete"
                        @click="scope.row.visible = true"
                      >删除</el-button>
                    </template>
                  </el-popover>
                </template>
              </el-table-column>
            </el-table>

            <div class="gva-pagination" style="justify-content: space-between;">
              <div class="gva-btn-list" style="margin-top: 12px;">
                <el-button
                  type="primary"
                  icon="plus"
                  size="default"
                  @click="dialogFormVisible2 = true"
                >新增行</el-button>
                <el-button
                  type="primary"
                  icon="delete"
                  size="default"
                  @click="DelRows"
                >删除</el-button>

                <el-button
                  type="primary"
                  icon="printer"
                  size="default"
                  @click="exportExcel"
                >导出excel</el-button>
                <el-button
                  type="warning"
                  icon="message"
                  size="default"
                  @click="send"
                >批量发送信息</el-button>
              </div>
              
              <div>
                <el-pagination
                :current-page="pageSetting.page"
                :page-size="pageSetting.pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :total="pageSetting.total"
                layout="total, sizes, prev, pager, next, jumper"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
              />
              </div>
            </div>
          </div>
        </el-tab-pane>

      </el-tabs>

    </div>

    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="新建项目"
    >
      <el-form
        :inline="false"
        :model="form"
        label-position="top"
      >
      <el-form-item label="项目名">
          <el-input
            v-model="form.project_name"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="模板ID">
          <el-popover
            :visible="visibleMbID"
            placement="bottom"
            :width="400"
            trigger="click"
          >
            <template #reference>
              <el-input
                slot="reference"
                v-model="form.template_id"
                @click="visibleMbID=true"
                @blur="visibleMbID = false"
              />
            </template>
            <el-table
              :data="tableData2"
              ref="singleTable"
              highlight-current-row
              style="width: 100%"
              @row-click="selectTemplate"
              
            >
              <el-table-column
                type="index"
                width="50"
              />
              <el-table-column
                prop="template_id"
                label="模板ID"
                width="120"
              />
              <el-table-column
                prop="template_content"
                label="模板内容"
              />
            </el-table>
          </el-popover>
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model="form.comment"
            autocomplete="off"
          />
        </el-form-item>

      </el-form>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="dialogFormVisible2"  title="新增行">
		<el-row :gutter="20">
		  <el-col v-for="obj, idx in win.project.param_num" :span="8">
			<el-input v-model="rowCommit.tpl_params[idx]" :placeholder="`参数${idx + 1}`" />
		  </el-col>
		</el-row>
		
		<el-row>
		  <el-input v-model="rowCommit.tpl_phones" :autosize="{ minRows: 5, maxRows: 15 }"
      type="textarea" placeholder="手机号列表" />
		</el-row>
  
		<template #footer>
		  <div class="dialog-footer">
			<el-button @click="rowCommit.tpl_phones = ''">清空</el-button>
			<el-button
			  type="warning"
			  @click="dialogFormVisible2 = false"
			>取 消</el-button>
			<el-button
			  type="primary"
			  @click="addrow"
			>确 定</el-button>
		  </div>
		</template>
	  </el-dialog>
  </div>
</template>
  
<script>
import { GetSmsList } from '@/api/txyun/sms'
export default {
  name: 'SendSms',
  data() {
    return {
      activeName:"", // tabs 选中
      multipleSelection:[], // 下方表格中 选中的row 元素
      rowCommit:{
        tpl_params:["","","","","","","","",""],// 输入参数列表
        tpl_phones:"",
        sms_project_id:"",
      },
      rows:[],
      visibleMbID: false,
      tableData2: [], //模板列表, 用于选则模板ID
      pageSetting: {
        page: 1,
        pageSize: 10,
        total:0,
      },
      dialogFormVisible:false,
      dialogFormVisible2:false,
      project_list:[], // 项目列表
      form:{
        project_name: '', //项目名
        comment: '',  //备注
        template_id: '', //模板ID
      },
      win:{
        selectTab: 0,
        temeplate: null,
        project: null,
      },
    }
  },
  props: {},
  components: {},
  computed: {},
  created() { 
    this.GetTemplateList() 
    this.getProjects()
    
  },
  methods: {
    async addrow(){
      // console.log(this.rowCommit)
      if(this.rowCommit.tpl_phones==""){
        return
      }
      let req = {
        tpl_phones:this.rowCommit.tpl_phones,
        sms_project_id: parseInt(this.win.project.ID),
        param1:this.rowCommit.tpl_params[0],
        param2:this.rowCommit.tpl_params[1],
        param3:this.rowCommit.tpl_params[2],
        param4:this.rowCommit.tpl_params[3],
        param5:this.rowCommit.tpl_params[4],
        param6:this.rowCommit.tpl_params[5],
        param7:this.rowCommit.tpl_params[6],
        param8:this.rowCommit.tpl_params[7],
        param9:this.rowCommit.tpl_params[8],
      }
      let res = await AddSmsProjectRow(req);
      // console.log(res)
      if(res.code === 0){
        ElMessage({ type: 'success', message: '添加成功' })
        this.rowCommit = {
        tpl_params:["","","","","","","","",""],// 输入参数列表
        tpl_phones:"",
        sms_project_id:""}
        this.dialogFormVisible2 = false;

        this.getProjectRows()
      }
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
      // console.log(val)
    },
    handleSizeChange(val){
      this.pageSetting.pageSize = val
      this.getProjectRows()
    },
    handleCurrentChange(val){
      this.pageSetting.page = val
      this.getProjectRows()
    },
    async exportExcel(){
      let ids = []
      for (let i = 0; i < this.multipleSelection.length; i++) {
        ids.push(this.multipleSelection[i].ID);
      }
      if(ids.length == 0){
        ElMessage({ type: 'warning', message: '请至少选择一项' })
        return 
      }
      let res = await ExportExcelSmsRows({ids:ids});
      if(res.code === 0){
        // console.log("excel",res.data.url)
        window.open(res.data.url)
        ElMessage({ type: 'success', message: '导出成功' })
      }
    },
    async send(){

      let ids = []
      for (let i = 0; i < this.multipleSelection.length; i++) {
        ids.push(this.multipleSelection[i].ID);
      }
      if(ids.length == 0){
        ElMessage({ type: 'warning', message: '请至少选择一项' })
        return 
      }
      let res = await SendSmsByRows({ids:ids});
      if(res.code === 0){
        ElNotification({
          title: '发信成功',
          message: '批量发送信息成功',
          duration: 5000,
        })
        // ElMessage({ type: 'success', message: '批量发信成功' })
      }
    },
    async DelRow(id){
      // console.log(id)
      let res = await DelSmsProjectRow({id:id});
      // console.log("DelRow",res)
      if(res.code === 0){
        ElMessage({ type: 'success', message: '删除成功' })
        this.getProjectRows()
      }
    },
    async DelRows(){
        let ids = []
				for (let i = 0; i < this.multipleSelection.length; i++) {
          ids.push(this.multipleSelection[i].ID);
        }
        if(ids.length == 0){
          return 
        }
        let res = await DelSmsProjectRows({ids:ids});
        if(res.code === 0){
          ElMessage({ type: 'success', message: '批量删除成功' })
          this.getProjectRows()
        }
    },
    async changeTab(tab){
      this.win.selectTab = tab.index;
      this.win.project = this.project_list[tab.index];
      this.win.temeplate = this.win.project.sms_template;

      await this.getProjectRows()
    },
    async GetTemplateList() {
      const res = await GetSmsList({page: 1,pageSize: 1000})
      this.tableData2 = res.data.list
    },
    async getProjects(){
      let table = await GetAllSmsProject()
      if (table.code === 0) {
        this.project_list = table.data.list;
        // console.log(this.project_list)
        if (table.data.list.length != 0) {
          this.activeName = table.data.list[0].ID;
          this.changeTab({ index: '0' })
        }
      }
    },
    /* 选中表格项回调 */
    selectTemplate(val) {
      console.log(val)
      if(val){
        if(val.template_id){
          this.form.template_id = val.template_id
          this.visibleMbID = false
        }
      }
      
    },
    async delSmsProject(){
      const res = await DelSmsProject({ "id": this.win.project.ID })
      if (res.code === 0) {
        this.getProjects()
      }
    },
    async enterDialog(){
      this.form.template_id = parseInt(this.form.template_id)
      let res = await AddSmsProject(this.form)
      if (res.code === 0) {
        this.closeDialog()
        this.getProjects()
        ElMessage({ type: 'success', message: '新增短信项目成功' })
      }
    },
    async getProjectRows(){
      let res = await SmsProjectRows({id:this.win.project.ID, ...this.pageSetting})
      // console.log(res)
      if(res.code === 0){
        this.rows = []
        this.rows = res.data.list;
        this.pageSetting.total = res.data.total;
        this.pageSetting.page = res.data.page;
        this.pageSetting.pageSize = res.data.pageSize;
        // console.log("res.data",res.data)
      }

    },
    // 关闭弹窗
    async closeDialog(){
      this.dialogFormVisible = false
      this.form = {
        project_name: '',
        comment: '',
        template_id: '',
      }
    }
  },
};
</script>


<script setup>
import { AddSmsProject, GetAllSmsProject, DelSmsProject,
  SmsProjectRows,AddSmsProjectRow,DelSmsProjectRow,
  DelSmsProjectRows,SendSmsByRows,ExportExcelSmsRows } from '@/api/txyun/sms.js'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
//  ElNotification->通知
import { ElMessage,ElNotification } from 'element-plus'
import { formatDate } from '@/utils/format'


const tableData = ref([])

</script>

<style>
.project_tag {
  height: 27px;
  font-size: 15px;
  margin-left: 2px;
  margin-right: 10px;
}
</style>
  