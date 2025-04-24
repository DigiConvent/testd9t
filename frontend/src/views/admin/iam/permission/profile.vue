<template>
   <PermissionProfile :id="id"></PermissionProfile>
</template>
<script lang="ts" setup>
import PermissionProfile from "@/components/pkg/iam/permission/profile.vue"
import { useRoute } from "vue-router"

const name = useRoute().params.name
const id = Array.isArray(name) ? name[0] : name
</script>
