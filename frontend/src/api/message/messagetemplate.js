import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('/business/noticeTemplate/list', { params: data })
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
