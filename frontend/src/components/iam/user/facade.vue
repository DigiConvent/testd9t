<template>
   membas
   <template v-if="users">
      <template v-for="_user in users" :key="_user.id">
         {{ _user.name }}
      </template>
   </template>
   <template v-else-if="user">
      {{ user.name }}
   </template>
   {{ users }}
</template>

<script lang="ts" setup>
import type { UserFacade } from "@/api/iam/user/types"

defineProps<{ user?: UserFacade; users?: UserFacade[] }>()
</script>
