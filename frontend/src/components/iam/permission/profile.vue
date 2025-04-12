<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else-if="profile">
      Loaded profile:
      <pre>{{ JSON.stringify(profile, null, 3) }}</pre>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { type PermissionProfile } from "@/api/iam/permission/types"
// import UserList from "@/components/iam/user/list.vue"
import { error } from "@/composables/toast"
import { ref } from "vue"

const loading = ref<boolean>(true)
const props = defineProps<{ id: string }>()
const profile = ref<PermissionProfile>()

async function load() {
   ;(await api.iam.permission.profile(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (data: PermissionProfile) => {
         profile.value = data
         loading.value = false
      },
   )
}

load()
</script>
