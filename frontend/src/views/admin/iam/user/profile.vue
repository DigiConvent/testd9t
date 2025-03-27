<template>
   <UserProfileComponent :id="id as string" />
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { UserProfile } from "@/api/iam/user/types"
import { error } from "@/composables/toast"
import { ref } from "vue"
import { useRoute } from "vue-router"
import UserProfileComponent from "@/components/iam/user/profile.vue"

const id = useRoute().params.id

const profile = ref<UserProfile | null>(null)

async function load_user_profile() {
   ;(await api.iam.user.profile(id as string)).fold(
      (err: string) => {
         error(err)
      },
      (pg: UserProfile) => {
         profile.value = pg
      },
   )
}

load_user_profile()
</script>
