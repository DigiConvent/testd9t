<template>
  <div>
    <h2 class="text-2xl">System</h2>
    <Toast />
    <div v-if="system_status">
      <Card>
        <template #header>
          <h3>{{ $t("sys.data") }}</h3>
        </template>
        <template #content>
          <Knob
            v-model="space.taken_in_perc"
            readonly
            :show-value="true"
            value-template="{value}%"
            :value-color="'#ff8000'"
          />
          <span>{{ $t("sys.total") }}: {{ (space.total / (1024 * 1024 * 1024)).toFixed(2) }}GB</span
          ><br />
          <span
            >{{ $t("sys.occupied") }}: {{ (space.taken / (1024 * 1024 * 1024)).toFixed(2) }}GB</span
          ><br />
          <span
            >{{ $t("sys.free") }}:
            {{ ((space.total - space.taken) / (1024 * 1024 * 1024)).toFixed(2) }}GB</span
          ><br />
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
                  <i v-if="record.done" class="pi pi-check-circle text-green-500"></i>
                  <i
                    v-else-if="record.is.value == ''"
                    class="pi pi-exclamation-circle text-red-400"
                  ></i>
                  <i v-else class="pi pi-question-circle text-yellow-400"></i>

                  <span class="font-bold whitespace-nowrap">{{ record.name }}</span>
                </span>
              </AccordionHeader>
              <AccordionContent>
                <div v-if="!record.done">
                  <InputGroup v-if="record.is.type != ''">
                    <Badge>{{ record.is.type }}</Badge>
                    <Badge severity="secondary">{{ record.is.domain }}</Badge>
                    <Badge severity="info">{{ record.is.shortened }}</Badge>
                    <!-- <Badge severity="danger" v-if="record.is.value != record.should.value">{{
                      record.should.shortened
                    }}</Badge> -->
                  </InputGroup>
                  <div v-else>{{ record.is }}</div>
                  {{ record.is.value }}<br />
                  {{ $t("sys.dns.invalid", { dns: record.name }) }}
                  <p class="m-0">
                    {{
                      $t("sys.dns.fix", { type: record.should.type, domain: record.should.domain })
                    }}
                  </p>
                  <div>
                    <InputGroup @click="copy_to_clipboard(record.should.value)">
                      <InputText :value="record.should.value" readonly class="w-full"></InputText>
                      <InputGroupAddon>
                        <i class="pi pi-copy"></i>
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
    </div>
    <router-view />
  </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { SystemStatus } from "@/api/sys/types"
import { useToast } from "primevue/usetoast"
import { ref } from "vue"
import { useI18n } from "vue-i18n"

const t = useI18n().t

const system_status = ref<SystemStatus>()
const toast = useToast()

const space = ref<{ total: number; taken: number; taken_in_perc: number }>({
  total: 0,
  taken: 0,
  taken_in_perc: 0
})

type DnsEntry = {
  type: string
  domain: string
  value: string
  shortened: string
}

const copy_to_clipboard = (input: string) => {
  navigator.clipboard.writeText(input)
  toast.add({
    severity: "info",
    detail:
      input.substring(0, 4) +
      "..." +
      input.substring(input.length - 4) +
      " " +
      t("sys.dns.copied_to_clipboard")
  })
}

const dns_checklist = ref<
  {
    name: string
    is: DnsEntry
    should: DnsEntry
    done: boolean
  }[]
>([])

api.sys.status().then((fold) => {
  fold.fold(
    (err: string) => {
      toast.add({
        severity: "error",
        summary: "Error",
        detail: err
      })
    },
    (data: SystemStatus) => {
      system_status.value = data
      space.value.taken = data.server.total_space - data.server.free_space
      space.value = {
        total: data.server.total_space,
        taken: data.server.total_space - data.server.free_space,
        taken_in_perc: 0
      }
      space.value.taken_in_perc = Math.round((space.value.taken / space.value.total) * 100)

      for (const record of ["mx", "dkim", "spf", "dns", "dmarc"]) {
        const is_segs = data.dns[record + "_is"].split("///")
        const should_segs = data.dns[record + "_should"].split("///")
        dns_checklist.value.push({
          name: record,
          should: {
            type: should_segs[0],
            domain: should_segs[1],
            value: should_segs[2],
            shortened: shortened_value(should_segs[2])
          },
          is: {
            type: is_segs[0],
            domain: is_segs[1],
            value: is_segs[2],
            shortened: shortened_value(is_segs[2])
          },
          done: data.dns[record + "_is"] == data.dns[record + "_should"]
        })
      }
    }
  )
})

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
</script>
