<template>
    <div>
        <el-dialog title="用户" :visible.sync="Visible">
            <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
                <el-form-item label="用户账号" prop="username">
                    <el-input v-model="searchForm.username" v-trim placeholder="请输入名称" @keypress.enter.native="search" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="el-icon-search" @click.stop="search">搜索</el-button>
                    <el-button type="primary" @click.stop="reset" icon="el-icon-refresh">重置</el-button>
                    <el-button type="primary" @click="$refs.AddUser.open(searchForm.userGroupId)" icon="el-icon-plus">增加用户</el-button>
                </el-form-item>
            </el-form>
            <el-table border v-loading="isWorking.search" :data="tableData.list" stripe>
                <el-table-column type="index" fixed="left" aling="center"></el-table-column>
                <el-table-column align="center" prop="username" label="用户账号 "  min-width="110px"></el-table-column>
                <el-table-column align="center" prop="realname" label="用户名称 "  min-width="110px"></el-table-column>
                <el-table-column align="center" prop="status_dictText" label="状态"  min-width="110px"></el-table-column>
                <el-table-column align="center" label="操作"  min-width="60px">
                    <template slot-scope="{row}">
                        <el-button  type="text" @click="deleteById(row.id,searchForm.userGroupId)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-dialog>
        <AddUser ref="AddUser" @success="search"></AddUser>
    </div>
</template>
<script>
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "@/components/base/BaseTable";
import Pagination from '@/components/common/Pagination.vue'
import AddUser from '@/components/system/userGroup/addUser'
export default {
    extends: BaseTable,
    components: {TableLayout,Pagination,AddUser},
    data(){
        return{
            searchForm:{
                userGroupId:'',
                username:'',
            },
            Visible:false,
            visible:false,
            UserList:[]
        }
        
    },
    created(){
        this.config({
        api: "/system/userGroup/user",
        "field.main": "username",
        });
        this.search();
    },
    methods:{
        open(title,row){
            this.Visible = true
            this.$nextTick(() =>{
                this.searchForm.userGroupId = row.id
                this.search()
            })
        },
        deleteById(id,ids){
            this.api.deleteById({id,ids}).then(res => {
                // console.log(res);
                this.$tip.success('删除成功')
                this.search()
            }).catch((err) => {
                
            });
        },
        addUser(){
            this.api.list({
                pageNo: 1,
                pageSize: 100,
                order: 'desc',
                column: 'createTime',
                userGroupId: this.searchForm.userGroupId
            }).then(res =>{
                this.UserList = res.records
                this.visible = true
            })
        }
    }
}
</script>
<style lang="scss" scoped>

</style>