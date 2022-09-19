import { defineStore } from "pinia";
import { request } from "../api";

export interface Meta {
    SiteName: string
}

export const createMetaStore = defineStore('meta', {
    state() {
        return {
            SiteName: ''
        }
    },
    actions: {
        async update() {
            const data: Meta = await request('/api/v1/meta')
            this.SiteName = data.SiteName
        }
    }
})