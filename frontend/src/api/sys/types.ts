export function to_system_status(data: any) : SystemStatus {
    console.log(data);
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
        server: {
            data_space: data.server.data_space,
            free_space: data.server.free_space,
            total_space: data.server.total_space,
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
        }
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
	server: {
		data_space: number
		free_space: number
		total_space: number
	}
	telegram_bot:  {
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