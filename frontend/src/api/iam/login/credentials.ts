import Either from "@/api/core/either"

export default async function credentials(emailaddress: string, password: string, connect_telegram: string) : Promise<Either<string, string>> {
    const url = "/api/iam/login/credentials"
    const body = JSON.stringify({ emailaddress, password, connectTelegram: connect_telegram });
    const request = await fetch(url, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        mode: "same-origin",
        body: body
    });

    const data = await request.json();
    
    if (request.ok) {
        return new Either<string, string>().right(data.token);
    } else {
        return new Either<string, string>().left(request.status + ": " + data['message']);
    }
}