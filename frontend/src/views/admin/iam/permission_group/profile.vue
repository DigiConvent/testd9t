<template>
   <PermissionGroupProfileC :id="id" />
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router"
import PermissionGroupProfileC from "@/components/iam/permission_group/profile.vue"

const ids = useRoute().params.id

const id = Array.isArray(ids) ? ids[0] : ids
</script>
