<template>
    <GlobalWindow
        :title="title"
        :visible.sync="visible"
        :confirm-working="isWorking"
        @confirm="confirm"
        @closeTrigger="$refs.form.resetFields()"
    >
        <el-form :model="form" ref="form" :rules="rules">
            <el-form-item label="规则名称" prop="ruleName">
                <el-input
                    v-model="form.ruleName"
                    placeholder="请输入规则名称"
                    v-trim
                    maxlength="50"
                />
            </el-form-item>
            <el-form-item label="规则字段" prop="ruleColumn">
                <el-select clearable
                    v-model="form.ruleColumn"
                    placeholder="请输入规则字段"
                    @change="ruleColunmbtn"
                >
                    <el-option
                        v-for="item in ruleColumnList"
                        :key="item.value"
                        :label="item.text"
                        :value="item.value"
                    ></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="规则值" prop="ruleValue">
                <el-select clearable
                    v-model="form.ruleValue"
                    multiple
                    placeholder="请选择规则值"
                >
                    <el-option
                        v-for="item in ruleValueList"
                        :key="item.id"
                        :label="item.name"
                        :value="item.id"
                    ></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="状态">
                <el-radio-group v-model="form.status">
                    <el-radio label="1">有效</el-radio>
                    <el-radio label="0">无效</el-radio>
                </el-radio-group>
            </el-form-item>
        </el-form>
    </GlobalWindow>
</template>

<script>
import BaseOpera from "@/components/base/BaseOpera";
import GlobalWindow from "@/components/common/GlobalWindow";
export default {
    name: "OperaRoleWindow",
    extends: BaseOpera,
    components: { GlobalWindow },
    data() {
        return {
            // 原角色码
            originRoleCode: "",
            // 表单数据
            form: {
                id: undefined,
                ruleName: "",
                ruleColumn: "",
                ruleValue: "",
                status: "1",
                permissionId: "",
                ruleValueText: "",
                ruleConditions: "IN",
            },
            // 验证规则
            rules: {
                ruleName: [{ required: true, message: "请输入规则名称" }],
                ruleColumn: [{ required: true, message: "请选择规则字段" }],
                ruleValue: [{ required: true, message: "请选择规则值" }],
            },
            ruleColumnList: [],
            ruleValueList: [],
        };
    },
    methods: {
        open(title, target) {
            this.title = title;
            this.visible = true;
            // 新建
            if (target == null) {
                this.$nextTick(() => {
                    this.title = '新建'
                    this.form.permissionId = title
                    this.$refs.form.resetFields();
                    this.form[this.configData["field.id"]] = undefined;
                });
                return;
            }

            // 编辑
            this.$nextTick(() => {
                this.originRoleCode = target.code;
                for (const key in this.form) {
                    this.form[key] = target[key];
                }
                this.form.ruleValue = target.ruleValue.split(',')
                this.ruleColunmbtn()
            });
        },
        confirm() {
            if (
                this.form.id == null ||
                this.form.id === "" ||
                this.form.id === undefined
            ) {
                this.$refs.form.validate((valid) => {
                    if (!valid) {
                        return
                    }
                    let arr = [];
                    let ruleValueText = "";
                    let ruleValue = this.form.ruleValue.join(",");
                    this.form.ruleValue.forEach((item) => {
                        this.ruleValueList.forEach((lable) => {
                            if (lable.id === item) {
                                arr.push(lable.name);
                            }
                        });
                        ruleValueText = arr.join(",");
                    });
                    this.form.ruleValueText = ruleValueText
                    this.api.addPermissionRule({
                        ruleName: this.form.ruleName,
                        ruleColumn: this.form.ruleColumn,
                        ruleValue: ruleValue,
                        status: this.form.status,
                        permissionId: this.form.permissionId,
                        ruleValueText: ruleValueText,
                        ruleConditions: "IN",
                    }).then(res =>{
                        this.$tip.success('新建成功')
                        this.$emit('success')
                    })
                    this.visible = false
                    return;
                })
            }
            if (this.originRoleCode === this.form.code) {
                this.__confirmEdit();
                return;
            }
            
        },
        ruleColunmbtn() {
            this.api.queryByField({
                    field: this.form.ruleColumn,
                })
                .then((res) => {
                    this.ruleValueList = res;
                });
        },
        __confirmEdit(){
            let arr = [];
            let ruleValueText = "";
            let ruleValue = this.form.ruleValue.join(",");
            this.form.ruleValue.forEach((item) => {
                this.ruleValueList.forEach((lable) => {
                    if (lable.id === item) {
                        arr.push(lable.name);
                    }
                });
                ruleValueText = arr.join(",");
            });
            this.form.ruleValueText = ruleValueText
            this.api.editPermissionRule({
                id:this.form.id,
                ruleName: this.form.ruleName,
                ruleColumn: this.form.ruleColumn,
                ruleValue: ruleValue,
                status: this.form.status,
                permissionId: this.form.permissionId,
                ruleValueText: ruleValueText,
                ruleConditions: "IN",
            }).then(res =>{
                this.$tip.success('修改成功')
                this.$emit('success')
            })
            this.visible = false
        }
    },
    created() {
        this.config({
            api: "/system/userGroup/Manifest",
        });
        this.api.permission_field().then((res) => {
            this.ruleColumnList = res;
        });
        // this.$nextTick(() =>{
        //     this.ruleColunmbtn()
        // })
    },
    // watch:{
    //     'form.ruleColumn'(val, oldval){
    //         this.ruleColunmbtn()
    //     }
    // }
};
</script>
