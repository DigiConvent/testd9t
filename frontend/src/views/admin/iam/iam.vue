<template>
   <div class="p-4">
      iam
      <router-view v-slot="{ Component, route }">
         <PermissionGroupList v-if="route.name == 'iam'"></PermissionGroupList>
         <component :is="Component" :key="route.path" />
      </router-view>
   </div>
</template>

<script lang="ts" setup>
import PermissionGroupList from "@/views/admin/iam/permission_group/list.vue"
</script>
