<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else-if="profile">
      <h1 class="text-xl">{{ profile.user.first_name }} {{ profile.user.last_name }}</h1>
      <UserRead :id="profile.user.id" />
      <template v-if="auth.has_permission('iam.user.update') || user_id == profile.user.id">
         <UserUpdateForm :id="profile.user.id" />
      </template>
      <template v-else-if="auth.has_permission('iam.user.view')">
         You can maybe only view this user
      </template>
      <template v-else>No</template>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { UserProfile } from "@/api/iam/user/types"
import JwtAuthenticator from "@/auth/jwt"
import { error } from "@/composables/toast"
import { ref } from "vue"
import UserUpdateForm from "@/components/iam/user/update.vue"
import UserRead from "@/components/iam/user/read.vue"

const props = defineProps<{ id: string }>()
const loading = ref(true)

const profile = ref<UserProfile | null>(null)
const auth = JwtAuthenticator.get_instance()
const user_id = auth.get_token()?.id

async function load_user_profile() {
   ;(await api.iam.user.profile(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (pg: UserProfile) => {
         profile.value = pg
         loading.value = false
      },
   )
}

load_user_profile()
</script>
