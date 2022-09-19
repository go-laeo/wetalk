<script lang="ts" setup>
import { reactive } from 'vue';
import { createForceStore } from '../../stores/force';
import { useProgrammatic } from '@oruga-ui/oruga-next';

const { oruga } = useProgrammatic()
const force = createForceStore()
const emit = defineEmits(['close'])

const state = reactive({
    Name: '',
    Intro: ''
})
const submit = async () => {
    try {
        await force.createGroup(state)
        emit('close')
    } catch (e: any) {
        oruga.notification.open({
            message: e.message,
            position: 'top',
            variant: 'danger',
            duration: 2000,
        })
    }
}
</script>

<template>
    <form @submit.prevent="submit">
        <section>
            <label for="">圈子名称</label>
            <input v-model.trim="state.Name" type="text" required>
        </section>
        <section>
            <label for="">简介</label>
            <textarea v-model.trim="state.Intro" cols="30" rows="6"></textarea>
        </section>
        <section>
            <label for="">费用</label>
            <input type="text" value="10000" disabled>
        </section>
        <section>
            <button type="submit" :disabled="state.Name.length === 0">确定</button>
        </section>
    </form>
</template>

<style scoped>
    form {
        @apply bg-white p-4 flex flex-col gap-4
    }
    section {
        @apply flex flex-col gap-2
    }
    input,textarea {
        @apply border-none rounded outline-none bg-gray-200 resize-none
    }
    button {
        @apply bg-blue-500 hover:bg-blue-400 text-white p-2 rounded shadow transition disabled:bg-blue-300 disabled:cursor-not-allowed
    }
</style>
