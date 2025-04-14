<template>
   <PermissionGroupTreeView
      :refresh="refresh"
      @click="handle_click($event)"
   ></PermissionGroupTreeView>
   <Menu ref="show_menu" :model="menu_items" :popup="true">
      <template #item="{ item }">
         <router-link :to="item.route" class="inline-block w-full p-2">
            <Fa :icon="item.icon" class="fa-fw mr" /> {{ item.label }}
         </router-link>
      </template>
   </Menu>
</template>

<script lang="ts" setup>
import PermissionGroupTreeView from "@/components/iam/permission_group/list.vue"
import { ref } from "vue"
import JwtAuthenticator from "@/auth/jwt"
import { useI18n } from "vue-i18n"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"

const pg = ref<string>()

const show_menu = ref()
const handle_click = (arg: { event: any; pg: PermissionGroupFacade }) => {
   pg.value = arg.pg.id
   generate_menu_items(arg.pg)
   show_menu.value.toggle(arg.event)
}

const t = useI18n().t

const auth = JwtAuthenticator.get_instance()
const menu_items = ref<({ label: string; icon: string; route?: any } | { separator: boolean })[]>(
   [],
)

const generate_menu_items = (pg: PermissionGroupFacade) => {
   menu_items.value = []
   if (pg.meta == "role") {
      generate_role_menu(pg)
   } else if (pg.meta == "status") {
      generate_status_menu(pg)
   } else {
      generate_group_menu(pg)
   }
}

const refresh = ref(0)

function generate_group_menu(pg: PermissionGroupFacade) {
   menu_items.value.push({
      label: t("actions.view", { entity: t("iam.pg.pg") }),
      icon: "eye",
      route: { name: "admin.iam.permission_group.profile", params: { id: pg.id } },
   })
   if (auth.has_permission("iam.permission_group.write")) {
      menu_items.value.push({
         label: t("actions.edit", { entity: t("iam.pg.pg") }),
         icon: "pencil",
         route: { name: "admin.iam.permission_group.update", params: { id: pg.id } },
      })
      menu_items.value.push({
         label: t("actions.add", { entity: t("iam.pg.pg") }),
         icon: "folders",
         route: { name: "admin.iam.permission_group.create", params: { parent: pg.id } },
      })
      menu_items.value.push({
         label: t("actions.delete", { entity: t("iam.pg.pg") }),
         icon: "trash",
         route: { name: "admin.iam.permission_group.delete", params: { id: pg.id } },
      })
   }
   if (auth.has_permission("iam.user_role.write")) {
      menu_items.value.push({
         separator: true,
      })
      menu_items.value.push({
         label: t("actions.add", { entity: t("iam.ur.ur") }),
         icon: "user-shield",
         route: { name: "admin.iam.user_role.create", params: { parent: pg.id } },
      })
   }
   if (auth.has_permission("iam.user_status.write")) {
      menu_items.value.push({
         separator: true,
      })
      menu_items.value.push({
         label: t("actions.add", { entity: t("iam.us.us") }),
         icon: "user-tag",
         route: { name: "admin.iam.user_status.create", params: { parent: pg.id } },
      })
   }
}

function generate_role_menu(pg: PermissionGroupFacade) {
   menu_items.value.push({
      label: t("actions.view", { entity: t("iam.ur.ur") }),
      icon: "eye",
      route: { name: "admin.iam.user_role.profile", params: { id: pg.id } },
   })
   menu_items.value.push({
      label: t("actions.edit", { entity: t("iam.ur.ur") }),
      icon: "pencil",
      route: { name: "admin.iam.user_role.update", params: { id: pg.id } },
   })
   menu_items.value.push({
      label: t("actions.delete", { entity: t("iam.ur.ur") }),
      icon: "trash",
      route: { name: "admin.iam.permission_group.delete", params: { id: pg.id } },
   })
}

function generate_status_menu(pg: PermissionGroupFacade) {
   menu_items.value.push({
      label: t("actions.view", { entity: t("iam.us.us") }),
      icon: "eye",
      route: { name: "admin.iam.user_status.profile", params: { id: pg.id } },
   })
   menu_items.value.push({
      label: t("actions.edit", { entity: t("iam.us.us") }),
      icon: "pencil",
      route: { name: "admin.iam.user_status.update", params: { id: pg.id } },
   })
   menu_items.value.push({
      label: t("actions.delete", { entity: t("iam.us.us") }),
      icon: "trash",
      route: { name: "admin.iam.permission_group.delete", params: { id: pg.id } },
   })
}
</script>
