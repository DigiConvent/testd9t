import Either from "./either";

interface FromJSON<T> {
  (data: any): T;
}

export async function apiGetRequest<T>(url: string): Promise<Either<string, T>> {
  const result = new Either<string, T>();
  const request = await fetch(url, {
    method: "GET",
    mode: "same-origin",
    headers: {
      "Authorization": "" + (localStorage.getItem('jwt') || "")
    }
  });

  const data = await request.json();

  if (request.ok) {
    return result.right(data);
  } else {
    return result.left(request.status + ": " + data['message']);
  }
}
export async function apiPostRequest<T>(url: string, body: any, formatData: FromJSON<T>): Promise<Either<string, T>> {
  const result = new Either<string, T>();
  body = JSON.stringify(body);
  const request = await fetch(url, {
    method: "POST",
    headers: { "Content-Type": "application/json", "Authorization": localStorage.getItem('jwt') || "" },
    mode: "same-origin",
    body: body
  });

  let data: any = {};
  if (request.status >= 200 && request.status < 300) {
    try {
      data = await request.json();
    } catch (e: any) {
      console.error(e);
    }
  }

  if (request.ok) {
    const formattedData = formatData(data);
    return result.right(formattedData);
  } else {
    return result.left(request.status + ": " + data['message']);
  }
}

export async function apiMultipartPostRequest<T>(url: string, body: any, callback?: FromJSON<T>): Promise<Either<string, T>> {
  const result = new Either<string, T>();
  const request = await fetch(url, {
    method: "POST",
    headers: { "Authorization": "" + (localStorage.getItem('jwt') || "") },
    mode: "same-origin",
    body: body
  });

  let data: any = {};
  if (request.status != 204) {
    try {
      data = await request.json();
    } catch (e: any) {
      console.error(e);
    }
  }

  if (request.ok) {
    return result.right(data);
  } else {
    return result.left(request.status + ": " + data['message']);
  }
}

export async function apiDeleteRequest(url: string): Promise<Either<string, void>> {
  const result = new Either<string, void>();
  const request = await fetch(url, {
    method: "DELETE",
    mode: "same-origin",
    headers: {
      "Authorization": "" + (localStorage.getItem('jwt') || "")
    }
  });

  if (request.ok) {
    return result.right();
  } else {
    const data = await request.json();
    return result.left(request.status + ": " + data['message']);
  }
}