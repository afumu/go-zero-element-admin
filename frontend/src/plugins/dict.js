
import cache from './cache'
export default {
  getAllDictItemByCode(code) {
    let allDictItem = cache.session.getJSON("dict")
    for (const allDictItemKey in allDictItem) {
      if (allDictItemKey === code) {
        return allDictItem[code]
      }
    }
  },
  getDictItemText(code, value) {
    let text
    let rv = this.getAllDictItemByCode(code)
    rv.forEach(element => {
      if (element.value === value) {
        text = element.text
      }
    })
    return text
  },
  getDictItemValue(code, text) {
    let value
    let rv = this.getAllDictItemByCode(code)
    rv.forEach(element => {
      if (element.text === text) {
        value = element.value
      }
    })
    return value
  }
}
