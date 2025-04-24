<template>
   <CreatePermissionGroup :parent="parent" />
</template>

<script lang="ts" setup>
import CreatePermissionGroup from "@/components/pkg/iam/permission_group/create.vue"
import { useRoute } from "vue-router"

const args = useRoute().params.parent
const parent = Array.isArray(args) ? args[0] : args
</script>
