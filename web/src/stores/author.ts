import { defineStore } from "pinia";
import { request } from "../api";

export interface ProfileData {
    ID: number
    Name: string
    Account: string
    AvatarURL: string
    Intro: string
}

export interface Group {
    ID: number
    Name: string
    Intro: string
}

export const createAuthorStore = defineStore('author', {
    actions: {
        async getProfile(): Promise<ProfileData> {
            return await request('/api/v1/profile')
        },
        async getGroupList(): Promise<Group[]> {
            return await request('/api/v1/groups')
        },
        async publish(content: string, group_id: number) {
            return await request('/api/v1/posts', {
                method: 'POST',
                body: JSON.stringify({ Content: content, GroupID: group_id })
            })
        },
        async updateProfile(name: string, intro: string) {
            await request('/api/v1/profile', {
                method: 'PUT',
                body: JSON.stringify({ Name: name, Intro: intro })
            })
        }
    }
})