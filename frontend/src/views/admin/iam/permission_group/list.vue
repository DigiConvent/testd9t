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
         <UpdatePermissionGroup v-model="edit_pg" @updated="handle_updated"></UpdatePermissionGroup>
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
            @created="handle_created"
         ></CreatePermissionGroup>
      </Drawer>
   </div>
</template>

<script lang="ts" setup>
import PermissionGroupTreeView from "@/components/iam/permission_group/list.vue"
import UpdatePermissionGroup from "@/components/iam/permission_group/update.vue"
import CreatePermissionGroup from "@/components/iam/permission_group/create.vue"
import { ref } from "vue"
import JwtAuthenticator from "@/auth/jwt"
import { useI18n } from "vue-i18n"
import router from "@/router"

const pg = ref<string>()

const show_menu = ref()
const handle_click = (arg: { event: any; id: string }) => {
   pg.value = arg.id
   generate_menu_items()
   show_menu.value.toggle(arg.event)
}

const handle_created = () => {
   add_pg_to_pg.value = null
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

const generate_menu_items = () => {
   menu_items.value = []
   if (auth.has_permission("iam.permission_group.read"))
      menu_items.value.push({
         label: t("iam.pg.read.title"),
         icon: "eye",
         command: () => {
            router.push({ name: "iam.pg.profile", params: { id: pg.value } })
         },
      })
   if (auth.has_permission("iam.permission_group.create"))
      menu_items.value.push({
         label: t("iam.pg.create.title"),
         icon: "plus",
         command: () => {
            add_pg_to_pg.value = pg.value
         },
      })
   if (auth.has_permission("iam.permission_group.update"))
      menu_items.value.push({
         label: t("iam.pg.update.title"),
         icon: "pencil",
         command: () => {
            edit_pg.value = pg.value
         },
      })
   // if (auth.has_permission("iam.permission_group.delete"))
   //    menu_items.value.push({
   //       label: t("iam.pg.delete.title"),
   //       icon: "trash",
   //       command: () => {},
   //    })
}

const refresh = ref(0)
</script>
