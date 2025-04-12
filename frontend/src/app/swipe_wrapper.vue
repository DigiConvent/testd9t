<template>
   <div class="app-container" style="height: 100vh; max-width: 100vw">
      <div
         v-if="swipe_direction == 'right' && swipe_position < 0"
         class="fixed top-[50%] z-10 rounded-full"
         :style="{
            opacity: swipe_progress / -0.8,
            transform: `translateX(${-swipe_position}px)`,
         }"
      >
         <div
            class="border p-1 border-gray-200 w-10 h-10 rounded-full absolute"
            :class="swipe_progress == -0.8 ? 'border-sky-800' : ''"
         ></div>
         <div class="bg-sky-500 text-white p-2 rounded-full relative mt-1 ml-1 w-8 h-8">
            <Fa icon="arrow-left" class="fa-fw absolute-center"></Fa>
         </div>
      </div>
      <div
         v-if="swipe_direction == 'left' && swipe_position > 0 && can_go_forward"
         class="fixed top-[50%] right-0 z-10 rounded-full"
         :style="{
            opacity: swipe_progress / 0.8,
            transform: `translateX(${-swipe_position}px)`,
         }"
      >
         <div
            class="border p-1 border-gray-200 w-10 h-10 rounded-full absolute"
            :class="swipe_progress == 0.8 ? 'border-sky-800' : ''"
         ></div>
         <div class="bg-sky-500 text-white p-2 rounded-full relative mt-1 ml-1 w-8 h-8">
            <Fa icon="arrow-right" class="fa-fw absolute-center"></Fa>
         </div>
      </div>

      <div
         ref="el"
         class="page-content"
         :class="{
            'swiping-left': is_swiping && swipe_direction === 'left',
            'swiping-right': is_swiping && swipe_direction === 'right',
         }"
      >
         <slot />
      </div>
   </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue"
import { useSwipe } from "@vueuse/core"
import { useRouter } from "vue-router"

const router = useRouter()
const el = ref(null)
const is_horizontal_swipe = ref(false)
const initial_touch = ref({ x: 0, y: 0 })
const swipe_direction = ref<null | "left" | "right">()
const direction_locked = ref(false)

const can_go_forward = computed(() => {
   return window.history.state.forward !== null
})

const { isSwiping: is_swiping, lengthX: length_x } = useSwipe(el, {
   threshold: 10,
   onSwipeStart(e) {
      initial_touch.value = {
         x: e.touches[0].clientX,
         y: e.touches[0].clientY,
      }

      is_horizontal_swipe.value = false
      direction_locked.value = false
   },
   onSwipe(e) {
      if (direction_locked.value) return
      const current_x = e.touches[0].clientX
      const current_y = e.touches[0].clientY
      const delta_x = current_x - initial_touch.value.x
      const delta_y = current_y - initial_touch.value.y

      if (Math.abs(delta_x) > 5 || Math.abs(delta_y) > 5) {
         if (Math.abs(delta_x) > Math.abs(delta_y) * 1.5) {
            if (
               initial_touch.value.x > window.innerWidth * 0.1 &&
               initial_touch.value.x < window.innerWidth * 0.9
            )
               return
            is_horizontal_swipe.value = true
            direction_locked.value = true
            swipe_direction.value = delta_x > 0 ? "right" : "left"
         } else if (Math.abs(delta_y) > Math.abs(delta_x) * 1.5) {
            is_horizontal_swipe.value = false
            direction_locked.value = true
         }
      }
   },
   onSwipeEnd() {
      if (swipe_direction.value === "right" && length_x.value < -60) {
         if (initial_touch.value.x < window.innerWidth * 0.9) router.go(-1)
      } else if (swipe_direction.value === "left" && length_x.value > 60) {
         if (initial_touch.value.x > window.innerWidth * 0.1 && can_go_forward) router.go(1)
      }
      swipe_direction.value = null
      is_horizontal_swipe.value = false
      direction_locked.value = false
   },
})

const swipe_progress = computed(() => {
   if (!is_swiping.value || !is_horizontal_swipe.value) return 0
   if (swipe_direction.value == "left") return Math.min(length_x.value / 100, 0.8)
   return Math.max(length_x.value / 100, -0.8)
})

const swipe_position = computed(() => {
   if (!is_swiping.value || !is_horizontal_swipe.value) return 0
   if (swipe_direction.value === "left") return Math.min(length_x.value / 2, 50)
   return Math.max(length_x.value / 2, -50)
})
</script>

<style>
.app-container {
   position: relative;
   width: 100%;
   height: 100%;
}

.page-content {
   width: 100%;
   height: 100%;
   touch-action: pan-y;
   transition: transform 0.15s ease;
}

.absolute-center {
   position: absolute;
   top: 50%;
   left: 50%;
   transform: translate(-50%, -50%);
}
</style>
