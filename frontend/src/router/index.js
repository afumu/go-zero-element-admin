import Vue from 'vue'
import VueRouter from 'vue-router'
import AppLayout from '@/layouts/AppLayout'
import { getUserInfo } from '@/api/system/common'
import { getAll } from '@/api/system/dict'
import  cache from '@/plugins/cache'
const Login = () => import('@/views/login')
const ErrorNoPermissions = () => import('@/views/no-permissions')
const Error404 = () => import('@/views/not-found')
Vue.use(VueRouter)

const routerMode = 'hash'
const router = new VueRouter({
  // base: process.env.VUE_APP_CONTEXT_PATH + (routerMode === 'hash' ? '#' : ''),
  mode: routerMode,
  routes: [
    // 登录
    {
      name: 'login',
      path: '/login',
      component: Login
    },
    // 无权限
    {
      name: 'no-permissions',
      path: '/no-permissions',
      component: ErrorNoPermissions
    },
    // 404
    {
      name: 'not-found',
      path: '/not-found',
      component: Error404
    },
    // 内容页（动态加载）
    {
      name: 'layout',
      path: '',
      component: AppLayout,
      children: []
    }
  ]
})
router.beforeEach((to, from, next) => {
  console.info("----------------------------before-------------------------")
  // 无权访问&404页面可直接访问
  if (to.name === 'no-permissions' || to.name === 'not-found') {
    next()
    return
  }

  //从vuex中获取用户信息，如果刷新页面，vuex数据失效
  const userInfo = router.app.$options.store.state.userInfo
  if (userInfo != null) {
    // 如果用户不存在权限
    if (userInfo.permissions.length === 0) {
      next({ name: 'no-permissions' })
      return
    }

    //处理字典被删除
/*    let dict = cache.session.getJSON("dict")
    if(dict == null){
      getAll()
      .then(date => {
        cache.session.setJSON("dict",date)
      })
      .catch(e => {
      })
    }*/

    next()
    return
  }
  // 登录页面处理
  if (to.name === 'login') {
    next()
    return
  }

  // 非登录页面，验证用户是否登录
  getUserInfo()
    .then(userInfo => {
      // console.log(userInfo);
      // 如果用户不存在权限
      if (userInfo.permissions.length === 0) {
        next({ name: 'no-permissions' })
        return
      }
      // 已登录，存储userInfo
      router.app.$store.commit('setUserInfo', userInfo)
      // 判断缓存中是否存在字典，如果不存在重新获取
      // let dict = cache.session.getJSON("dict")
      // if(dict==null){
      //   getAll().then(date=>{
      //     cache.session.setJSON("dict",date)
      //   })
      // }
      next()
    })
    .catch(e => {
      console.log(e);
      // 未登录，跳转至登录页
      next({ name: 'login' })
    })
})

export default router
