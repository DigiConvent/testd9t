<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else-if="profile">
      <h1 class="text-xl">{{ profile.user.first_name }} {{ profile.user.last_name }}</h1>
      <UserRead :id="profile.user.id" />
      <NeedsPermission permission="iam.user.set_enabled">
         <FormSwitch
            v-model="profile.user.enabled"
            :label_on="$t('iam.user.enabled')"
            :label_off="$t('iam.user.disabled')"
            icon_on="user-unlock"
            icon_off="user-lock"
            :loading="loading_enabled"
            @update:model-value="set_enabled"
         />
      </NeedsPermission>
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
import { error, success } from "@/composables/toast"
import { ref } from "vue"
import UserUpdateForm from "@/components/iam/user/update.vue"
import UserRead from "@/components/iam/user/read.vue"
import FormSwitch from "@/components/form/switch.vue"
import { useI18n } from "vue-i18n"

const props = defineProps<{ id: string }>()
const loading = ref(true)
const loading_enabled = ref(false)

const profile = ref<UserProfile | null>(null)
const auth = JwtAuthenticator.get_instance()
const user_id = auth.get_token()?.id

const t = useI18n().t

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

async function set_enabled(value: boolean) {
   loading_enabled.value = true
   ;(await api.iam.user.set_enabled(props.id, value)).fold(
      () => {},
      () => {
         success(t("iam.user.set_enabled.success"))
      },
   )
   loading_enabled.value = false
}
</script>
