<template>
   <NeedsPermission permission="iam.user.create">
      <Form class="flex flex-col gap-4 w-full" @submit="handle_submit">
         <h2 class="text-lg">{{ t("iam.user.create.title") }}</h2>
         <FormTextInput v-model="email" label="iam.user.create" name="email" />
         <FormTextInput v-model="first_name" label="iam.user.create" name="first_name" />
         <FormTextInput v-model="last_name" label="iam.user.create" name="last_name" />
         <div class="grid grid-cols-2 gap-4">
            <FloatLabel variant="in">
               <UserStatusPicker id="user_status" v-model="user_status"></UserStatusPicker>
               <label for="user_status">{{ $t("iam.user_status.picker.placeholder") }}</label>
            </FloatLabel>
            <FormMaskInput
               v-model="user_status_start"
               label="iam.user.create"
               name="user_status_start"
               mask="99/99/9999"
               slot-char="DD/MM/YYYY"
            />
         </div>
         <Button type="submit" severity="secondary" :label="$t('iam.user.create.submit')" />
      </Form>
   </NeedsPermission>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import * as v from "valibot"
import { useI18n } from "vue-i18n"
import { api } from "@/api"
import UserStatusPicker from "../user_status/picker.vue"
import { error, success } from "@/composables/toast"
import FormTextInput from "@/components/form/text_input.vue"
import FormMaskInput from "@/components/form/mask_input.vue"
import router from "@/router"

const t = useI18n().t

const email = ref<string>("")
const first_name = ref<string>("")
const last_name = ref<string>("")
const user_status = ref<string>("")
const user_status_start = ref<string>("")

const email_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.user.create.required", { field: t("iam.user.create.email") })),
   v.email(t("iam.user.create.invalid", { field: t("iam.user.create.email") })),
   v.toLowerCase(),
)

const first_name_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.user.create.required", { field: t("iam.user.create.first_name") })),
)
const last_name_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.user.create.required", { field: t("iam.user.create.last_name") })),
)

const user_create = v.object({
   email: email_check,
   first_name: first_name_check,
   last_name: last_name_check,
   user_status: v.string(),
   user_status_start: v.string(),
})

const handle_submit = async () => {
   const re = v.safeParse(user_create, {
      email: email.value,
      first_name: first_name.value,
      last_name: last_name.value,
      user_status: user_status.value,
      user_status_start: user_status_start.value,
   })

   if (re.success) {
      ;(
         await api.iam.user.create({
            emailaddress: re.output["email"],
            first_name: re.output["first_name"],
            last_name: re.output["last_name"],
         })
      ).fold(
         (l) => {
            error(l)
         },
         (user_id: string) => {
            router.replace({ name: "iam.user.profile", params: { id: user_id } })
            success(t("iam.user.create.success"), "")
         },
      )
   }
}
</script>
