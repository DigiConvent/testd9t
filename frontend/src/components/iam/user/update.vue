<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else-if="user_read" class="card flex justify-center">
      <Form class="flex flex-col gap-4 w-full" @submit="handle_submit">
         <h2>{{ t("iam.user.update.title") }}</h2>
         <FormTextInput v-model="user_read.emailaddress" label="iam.user.create" name="email" />
         <FormTextInput v-model="user_read.first_name" label="iam.user.create" name="first_name" />
         <FormTextInput v-model="user_read.last_name" label="iam.user.create" name="last_name" />
         <Button
            type="submit"
            severity="secondary"
            :label="$t('iam.user.update.submit')"
            @click="handle_submit"
         />
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { useI18n } from "vue-i18n"
import FormTextInput from "@/components/form/text_input.vue"

import type { UserIdOrUserRead, UserRead } from "@/api/iam/user/types"
import { api } from "@/api"
import { error } from "@/composables/toast"
import { ref } from "vue"

const props = defineProps<UserIdOrUserRead>()

const t = useI18n().t

const loading = ref(false)
const user_read = ref<UserRead>()
async function load_user() {
   loading.value = true
   if (props.id === undefined) {
      user_read.value = props.data
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
   ;(await api.iam.user.update(user_read.value!, props.id)).fold(
      (err: string) => {
         error(err)
      },
      () => {},
   )
}
</script>
