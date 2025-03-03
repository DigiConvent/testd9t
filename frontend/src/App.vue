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
    <span>{{ telegramLoginStatus }}</span>
    <div v-if="tgWebApp != null">
      <pre>
        {{ JSON.stringify(tgWebApp, null, 2) }}
      </pre>
    </div>
  </header>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import api from './api'
import type { SystemStatus } from './api/sys/sys.get_status';
import getWebApp, { type TelegramWebApp } from './auth/telegram';

let status = ref<SystemStatus>()
let tgWebApp = ref<TelegramWebApp>()
const telegramLoginStatus = ref<string>('')

api.sys.getStatus().then(result => {
  tgWebApp.value = getWebApp();
  if (tgWebApp.value.initData == "") {
    telegramLoginStatus.value = "Not using telegram as auth"
  } else {
    telegramLoginStatus.value = "Using telegram as auth"
  }
  result.fold(() => {
    console.log('API is down')
  }, (data) => {
    console.log('API is up:' + data)
    status.value = data;
  })
});


</script>