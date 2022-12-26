import request from '@/utils/request'

// 查询
export function fetchList(data) {
    return request.get('/business/slideshow/list', {params: data})
}

// 新建
export function create(data) {
    return request.post('/business/slideshow/add',  data)
}

// 修改
export function updateEnable(data) {
    return request.put('/business/slideshow/enable', data)
}

// 删除
export function deleteById(id) {
    return request.delete('/business/slideshow/delete', {params: {'id': id}})
}

// 上传
export function deleteByIdInBatch(date) {
    return request.post('//system/sysFile/upload', {params: data})
}