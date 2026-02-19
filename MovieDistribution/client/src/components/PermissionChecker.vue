<template>
  <div>
    <h3>Check Permission</h3>

    <input v-model="name" placeholder="Distributor" />
    <input v-model="region" placeholder="Region Code" />

    <button @click="check">Check</button>

    <h2 v-if="result !== null">
      {{ result ? "✅ Allowed" : "❌ Denied" }}
    </h2>
  </div>
</template>

<script setup>
import { ref } from "vue"
import api from "../api"

const name = ref("")
const region = ref("")
const result = ref(null)

const check = async () => {
  const res = await api.post("/check", {
    name: name.value,
    region: region.value
  })
  result.value = res.data.allowed
}
</script>
