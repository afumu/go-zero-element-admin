<template>
  <TreeSelect
    :placeholder="placeholder"
    :value="value"
    :data="data"
    :clearable="clearable"
    :append-to-body="appendToBody"
    :inline="inline"
    :multiple="multiple"
    :flat="multiple"
    @input="$emit('input', $event)"
  />
</template>

<script>
import TreeSelect from './TreeSelect'
import { fetchTree } from '@/api/base/materialType'

export default {
  name: 'MaterialTypeSelect',
  components: { TreeSelect },
  props: {
    value: {},
    inline: {
      default: true
    },
    multiple: {
      default: false
    },
    placeholder: {
      default: '请选择类型'
    },
    // 是否可清空
    clearable: {
      default: false
    },
    appendToBody: {
      default: false
    },
    // 需被排除的部门ID
    excludeId: {}
  },
  data () {
    return {
      data: []
    }
  },
  watch: {
    excludeId () {
      this.fetchData()
    }
  },
  methods: {
    // 获取所有类型
    fetchData () {
      fetchTree()
        .then(records => {
          this.data = []
          this.__fillData(this.data, records)
        })
        .catch(e => {
          this.$tip.apiFailed(e)
        })
    },
    // 填充类型树
    __fillData (list, pool) {
      for (const type of pool) {
        if (type.id === this.excludeId) {
          continue
        }
        const typeNode = {
          id: type.id,
          label: type.name
        }
        list.push(typeNode)
        if (type.children != null && type.children.length > 0) {
          typeNode.children = []
          this.__fillData(typeNode.children, type.children)
          if (typeNode.children.length === 0) {
            typeNode.children = undefined
          }
        }
      }
    }
  },
  created () {
    this.fetchData()
  }
}
</script>
