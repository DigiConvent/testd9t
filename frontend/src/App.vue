<template>
  <header>
    <div v-if="status != null">
      <table>
        <tr>
          <td>Version:</td>
          <td>{{ status.version }}</td>
        </tr>
        <tr>
          <td>Uptime:</td>
          <td>{{ status.online_since }}</td>
        </tr>
        <tr>
          <td>Disk status:</td>
          <td>{{ (100 / status.total_space) * (status.total_space - status.free_space) }}%</td>
        </tr>
      </table>
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