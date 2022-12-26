<template>
  <!-- <el-dialog -->
  <el-drawer
    class="location-window"
    :title="title"
    status-icon
    :visible="visible"
    :size="width"
    :destroy-on-close="true"
    :append-to-body="true"
    :wrapperClosable="false"
    :close-on-press-escape="false"
    @close="close"
  >
<!--    <div class="window__body">-->
<!--      <slot></slot>-->
<!--    </div>-->
<!--    <div v-if="withFooter" class="window__footer">-->
<!--      <slot name="footer">-->
<!--        <el-button @click="confirm" :loading="confirmWorking" type="primary"-->
<!--          >确定</el-button-->
<!--        >-->
<!--        <el-button @click="close">取消</el-button>-->
<!--      </slot>-->
<!--    </div>-->
    <!-- </el-dialog> -->
    <el-dialog title="选择" :visible.sync="Visible" top="8vh">
      <el-form ref="searchFormT" :model="searchForm" inline>
        <el-form-item label="地区名字" prop="name">
          <el-input v-model="searchForm.name" @change="searchCg"></el-input>
        </el-form-item>
        <el-form-item label="地区层级" prop="level">
          <el-input v-model="searchForm.level" @change="searchCg"></el-input>
        </el-form-item>
        <el-form-item label="地区范围" prop="parentId">
          <LocationSelect :city-id.sync="searchForm.parentId" placeholder="请选择地区范围" :level="2" clearable @change="searchCg"/>
        </el-form-item>
        <el-button type="primary"  icon="el-icon-search" @click.stop="add">新增</el-button>
        <el-button @click.stop="reset"  icon="el-icon-refresh">重置</el-button>
      </el-form>
      <el-table  border v-loading="isWorking" height="250px" :data="locationList" stripe @sort-change="handleSortChange" @selection-change="handleSelectionChange">
        <el-table-column ref="tb" type="selection" width="55" align="center"></el-table-column>
        <el-table-column prop="name" label="地区名称"  align="center" min-width="70px"></el-table-column>
        <el-table-column prop="shortName" label="地区简称"  align="center" min-width="70px"></el-table-column>
        <el-table-column prop="fullName" label="全称"  align="center" min-width="70px"></el-table-column>
        <el-table-column prop="level" label="地区层级"  align="center" min-width="70px"></el-table-column>
        <el-table-column prop="areaCode" label="区号"  align="center" min-width="70px"></el-table-column>
      </el-table>
      <span slot="footer" class="dialog-footer">
           <el-button @click="VisibleT = false">取 消</el-button>
           <el-button type="primary" @click="Ok" >确定</el-button>
        </span>
    </el-dialog>
  </el-drawer>
</template>

<script>
export default {
  name: 'SelectlocationWindow',
  props: {
    width: {
      type: String,
      default: '36%',
    },
    // 是否包含底部操作
    withFooter: {
      type: Boolean,
      default: true,
    },
    // 确认按钮loading状态
    confirmWorking: {
      type: Boolean,
      default: false,
    },
    // 标题
    title: {
      type: String,
      default: '',
    },
    // 是否展示
    visible: {
      type: Boolean,
      required: true,
    },
    searchForm:{
      parentId:'',
      name:'',
      level:'',
      firstLetter:'',
      pinyin:''
    },
  },
  methods: {
    confirm() {
      this.$emit('confirm')
    },
    close() {
      this.$emit('update:visible', false)
      this.$emit("closeTrigger")
    },
    reset(){
      this.searchForm='',
          this.searchCg()
    },
    Ok(){
      this.VisibleT = false
      console.log(this.form.name)
    },
    openlocation() {
      this.VisibleT = true
      this.searchCg()
    },
    searchCg(){
      this.api.LocationList({
        level:this.searchForm.level,
        name:this.searchForm.name,
        shortName:this.searchForm.shortName,
        firstLetter: this.searchForm.firstLetter,
        parentId:this.searchForm.parentId
      }).then(res=>{
        this.locationList=res.records

      })
    }, // 行选中处理
    handleSelectionChange(selection) {
      // this.selected = selectedRows
      this.checkedGh=selection[0].fullName
      this.forms.name=selection[0].fullName
      this.form.name=selection[0].id
      if (selection.length>1){
        for(let i=0;i<selection.length;i++){
          this.$nextTick(()=>{
            this.$refs.tb[i].clearSelection()
          })
        }


        // this.$refs.tb[0].toggleRowSelection(selection.pop());
        // console.log(this.$refs.)
      }

    },
    // 排序
    handleSortChange(sortData) {
      this.tableData.sorts = []
      if (sortData.order != null) {
        this.tableData.sorts.push({
          property: sortData.column.sortBy,
          direction: sortData.order === 'descending' ? 'DESC' : 'ASC',
        })
      }
      this.handlePageChange()
    },
    // 页码变更处理
    handlePageChange(pageIndex) {
      this.__checkApi()
      this.tableData.pagination.pageIndex =
          pageIndex || this.tableData.pagination.pageIndex
      this.isWorking.search = true
      this.api
          .LocationList({
            order: 'desc',
            column: 'createTime',
            pageNo: this.tableData.pagination.pageIndex,
            pageSize: this.tableData.pagination.pageSize,
            ...this.searchForm,
          })
          .then((data) => {
            this.tableData.list = data.records
            this.tableData.pagination.total = data.total
          })
          .catch((e) => {
            this.$tip.apiFailed(e)
          })
          .finally(() => {
            this.isWorking.search = false
          })
    },
    open (title, target) {
      this.title = title
      this.visible = true
      // 新建
      if (target == null) {
        this.name='',
            this.$nextTick(() => {
              this.$refs.form.resetFields()
              this.form[this.configData['field.id']] = null
            })
        return
      }
      // 编辑
      this.$nextTick(() => {
        for (const key in this.form) {
          this.form[key] = target[key]
          this.name=target.name_dictText
        }
      })
    },
  },
}
</script>

<style scoped lang="scss">
//@import '@/assets/style/variables.scss';
//// 输入框高度
//$input-height: 32px;
//.global-window {
//  // 头部
//  /deep/ .el-dialog__header {
//    border-bottom: 1px solid #eee;
//  }
//  // 内容
//  /deep/ .el-dialog__body {
//    padding: 0;
//  }
//  /deep/ .window__body {
//    height: calc(100vh - 140px);
//    overflow-y: auto;
//    padding: 12px 16px;
//    // 标签
//    .el-form-item__label {
//      float: none;
//    }
//    // 元素宽度为100%
//    .el-form-item__content > * {
//      width: 100%;
//    }
//  }
//  // 尾部
//  /deep/ .window__footer {
//    user-select: none;
//    border-top: 1px solid #eee;
//    height: 60px;
//    line-height: 60px;
//    text-align: center;
//    position: relative;
//  }
//
//  /deep/ .el-drawer__header {
//    margin-bottom: 20px !important;
//  }
//}
</style>
