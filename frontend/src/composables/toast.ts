import toasteventbus from "primevue/toasteventbus"

export function info(summary: string, detail: string = "", life: number = 3000) {
   toasteventbus.emit("add", { severity: "info", summary, detail, life })
}

export function warn(summary: string, detail: string = "", life: number = 3000) {
   toasteventbus.emit("add", { severity: "warn", summary, detail, life })
}

export function success(summary: string, detail: string = "", life: number = 3000) {
   toasteventbus.emit("add", { severity: "success", summary, detail, life })
}

export function error(summary: string, detail: string = "", life: number = 3000) {
   toasteventbus.emit("add", { severity: "error", summary, detail, life })
}
