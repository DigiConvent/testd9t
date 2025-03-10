<template>
   <div>
      <Toast />
      <div v-if="!logged_in" class="card flex justify-center">
         <Form class="flex flex-col gap-4 w-full sm:w-56" @submit="handle_submit">
            <FormTextInput
               v-model="email"
               label="iam.auth.login_form"
               :error="errors.email"
               name="email"
            />
            <FormPasswordInput
               v-model="password"
               label="iam.auth.login_form"
               :error="errors.password"
               name="password"
            />
            <Button type="submit" severity="secondary" :label="$t('iam.auth.login_form.submit')" />
         </Form>
      </div>
      <div v-else class="card flex justify-center">
         <Button @click="JwtAuthenticator.get_instance().logout()">{{
            $t("iam.auth.login_form.logout")
         }}</Button>
      </div>
   </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { useToast } from "primevue/usetoast"
import FormTextInput from "@/components/form/text_input.vue"
import FormPasswordInput from "@/components/form/password_input.vue"
import Button from "primevue/button"
import Toast from "primevue/toast"
import JwtAuthenticator from "@/auth/jwt"
import * as v from "valibot"
import { useI18n } from "vue-i18n"
import get_web_app from "@/auth/telegram"
import router from "@/router"
const t = useI18n().t

const email = ref<string>("")
const password = ref<string>("")
const errors = ref<{ email: string; password: string }>({
   email: "",
   password: "",
})

const toast = useToast()

const email_validation = v.pipe(v.string(), v.nonEmpty(t("iam.auth.login_form.emailRequired")))
const password_validation = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.auth.login_form.passwordRequired")),
)

const login_form = v.object({
   email: email_validation,
   password: password_validation,
})

const handle_submit = async () => {
   const re = v.safeParse(login_form, { email: email.value, password: password.value })

   if (re.success) {
      const success = await JwtAuthenticator.get_instance().login_using_credentials(
         re.output["email"],
         re.output["password"],
      )

      if (success) {
         // we used credentials so now we can check if a telegram user is logged in
         const data = get_web_app().initData
         if (data != "") {
            const telegram_id = get_web_app().initDataUnsafe.user.id
            const saved_telegram_id = JwtAuthenticator.get_instance().get_token()?.tgid

            if (telegram_id != saved_telegram_id) {
               router.push({ name: "connect-telegram-user" })
            }
         }
         toast.add({
            severity: "success",
            summary: t("iam.auth.login_form.login_successful"),
            life: 3000,
         })
      }
   } else {
      for (const issue of re.issues) {
         toast.add({
            severity: "error",
            summary: issue.message,
            life: 3000,
         })
      }
   }
}

const logged_in = JwtAuthenticator.get_instance().is_authenticated
</script>
