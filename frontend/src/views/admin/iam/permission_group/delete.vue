<template>
   <DeletePermissionGroup :id="id" />
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router"
import DeletePermissionGroup from "@/components/iam/permission_group/delete.vue"

const ids = useRoute().params.id
const id = Array.isArray(ids) ? ids[0] : ids
</script>
