<template>
   <div>
      <span v-permission="'admin'" class="inline">
         <ToggleSwitch v-model="debug"></ToggleSwitch>
      </span>
      <div v-for="route of sorted_routes" :key="route.path">
         <router-link
            v-if="!route.path.includes(':') && route.components != undefined"
            :to="route.path"
            class="underline text-sky-700"
            >{{ route.path }}</router-link
         >
         <span v-else>{{ route.path }}</span>
      </div>
   </div>
</template>

<script lang="ts" setup>
import router from "@/router"
import { computed, nextTick } from "vue"

const debug = computed({
   get: () => window.debug,
   set: (value) => {
      window.debug = value
      nextTick()
   },
})
const sorted_routes = computed(() => {
   return router.getRoutes().sort((a, b) => {
      return a.path.localeCompare(b.path)
   })
})
console.log(sorted_routes)
</script>
