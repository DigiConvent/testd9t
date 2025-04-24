<template>
   <div>
      <Form
         v-permission="'iam.user.write'"
         class="flex flex-col gap-4 w-full"
         @submit="handle_submit"
      >
         <h2 class="text-lg">{{ t("iam.user.create.title") }}</h2>
         <FormTextInput v-model="user_create.emailaddress" label="iam.user.fields" name="email" />
         <FormTextInput
            v-model="user_create.first_name"
            label="iam.user.fields"
            name="first_name"
         />
         <FormTextInput v-model="user_create.last_name" label="iam.user.fields" name="last_name" />
         <div class="grid grid-cols-2 gap-4">
            <FloatLabel variant="in">
               <UserStatusPicker
                  id="user_status"
                  v-model="user_create.user_status"
               ></UserStatusPicker>
               <label for="user_status">{{ $t("iam.us.picker.placeholder") }}</label>
            </FloatLabel>
            <FormMaskInput
               v-model="user_status_start"
               label="iam.user.fields"
               name="user_status_start"
               mask="99/99/9999"
               slot-char="DD/MM/YYYY"
            />
         </div>
         <Button type="submit" severity="secondary" :label="$t('iam.user.create.submit')" />
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { useI18n } from "vue-i18n"
import { api } from "@/api"
import UserStatusPicker from "../user_status/picker.vue"
import { error, success } from "@/composables/toast"
import FormTextInput from "@/components/form/text_input.vue"
import FormMaskInput from "@/components/form/mask_input.vue"
import router from "@/router"
import { type UserCreate } from "@/api/iam/user/types"

const t = useI18n().t

const user_status_start = ref<string>("")

const user_create = ref<UserCreate>({
   emailaddress: "",
   first_name: "",
   last_name: "",
   user_status: "",
   when: new Date(),
})

const handle_submit = async () => {
   const date = user_status_start.value.split("/")
   // make sure this is timezone aware otherwise it will shift in the database
   const format = `${date[0]}-${date[1]}-${date[2]}T00:00:00Z`
   const since = new Date(format)

   user_create.value.when = since
   ;(await api.iam.user.create(user_create.value)).fold(
      (l) => {
         error(l)
      },
      (user_id: string) => {
         router.replace({ name: "admin.iam.user.profile", params: { id: user_id } })
         success(t("iam.user.create.success"), "")
      },
   )
}
</script>
