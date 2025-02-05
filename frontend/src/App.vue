<template>
  <header>
    <div v-if="status != null">
      {{ status }}
    </div>
  </header>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import api from './api'
import type { SystemStatus } from './api/sys/sys.get_status';

let status = ref<SystemStatus>()

api.sys.getStatus().then(result => {
  result.fold(() => {
    console.log('API is down')
  }, (data) => {
    console.log('API is up:'+ data)
    status.value = data;
  })
});
</script>