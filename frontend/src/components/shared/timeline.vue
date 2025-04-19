<template>
   <div>
      <canvas
         ref="canvas"
         class="w-full border-gray-500 border p-4 rounded"
         :width="props.width"
         :height="props.height"
         @wheel="handle_wheel"
      ></canvas>
      <div>
         <Badge
            v-for="line in visible_lines"
            :key="line.id"
            :style="`background-color: ${line.colour}`"
         >
            {{ line.name }}
         </Badge>
      </div>
   </div>
</template>

<script lang="ts" setup>
import { get_colour } from "@/utils/colour"
import { onMounted, ref } from "vue"
import { useI18n } from "vue-i18n"
const t = useI18n().t

const props = defineProps<{ data: LifeLine[]; width: number; height: number; show_now?: boolean }>()

type LifeLine = {
   label: string
   start: Date
   end: Date
}
type CanvasLine = {
   label: string
   colour: string
   start: number
   end: number
}

const canvas = ref<HTMLCanvasElement | null>(null)
let ctx: CanvasRenderingContext2D | null = null

const visible_lines = ref<{ id: number; name: string; colour: string }[]>([])
const lines = ref<CanvasLine[]>([])
const center = ref(Date.now())

function h(): number {
   return canvas.value?.height || 0
}

let scale = 1
function handle_wheel(event: WheelEvent) {
   const scroll_down = event.deltaY > 0
   if (event.ctrlKey) {
      event.preventDefault()
      if (scroll_down) {
         if (scale == 4) return
         scale += 1
      } else {
         if (scale == -18) return
         scale -= 1
      }
   } else {
      const diff = center.value - unmap(Math.abs(event.deltaY))
      const earliest = lines.value.map((l) => l.start).reduce((a, b) => Math.min(a, b))
      const latest = lines.value.map((l) => l.end).reduce((a, b) => Math.max(a, b))
      if (scroll_down) {
         const new_center = center.value + diff
         if (latest > new_center) {
            center.value = new_center
            event.preventDefault()
         }
      } else {
         const new_center = center.value - diff
         if (earliest < new_center) {
            center.value = new_center
            event.preventDefault()
         }
      }
   }
   draw()
}

setInterval(() => {
   draw()
}, 250)

const day = 24 * 60 * 60 * 1000
function unmap(value: number): number {
   const y = [
      center.value - day * 182 * Math.pow(2, scale),
      center.value + day * 183 * Math.pow(2, scale),
   ]
   const out_min = Math.min(y[0], y[1])
   const out_max = Math.max(y[0], y[1])
   const in_min = 0
   const in_max = h()
   if (in_min === in_max) {
      throw new Error("Input range cannot be zero-width.")
   }

   if (value < in_min) {
      value = in_min
   }
   if (value > in_max) {
      value = in_max
   }

   const clamped = Math.min(Math.max(value, in_min), in_max)

   const result = out_min + ((clamped - in_min) * (out_max - out_min)) / (in_max - in_min)
   return result
}
function map(value: number): number {
   const y = [
      center.value - day * 182 * Math.pow(2, scale),
      center.value + day * 183 * Math.pow(2, scale),
   ]
   const in_min = Math.min(y[0], y[1])
   const in_max = Math.max(y[0], y[1])
   const out_min = 0
   const out_max = h()
   if (in_min === in_max) {
      throw new Error("Input range cannot be zero-width.")
   }

   if (value < in_min) {
      value = in_min
   }
   if (value > in_max) {
      value = in_max
   }

   const clamped = Math.min(Math.max(value, in_min), in_max)

   const result = out_min + ((clamped - in_min) * (out_max - out_min)) / (in_max - in_min)
   return result
}

function t_start() {
   return center.value - day * 182 * Math.pow(2, scale)
}
function t_end() {
   return center.value + day * 183 * Math.pow(2, scale)
}

function draw() {
   if (ctx == null) return
   if (canvas.value == null) return
   if (lanes.value == null) return
   ctx.clearRect(0, 0, canvas.value.width, canvas.value.height)

   draw_timeline()
   visible_lines.value = []
   for (let i = 0; i < lanes.value.length; i++) {
      for (const line of lanes.value[i].lines) {
         draw_line(line, i)
      }
   }
}

