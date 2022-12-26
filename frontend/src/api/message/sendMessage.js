import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('/sys/annountCement/list', { params: data })
}

// 新建
export function create (data) {
  return request.post('/sys/annountCement/add', data, {
    trim: true
  })
}

// 修改
export function updateById (data) {
  return request.put('/sys/annountCement/edit', data, {
    trim: true
  })
}

// 删除
export function deleteById (id) {
  return request.delete('/sys/annountCement/delete', {params: {
    'id': id
  }
  })
}

//发布通知
export function release (id) {
  return request.get('/sys/annountCement/doReleaseData', { params: {id: id} })
}

// 撤销
export function reovke (id) {
  return request.get('/sys/annountCement/doReovkeData', { params: {id: id} })
}

export function queryByIds (ids) {
  return request.get('/sys/user/queryByIds', { params: {userIds: ids} })
}