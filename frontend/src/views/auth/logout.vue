<template>
   <div>
      {{ $t("iam.auth.logged_out") }}
   </div>
</template>

<script lang="ts" setup>
import JwtAuthenticator from "@/auth/jwt"
import router from "@/router"

function logout() {
   JwtAuthenticator.get_instance().logout()
   router.replace("/")
}
logout()
</script>
