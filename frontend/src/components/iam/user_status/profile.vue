<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
</template>
<script lang="ts" setup>
import { ref } from "vue"

const loading = ref<boolean>(true)
</script>
