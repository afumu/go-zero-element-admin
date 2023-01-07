<!--  -->
<template>
  <div class="main">
    <el-dropdown trigger="click">
      <el-badge :value="messageNum" :hidden="messageNum == 0">
        <span class="kj-icon">
          <i class="el-icon-bell"></i>
        </span>
      </el-badge>
      <el-dropdown-menu slot="dropdown">

        <el-tabs v-model="activeName" style="padding: 0 20px">
          <el-tab-pane :label="msg1Title" name="first">
            <ol class="content">
              <li :key="index" v-for="(record, index) in announcement1" style="margin-bottom:8px">
                <div style="width: 80%">
                  <h4><a @click="openMessage(record,index)">{{ record.titile }}</a></h4>
                  <!-- <p @click="openMessage(record,index)" style="color: rgba(0,0,0,.45);margin-bottom: 0px;font-size: 14px">{{ record.msgAbstract }}</p> -->
                </div>
                <div style="text-align: right;">
                  <el-tag v-if="record.priority === 'L'" type="info">一般消息</el-tag>
                  <el-tag v-if="record.priority === 'M'" type="warning">重要消息</el-tag>
                  <el-tag v-if="record.priority === 'H'" type="danger">紧急消息</el-tag>
                </div>
                <div>
                  <p style="color: rgba(0,0,0,.45);font-size: 14px;margin-top: 5px !important">
                    {{record.sendTime}} 发布
                  </p>
                </div>
              </li>
            </ol>
            <div style="margin: 10px;text-align: center">
              <el-button @click="toNext()" type="dashed" block>查看更多</el-button>
            </div>
          </el-tab-pane>
          <el-tab-pane :label="msg2Title" name="second">
            <ol class="content">
              <li :key="index" v-for="(record, index) in announcement2" style="margin-bottom:8px">
                <div style="width: 80%">
                  <h4><a @click="openMessage(record,index)">{{ record.titile }}</a></h4>
                  <!-- <p @click="openMessage(record,index)" style="color: rgba(0,0,0,.45);margin-bottom: 0px;font-size: 14px">{{ record.msgAbstract }}</p> -->
                </div>
                <div style="text-align: right;">
                  <el-tag  v-if="record.priority === 'L'" type="info">一般消息</el-tag>
                  <el-tag  v-if="record.priority === 'M'" type="warning">重要消息</el-tag>
                  <el-tag  v-if="record.priority === 'H'" type="danger">紧急消息</el-tag>
                </div>
                <div>
                  <p  style="color: rgba(0,0,0,.45);font-size: 14px;margin-top: 5px !important">
                    {{record.sendTime}} 发布
                  </p>
                </div>
              </li>
            </ol>

            <div style="margin: 10px;text-align: center">
              <el-button @click="toNext()" type="dashed" block>查看更多</el-button>
            </div>
          </el-tab-pane>
        </el-tabs>
      </el-dropdown-menu>
    </el-dropdown>

    <el-dialog
      title="通知消息"
      :visible.sync="visible"
      width="50%">
      <h2 style="color: rgba(0,0,0,.85);">{{message.titile}}</h2>
      <p style="color: rgba(0,0,0,.45);">发布人：{{message.sender}}  发布时间：{{message.sendTime}}</p>
      <el-divider></el-divider>

      <div v-html="message.msgContent" class="ql-editor" style="padding: 12px 0px"></div>

      <span slot="footer" class="dialog-footer">
        <el-button @click="visible = false">关 闭</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
