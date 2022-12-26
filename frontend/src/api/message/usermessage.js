import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('/sys/sysAnnouncementSend/getMyAnnouncementSend', { params: data })
}

//全部已读
export function readAll () {
  return request.put('/sys/sysAnnouncementSend/readAll')
}


// 新建
export function create (data) {
  return request.post('/business/noticeTemplate/add', data, {
    trim: true
  })
}

// 修改
export function updateById (data) {
  return request.put('/business/noticeTemplate/edit', data, {
    trim: true
  })
}

// 删除
export function deleteById (id) {
  return request.delete('/business/noticeTemplate/delete', {params: {
    'id': id
  }
  })
}

// 获取消息
export function listCementByUser () {
  return request.get('/sys/annountCement/listByUser')
}

// 添加阅读记录
export function readNotic (anntId) {
  return request.get('/sys/annountCement/readNotic',{ params: {"anntId": anntId}})
}