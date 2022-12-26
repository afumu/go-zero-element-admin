import request from '@/utils/request'

// 查询
export function fetchList(data) {
    return request.get('/sys/user/list', {params: data})
}
// 查询
export function addUser(data) {
    return request.post('/system/userGroup/addUser', data)
}

