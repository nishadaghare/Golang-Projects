<template>
  <div>
    <v-select
      :options="regions"
      v-model="selected"
      placeholder="Search Region Code"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"
import api from "../api"
import vSelect from "vue-select"
import "vue-select/dist/vue-select.css"

const selected = ref(null)
const regions = ref([])

onMounted(async () => {
  const res = await api.get("/regions")
  regions.value = res.data
})

defineExpose({ selected })
</script>