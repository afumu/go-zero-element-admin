<template>
    <el-dialog title="增加用户" :visible.sync="visible">
        <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
            <el-form-item label="用户账号" prop="id">
                <el-select clearable v-model="searchForm.id" v-trim placeholder="请输入名称" >
                    <el-option v-for="item in tableData.list" :key="item.id" :label="item.username" :value="item.id"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" icon="el-icon-search" @click.stop="search">搜索</el-button>
                <el-button type="primary" @click.stop="reset" icon="el-icon-refresh">重置</el-button>
            </el-form-item>
        </el-form>
        <el-table border v-loading="isWorking.search" :data="tableData.list" stripe  @sort-change="handleSortChange" @selection-change="handleSelectionChange">
            <el-table-column type="selection" fixed="left" aling="center"></el-table-column>
            <el-table-column align="center" prop="username" label="用户账号  " fixed="left" min-width="110px"></el-table-column>
            <el-table-column align="center" prop="realname" label="用户名称 " fixed="left" min-width="110px"></el-table-column>
            <el-table-column align="center" prop="sex_dictText" label="性别 " fixed="left" min-width="110px"></el-table-column>
            <el-table-column align="center" prop="phone" label="电话 " fixed="left" min-width="110px"></el-table-column>
            <el-table-column align="center" prop="departIds_dictText" label="部门 " fixed="left" min-width="110px"></el-table-column>
        </el-table>
        <span slot="footer" class="dialog-footer">
            <el-button @click="visible = false">取 消</el-button>
            <el-button type="primary" @click="addUsers">确 定</el-button>
        </span>
        <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
    </el-dialog>
</template>
<script>
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "@/components/base/BaseTable";
import Pagination from '@/components/common/Pagination.vue'
export default {
    extends: BaseTable,
    components: {TableLayout,Pagination},
    methods:{
        open(title){
            this.visible = true
            this.$nextTick(() =>{
                this.userGroupId = title
            })
            this.search()
        },
        addUsers(){
            let list =  this.tableData.selectedRows
            let Idlist = []
            if (list.length<1) {
                this.$tip.warning('请选择数据')
                return
            }
            list.forEach(item => {
                Idlist.push(item.id)
            }); 
            this.api.addUser({
                userGroupId: this.userGroupId,
                userIdList:Idlist
            }).then(res =>{
                this.$tip.success('添加成功')
                this.visible = false
                this.$emit('success')
            })
        }
    },
    data(){
        return{
           visible: false ,
           searchForm:{
               id: '',
           },
           userGroupId: '',
        }
    },
    created(){
        this.config({
        api: "/system/userGroup/addUser",
        });
    }
}
</script>
<style lang="scss" scoped>

</style>