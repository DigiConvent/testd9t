<template>
   <div>
      <PermissionGroupTreeView
         :refresh="refresh"
         @click="handle_click($event)"
      ></PermissionGroupTreeView>
      <Menu ref="show_menu" :model="menu_items" :popup="true">
         <template #item="{ item }">
            <div class="hover:bg-gray-100 p-2 cursor-pointer">
               <Fa :icon="item.icon" class="fa-fw mr" /> {{ item.label }}
            </div>
         </template>
      </Menu>
      <Drawer
         :visible="edit_pg != null"
         modal
         style="width: 100%; height: 100%"
         position="top"
         @update:visible="edit_pg = null"
      >
         <UpdatePermissionGroup :data="edit_pg" @updated="handle_updated"></UpdatePermissionGroup>
      </Drawer>
      <Drawer
         :visible="add_pg_to_pg"
         position="bottom"
         modal
         style="width: 100%; height: 100%"
         @update:visible="add_pg_to_pg = null"
      >
         <CreatePermissionGroup
            :parent="add_pg_to_pg"
            @created="handle_created_pg"
         ></CreatePermissionGroup>
      </Drawer>
      <Drawer
         :visible="add_us_to_pg"
         position="right"
         modal
         style="width: 100%; height: 100%"
         @update:visible="add_us_to_pg = null"
      >
         <CreateUserStatus :parent="add_us_to_pg" @created="handle_created_us"></CreateUserStatus>
      </Drawer>
   </div>
</template>

<script lang="ts" setup>
import PermissionGroupTreeView from "@/components/iam/permission_group/list.vue"
import UpdatePermissionGroup from "@/components/iam/permission_group/update.vue"
import CreatePermissionGroup from "@/components/iam/permission_group/create.vue"
import CreateUserStatus from "@/components/iam/user_status/create.vue"
import { ref } from "vue"
import JwtAuthenticator from "@/auth/jwt"
import { useI18n } from "vue-i18n"
import router from "@/router"
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"

const pg = ref<string>()

const show_menu = ref()
const handle_click = (arg: { event: any; pg: PermissionGroupFacade }) => {
   pg.value = arg.pg.id
   generate_menu_items(arg.pg)
   show_menu.value.toggle(arg.event)
}

const handle_created_pg = () => {
   add_pg_to_pg.value = null
   refresh.value += 1
}

const add_us_to_pg = ref()
const handle_created_us = () => {
   add_us_to_pg.value = null
   refresh.value += 1
}
const handle_updated = () => {
   edit_pg.value = null
   refresh.value += 1
}

const t = useI18n().t

const auth = JwtAuthenticator.get_instance()
const menu_items = ref()
const add_pg_to_pg = ref()
const edit_pg = ref()

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
   menu_items.value = []
   menu_items.value.push({
      label: t("iam.pg.read.title"),
      icon: "eye",
      command: () => {
         router.push({ name: "admin.iam.permission_group.profile", params: { id: pg.id } })
      },
   })
   if (auth.has_permission("iam.permission_group.update"))
      menu_items.value.push({
         label: t("iam.pg.update.title"),
         icon: "pencil",
         command: () => {
            edit_pg.value = pg.id
         },
      })
   menu_items.value.push({
      label: t("iam.pg.create.title"),
      icon: "plus",
      command: () => {
         add_pg_to_pg.value = pg.id
      },
   })
   menu_items.value.push({
      label: t("iam.user_status.create.title"),
      icon: "plus",
      command: () => {
         add_us_to_pg.value = pg.id
      },
   })
}

function generate_role_menu(pg: PermissionGroupFacade) {
   menu_items.value = []
   menu_items.value.push({
      label: t("iam.pg.role.title"),
      icon: "magnifying-glass",
      command: () => {
         router.push({ name: "admin.iam.user_role.profile", params: { id: pg.id } })
      },
   })
}

function generate_status_menu(pg: PermissionGroupFacade) {
   menu_items.value = []
   menu_items.value.push({
      label: t("iam.pg.status.title"),
      icon: "magnifying-glass",
      command: () => {
         router.push({ name: "admin.iam.user_status.profile", params: { id: pg.id } })
      },
   })
}
</script>
