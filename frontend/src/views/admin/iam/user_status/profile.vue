<template>
   <UserStatusProfile :id="id" />
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router"
import UserStatusProfile from "@/components/pkg/iam/user_status/profile.vue"
import { ref } from "vue"

const params = useRoute().params

const id = ref(Array.isArray(params.id) ? params.id[0] : params.id)
</script>
