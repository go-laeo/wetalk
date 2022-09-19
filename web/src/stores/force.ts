import { defineStore } from "pinia";
import { request } from "../api";

export interface GroupData {
    Name: string
}

export const createForceStore = defineStore('force', {
    actions: {
        async createGroup(data: GroupData) {
            await request('/api/v1/groups', {
                method: 'POST',
                body: JSON.stringify(data)
            })
        }
    }
})