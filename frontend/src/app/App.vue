<template>
  <div>
    <Menubar :model="items">
      <template #start>
        <router-link to="/">Logo</router-link>
      </template>
      <template #item="{ item, props, hasSubmenu, root }">
        <a v-if="item.hasSubmenu" v-bind="props.action">{{ item.label }}</a>
        <router-link v-else :to="item.route" v-bind="props.action" class="p-0">
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
              { 'pi-angle-down': root, 'pi-angle-right': !root }
            ]"
          ></i>
        </router-link>
      </template>
      <template #end>
        <span v-if="logged_in">
          <Button
            :label="$t('iam.auth.login_form.logout')"
            severity="secondary"
            @click="auth.logout()"
          ></Button>
        </span>
        <div v-else>
          <Button
            :label="$t('iam.auth.login_form.title')"
            severity="secondary"
            @click="op.show($event)"
          ></Button>
          <Popover ref="op">
            <LoginForm></LoginForm>
          </Popover>
        </div>
      </template>
    </Menubar>
    <header>
      <router-view v-slot="{ Component, route }">
        <component :is="Component" :key="route.path" />
      </router-view>
    </header>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import JwtAuthenticator from "../auth/jwt"
import LoginForm from "@/components/iam/auth/login/credentials.vue"
import type { MenuItem } from "primevue/menuitem"
import { useI18n } from "vue-i18n"
const op = ref()
const auth = JwtAuthenticator.get_instance()
const logged_in = auth.is_authenticated

const items = ref<MenuItem[]>([])
const admin_items = ref<MenuItem[]>([])

const t = useI18n().t

if (auth.has_permission("iam")) {
  admin_items.value.push({
    label: t("iam.title"),
    route: "/admin/iam"
  })
}
if (auth.has_permission("sys")) {
  admin_items.value.push({
    label: t("sys.title"),
    route: "/admin/sys"
  })
}

if (auth.has_permission("iam.user.list")) {
  items.value.push({
    label: t("iam.user.list.title"),
    route: "/admin/iam/user",
    badge: 1
  })
}

if (admin_items.value.length > 0) {
  items.value.push({
    label: t("admin.title"),
    route: "",
    items: admin_items.value
  })
}
</script>
