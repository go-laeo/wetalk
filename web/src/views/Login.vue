<script lang="ts" setup>
import { getCurrentInstance, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { createAuthStore } from '../stores/auth';
import { useProgrammatic } from '@oruga-ui/oruga-next';

const router = useRouter()
const auth = createAuthStore()

const state = reactive({
    account: '',
    password: ''
})

const { oruga } = useProgrammatic()

const submit = async () => {
    try {
        await auth.login(state.account, state.password)
        router.replace({ name: 'Home' })
    } catch (e: any) {
        oruga.notification.open({
            message: e.message,
            rootClass: 'toast-notification',
            position: 'top',
            variant: 'danger',
            duration: 5000,
        })
    }
}
</script>

<template>
    <div class=" flex-1 flex items-center justify-center">
        <div class=" max-w-screen-sm w-full shadow bg-white p-8 rounded">
            <div class=" text-slate-700 select-none">登录</div>
            <form class=" mt-4 flex flex-col gap-4" @submit.prevent="submit">
                <input v-model="state.account" type="text" class=" p-4 bg-gray-100 border-none rounded"
                    placeholder="ACCOUNT" required minlength="5">
                <input v-model="state.password" type="password" class=" p-4 bg-gray-100 border-none rounded"
                    placeholder="PASSWORD" required minlength="6">
                <button type="submit" class=" p-3 bg-blue-500 text-white text-lg hover:bg-blue-400 rounded">登录</button>
            </form>
        </div>
    </div>
</template>