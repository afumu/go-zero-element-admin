<template>
  <TableLayout :permissions="['system:department:query']">
  <!-- <TableLayout> -->
    <!-- 表格和分页 -->
    <div slot="search-form">
            <el-collapse v-model="activeNames" accordion>
                <el-collapse-item name="1" title="查询条件">
                    <template slot="title">
                        <div class="kj-title">
                            <span style="margin: 0 10px">查询条件</span>
                            <section>
                                <el-button type="primary" size="mini" icon="el-icon-search" @click.stop="search">搜索</el-button>
                                <el-button @click.stop="reset" size="mini" icon="el-icon-refresh">重置</el-button>
                            </section>
                        </div>
                    </template>
                    <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
                        <el-form-item label="职务编码" prop="code">
                            <search-input v-model="searchForm.code" @keypress.enter.native="search"></search-input>
                        </el-form-item>
                        <el-form-item label="职务名称" prop="name">
                            <search-input v-model="searchForm.name" @keypress.enter.native="search"></search-input>
                        </el-form-item>
                        <el-form-item label="职级" prop="postRank">
                            <el-select clearable v-model="searchForm.postRank"  placeholder="请选择职级" >
                                <el-option v-for="item in postRankList" :key="item.text" :label="item.text" :value="item.value"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-form>
                </el-collapse-item>
            </el-collapse>
        </div>
    <template v-slot:table-wrap>
        <ul class="toolbar" >
            <li><el-button type="primary" @click="$refs.OperajobManagementWindow.open('新建')" icon="el-icon-plus">新建</el-button>
            <el-button icon="el-icon-download" type="primary" @click="ExportFile('职务表')" >导出</el-button>
            <el-button icon="el-icon-upload" type="primary" @click="importBtn" >导入</el-button></li>
        </ul>
        <el-table border v-loading="isWorking.search" :data="tableData.list" stripe @selection-change="handleSelectionChange">
            <el-table-column type="selection" fixed="left" width="55"></el-table-column>
            <el-table-column prop="code" label="职务编码" fixed="left" min-width="160px"></el-table-column>
            <el-table-column prop="name" label="职务名称" fixed="left" min-width="100px"></el-table-column>
            <el-table-column prop="postRank_dictText" label="职级" fixed="left" min-width="100px"></el-table-column>
            <el-table-column label="操作" align="center" min-width="160" fixed="right">
            <template slot-scope="{row}">
                <el-button style="margin-left: 0;!important"  type="text" @click="$refs.OperajobManagementWindow.open('编辑', row)" icon="el-icon-edit">编辑</el-button>
                <el-button style="margin-left: 15px;!important" type="text" @click="deleteById(row)" icon="el-icon-delete">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
            <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
    </template>
    <el-dialog title="导入" :visible.sync="dialogVisible" width="25%">
            <el-alert style="margin-bottom: 3vh" type="warning">
                <div>
                    请先下载模板，填入数据后，通过选择文件上传数据:<a style="color:orange" href="/template/职务表.xls">下载模板</a>
                </div>
            </el-alert>
            <el-upload
                name="file"
                :action="'/sys/position/importExcel'"
                :headers="token"
                :on-change="handleImportExcel"
                :show-file-list="false"
                :on-success="uploadSuccess"
            >
                <el-button size="small" type="primary">点击上传</el-button>
            </el-upload>
            <span slot="footer" class="dialog-footer"> </span>
        </el-dialog>
    <!-- 新建/修改 -->
    <OperajobManagementWindow ref="OperajobManagementWindow" @success="handlePageChange(tableData.pagination.pageIndex)" />
  </TableLayout>
</template>

<script>
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "@/components/base/BaseTable";
import Pagination from '@/components/common/Pagination.vue'
import OperajobManagementWindow from "@/components/system/jobManagement/OperajobManagementWindow";
import {importFile, exportFile} from '@/utils/util'
import Cookies from 'js-cookie'
export default {
  name: "SystemDepartment",
  extends: BaseTable,
  components: { OperajobManagementWindow, TableLayout,Pagination },
  data() {
    return {
        // 搜索
        searchForm: {
                code:'',
                name:'',
                postRank:'',
        },
        activeNames:'1',
        token: {'X-Access-Token': Cookies.get('X-Access-Token')},
        dialogVisible:false,
        postRankList:[],
        };
    },
    methods: {
        ExportFile(data){
            exportFile('/sys/position/exportXls',data)
        },
        importBtn() {
            this.dialogVisible = true
        },
      uploadSuccess (res, file) {
        this.resMessage = file.name
        this.docFileId = res.message
        this.docName = file.name
        if (file.response.code == 200) {
          this.$message.success(
              `${file.name} 文件上传成功`
          )
        } else if (file.response.code == 201) {
          this.$message.warning(`${file.name, file.response.message} `)
        } else {
          this.$message.error(`${file.name, file.response.message} `)
        }
      },
        handleImportExcel(file) {
            this.search()
        },
    },
  created() {
    this.config({
      module: "职务",
      api: "/system/jobManagement",
      // 主字段
      "field.main": "name",
    });
    this.search();
    this.api.position_rank()
        .then(res =>{
            this.postRankList = res
        })
  },
};
</script>
<style lang="scss" scoped>
.table-layout {
  /deep/ .table-content {
    margin-top: 0;
  }
}
</style>
