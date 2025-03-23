<template>
   <div>
      <PermissionGroupTreeView @click="handle_click($event)"></PermissionGroupTreeView>

      <!-- <PermissionPicker :multiple="true" @picked="handle_picked"></PermissionPicker> -->
      <div class="grid grid-cols-2 gap-4">
         <UpdatePermissionGroup v-if="pg != null" v-model="pg"></UpdatePermissionGroup>
         <CreatePermissionGroup @created="handle_created"></CreatePermissionGroup>
      </div>
   </div>
</template>

<script lang="ts" setup>
import PermissionGroupTreeView from "@/components/iam/permission_group/list.vue"
import UpdatePermissionGroup from "@/components/iam/permission_group/update.vue"
import CreatePermissionGroup from "@/components/iam/permission_group/create.vue"
import { ref } from "vue"

const pg = ref<string>()

const handle_click = (node: any) => {
   pg.value = node
}

const handle_created = () => {}
</script>
