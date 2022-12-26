import Vue from 'vue'
import Vuex from 'vuex'
import {Message} from 'element-ui'
Vue.use(Vuex)

const state = {
  // 登录用户信息
  userInfo: null,
  // 首页
  homePage: null,
  // 菜单
  menuData: {
    // 菜单列表
    list: [],
    // 是否收起
    collapse: false
  },
  messageNum: 0,
  paths:[],
  Switch:false
}

const mutations = {
  // 切换菜单状态
  switchCollapseMenu(state, value) {
    if (value != null) {
      state.menuData.collapse = value
    } else {
      state.menuData.collapse = !state.menuData.collapse
    }
    window.localStorage.setItem('MENU_STATUS', state.menuData.collapse)
  },
  // 设置已登录的用户信息
  setUserInfo: (state, data) => {
    state.userInfo = data
  },
  // 设置首页路由信息
  setHomePage(state, homePage) {
    state.homePage = homePage
  },
  // 重置菜单
  resetMenus: (state) => {
    state.menuData.list = []
  },
  // 设置菜单
  SET_MENU: (state, meun) => {
    state.menuData.list = meun
  },
  SET_MESSAGENUM: (state, messageNum) => {
    state.messageNum = messageNum.success ? messageNum.result : messageNum
  },
  //历史菜单
  paths:(state,data) =>{
    for(var i=0; i<state.paths.length; i++){
        if (state.paths[i].title == data.title) {
            return
        }
    }
    state.paths.push(data)
    
  },
  //删除历史菜单
  remove(state,data){
    for(var i=0; i<state.paths.length; i++){
        if (state.paths.length == 1) {
            Message({
                message:'最后一页了, 不能在关闭了',
                type:'warning'
            })
            return
        }
        if (state.paths[i].title == data) {
            state.paths.splice(i,1)
        }
    }
  },
  switch_pattern(state,data){
      state.Switch = data
  }
}
const actions = {}
const getters = {}
export default new Vuex.Store({
  state,
  mutations,
  actions,
  getters
})