function draw_timeline() {
   if (ctx == null) return
   if (canvas.value == null) return

   const earliest_year = new Date(t_start()).getFullYear()
   const latest_year = new Date(t_end()).getFullYear()
   for (let i = earliest_year; i <= latest_year; i++) {
      const year_start = new Date(`${i}-01-01T00:00:00Z`).getTime()
      const month_duration = map(t_start() + day * 31)
      if (scale < 3) {
         for (let j = 0; j < 12; j++) {
            ctx.strokeStyle = get_fade_colour(2 - scale)
            const month_start = new Date(
               `${i}-${(j + 1).toString().padStart(2, "0")}-01T00:00:00Z`,
            ).getTime()
            if (month_start < t_start() - day * 31 || month_start > t_end()) {
               continue
            }

            if (j % 2 === 0) {
               ctx.fillStyle = "#fafafa"
            } else {
               ctx.fillStyle = "white"
            }
            ctx.fillRect(50, map(month_start), canvas.value.width - 50, month_duration)

            if (scale < -1) {
               // draw days
               const day_duration = map(t_start() + day)
               const days_of_month = new Date(i, j + 1, 0).getDate()
               for (let k = 0; k < days_of_month; k++) {
                  const day_start = new Date(
                     `${i}-${(j + 1).toString().padStart(2, "0")}-${(k + 1).toString().padStart(2, "0")}T00:00:00Z`,
                  ).getTime()
                  if (day_start < t_start() - day || day_start > t_end()) {
                     continue
                  }

                  if (scale < -5) {
                     // draw hours
                     const hour_duration = map(t_start() + day / 24)
                     for (let l = 0; l < 24; l++) {
                        const hour_start = new Date(
                           `${i}-${(j + 1).toString().padStart(2, "0")}-${(k + 1)
                              .toString()
                              .padStart(2, "0")}T${l.toString().padStart(2, "0")}:00:00Z`,
                        ).getTime()
                        if (hour_start < t_start() - day / 24 || hour_start > t_end()) {
                           continue
                        }

                        ctx.beginPath()
                        ctx.moveTo(150, map(hour_start))
                        ctx.lineTo(canvas.value.width, map(hour_start))
                        ctx.closePath()
                        ctx.strokeStyle = get_fade_colour(-6 - scale)
                        ctx.stroke()
                        ctx.font = `${Math.min(15, hour_duration / 2)}px sans-serif`
                        ctx.fillStyle = get_fade_colour(-6 - scale)
                        ctx.fillText(`${l}:00`, 150, map(hour_start) + 3)
                     }
                  }
                  ctx.strokeStyle = get_fade_colour(-2 - scale)
                  ctx.beginPath()
                  ctx.moveTo(100, map(day_start))
                  ctx.lineTo(canvas.value.width, map(day_start))
                  ctx.closePath()
                  ctx.stroke()
                  ctx.font = `${Math.min(15, day_duration / 2)}px sans-serif`
                  ctx.fillStyle = get_fade_colour(-2 - scale)
                  ctx.fillText(`${k + 1}`, 100, map(day_start) + 3)
               }
            }
            ctx.strokeStyle = get_fade_colour(2 - scale)
            ctx.beginPath()
            ctx.moveTo(50, map(month_start))
            ctx.lineTo(canvas.value.width, map(month_start))
            ctx.closePath()
            ctx.stroke()

            const font_size = Math.min(15, month_duration / 2)
            ctx.font = `${font_size}px sans-serif`
            const month = t(`time.month.${j}`)
            ctx.fillStyle = get_fade_colour(2 - scale)
            ctx.fillText(`${month}`, 50, map(month_start) + 3)
         }
      }

      ctx.fillStyle = "black"
      ctx.font = "15px sans-serif"
      ctx.textBaseline = "top"
      ctx.strokeStyle = "darkgrey"
      ctx.beginPath()
      ctx.moveTo(0, map(year_start))
      ctx.lineTo(canvas.value.width, map(year_start))
      ctx.strokeStyle = "#000"
      ctx.fillText(`${i}`, 0, map(year_start) + 3)
      ctx.stroke()
   }

   if (props.show_now && Date.now() > t_start() && Date.now() < t_end()) {
      ctx.beginPath()
      ctx.moveTo(0, map(Date.now()))
      ctx.lineTo(canvas.value.width, map(Date.now()))
      ctx.strokeStyle = "red"
      ctx.stroke()
      ctx.fillStyle = "red"
      ctx.font = "15px sans-serif"
      ctx.textBaseline = "top"
      ctx.fillText(t("time.now"), 0, map(Date.now()) + 3)

      if (scale < -8) {
         ctx.fillStyle = get_fade_colour(-8 - scale)
         const second_duration = map(t_start() + 1000 * 60)
         console.log(second_duration)
         ctx.font = `${Math.min(15, second_duration / 2)}px sans-serif`
         ctx.fillText(
            new Date().getHours().toString().padStart(2, "0") +
               ":" +
               new Date().getMinutes().toString().padStart(2, "0") +
               ":" +
               new Date().getSeconds().toString().padStart(2, "0"),
            0,
            map(Date.now()) + 20,
         )
      }
   }
}

