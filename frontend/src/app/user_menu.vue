<template>
   <div>
      <Button
         type="button"
         aria-controls="overlay-menu"
         aria-haspopup="true"
         severity="secondary"
         @click="show_user_menu.show($event)"
         >{{
            t("iam.auth.user_menu.logged_in_as", {
               user: last_name,
            })
         }}
         <Fa icon="user" />
      </Button>
      <Menu ref="show_user_menu" :model="items" :popup="true" class="justify-center">
         <template #item="{ item, props }">
            <router-link v-if="item.to" v-slot="{ href, navigate }" :to="{ name: item.to }" custom>
               <a v-ripple :href="href" v-bind="props.action" @click="navigate">
                  <Fa :icon="item.icon" class="fa-fw" />
                  <span class="ml-2">{{ item.label }}</span>
               </a>
            </router-link>
         </template>
      </Menu>
   </div>
</template>

<script lang="ts" setup>
import JwtAuthenticator from "@/auth/jwt"
import { ref } from "vue"
import { useI18n } from "vue-i18n"

const t = useI18n().t

const show_user_menu = ref()

const last_name = ref("")
const user = JwtAuthenticator.get_instance().get_token()
if (user != null) {
   last_name.value = user.user.last_name
}

const items = ref([
   {
      label: t("iam.auth.user_menu.my_profile"),
      icon: "id-card",
      to: "user.profile",
   },
   {
      label: t("iam.auth.user_menu.logout"),
      icon: "arrow-right-from-bracket",
      to: "logout",
   },
])
</script>
