import request from '@/utils/request'

// 查询列表
export function fetchList (data) {
  return request.get('/sys/position/list',{ params: data })
}

// 新建
export function create (data) {
  return request.post('/sys/position/add', data)
}

// 修改
export function updateById (data) {
  return request.put('/sys/position/edit', { params: data })
}

// 删除
export function deleteById (data) {
  return request.delete('/sys/position/delete',{params: {
    'id': data
  }})
}

// 批量删除
export function deleteByIdInBatch (ids) {
  return request.delete('/sys/position/deleteBatch', {
    params: {
      ids
    }
  })
}

