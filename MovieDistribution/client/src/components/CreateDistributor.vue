<template>
  <div class="card">
    <h2>Create Distributor</h2>

    <input v-model="name" placeholder="Distributor Name" />
    <input v-model="parent" placeholder="Parent (optional)" />

    <button @click="create">Create</button>
  </div>
</template>

<script setup>
import { ref } from "vue"
import api from "../api"

const name = ref("")
const parent = ref("")

async function create() {
  if (!name.value) {
    alert("Enter distributor name")
    return
  }

  try {
    await api.post("/distributor", {
      name: name.value,
      parent: parent.value
    })

    alert("Distributor created ✅")

    name.value = ""
    parent.value = ""

  } catch (err) {
    alert(err.response?.data?.error || "Server error")
  }
}
</script>

<style scoped>
.card {
  border: 1px solid #ddd;
  padding: 15px;
  margin-bottom: 20px;
}
</style>