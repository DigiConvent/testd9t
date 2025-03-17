<template>
   <div class="p-4">
      <h2 class="text-2xl">{{ $t("iam.title") }}</h2>
      <router-view v-slot="{ Component, route }">
         <component :is="Component" :key="route.path" />
      </router-view>
      <PermissionGroupList></PermissionGroupList>
   </div>
</template>

<script lang="ts" setup>
import PermissionGroupList from "@/views/admin/iam/permission_group/list.vue"
</script>
