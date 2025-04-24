<template>
   <UserUpdate :id="id" />
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router"
import UserUpdate from "@/components/pkg/iam/user/update.vue"
import { ref } from "vue"

const id = ref(useRoute().params.id as string)
</script>
