<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else-if="ur" class="card flex justify-center">
      <Form class="flex flex-col gap-4 w-full" @submit="handle_submit">
         <h2>{{ t("iam.user.create.title") }}</h2>
         <FormTextInput v-model="ur.first_name" label="iam.user.create" name="first_name" />
         <FormTextInput v-model="ur.last_name" label="iam.user.create" name="last_name" />
         <Button type="submit" severity="secondary" :label="$t('iam.user.update.submit')" />
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { useI18n } from "vue-i18n"
import FormTextInput from "@/components/form/text_input.vue"

import type { UserRead } from "@/api/iam/user/types"
import { api } from "@/api"
import { error } from "@/composables/toast"
import { ref } from "vue"

const props = defineProps<{ id: string }>()

const t = useI18n().t

const loading = ref(false)
const ur = ref<UserRead>()
async function load_user() {
   ;(await api.iam.user.get(props.id)).fold(
      (err: string) => {
         error(err)
      },
      (user_read: UserRead) => {
         ur.value = user_read
         loading.value = false
      },
   )
}

load_user()
const handle_submit = async () => {}
</script>
