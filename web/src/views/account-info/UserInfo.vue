<script lang="ts" setup>
import { onMounted, reactive, ref } from 'vue';
import { createAuthorStore, ProfileData } from '../../stores/author';
import { useProgrammatic } from '@oruga-ui/oruga-next';

const { oruga } = useProgrammatic()
const author = createAuthorStore()
const profile = ref<ProfileData>()
const refreshProfile = async () => profile.value = await author.getProfile()
await refreshProfile()

const state = reactive({
    name: profile.value?.Name,
    intro: profile.value?.Intro,
    avatarURL: profile.value?.AvatarURL
})
const submit = async () => {
    try {
        author.updateProfile(state.name!, state.intro!)
        oruga.notification.open({
            message: 'Update successful!',
            position: 'top'
        })
    } catch (e: any) {
        oruga.notification.open({
            message: e.message,
            position: 'top',
            variant: 'danger'
        })
    }
}
</script>

<template>
    <form class=" flex-1 flex flex-col gap-4" @submit.prevent="submit">
        <section>
            <label for="">昵称</label>
            <input type="text" v-model="state.name" required>
        </section>
        <section>
            <label for="">账号</label>
            <input type="text" :value="profile?.Account" readonly>
        </section>
        <section>
            <label for="">自我介绍</label>
            <textarea v-model="state.intro" cols="30" rows="6" class=" resize-none"></textarea>
        </section>
        <section>
            <label for=""></label>
            <button type="submit">更新</button>
        </section>
    </form>
</template>

<style scoped>
    section {
        @apply flex flex-row gap-4 items-center
    }
    label {
        @apply w-16
    }
    input, textarea {
        @apply border-none rounded bg-gray-100 read-only:bg-gray-200 read-only:cursor-not-allowed read-only:focus:border-none read-only:focus:shadow-none read-only:focus:ring-0 read-only:text-slate-400
    }
    button[type="submit"] {
        @apply bg-blue-500 hover:bg-blue-400 text-white px-4 py-2 rounded disabled:bg-blue-200 disabled:cursor-not-allowed
    }
</style>
