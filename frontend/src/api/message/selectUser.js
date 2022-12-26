import request from '@/utils/request'
// 选择用户 - 新建通知

// 查询用户
export function fetchList (data) {
  return request.get('/sys/user/list', { params: data })
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