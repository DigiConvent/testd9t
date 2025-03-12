<template>
   <div>
      <i class="pi pi-"></i>
   </div>
</template>

<script lang="ts" setup>
import JwtAuthenticator from "@/auth/jwt"

function logout() {
   JwtAuthenticator.get_instance().logout()
}
logout()
</script>
