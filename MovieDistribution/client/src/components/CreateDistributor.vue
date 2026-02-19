<template>
  <div>
    <h3>Create Distributor</h3>

    <input v-model="name" placeholder="Name" />
    <input v-model="parent" placeholder="Parent (optional)" />

    <button @click="create">Create</button>
  </div>
</template>

<script setup>
import { ref } from "vue"
import api from "../api"

const name = ref("")
const parent = ref("")

const create = async () => {
  await api.post("/distributor", {
    name: name.value,
    parent: parent.value
  })
  alert("Created")
}
</script>
