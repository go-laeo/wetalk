<script lang="ts" setup>
import { reactive } from 'vue';
import { createAuthStore } from '../../stores/auth';
import { useProgrammatic } from '@oruga-ui/oruga-next';

const { oruga } = useProgrammatic()
const auth = createAuthStore()

const state = reactive({
    password: '',
    password2: '',
    password3: ''
})

const submit = async () => {
    try {
        await auth.updatePassword(state.password, state.password2)
        state.password = ''
        state.password2 = ''
        state.password3 = ''
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
            <label for="">当前密码</label>
            <input type="text" v-model="state.password" required minlength="6">
        </section>
        <section>
            <label for="">新密码</label>
            <input type="password" v-model="state.password2" required minlength="6">
        </section>
        <section>
            <label for="">密码确认</label>
            <input type="password" v-model="state.password3" required minlength="6">
        </section>
        <section>
            <label for=""></label>
            <button type="submit" :disabled="state.password2 !== state.password3">更新</button>
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
