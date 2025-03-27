<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else-if="ur" class="card flex justify-center">
      <Form class="flex flex-col gap-4 w-full" @submit="handle_submit">
         <h2>{{ t("iam.user.create.title") }}</h2>
         <FormTextInput v-model="ur.first_name" label="iam.user.create" name="first_name" />
         <FormTextInput v-model="ur.last_name" label="iam.user.create" name="last_name" />
         <FormMaskInput
            id="date_of_birth"
            v-model="ur.date_of_birth"
            label="iam.user.create"
            name="date_of_birth"
            mask="99/99/9999"
            slot-char="dd/mm/yyyy"
         />
         <Button type="submit" severity="secondary" :label="$t('iam.user.update.submit')" />
      </Form>
   </div>
</template>

<script lang="ts" setup>
import * as v from "valibot"
import { useI18n } from "vue-i18n"
import FormTextInput from "@/components/form/text_input.vue"
import FormMaskInput from "@/components/form/mask_input.vue"

import type { UserRead } from "@/api/iam/user/types"
import { api } from "@/api"
import { error } from "@/composables/toast"
import { ref, watch } from "vue"

const props = defineProps<{ id: string }>()

const t = useI18n().t

const date_of_birth_check = v.pipe(
   v.custom(
      (value: unknown) => {
         const regex = /^\d{2}\/\d{2}\/\d{4}$/
         if (!regex.test(value as string)) {
            return false
         }
         const segments = (value as string).split("/")
         const day = parseInt(segments[0])
         const month = parseInt(segments[1])
         const year = parseInt(segments[2])
         const date = new Date(year, month - 1, day)
         if (
            date.getFullYear() !== year ||
            date.getMonth() !== month - 1 ||
            date.getDate() !== day
         ) {
            return false
         }
         return true
      },
      t("iam.user.create.invalid", { field: t("iam.user.create.date_of_birth") }),
   ),
)

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

watch(
   ur,
   () => {
      try {
         v.parse(date_of_birth_check, ur.value?.date_of_birth)
      } catch {
         //illegal
      }
   },
   { deep: true },
)

load_user()
const handle_submit = async () => {}
</script>
