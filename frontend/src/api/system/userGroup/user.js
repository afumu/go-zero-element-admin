import request from '@/utils/request'

// 查询
export function fetchList(data) {
    return request.get('/system/userGroup/userGroupList', {params: data})
}
// 查询
export function list(data) {
    return request.get('/sys/user/list', {params: data})
}
// 删除
export function deleteById(data) {
    return request.delete('/system/userGroup/deleteUserGroupDetail', {params: {'id': 0,'userGroupId':data.ids,'userId':data.id}})
}


