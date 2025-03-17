<template>
   <div>
      <PermissionGroupList @click="console.log($event)"></PermissionGroupList>

      <!-- <PermissionPicker :multiple="true" @picked="handle_picked"></PermissionPicker> -->
   </div>
</template>

<script lang="ts" setup>
import PermissionGroupList from "@/components/iam/permission_group/list.vue"
</script>
