<script lang="ts" setup>
import { HashtagIcon } from '@heroicons/vue/24/outline';
import { useProgrammatic } from '@oruga-ui/oruga-next';
import CreateGroupVue from './force-market/CreateGroup.vue';

const { oruga } = useProgrammatic()

function open() {
    oruga.modal.open({
        component: CreateGroupVue,
        trapFocus: true,
        contentClass: 'z-50'
    })
}
</script>

<template>
    <div class=" mt-8 px-2 text-sm text-slate-500 border-l-4 border-slate-700">明码标价，童叟无欺</div>
    <div class=" mt-5 flex flex-row gap-4">
        <div class="card" @click="open">
            <HashtagIcon class=" w-8 aspect-square" />
            <div>添加群组</div>
            <div class=" text-sm">消耗积分 10k</div>
        </div>
    </div>
</template>

<style scoped>
    .card {
        @apply bg-white w-32 aspect-square rounded-lg shadow p-2 text-base text-slate-600 flex flex-col items-center justify-center gap-2 cursor-pointer transition hover:bg-blue-100
    }
</style>
