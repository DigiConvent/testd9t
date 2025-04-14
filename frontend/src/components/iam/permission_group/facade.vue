<template>
   <template v-if="permission_groups">
      <template v-for="pg in permission_groups" :key="pg.id">
         <a
            v-if="has_on_click"
            class="cursor-pointer select-none"
            @click="handle_click($event, pg)"
         >
            <Fa :icon="get_icon(pg.meta)" class="fa-fw" />
            {{ pg.name }} {{ pg.abbr != "" ? `(${pg.abbr})` : "" }}
         </a>
         <router-link
            v-else
            :to="{ name: 'admin.iam.permission_group.profile', params: { id: pg.id } }"
            class="block"
         >
            <Fa :icon="get_icon(pg.meta)" class="fa-fw" />
            {{ pg.name }} {{ pg.abbr != "" ? `(${pg.abbr})` : "" }}
         </router-link>
      </template>
   </template>
   <template v-else>
      {{ permission_group }}
   </template>
</template>

<script lang="ts" setup>
import type { PermissionGroupFacade } from "@/api/iam/permission_group/types"
import { getCurrentInstance } from "vue"

type PermissionGroupFacadeProps =
   | { permission_group?: undefined; permission_groups: PermissionGroupFacade[] }
   | { permission_group: PermissionGroupFacade; permission_groups?: undefined }

defineProps<PermissionGroupFacadeProps>()
function get_icon(type: string | null): string {
   if (type == null || type == "") return "folders"
   if (type == "role") return "user-shield"
   if (type == "status") return "user-tag"
   return "folders"
}

const emits = defineEmits<{
   (e: "pick", payload: { event: Event; pg: PermissionGroupFacade }): void
}>()

const has_on_click = getCurrentInstance()?.vnode.props!.onPick != null

function handle_click(event: Event, pg: PermissionGroupFacade) {
   if (has_on_click) emits("pick", { event, pg })
}
</script>
