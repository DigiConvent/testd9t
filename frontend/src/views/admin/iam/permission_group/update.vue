<template>
   <UpdatePermissionGroup :id="id" />
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router"
import UpdatePermissionGroup from "@/components/iam/permission_group/update.vue"

const ids = useRoute().params.id

const id = Array.isArray(ids) ? ids[0] : ids
</script>