draw()

type Lane = {
   lines: CanvasLine[]
}
const lanes = ref<Lane[]>([])

function init() {
   let max_overlap = 0
   for (let i = 0; i < props.data.length; i++) {
      lines.value.push({
         label: props.data[i].label,
         start: props.data[i].start.getTime(),
         end: props.data[i].end.getTime(),
         colour: get_colour(i),
      })
      let overlap = 1
      for (let j = i + 1; j < props.data.length; j++) {
         if (
            (props.data[i].start < props.data[j].start &&
               props.data[j].start < props.data[i].end) ||
            (props.data[i].start < props.data[j].end && props.data[j].end < props.data[i].end)
         ) {
            overlap += 1
         }
      }

      if (overlap > max_overlap) {
         max_overlap = overlap
      }
   }
   for (let i = 0; i < max_overlap; i++) {
      lanes.value.push({
         lines: [],
      })
   }

   for (let i = 0; i < props.data.length; i++) {
      const line = lines.value[i]
      for (let j = 0; j < max_overlap; j++) {
         let collision = false
         if (!collision) {
            for (const lane_line of lanes.value[j].lines) {
               if (!(line.start > lane_line.end || lane_line.start > line.end)) {
                  collision = true
                  break
               }
            }
         }

         if (!collision) {
            lanes.value[j].lines.push(line)
            break
         }
      }
   }
}

init()

onMounted(() => {
   ctx = (canvas.value! as HTMLCanvasElement).getContext("2d")
   draw()
})

function draw_line(line: CanvasLine, i: number) {
   if (ctx === null) return
   const font_size = 15
   const lane_width = canvas.value!.width / (lanes.value.length + 1)
   const start = map(line.start)
   if (start == h()) return
   const end = map(line.end)
   if (end == 0) return
   visible_lines.value.push({
      id: i,
      name: line.label,
      colour: line.colour,
   })
   ctx.fillStyle = line.colour
   ctx.fillRect((i + 1) * lane_width, start, 10, end - start)
   ctx.fillStyle = "black"
   ctx.font = `${font_size}px sans-serif`
   ctx.textBaseline = "top"
   const w = ctx.measureText(`${line.label}`)
   ctx.fillText(`${line.label}`, (i + 1) * lane_width - 2 - w.width, Math.max(2, start))
   ctx.fillText(
      `${line.label}`,
      (i + 1) * lane_width + 2 + 10,
      Math.min(end - font_size, h() - font_size),
   )
}

function get_fade_colour(index: number) {
   const fades = ["#fafafa", "#eee", "#ccc", "#aaa", "#777", "#000"]
   if (index < 0) index *= -1
   if (index > fades.length - 1) index = fades.length - 1
   const res = fades[index]
   return res
}
</script>
