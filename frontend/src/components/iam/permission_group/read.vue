<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <NeedsPermission permission="iam.permission_group.read">
      <div>read permission group - to be implemented</div>
   </NeedsPermission>
</template>

<script lang="ts" setup>
import { ref } from "vue"

const loading = ref(true)
</script>
