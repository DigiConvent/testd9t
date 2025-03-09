import { nextTick } from 'vue'
import { createI18n } from 'vue-i18n'

export const support_locales = ['de', 'en']

export function setupI18n(options = { locale: 'de' }) {
  const i18n = createI18n(options)
  setI18nLanguage(i18n, options.locale)
  return i18n
}

export function setI18nLanguage(i18n: any, locale: string) {
  if (i18n.mode === 'legacy') {
    i18n.global.locale = locale
  } else {
    i18n.global.locale.value = locale
  }
  
  document.querySelector('html')!.setAttribute('lang', locale)
}

export async function loadLocaleMessages(i18n: any, locale: string) {
  const messages = await import(
    `./locales/${locale}.json`
  )

  i18n.global.setLocaleMessage(locale, messages.default)

  return nextTick()
}