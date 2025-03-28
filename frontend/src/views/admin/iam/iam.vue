<template>
   <div class="p-4">
      <h2 class="text-2xl">{{ $t("iam.title") }}</h2>
      <router-view v-slot="{ Component, route }">
         <PermissionGroupList v-if="route.name == 'iam'"></PermissionGroupList>
         <component :is="Component" :key="route.path" />
      </router-view>
      <!-- <Card>
         <template #title>
            {{ $t("iam.pg.list.title") }}
         </template>
         <template #content>
         </template>
      </Card> -->
   </div>
</template>

<script lang="ts" setup>
import PermissionGroupList from "@/views/admin/iam/permission_group/list.vue"
</script>
