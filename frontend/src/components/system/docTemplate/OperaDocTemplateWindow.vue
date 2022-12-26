<template>
    <GlobalWindow
        :title="title"
        :visible.sync="visible"
        :confirm-working="isWorking"
        @confirm="confirm"
    >
        <el-form :model="form" ref="form" :rules="rules">
            <table>
                <tr>
                    <td>
                        <el-form-item v-if="this.form.id == null || this.form.id === ''">
                            <el-input
                                style="width:280px;margin:10px 0px 10px 0px"
                                :placeholder='resMessage'
                                v-model="resMessage"
                                :disabled="true"
                            ></el-input>
                        </el-form-item>
                    </td>
                    <td>
                        <el-upload
                            :action="this.$consts.fileConstant.uploadTemplateFilePath"
                            :headers="importHeaders"
                            :show-file-list="false"
                            :on-success="uploadSuccess"
                        >
                            <el-button size="small" type="primary" v-if="this.form.id == null || this.form.id === ''"
                                       style="height: 28px;margin:10px 0px 28px 0px">
                                点击上传
                            </el-button>
                        </el-upload>
                    </td>
                </tr>
                <tr>
                    <el-form-item label="文件名" prop="docName" required>
                        <template>
                            <el-input v-model="form.docName" placeholder="请输入文件名"></el-input>
                        </template>
                    </el-form-item>
                </tr>
            </table>
        </el-form>
    </GlobalWindow>
</template>

<script>
    import BaseOpera from '@/components/base/BaseOpera'
    import GlobalWindow from '@/components/common/GlobalWindow'
    import Cookies from 'js-cookie'

    export default {
        name: "OperaDocTemplateWindow",
        extends: BaseOpera,
        components: {GlobalWindow},
        data() {
            return {
                resMessage: "请上传模板文档",
                importHeaders: {"X-Access-Token": Cookies.get('X-Access-Token')},
                form: {
                    id: '',
                    docFileId: '',
                    docName: '',
                }, rules: {
                    docName: [
                        {required: true, message: '请输入文件名'}
                    ],
                }
            }
        }, methods: {
            open(title, target) {
                this.title = title
                this.visible = true
                // 新建
                this.resMessage = "请上传模板文档"
                this.form.docFileId = null
                this.form.id = null
                this.form.docName = null
                if (target == null) {
                    this.$nextTick(() => {
                        this.$refs.form.resetFields()
                        this.form[this.configData['field.id']] = null
                    })
                    return
                }
                // 编辑
                this.$nextTick(() => {
                    this.id = target.id
                    for (const key in this.form) {
                        this.form[key] = target[key]
                    }
                })
            },
            confirm() {
                if (this.form.id == null || this.form.id === '') {
                    this.__confirmCreate()
                } else {
                    this.__confirmEdit()
                }
            }, uploadSuccess(res, file) {
                this.resMessage = file.name
                this.form.docFileId = res.message
                this.form.docName = file.name
                this.resMessage = this.form.docName
            }
        },
        created() {
            this.config({
                api: '/system/docTemplate',
                'field.id': 'id',
                'field.main': 'docName',
            })
        }
    }
</script>

<style scoped>

</style>