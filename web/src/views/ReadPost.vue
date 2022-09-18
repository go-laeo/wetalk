<script lang="ts" setup>
import { HeartIcon, AtSymbolIcon, MusicalNoteIcon } from '@heroicons/vue/24/outline';
import { HeartIcon as HeartIconSolid } from '@heroicons/vue/24/solid';
import { onMounted, reactive, ref } from 'vue';
import { Comment, createSearchStore, Post } from '../stores/search';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'
import { useProgrammatic } from '@oruga-ui/oruga-next'
import { createAuthStore } from '../stores/auth';

dayjs.locale('zh-cn')
dayjs.extend(relativeTime)

const { oruga } = useProgrammatic()
const auth = createAuthStore()
const props = defineProps<{ id: number }>()
const search = createSearchStore()

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

const it = ref<Post>()
const refresh = async () => it.value = await search.getPost(props.id)
onMounted(refresh)

const state = reactive({
    content: ''
})
const submit = async () => {
    try {
        await search.createComment(it.value!.ID, state.content)
        state.content = ''
        await refreshComments()
    } catch (e: any) {
        oruga.notification.open({
            message: e.message,
            position: 'top',
            variant: 'danger'
        })
    }
}

const comments = ref<Comment[]>([])
const refreshComments = async () => comments.value = await search.getPostCommentList(props.id)
onMounted(refreshComments)
</script>

<template>
    <div class=" mt-4 flex flex-col gap-4" v-if="it">
        <div class=" flex flex-col gap-6 p-4 bg-white rounded-lg">
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
                        <AtSymbolIcon class=" w-4 h-4" />
                        {{ it.Group.Name }}
                    </div>
                </div>
                <div class="flex flex-row items-center gap-4">
                    <button>
                        <HeartIconSolid class=" w-4 h-4" v-if="it.Favorited" />
                        <HeartIcon class=" w-4 h-4" v-else @click="setFavorite(it!)" />
                    </button>
                </div>
            </div>
        </div>
        <hr>
        <textarea v-model.trim="state.content" cols="30" rows="6" placeholder="emm..."
            class=" outline-none rounded-lg border-none resize-none read-only:bg-transparent read-only:border-gray-200 read-only:cursor-not-allowed"
            :readonly="!auth.token"></textarea>
        <div class=" text-right">
            <button type="submit"
                class=" border bg-blue-600 text-white px-4 py-2 rounded-lg select-none transition disabled:cursor-not-allowed disabled:bg-transparent disabled:border-gray-200 disabled:text-slate-700"
                :disabled="!state.content" @click="submit">发布</button>
        </div>
        <hr>
        <div class=" pb-8 flex flex-col gap-4">
            <div class=" flex flex-col gap-6 p-4 bg-white rounded-lg" v-for="com in comments" :key="com.ID">
                <div class=" flex flex-row items-center gap-2">
                    <div class=" w-16 h-16 overflow-hidden bg-gray-200 rounded-full">
                        <img v-if="com.Author.AvatarURL" :src="com.Author.AvatarURL">
                        <MusicalNoteIcon class=" p-2 text-slate-600" v-else />
                    </div>
                    <div>
                        <div class=" font-bold">{{ com.Author.Name }}</div>
                        <div class=" text-xs">{{ com.Author.Intro || '^-^' }}</div>
                    </div>
                </div>
                <div class="line-clamp-5 text-ellipsis whitespace-pre-wrap">
                    {{ com.Content }}
                </div>
                <div class=" flex flex-col gap-2 md:flex-row md:items-center md:justify-between">
                    <div class=" flex flex-row items-center gap-4">
                        <div class=" text-slate-700 select-none text-xs">{{ dayjs(com.CreatedAt).fromNow() }}</div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>