import { mapState, mapMutations } from 'vuex'
import BaseOpera from '@/components/base/BaseOpera'
export default {
  extends: BaseOpera,
  data () {
    return {
      activeName: 'first',
      announcement1: [],
      announcement2: [],
      msg1Count: '0',
      msg2Count: '0',
      msg1Title: '通知公告(0)',
      msg2Title: '系统消息(0)',
      message: {},
      visible: false
    }
  },
  methods: {
    ...mapMutations(['setUserInfo', 'switchCollapseMenu']),
    toNext () {
      this.$router.push({
        path: '/message/userMessage'
      })
    },
    // 读消息
    openMessage (message, index) {
      this.message = message
      console.log(index)
      if (message.msgCategory == 1) { // 通知
        this.announcement1.splice(index, 1)
        this.msg1Count--
        this.msg1Title = '通知公告(' + this.msg1Count + ')'
        console.log(this.announcement1)
      } else if (message.msgCategory == 2) {
        this.announcement2.splice(index, 1)
        this.msg2Count--
        this.msg2Title = '系统消息(' + this.msg2Count + ')'
      }
      this.$store.commit('SET_MESSAGENUM', this.messageNum - 1)
      this.api.readNotic(message.id)
      this.visible = true
    },

    loadData () {
      try {
        // 获取系统消息
        this.api.listCementByUser().then((res) => {
          if (res) {
            res.anntMsgList.forEach(e => {
              this.announcement1.push(e)
            });
            this.msg1Count = res.anntMsgTotal
            this.msg1Title = '通知公告(' + res.anntMsgTotal + ')'
            res.sysMsgList.forEach(e => {
              this.announcement2.push(e)
            });
            this.msg2Count = res.sysMsgTotal
            this.msg2Title = '系统消息(' + res.sysMsgTotal + ')'
            this.$store.commit('SET_MESSAGENUM', res.anntMsgTotal + res.sysMsgTotal)
            console.log(this.announcement1)
            console.log(this.announcement2)
          }
        }).catch(error => {
          console.log('系统消息通知异常', error)// 这行打印permissionName is undefined
          this.stopTimer = true
          console.log('清理timer')
        })
      } catch (err) {
        this.stopTimer = true
        console.log('通知异常', err)
      }
    },

    //  websocket相关操作
    initWebSocket: function () {
      // WebSocket与普通的请求所用协议有所不同，ws等同于http，wss等同于https
      var url = window._CONFIG.domianURL.replace('https://', 'wss://').replace('http://', 'ws://') + '/websocket/' + this.userInfo.id
      // console.log(url);
      this.websock = new WebSocket(url)
      this.websock.onopen = this.websocketOnopen
      this.websock.onerror = this.websocketOnerror
      this.websock.onmessage = this.websocketOnmessage
      this.websock.onclose = this.websocketOnclose
    },
    websocketOnopen: function () {
      console.log('WebSocket连接成功')
      // 心跳检测重置
      // this.heartCheck.reset().start();
    },
    websocketOnerror: function (e) {
      console.log('WebSocket连接发生错误')
      this.reconnect()
    },
    websocketOnmessage: function (e) {
      this.$store.commit('SET_MESSAGENUM', this.messageNum + 1)
      console.log('-----接收消息-------', e.data)
      var data = eval('(' + e.data + ')') // 解析对象
      if (data.msgCategory == 1) { // 通知
        this.announcement1.push(data)
        this.msg1Count++
        this.msg1Title = '通知(' + this.msg1Count + ')'
      } else if (data.msgCategory == 2) {
        this.announcement2.push(data)
        this.msg2Count++
        this.msg2Title = '系统消息(' + this.msg2Count + ')'
      }

      // 心跳检测重置
      // this.heartCheck.reset().start();
    },
    websocketOnclose: function (e) {
      console.log('connection closed (' + e + ')')
      if (e) {
        console.log('connection closed (' + e.code + ')')
      }
      this.reconnect()
    },
    websocketSend (text) { // 数据发送
      try {
        this.websock.send(text)
      } catch (err) {
        console.log('send failed (' + err.code + ')')
      }
    },

    openNotification (data) {
      var text = data.msgTxt
      const key = `open${Date.now()}`
      this.$notification.open({
        message: '消息提醒',
        placement: 'bottomRight',
        description: text,
        key,
        btn: (h) => {
          return h('a-button', {
            props: {
              type: 'primary',
              size: 'small'
            },
            on: {
              click: () => this.showDetail(key, data)
            }
          }, '查看详情')
        }
      })
    },

    reconnect () {
      var that = this
      if (that.lockReconnect) return
      that.lockReconnect = true
      // 没连接上会一直重连，设置延迟避免请求过多
      setTimeout(function () {
        console.info('尝试重连...')
        that.initWebSocket()
        that.lockReconnect = false
      }, 5000)
    },
    heartCheckFun () {
      var that = this
      // 心跳检测,每20s心跳一次
      that.heartCheck = {
        timeout: 20000,
        timeoutObj: null,
        serverTimeoutObj: null,
        reset: function () {
          clearTimeout(this.timeoutObj)
          // clearTimeout(this.serverTimeoutObj);
          return this
        },
        start: function () {
          var self = this
          this.timeoutObj = setTimeout(function () {
            // 这里发送一个心跳，后端收到后，返回一个心跳消息，
            // onmessage拿到返回的心跳就说明连接正常
            that.websocketSend('HeartBeat')
            console.info('客户端发送心跳')
            // self.serverTimeoutObj = setTimeout(function(){//如果超过一定时间还没重置，说明后端主动断开了
            //  that.websock.close();//如果onclose会执行reconnect，我们执行ws.close()就行了.如果直接执行reconnect 会触发onclose导致重连两次
            // }, self.timeout)
          }, this.timeout)
        }
      }
    }

  },
  computed: {
    ...mapState(['menuData', 'userInfo', 'messageNum'])
  },
  // 生命周期 - 创建完成（访问当前this实例）
  created () {
    this.config({
      module: '我的消息',
      api: '/message/usermessage'
    })
    // this.initWebSocket()
    // this.loadData()
    // this.announcement1.push({
    //   titile: "标题1",
    //   msgContent:"<p>内容内容内容</p>",
    //   msgCategory: 1,
    //   sender: "admin",
    //   priority: "M",
    //   sendTime: "2021-10-28 19:30:04",
    //   msgAbstract: "摘要摘要",
    // });
    // this.announcement1.push({
    //   titile: "标题1",
    //   msgContent:"<p>内容内容内容</p>",
    //   sender: "admin",
    //   msgCategory: 1,
    //   priority: "M",
    //   sendTime: "2021-10-28 19:30:04",
    //   msgAbstract: "摘要摘要",
    // });
    // this.$store.commit("SET_MESSAGENUM", 2);
  },
  destroyed () {
    this.websocketOnclose()
  }
}
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */

.main{
  display: inline;
  text-align: left;
  margin-right: 30px;
}
.content{
  padding-left: 20px;
  overflow: auto;
  h4{
    margin-top: 0 !important;
    color: rgba(0,0,0,.65) !important;
    a{
      cursor:pointer;
    }
  }

}
ol{
    list-style:none;
}

</style>
