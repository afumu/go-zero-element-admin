import { exportFileByCondition } from '@/api/system/common'

// 去空
export function trim (data) {
    if (data == null) {
      return null
    }
    if (typeof data === 'string') {
      return data.trim()
    }
    if (data instanceof Array) {
      for (const item of data) {
        trim(item)
      }
    }
    if (typeof data === 'object') {
      for (const key in data) {
        data[key] = trim(data[key])
      }
    }
    return data
  }
// 导出文件
export function exportFile (url, fileName, condition) {
    if (!fileName || typeof fileName != 'string') {
      fileName = '导出文件'
    }
    exportFileByCondition(url, condition).then(data => {
      if (!data) {
        message.warning('文件下载失败')
        return
      }
      if (typeof window.navigator.msSaveBlob !== 'undefined') {
        window.navigator.msSaveBlob(new Blob([data.data], { type: 'application/vnd.ms-excel' }), fileName + '.xls')
      } else {
        let url = window.URL.createObjectURL(new Blob([data.data], { type: 'application/vnd.ms-excel' }))
        let link = document.createElement('a')
        link.style.display = 'none'
        link.href = url
        link.setAttribute('download', fileName + '.xls')
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link) //下载完成移除元素
      }
    })
  }
  export function timestampToTime(timestamp) {
    var  date =  new  Date(timestamp); //时间戳为10位需*1000，时间戳为13位的话不需乘1000
    var Y = date.getFullYear() +  '-' ;
    var M = (date.getMonth()+1 < 10 ?  '0' +(date.getMonth()+1) : date.getMonth()+1) +  '-' ;
    var D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate() )+  ' ' ;
    var h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours() )+  ':' ;
    var m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes() ) +  ':' ;
    var s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
    return  Y+M+D+h+m+s;
}
/**函数节流方法
 @param Function fn 延时调用函数
 @param Number dalay 延迟多长时间
 @param Number atleast 至少多长时间触发一次
 @return Function 延迟执行的方法
 **/
 export function throttle(fn, delay, atleast) {
   let timer = null;
   let previous = null;
   return function () {
     var now = +new Date();
     if (!previous) previous = now;
     if (atleast && now - previous > atleast) {
       fn();
       // 重置上一次开始时间为本次结束时间
       previous = now;
       clearTimeout(timer);
     } else {
       clearTimeout(timer);
       timer = setTimeout(function () {
         fn();
         previous = null;
       }, delay);
     }
   }
 }
