// list: '/sys/user/queryByOrgCodeForAddressList',
//           listByPosition: '/sys/position/list'
// /sys/sysDepart/queryTreeList
import request from '@/utils/request'
export function fetchList (data) {
    return request.get('/sys/user/queryByOrgCodeForAddressList', { params: data })
}
export function queryTreeList (data) {
    return request.get('/sys/sysDepart/queryTreeList', { params: data })
}
export function list (data) {
    return request.get('/sys/position/list', { params: data })
}