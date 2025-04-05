<template>
   <CreateUserStatus :parent="parent" />
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router"
import CreateUserStatus from "@/components/iam/user_status/create.vue"
import { ref } from "vue"

const params = useRoute().params

const parent = ref(Array.isArray(params.parent) ? params.parent[0] : params.parent)
</script>
