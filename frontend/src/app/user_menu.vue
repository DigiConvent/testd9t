<template>
   {{
      t("iam.auth.user_menu.logged_in_as", {
         user: JwtAuthenticator.get_instance().get_token()!.user.last_name,
      })
   }}
   <Button
      type="button"
      icon="pi pi-user"
      aria-controls="overlay-menu"
      aria-haspopup="true"
      severity="secondary"
      @click="show_user_menu.show($event)"
   ></Button>
   <Menu ref="show_user_menu" :model="items" :popup="true" class="justify-center">
      <template #item="{ item, props }">
         <router-link v-if="item.to" v-slot="{ href, navigate }" :to="{ name: item.to }" custom>
            <a v-ripple :href="href" v-bind="props.action" @click="navigate">
               <span :class="item.icon" />
               <span class="ml-2">{{ item.label }}</span>
            </a>
         </router-link>
      </template>
   </Menu>
</template>

<script lang="ts" setup>
import JwtAuthenticator from "@/auth/jwt"
import { ref } from "vue"
import { useI18n } from "vue-i18n"

const t = useI18n().t

const show_user_menu = ref()

const items = ref([
   {
      label: "Settings",
      icon: "pi pi-cog",
      to: "user.settings",
   },
   {
      label: t("iam.auth.user_menu.my_profile"),
      icon: "pi pi-id-card",
      to: "user.profile",
   },
   {
      label: "Logout",
      icon: "pi pi-sign-out",
      to: "logout",
   },
])
</script>
