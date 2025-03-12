export interface TelegramWebApp {
   initData: string
   initDataUnsafe: {
      query_id: string
      user: {
         id: number
         is_bot: boolean
         first_name: string
         last_name: string
         username: string
         language_code: string
         photo_url: string
         allow_write_to_pm: boolean
      }
      chat: {
         id: number
         type: string
         title: string
         username: string
         photo_url: string
      }
      chat_type: string
      chat_instance: string
      hash: string
      signature: string
   }
   version: string
   platform: string
   colorScheme: "light" | "dark"

   isExpanded: boolean
   HapticFeedback: {
      impactOccurred(style: "light" | "medium" | "heavy" | "rigid" | "soft"): void
      notificationOccurred(type: "error" | "warning" | "success"): void
      selectionChanged(): void
   }
   downloadFile(
      params: {
         url: string
         file_name: string
      },
      callback: any,
   ): void

   expand(): void
   close(): void
}

declare global {
   interface Window {
      Telegram: any
   }
}

export function is_mini_app(): boolean {
   return window["Telegram"] != undefined && window.Telegram.WebApp.initData != ""
}

export default function get_web_app(): TelegramWebApp {
   if (window.Telegram != undefined) {
      return window.Telegram!.WebApp as TelegramWebApp
   }
   return {} as TelegramWebApp
}
