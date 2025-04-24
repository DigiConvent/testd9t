<template>
   <UpdateUserRole :id="id" />
</template>

<script lang="ts" setup>
import UpdateUserRole from "@/components/pkg/iam/user_role/update.vue"
import { useRoute } from "vue-router"

const params = useRoute().params.id
const id = Array.isArray(params) ? params[0] : params
</script>
