<template>
   <template v-if="users">
      <template v-for="_user in users" :key="_user.id">
         <router-link
            :to="{ name: 'admin.iam.user.profile', params: { id: _user.id } }"
            class="block"
            >{{ _user.status_name }} {{ _user.name }}
         </router-link>
      </template>
   </template>
   <template v-else-if="user">
      {{ user.name }}
   </template>
</template>

<script lang="ts" setup>
import type { UserFacade } from "@/api/iam/user/types"

defineProps<{ user?: UserFacade; users?: UserFacade[] }>()
</script>
