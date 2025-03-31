<template>
   <div>
      <ProgressBar v-if="loading" mode="indeterminate" />
      <div v-else-if="system_status" :permission="'sys.status.read'">
         <Card class="p-4">
            <template #title>
               <h2 class="text-2xl mb-4">System</h2>
               <div>{{ $t("sys.built_at", { date: system_status.version.built_at }) }}</div>
               <div>{{ $t("sys.online_since", { date: system_status.version.online_since }) }}</div>
            </template>
            <template #content>
               <MeterGroup
                  :value="installation_disk_usage"
                  label-position="start"
                  :max="space.total_home"
               >
                  <template #label="{ value }">
                     <div class="flex flex-wrap gap-4">
                        <template v-for="val of value" :key="val.label">
                           <Card
                              class="flex-1 border border-surface shadow-none"
                              :style="`border-color: ${val.color}`"
                           >
                              <template #content>
                                 <div class="flex justify-between gap-8">
                                    <div class="flex flex-col gap-1">
                                       <span
                                          class="text-surface-500 dark:text-surface-400 text-sm"
                                          >{{ val.label }}</span
                                       >
                                       <span class="font-bold text-lg">{{
                                          format_bytes(val.value)
                                       }}</span>
                                    </div>
                                    <div
                                       class="w-8 h-8 rounded-full inline-flex justify-center items-center text-center"
                                       :style="`background-color: ${val.color};color: #ffffff;`"
                                    >
                                       <i :class="val.icon" />
                                    </div>
                                 </div>
                              </template>
                           </Card>
                        </template>
                     </div>
                  </template>
                  <template #meter="slotProps">
                     <span
                        :class="slotProps.class"
                        :style="`background-color: ${slotProps.value.color}; width: ${slotProps.size}; border-left: ${slotProps.index == 0 ? '0px' : '3px white solid;'}`"
                     />
                  </template>
                  <template #end="{ totalPercent }">
                     <div class="flex justify-between mt-4 mb-2 relative">
                        <span :style="{ width: totalPercent + '%' }"></span>
                        <span class="font-medium"
                           >{{ format_bytes(space.total_home) }}/{{
                              format_bytes(space.total_server - space.rest)
                           }}
                           <br />
                           ({{
                              (
                                 (100 / (space.total_server - space.rest)) *
                                 space.total_home
                              ).toFixed(3)
                           }}%)
                        </span>
                     </div>
                  </template>
               </MeterGroup>
            </template>
         </Card>
         <Card>
            <template #title>
               <h2 class="text-2xl">{{ $t("sys.dns.abbr") }}</h2>
            </template>
            <template #subtitle>{{ $t("sys.dns.title") }}</template>
            <template #content>
               <Accordion value="0">
                  <AccordionPanel v-for="(record, i) of dns_checklist" :key="i" :value="i">
                     <AccordionHeader>
                        <span class="flex items-center gap-2 w-full">
                           <Fa v-if="record.done" icon="circle-check" class="text-green-500" />
                           <Fa
                              v-else-if="record.is.value == ''"
                              icon="circle-exclamation"
                              class="text-red-400"
                           />
                           <Fa v-else icon="circle-question" class="text-yellow-400" />
                           <span class="font-bold whitespace-nowrap flex gap-4"
                              >{{ record.name }}
                              <InputGroup v-if="record.done">
                                 <Badge>{{ record.is.type }}</Badge>
                                 <Badge severity="secondary">{{ record.is.domain }}</Badge>
                                 <Badge severity="info">{{ record.is.shortened }}</Badge>
                              </InputGroup></span
                           >
                        </span>
                     </AccordionHeader>
                     <AccordionContent>
                        <div v-if="!record.done">
                           <InputGroup v-if="record.is.type != ''">
                              <Badge>{{ record.is.type }}</Badge>
                              <Badge severity="secondary">{{ record.is.domain }}</Badge>
                              <Badge severity="info">{{ record.is.shortened }}</Badge>
                           </InputGroup>
                           <div v-else>{{ record.is }}</div>
                           {{ record.is.value }}<br />
                           {{ $t("sys.dns.invalid", { dns: record.name }) }}
                           <p class="m-0">
                              {{
                                 $t("sys.dns.fix", {
                                    type: record.should.type,
                                    domain: record.should.domain,
                                 })
                              }}
                           </p>
                           <div>
                              <InputGroup @click="copy_to_clipboard(record.should.value)">
                                 <InputText
                                    :value="record.should.value"
                                    readonly
                                    class="w-full"
                                 ></InputText>
                                 <InputGroupAddon>
                                    <Fa icon="copy" />
                                 </InputGroupAddon>
                              </InputGroup>
                           </div>
                        </div>
                        <div v-else>
                           {{ $t("sys.dns.valid", { dns: record.name }) }}
                        </div>
                     </AccordionContent>
                  </AccordionPanel>
               </Accordion>
            </template>
         </Card>
         <div class="grid grid-cols-2 gap-4 mt-4">
            <div class="">
               <LogoUpload variant="small" :label="$t('sys.upload_logo.small_label')" />
            </div>
            <div class="">
               <LogoUpload variant="large" :label="$t('sys.upload_logo.large_label')" />
            </div>
         </div>
      </div>
   </div>
