<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else-if="profile" v-permission="'iam.user_role.read'">This is a user wowe</div>
</template>
<script lang="ts" setup>
import { api } from "@/api"
import type { UserRoleProfile } from "@/api/iam/user_role/types"
import { error } from "@/composables/toast"
import { ref } from "vue"

const loading = ref<boolean>(true)

const props = defineProps<{ id: string }>()
const profile = ref<UserRoleProfile>()

async function load() {
   loading.value = true
   ;(await api.iam.user_role.profile(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (data: UserRoleProfile) => {
         profile.value = data
         loading.value = false
      },
   )
   loading.value = false
}

load()
</script>
