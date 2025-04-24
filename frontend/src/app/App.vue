<template>
   <SwipeWrapper>
      <Menubar v-if="logged_in" :model="items">
         <template #start>
            <router-link :to="{ name: 'home' }">
               <img
                  class="h-16 w-16 rounded-full"
                  :src="`/assets/small.jpg?v=${new Date().getTime()}`"
               />
            </router-link>
         </template>
         <template #item="{ item, props, hasSubmenu, root }">
            <a v-if="item.hasSubmenu" v-bind="props.action">
               <Fa :icon="item.icon" class="fa-fw" />
               {{ item.label }}
            </a>
            <router-link
               v-else
               :to="{ name: item.route, params: {} }"
               v-bind="props.action"
               class="p-0"
            >
               <span>
                  <Fa :icon="item.icon" class="fa-fw" />
                  {{ item.label }}
               </span>
               <Fa v-if="hasSubmenu" :icon="`${root ? 'caret-down' : 'caret-up'}`" />
            </router-link>
         </template>
         <template #end>
            <UserMenu></UserMenu>
         </template>
      </Menubar>
      <div v-else class="absolute">
         <div class="right-0">
            <Button @click="show_login_form = true"><Fa icon="user" /></Button>
            <Dialog v-model:visible="show_login_form" modal>
               <LoginForm></LoginForm>
            </Dialog>
         </div>
      </div>
      <header>
         <router-view v-slot="{ Component, route }">
            <component :is="Component" :key="route.path" />
         </router-view>
      </header>
      <footer class="mt-5">
         <a href="https://github.com/DigiConvent/testd9t" target="_blank"
            ><Fa icon="fab fa-github"></Fa
         ></a>
      </footer>
      <div v-permission="'admin'">
         <router-link :to="{ name: 'app.meta' }">Dev</router-link>
      </div>
      <Toast></Toast>
   </SwipeWrapper>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from "vue"
import JwtAuthenticator from "../auth/jwt"
import LoginForm from "@/components/pkg/iam/auth/login/credentials.vue"
import type { MenuItem } from "primevue/menuitem"
import { useI18n } from "vue-i18n"
import UserMenu from "./user_menu.vue"
import SwipeWrapper from "./swipe_wrapper.vue"

const t = useI18n().t

const auth = JwtAuthenticator.get_instance()
const logged_in = auth.is_authenticated

const show_login_form = ref(false)

const items = ref<MenuItem[]>([])

watch(auth.is_authenticated, () => {
   generate_menu_items()
})

const admin_items = ref<MenuItem[]>([])
function generate_menu_items() {
   items.value = []

   admin_items.value = []
   if (auth.has_permission("iam")) {
      admin_items.value.push({
         label: t("iam.title"),
         route: "admin.iam.permission_group.list",
         icon: "shield",
      })

      admin_items.value.push({
         label: t("iam.user.list.title"),
         route: "admin.iam.user.list",
         icon: "users",
      })
   }
   if (auth.has_permission("sys")) {
      admin_items.value.push({
         label: t("sys.title"),
         route: "admin.sys.overview",
         icon: "hard-drive",
      })
   }

   if (admin_items.value.length > 0) {
      items.value.push({
         label: t("admin.title"),
         route: "",
         items: admin_items.value,
         icon: "cogs",
      })
   }
}

onMounted(() => {
   generate_menu_items()
})
</script>