</template>

<script lang="ts" setup>
import LogoUpload from "@/components/sys/upload_logo.vue"
import { api } from "@/api"
import type { SystemStatus } from "@/api/sys/types"
import { error, info } from "@/composables/toast"
import { ref } from "vue"
import { useI18n } from "vue-i18n"

const t = useI18n().t

const loading = ref(true)
const system_status = ref<SystemStatus>()

const space = ref<{ total_server: number; total_home: number; rest: number }>({
   total_server: 0,
   total_home: 0,
   rest: 0,
})

type DnsEntry = {
   type: string
   domain: string
   value: string
   shortened: string
}

const copy_to_clipboard = (input: string) => {
   navigator.clipboard.writeText(input)
   info(
      input.substring(0, 4) + "..." + input.substring(input.length - 4),
      t("sys.dns.copied_to_clipboard"),
   )
}

const dns_checklist = ref<
   {
      name: string
      is: DnsEntry
      should: DnsEntry
      done: boolean
   }[]
>([])

const installation_disk_usage = ref<
   { label: string; color: string; value: number; icon: string }[]
>([])

async function load_system_status() {
   loading.value = true
   ;(await api.sys.status()).fold(
      (error_message: string) => {
         error(error_message)
      },
      (data: SystemStatus) => {
         system_status.value = data
         installation_disk_usage.value = fill_colours_and_sort([
            {
               label: "Backend",
               color: "",
               value: data.space.program.backend,
               icon: "server",
            },
            {
               label: "Frontend",
               color: "",
               value: data.space.program.frontend,
               icon: "browser",
            },
            { label: "Iam", color: "", value: data.space.data.iam, icon: "id-card-clip" },
            { label: "Sys", color: "", value: data.space.data.sys, icon: "hard-drive" },
            { label: "Post", color: "", value: data.space.data.post, icon: "mailbox" },
            {
               label: "Certificates",
               color: "",
               value: data.space.data.certificates,
               icon: "file-certificate",
            },
         ])

         space.value = {
            total_server: data.space.total_server,
            total_home: data.space.total_home,
            rest: data.space.os,
         }

         for (const record of ["mx", "dkim", "spf", "dns", "dmarc"]) {
            const j = JSON.parse(JSON.stringify(data.dns))
            const is_segs = j[record + "_is"].split("///")
            const should_segs = j[record + "_should"].split("///")
            dns_checklist.value.push({
               name: record,
               should: {
                  type: should_segs[0],
                  domain: should_segs[1],
                  value: should_segs[2],
                  shortened: shortened_value(should_segs[2]),
               },
               is: {
                  type: is_segs[0],
                  domain: is_segs[1],
                  value: is_segs[2],
                  shortened: shortened_value(is_segs[2]),
               },
               done: j[record + "_is"] == j[record + "_should"],
            })
         }
         loading.value = false
      },
   )
}

const shortened_value = (value: string) => {
   const result: string[] = []
   const segs = value.split("; ")
   for (const seg of segs) {
      if (seg.length > 20) {
         result.push(seg.substring(0, 4) + "..." + seg.substring(seg.length - 4))
      } else {
         result.push(seg)
      }
   }
   return result.join("; ")
}

const colours = [
   "#ef4444",
   "#f59e0b",
   "#84cc16",
   "#10b981",
   "#06b6d4",
   "#3b82f6",
   "#8b5cf6",
   "#d946ef",
   "#f43f5e",
   "#f97316",
   "#eab308",
   "#22c55e",
   "#14b8a6",
   "#0ea5e9",
   "#6366f1",
   "#a855f7",
   "#ec4899",
]
function fill_colours_and_sort<T>(
   entries: { color: string; value: number; [key: string]: any }[] & T[],
): T[] {
   for (let i = 0; i < entries.length; i++) {
      entries[i].color = colours[i % colours.length]
   }
   // sort
   entries.sort((a, b) => {
      return b.value - a.value
   })

   return entries
}

// https://gist.github.com/zentala/1e6f72438796d74531803cc3833c039c
function format_bytes(bytes: number) {
   if (bytes == 0) return "0 Bytes"
   const k = 1024,
      sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"],
      i = Math.floor(Math.log(bytes) / Math.log(k))
   return parseFloat((bytes / Math.pow(k, i)).toFixed(0)) + sizes[i]
}

load_system_status()
</script>
