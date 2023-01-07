import request from '@/utils/request'
import axios from 'axios'
import Cookies from 'js-cookie'


// 根据密码登录
export function loginByPassword(data) {
  return request.post('/sys/user/login', data)
}

// 获取已登录的用户信息
export function getUserInfo() {
  return request.get('/sys/user/queryLoginUser')
}

// 校验值是否已经存在
export function check(data) {
  return request.get('/sys/duplicate/check', { params: data })
}

// 登出
export function logout(data) {
  return request.get('/sys/logout', data)
}

// 修改密码
export function updatePwd(data) {
  return request.put('/sys/user/updatePassword', data)
}

// 获取菜单
export function getMenu(data) {
  return request.get('/sys/menu/queryUserMenu', { params: data })
}

// 导出文件
export function exportFileByCondition(url, data) {
  return axios({
    baseURL: process.env.VUE_APP_API_PREFIX,
    url: url,
    params: data,
    method: 'get',
    responseType: 'blob',
    headers: {
      "X-Access-Token": Cookies.get('X-Access-Token')
  }
  })
}

// 刷新缓存
export function refleshCache() {
  return request.get('/sys/dict/refleshCache')
}
