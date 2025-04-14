<template>
   <UserRoleProfile :id="id" />
</template>

<script lang="ts" setup>
import UserRoleProfile from "@/components/iam/user_role/profile.vue"
import { useRoute } from "vue-router"

const ids = useRoute().params.id
const id = Array.isArray(ids) ? ids[0] : ids
</script>
