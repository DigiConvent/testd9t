<template>
   <div class="card flex justify-center">
      <Menu :model="items" class="w-full md:w-60">
         <template #start></template>
         <template #item="{ item }">
            <router-link
               v-ripple
               class="flex items-center"
               :to="item.to ? { name: item.to, params: {} } : {}"
            >
               <span :class="item.icon" />
               <span>{{ item.label }}</span>
               <Badge v-if="item.badge" class="ml-auto" :value="item.badge" />
               <span
                  v-if="item.shortcut"
                  class="ml-auto border border-surface rounded bg-emphasis text-muted-color text-xs p-1"
                  >{{ item.shortcut }}</span
               >
            </router-link>
         </template>
         <template #end>
            <button
               v-ripple
               class="relative overflow-hidden w-full border-0 bg-transparent flex items-start p-2 pl-4 hover:bg-surface-100 dark:hover:bg-surface-800 rounded-none cursor-pointer transition-colors duration-200"
            >
               <Avatar
                  image="https://primefaces.org/cdn/primevue/images/avatar/amyelsner.png"
                  class="mr-2"
                  shape="circle"
               />
               <span class="inline-flex flex-col items-start">
                  <span class="font-bold">Amy Elsner</span>
                  <span class="text-sm">Admin</span>
               </span>
            </button>
         </template>
      </Menu>
   </div>
</template>

<script lang="ts" setup>
import JwtAuthenticator from "@/auth/jwt"
import { ref } from "vue"
import { useI18n } from "vue-i18n"

const t = useI18n().t

const items = ref([
   {
      label: t("iam.auth.user_menu.logged_in_as", {
         user: JwtAuthenticator.get_instance().get_token()!.user.last_name,
      }),
      items: [
         {
            label: "Settings",
            icon: "pi pi-cog",
         },
         {
            label: "Messages",
            icon: "pi pi-inbox",
            badge: 2,
         },
         {
            label: "Logout",
            icon: "pi pi-sign-out",
            to: "logout",
         },
      ],
   },
   {
      separator: true,
   },
])
</script>
