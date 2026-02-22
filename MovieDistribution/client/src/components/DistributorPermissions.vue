<template>
  <div class="card">
    <h2>Distributor Permissions</h2>

    <input v-model="name" placeholder="Distributor Name" />

    <RegionSearch ref="regionSearch" />

    <div class="buttons">
      <button @click="add('include')">Add INCLUDE</button>
      <button @click="add('exclude')">Add EXCLUDE</button>
    </div>

    <h3>Includes</h3>
    <ul>
      <li v-for="r in includes" :key="r">{{ r }}</li>
    </ul>

    <h3>Excludes</h3>
    <ul>
      <li v-for="r in excludes" :key="r">{{ r }}</li>
    </ul>
  </div>
</template>

<script setup>
import { ref } from "vue"
import api from "../api"
import RegionSearch from "./RegionSearch.vue"

const name = ref("")
const includes = ref([])
const excludes = ref([])

const regionSearch = ref()

async function add(type) {
  const region = regionSearch.value.selected

  if (!region) return alert("Select region")

  await api.post("/permission", {
    name: name.value,
    type,
    region
  })

  if (type === "include") includes.value.push(region)
  else excludes.value.push(region)
}
</script>