<template>
  <div>
    <h2>Region Tree</h2>

    <Tree :nodes="treeData" />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"
import api from "../api"
import Tree from "vue3-tree-vue"

const treeData = ref([])

onMounted(async () => {
  const res = await api.get("/regions")

  // Simple demo tree
  treeData.value = res.data.map(r => ({
    id: r,
    label: r
  }))
})
</script>