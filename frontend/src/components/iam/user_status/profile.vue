<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <div v-else-if="profile" v-permission="'iam.user_status.read'">
         <pre>{{ JSON.stringify(profile, null, 3) }}</pre>
      </div>
   </div>
</template>
<script lang="ts" setup>
import { api } from "@/api"
import type { UserStatusProfile } from "@/api/iam/user_status/types"
import { error } from "@/composables/toast"
import { ref } from "vue"

const props = defineProps<{ id: string }>()
const loading = ref<boolean>(true)
const profile = ref<UserStatusProfile>()

async function load() {
   ;(await api.iam.user_status.profile(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (data: UserStatusProfile) => {
         profile.value = data
         loading.value = false
      },
   )
}

load()
</script>
