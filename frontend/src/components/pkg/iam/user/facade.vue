<template>
   <div>
      <template v-if="users">
         <template v-for="_user in users" :key="_user.id">
            <router-link
               :to="{ name: 'admin.iam.user.profile', params: { id: _user.id } }"
               class="block"
               ><Fa :icon="_user.role_id != '' ? 'user-shield' : 'user'" /> {{ _user.first_name }}
               {{ _user.last_name }}
            </router-link>
         </template>
      </template>
      <template v-else>
         {{ user }}
      </template>
   </div>
</template>

<script lang="ts" setup>
import type { UserFacade } from "@/api/iam/user/types"

type UserFacadeProps =
   | { user?: undefined; users: UserFacade[] }
   | { user: UserFacade; users?: undefined }

defineProps<UserFacadeProps>()
</script>
