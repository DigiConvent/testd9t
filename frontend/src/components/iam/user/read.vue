<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <NeedsPermission v-else-if="user_read" :permission="'iam.user.read'">View user</NeedsPermission>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { UserRead } from "@/api/iam/user/types"
import { error } from "@/composables/toast"
import { ref } from "vue"

const props = defineProps<{ id: string }>()
const loading = ref(true)
const user_read = ref<UserRead | null>(null)

async function load_user() {
   ;(await api.iam.user.get(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (data: UserRead) => {
         user_read.value = data
         loading.value = false
      },
   )
}
load_user()
</script>
