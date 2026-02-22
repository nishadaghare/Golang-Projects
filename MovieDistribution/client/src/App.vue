<template>
  <h1>Movie Distribution Permission System</h1>

  <CreateDistributor />

  <DistributorPermissions />

  <PermissionChecker />
  <RegionTree />

</template>

<script setup>
import CreateDistributor from "./components/CreateDistributor.vue"
import DistributorPermissions from "./components/DistributorPermissions.vue"
import PermissionChecker from "./components/PermissionChecker.vue"
import RegionTree from "./components/RegionTree.vue"

</script>