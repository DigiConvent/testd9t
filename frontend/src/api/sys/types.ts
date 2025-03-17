export function to_system_status(data: any): SystemStatus {
   console.dir(data)
   return {
      dns: {
         dkim_is: data.dns.dkim_is,
         dkim_should: data.dns.dkim_should,
         dmarc_is: data.dns.dmarc_is,
         dmarc_should: data.dns.dmarc_should,
         dns_is: data.dns.dns_is,
         dns_should: data.dns.dns_should,
         mx_is: data.dns.mx_is,
         mx_should: data.dns.mx_should,
         spf_is: data.dns.spf_is,
         spf_should: data.dns.spf_should,
      },
      space: {
         free: data.space.free,
         total_server: data.space.total_server,
         total_home: data.space.total_home,
         os: data.space.os,
         data: {
            certificates: data.space.data.certificates,
            iam: data.space.data.iam,
            sys: data.space.data.sys,
            post: data.space.data.post,
         },
         program: {
            backend: data.space.program.backend,
            frontend: data.space.program.frontend,
         },
      },
      telegram_bot: {
         telegram_bot_hint: data.telegram_bot.telegram_bot_hint,
         telegram_bot_status: data.telegram_bot.telegram_bot_status,
      },
      version: {
         built_at: data.version.built_at,
         online_since: data.version.online_since,
         database_version: data.version.database_version,
         version: data.version.version,
      },
   }
}

export type SystemStatus = {
   dns: {
      dkim_is: string
      dkim_should: string
      dmarc_is: string
      dmarc_should: string
      dns_is: string
      dns_should: string
      mx_is: string
      mx_should: string
      spf_is: string
      spf_should: string
   }
   space: {
      total_home: number
      total_server: number
      free: number
      data: {
         certificates: number
         iam: number
         sys: number
         post: number
      }
      program: {
         backend: number
         frontend: number
      }
      os: number
   }
   telegram_bot: {
      telegram_bot_hint: string
      telegram_bot_status: string
   }
   version: {
      built_at: string
      online_since: string
      database_version: string
      version: string
   }
}
