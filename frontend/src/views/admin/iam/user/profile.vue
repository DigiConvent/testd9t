<template>
   <!-- this is external profile watching -->
   <UserProfileComponent :id="id" />
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router"
import UserProfileComponent from "@/components/iam/user/profile.vue"
import { ref } from "vue"

const id = ref(useRoute().params.id as string)
</script>
