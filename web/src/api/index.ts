import { createAuthStore } from "../stores/auth"

export interface ServerError {
    code: number,
    message: string
}

export async function request<T>(input: RequestInfo | URL, init?: RequestInit): Promise<T> {
    const auth = createAuthStore()

    init = init ? init : {}
    init.headers = Object.assign({Authorization: auth.token}, init.headers)

    const resp = await fetch(input, init)
    if (resp.ok) {
        return await resp.json()
    }

    const e: ServerError = await resp.json()

    switch (resp.status) {
        case 401:
            auth.logout()
            break;
    }

    throw e
}