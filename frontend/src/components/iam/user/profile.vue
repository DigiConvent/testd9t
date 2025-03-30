<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div
      v-else-if="profile"
      v-permission="'iam.user.read'"
      v-permission.except="true"
      class="flex flex-col gap-4"
   >
      <ReadUser :id="undefined" :data="profile.user" />
      <FormSwitch
         v-model="profile.user.enabled"
         v-permission="'iam.user.set_enabled'"
         :label_on="$t('iam.user.enabled')"
         :label_off="$t('iam.user.disabled')"
         icon_on="user-unlock"
         icon_off="user-lock"
         :loading="loading_enabled"
         @update:model-value="set_enabled"
      />
      <SetPassword
         :id="is_loggedin_user ? 'me' : profile.user.id"
         v-permission="'iam.user.set_password'"
         v-permission.except="is_loggedin_user"
         @success="password_set"
      ></SetPassword>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { UserProfile } from "@/api/iam/user/types"
import JwtAuthenticator from "@/auth/jwt"
import { error, success } from "@/composables/toast"
import { computed, ref } from "vue"
import ReadUser from "@/components/iam/user/read.vue"
import FormSwitch from "@/components/form/switch.vue"
import SetPassword from "@/components/iam/user/set_password.vue"
import { useI18n } from "vue-i18n"

const props = defineProps<{ id?: string }>()
const loading = ref(true)
const loading_enabled = ref(false)

const profile = ref<UserProfile | null>(null)
const auth = JwtAuthenticator.get_instance()

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
   ;(await api.iam.user.set_enabled(props.id!, value)).fold(
      () => {},
      () => {
         success(t("iam.user.set_enabled.success"))
      },
   )
   loading_enabled.value = false
}

const is_loggedin_user = computed(() => {
   return profile.value != null && profile.value.user.id == auth.get_token()?.id
})
function password_set() {
   success(t("iam.user.set_password.success"))
}
</script>
