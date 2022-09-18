import { defineStore } from "pinia";
import { request } from "../api";
import { Group, ProfileData } from "./author";

export interface Post {
    ID: number
    Content: string
    CreatedAt: string
    Author: ProfileData
    Group?: Group
    Favorited: boolean
    CommentCount: number
    Comments?: Comment[]
}

export interface Comment {
    ID: number
    Content: string
    CreatedAt: string
    Author: ProfileData
}

export const createSearchStore = defineStore('search', {
    actions: {
        async getTopPostList(): Promise<Post[]> {
            return await request('/api/v1/posts')
        },
        async setFavorite(post_id: number) {
            await request(`/api/v1/posts/${post_id}/favorite_users`, {
                method: 'POST'
            })
        },
        async getPost(id: number): Promise<Post> {
            return await request(`/api/v1/posts/${id}`)
        },
        async getPostCommentList(id: number): Promise<Comment[]> {
            return await request(`/api/v1/posts/${id}/comments`)
        },
        async createComment(post_id: number, content: string) {
            await request(`/api/v1/posts/${post_id}/comments`, {
                method: 'POST',
                body: JSON.stringify({ Content: content })
            })
        }
    }
})