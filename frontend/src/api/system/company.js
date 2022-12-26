import request from '@/utils/request'

// 查询
export function fetchList (data) {
  return request.get('/sys/company/list', { params: data })
}

// 新建
export function create (data) {
  return request.post('/sys/company/add', data, {
    trim: true
  })
}

// 修改
export function updateById (data) {
  return request.put('/sys/company/edit', data, {
    trim: true
  })
}

// 删除
export function deleteById (id) {
  return request.delete('/sys/company/delete', {params: {
    'id': id
  }
  })
}
