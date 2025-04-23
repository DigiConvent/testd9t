<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
      <div v-else-if="user_read" class="card flex justify-center">
         <Form class="flex flex-col gap-4 w-full" @submit="handle_submit">
            <h2>{{ t("iam.user.update.title") }}</h2>
            <FormTextInput v-model="user_read.emailaddress" label="iam.user.fields" name="email" />
            <FormTextInput
               v-model="user_read.first_name"
               label="iam.user.fields"
               name="first_name"
            />
            <FormTextInput v-model="user_read.last_name" label="iam.user.fields" name="last_name" />
            <Button type="submit" severity="secondary" :label="$t('iam.user.update.submit')" />
         </Form>
      </div>
   </div>
</template>

<script lang="ts" setup>
import { useI18n } from "vue-i18n"
import FormTextInput from "@/components/form/text_input.vue"

import type { UserRead } from "@/api/iam/user/types"
import { api } from "@/api"
import { error, success } from "@/composables/toast"
import { ref } from "vue"
import type { IdOrData } from "@/components/form/form"
import router from "@/router"

const props = defineProps<IdOrData<UserRead>>()

const t = useI18n().t

const loading = ref(false)
const user_read = ref<UserRead>()
async function load_user() {
   loading.value = true
   if (props.id === undefined) {
      user_read.value = props.data!
      loading.value = false
      return
   }

   ;(await api.iam.user.get(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (user_read_data: UserRead) => {
         user_read.value = user_read_data
         loading.value = false
      },
   )
}

load_user()
async function handle_submit() {
   ;(await api.iam.user.update(props.id, user_read.value)).fold(
      (err: string) => {
         error(t("feedback.-.update", { entity: t("iam.user.user") }), err)
      },
      () => {
         success(t("feedback.+.update", { entity: t("iam.user.user") }))
         router.back()
      },
   )
}
</script>
