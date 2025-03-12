<template>
   <div v-if="conflict">
      <h2>{{ $t("iam.auth.connect_telegram_user.title") }}</h2>
      <h2>{{ $t("iam.auth.connect_telegram_user.description") }}</h2>
      <Button @click="connect_telegram_user">{{
         $t("iam.auth.connect_telegram_user.button")
      }}</Button>
      <Dialog v-model:visible="show_message" modal header="Restart">
         <p>{{ $t("iam.auth.connect_telegram_user.restart_required") }}</p>
         <ProgressBar :value="countup" :show-value="false"></ProgressBar>
         <Button @click="restart">{{ $t("iam.auth.connect_telegram_user.restart") }}</Button>
      </Dialog>
   </div>
</template>

<script lang="ts" setup>
import JwtAuthenticator from "@/auth/jwt"
import get_web_app from "@/auth/telegram"
import { ref } from "vue"

const conflict = ref(false)

function get_status() {
   const current_telegram_id = get_web_app().initDataUnsafe.user.id
   const saved_telegram_id = JwtAuthenticator.get_instance().get_token()?.tgid

   if (current_telegram_id == saved_telegram_id) {
      conflict.value = false
   } else {
      conflict.value = true
   }
}

get_status()

const countup = ref(0)
const show_message = ref(false)
async function connect_telegram_user() {
   const success = await JwtAuthenticator.get_instance().connect_telegram_user()
   if (success) {
      show_message.value = true
      countup.value = 0
      setInterval(() => {
         countup.value += 20
         if (countup.value >= 100) {
            restart()
         }
      }, 1000)
   }
}

function restart() {
   get_web_app().close()
}
</script>
