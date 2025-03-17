<template>
   <div>
      <div class="card flex justify-center">
         <Form class="flex flex-col gap-4 w-full" @submit="handle_submit">
            <h2>{{ t("iam.user.create.title") }}</h2>
            <div class="flex flex-col gap-1">
               <FloatLabel variant="in">
                  <InputText id="email" v-model="email" name="email" type="text" fluid />
                  <label for="email">{{ $t("iam.user.create.email") }}</label>
               </FloatLabel>
               <Message v-if="errors.email">{{ errors.email }}</Message>
            </div>
            <div class="flex flex-col gap-1">
               <FloatLabel variant="in">
                  <InputText
                     id="first_name"
                     v-model="first_name"
                     name="first_name"
                     type="text"
                     fluid
                  />
                  <label for="first_name">{{ $t("iam.user.create.first_name") }}</label>
               </FloatLabel>
               <Message v-if="errors.first_name">{{ errors.first_name }}</Message>
            </div>
            <div class="flex flex-col gap-1">
               <FloatLabel variant="in">
                  <InputText
                     id="last_name"
                     v-model="last_name"
                     name="last_name"
                     type="text"
                     fluid
                  />
                  <label for="last_name">{{ $t("iam.user.create.last_name") }}</label>
               </FloatLabel>
               <Message v-if="errors.first_name">{{ errors.first_name }}</Message>
            </div>
            <div class="flex flex-col gap-1">
               <FloatLabel variant="in">
                  <InputMask
                     id="date_of_birth"
                     v-model="date_of_birth"
                     mask="99/99/9999"
                     slot-char="dd/mm/yyyy"
                     @value-change="check_date_of_birth"
                  />
                  <label for="date_of_birth">{{ $t("iam.user.create.date_of_birth") }}</label>
               </FloatLabel>
               <Message
                  v-if="errors.date_of_birth"
                  severity="error"
                  size="small"
                  variant="simple"
                  >{{ errors.date_of_birth }}</Message
               >
            </div>
            <div class="flex flex-col gap-1">
               <FloatLabel variant="in">
                  <UserStatusPicker id="user_status" v-model="user_status"></UserStatusPicker>
                  <label for="user_status">{{ $t("iam.user_status.picker.placeholder") }}</label>
               </FloatLabel>
               <Message v-if="errors.user_status" severity="error" size="small" variant="simple">{{
                  errors.user_status
               }}</Message>
            </div>
            <Button type="submit" severity="secondary" :label="$t('iam.user.create.submit')" />
         </Form>
      </div>
   </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import * as v from "valibot"
import { useI18n } from "vue-i18n"
import { api } from "@/api"
import UserStatusPicker from "../user_status/picker.vue"
import { error, success } from "@/composables/toast"

const t = useI18n().t

const email = ref<string>("")
const first_name = ref<string>("")
const last_name = ref<string>("")
const date_of_birth = ref<string>("")
const user_status = ref<string>("")

const errors = ref<{
   email: string
   first_name: string
   last_name: string
   date_of_birth: string
   user_status: ""
}>({
   email: "",
   first_name: "",
   last_name: "",
   date_of_birth: "",
   user_status: "",
})

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
   date_of_birth: date_of_birth_check,
})

const check_date_of_birth = (date: string) => {
   try {
      v.parse(date_of_birth_check, date)
      errors.value.date_of_birth = ""
   } catch (e: any) {
      errors.value.date_of_birth = e.message
   }
}

const handle_submit = async () => {
   const re = v.safeParse(user_create, {
      email: email.value,
      first_name: first_name.value,
      last_name: last_name.value,
      date_of_birth: date_of_birth.value,
   })

   if (re.success) {
      ;(
         await api.iam.user.create({
            emailaddress: re.output["email"],
            first_name: re.output["first_name"],
            last_name: re.output["last_name"],
            date_of_birth: new Date(re.output["date_of_birth"] as string).toISOString(),
         })
      ).fold(
         (l) => {
            error(l)
         },
         (user_id: string) => {
            success("User: " + user_id)
         },
      )
   }
}
</script>
