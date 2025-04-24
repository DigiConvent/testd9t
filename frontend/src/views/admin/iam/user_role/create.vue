<template>
   <CreateUserRole :parent="parent" />
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router"
import CreateUserRole from "@/components/pkg/iam/user_role/create.vue"
import { ref } from "vue"

const params = useRoute().params

const parent = ref(Array.isArray(params.parent) ? params.parent[0] : params.parent)
</script>
