<template>
   <div>
      <Toast></Toast>
      <Menubar v-if="logged_in" :model="items">
         <template #start>
            <router-link :to="{ name: 'home' }">
               <img
                  class="h-16 w-16 rounded-full"
                  :src="`/assets/logo/small.jpg?v=${new Date().getTime()}`"
               />
            </router-link>
         </template>
         <template #item="{ item, props, hasSubmenu, root }">
            <a v-if="item.hasSubmenu" v-bind="props.action">{{ item.label }}</a>
            <router-link
               v-else
               :to="{ name: item.route, params: {} }"
               v-bind="props.action"
               class="p-0"
            >
               <span>{{ item.label }}</span>
               <span
                  v-if="item.shortcut"
                  class="ml-auto border border-surface rounded bg-emphasis text-muted-color text-xs p-1"
                  >{{ item.shortcut }}</span
               >
               <i
                  v-if="hasSubmenu"
                  :class="[
                     'pi pi-angle-down ml-auto',
                     { 'pi-angle-down': root, 'pi-angle-right': !root },
                  ]"
               ></i>
            </router-link>
         </template>
         <template #end>
            <UserMenu></UserMenu>
         </template>
      </Menubar>
      <div v-else class="absolute">
         <div class="right-0">
            <Button @click="show_login_form = true"><i class="pi pi-user"></i></Button>
            <LoginForm v-if="show_login_form"></LoginForm>
         </div>
      </div>
      <header>
         <router-view v-slot="{ Component, route }">
            <component :is="Component" :key="route.path" />
         </router-view>
      </header>
      <footer class="fixed bottom-0">
         <a href="https://github.com/DigiConvent/testd9t" target="_blank"
            ><i class="pi pi-github"></i
         ></a>
      </footer>
   </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from "vue"
import JwtAuthenticator from "../auth/jwt"
import LoginForm from "@/components/iam/auth/login/credentials.vue"
import type { MenuItem } from "primevue/menuitem"
import { useI18n } from "vue-i18n"
import UserMenu from "./user_menu.vue"

const auth = JwtAuthenticator.get_instance()
const logged_in = auth.is_authenticated

const t = useI18n().t

const show_login_form = ref(false)

const items = ref<MenuItem[]>([])

watch(logged_in, () => {
   generate_menu_items()
})

const admin_items = ref<MenuItem[]>([])
function generate_menu_items() {
   items.value = []

   admin_items.value = []
   if (auth.has_permission("iam")) {
      admin_items.value.push({
         label: t("iam.title"),
         route: "iam",
      })
   }
   if (auth.has_permission("sys")) {
      admin_items.value.push({
         label: t("sys.title"),
         route: "sys",
      })
   }

   if (auth.has_permission("iam.user.list")) {
      items.value.push({
         label: t("iam.user.list.title"),
         route: "iam.user.list",
         badge: 1,
      })
   }

   if (admin_items.value.length > 0) {
      items.value.push({
         label: t("admin.title"),
         route: "",
         items: admin_items.value,
      })
   }
}

onMounted(() => {
   generate_menu_items()
})
</script>
