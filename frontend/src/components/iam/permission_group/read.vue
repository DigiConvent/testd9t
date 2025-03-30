<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else v-permission="'iam.permission_group.read'">
      read permission group - to be implemented
   </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"

const loading = ref(true)
</script>
