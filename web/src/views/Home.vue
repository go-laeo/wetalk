<script lang="ts" setup>
import { HeartIcon, ChatBubbleLeftEllipsisIcon, HashtagIcon, MusicalNoteIcon } from '@heroicons/vue/24/outline';
import { HeartIcon as HeartIconSolid } from '@heroicons/vue/24/solid';
import { onMounted, ref } from 'vue';
import { createSearchStore, Post } from '../stores/search';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'
import { useProgrammatic } from '@oruga-ui/oruga-next'

dayjs.locale('zh-cn')
dayjs.extend(relativeTime)

const { oruga } = useProgrammatic()
const search = createSearchStore()
const posts = ref<Post[]>([])
const refresh = async () => posts.value = await search.getTopPostList()

const setFavorite = async (it: Post) => {
    try {
        if (it.Favorited) {
            return
        }
        await search.setFavorite(it.ID)
        await refresh()
    } catch (e: any) {
        oruga.notification.open({
            message: e.message,
            position: 'top',
            variant: 'warning',
        })
    }
}

onMounted(refresh)
</script>

<template>
    <!-- <div class=" bg-gray-200 p-4 md:rounded-2xl flex flex-row gap-4 text-sm text-gray-700 flex-wrap">
        <router-link v-for="n in Array(20)" :key="n" :to="{ name: 'Home'}"
            class=" block bg-gray-100 px-2 py-1 rounded transition hover:bg-blue-400 hover:text-white">
            圈子1
        </router-link>
    </div> -->
    <div class=" p-5 text-base select-none">全站热门</div>
    <div class="  flex flex-col gap-4">
        <div class=" flex flex-col gap-6 p-4 bg-white rounded-lg" v-for="it in posts" :key="it.ID">
            <div class=" flex flex-row items-center gap-2">
                <div class=" w-16 h-16 overflow-hidden bg-gray-200 rounded-full">
                    <img v-if="it.Author.AvatarURL" :src="it.Author.AvatarURL">
                    <MusicalNoteIcon class=" p-2 text-slate-600" v-else />
                </div>
                <div>
                    <div class=" font-bold">{{ it.Author.Name }}</div>
                    <div class=" text-xs">{{ it.Author.Intro || '^-^' }}</div>
                </div>
            </div>
            <div class="line-clamp-5 text-ellipsis whitespace-pre-wrap">
                {{ it.Content }}
            </div>
            <div class=" flex flex-col gap-2 md:flex-row md:items-center md:justify-between">
                <div class=" flex flex-row items-center gap-4">
                    <div class=" text-slate-700 select-none text-xs">{{ dayjs(it.CreatedAt).fromNow() }}</div>
                    <div v-if="it.Group"
                        class=" bg-gray-200 text-xs text-slate-900 rounded-3xl p-1 px-4 select-none flex flex-row items-center">
                        <HashtagIcon class=" w-4 h-4" />
                        {{ it.Group.Name }}
                    </div>
                </div>
                <div class="flex flex-row items-center gap-4">
                    <button>
                        <HeartIconSolid class=" w-4 h-4" v-if="it.Favorited" />
                        <HeartIcon class=" w-4 h-4" v-else @click="setFavorite(it)" />
                    </button>
                    <router-link :to="{ name: 'ReadPost', params: { id: it.ID } }"
                        class=" flex flex-row items-center gap-1">
                        <ChatBubbleLeftEllipsisIcon class=" w-4 h-4" />
                        <span class=" text-slate-700 text-sm">{{ it.CommentCount }}</span>
                    </router-link>
                </div>
            </div>
        </div>
    </div>
</template>