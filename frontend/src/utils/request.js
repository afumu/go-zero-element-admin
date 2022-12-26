import axios from 'axios'
import Cookies from 'js-cookie'
import { trim } from './util'
import cache from '../plugins/cache'
import router from '../router/index'
import { msg } from 'element-ui'

axios.defaults.headers.common['Content-Type'] = 'application/json;charset=UTF-8'
const axiosInstance = axios.create({
  baseURL: process.env.VUE_APP_API_PREFIX,
  // 请求超时时间
  timeout: 60000
})

// 新建请求拦截器
axiosInstance.interceptors.request.use(config => {
  // 参数去空格
  if (config.trim === true) {
    if (config.data != null) {
      config.data = trim(config.data)
    }
    if (config.params != null) {
      config.params = trim(config.params)
    }
  }
  // 导出处理
  if (config.download === true) {
    config.responseType = 'blob'
  }
  // 设置操作平台
  // config.headers['eva-platform'] = `pc-${pkg.version}`
  // 设置认证头
  const authToken = Cookies.get('X-Access-Token')
  if (authToken != null) {
    config.headers['X-Access-Token'] = authToken
  }
  return config
}, function (error) {
  return Promise.reject(error)
})

// 新建响应拦截器
axiosInstance.interceptors.response.use((response) => {
    console.log("response:",response);
  //如果是性能监控服务，直接返回相应数据
  if(response.config.url.search("actuator/") != -1){
    return response.data
  }
//   if(response.data.msg ==='该值可用！'){
//     if(!response.config.params.dataId){
//         return Promise.reject(new Error('该值系统不存在'))
//     }
//   }
  if(response.data.msg ==='该值不可用，系统中已存在！'){
    if(!response.config.params.dataId){
        return Promise.reject(new Error('该值可用'))
    }
  }

  if (response.data.code === 500) {
    // msg.error(response.data.msg)
    return Promise.reject(new Error(response.data.msg))
  }

  // 401
  if (response.data.code === 401) {
    router.replace({ path: '/login' });
    return Promise.reject(response.data.msg)
  }

  // 404
  if (response.status === 404) {
    return Promise.reject(new Error('找不到服务！'))
  }

  // 业务失败
  if (response.data.code !== 200) {
    return Promise.reject(response.data)
  }

  if (!response.data.data) {
     return response.data
  }

  return response.data.data
}, function (error) {
  //http error
  if (error.response.status === 401) {
    router.replace({ path: '/login' });
    return Promise.reject(new Error('未授权，请重新登录'))
  }
  if (error.response.status === 404) {
    return Promise.reject(new Error('很抱歉，资源未找到'))
  }
  if (error.response.status === 500) {
    if (error.response.data.msg === 'Token失效，请重新登录') {
      router.replace({ path: '/login' });
      return Promise.reject(new Error('身份认证失效，请重新登录'))
    } else {
      console.log(error.response)
      return Promise.reject(new Error('服务器出现错误，请稍后重试'))
    }
  }
  if (error.response.status === 504) {
    return Promise.reject(new Error('服务器响应超时，请稍后再试'))
  }
  return Promise.reject(error)
})

// 缓存请求结果
const buildCachePromise = (cacheKey, method, args, cacheImpl) => {
  return {
    __cacheImpl: cache[cacheImpl],
    __arguments: args,
    __result_promise: null,
    // 开启缓存
    cache () {
      const data = this.__cacheImpl.getJSON(cacheKey)
      if (data != null) {
        this.__result_promise = Promise.resolve(data)
      }
      if (this.__result_promise != null) {
        return this.__result_promise
      }
      return this
    },
    then () {
      return this.__access('then', arguments)
    },
    catch () {
      return this.__access('catch', arguments)
    },
    finally () {
      return this.__access('finally', arguments)
    },
    __access (methodName, args) {
      if (this.__result_promise != null) {
        return this.__result_promise
      }
      this.__result_promise = axiosInstance[method].apply(axiosInstance, this.__arguments)
      this.__result_promise.then(data => {
        this.__cacheImpl.setJSON(cacheKey, data)
        return data
      })
      return this.__result_promise[methodName].apply(this.__result_promise, args)
    }
  }
}
const methods = ['get', 'post', 'delete', 'put', 'head', 'options', 'patch', 'request']
axiosInstance.cache = function (cacheKey, isLocal = false) {
  if (cacheKey == null) {
    throw Error('Request cache key can not be null.')
  }
  const cacheAxiosInstance = {}
  for (const method of methods) {
    cacheAxiosInstance[method] = function () {
      return buildCachePromise(cacheKey, method, arguments, isLocal ? 'local' : 'session')
    }
  }
  return cacheAxiosInstance
}

export default axiosInstance
