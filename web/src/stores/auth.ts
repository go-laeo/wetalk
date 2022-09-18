import { defineStore } from "pinia";
import { request } from "../api";

export interface TokenData {
    Token: string
}

export const createAuthStore = defineStore('auth', {
    state() {
        return {
            _token: sessionStorage.getItem('token')
        }
    },
    getters: {
        token(): string {
            return this._token ? `Bearer ${this._token}` : ''
        }
    },
    actions: {
        async login(account: string, password: string) {
            const token: TokenData = await request('/api/v1/tokens', {
                method: 'POST',
                body: JSON.stringify({ Account: account, Password: password })
            })

            this._token = token.Token
            sessionStorage.setItem('token', token.Token)
        },
        logout() {
            this._token = null
            sessionStorage.removeItem('token')
            this.router.push({ name: 'Login' })
        },
        async createAccount(account: string, password: string, name: string) {
            await request('/api/v1/users', {
                method: 'POST',
                body: JSON.stringify({ Account: account, Password: password, Name: name })
            })
        },
        async updatePassword(password: string, password1: string) {
            await request('/api/v1/profile/password', {
                method: 'POST',
                body: JSON.stringify({ Password: password, Successor: password1 })
            })
        }
    }
})