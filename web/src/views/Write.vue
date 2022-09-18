<script lang="ts" setup>
import { onMounted, reactive, ref } from 'vue';
import { createAuthorStore, Group, ProfileData } from '../stores/author';
import {
    Listbox,
    ListboxButton,
    ListboxOptions,
    ListboxOption,
} from '@headlessui/vue'
import { ChevronUpDownIcon, CheckIcon } from '@heroicons/vue/24/outline';
import { useProgrammatic } from '@oruga-ui/oruga-next'
import { useRouter } from 'vue-router';

const author = createAuthorStore()

const defaultGroup = { ID: 0, Name: '不发送到任何圈子', Intro: '' }

const state = reactive<{
    content: string
    group: Group
}>({
    content: '',
    group: defaultGroup
})

const profile = ref<ProfileData>()
onMounted(async () => {
    profile.value = await author.getProfile()
})

const { oruga } = useProgrammatic()
const router = useRouter()

const submit = async () => {
    try {
        await author.publish(state.content, state.group.ID)
        router.replace({ name: 'Home' })
    } catch (e: any) {
        oruga.notification.open({
            message: e.message,
            variant: 'danger',
            position: 'top',
            duration: 5000
        })
    }
}
</script>

<template>
    <form @submit.prevent="submit" class=" w-full max-w-screen-md mx-auto bg-gray-100 px-14 py-8 flex flex-col gap-4">
        <textarea v-model="state.content" cols="30" rows="10" class=" border-none resize-y rounded-lg shadow" required
            minlength="1" placeholder="写点啥"></textarea>
        <Listbox v-model="state.group" as="div" class=" relative">
            <ListboxButton
                class="relative w-full cursor-default rounded-lg bg-white py-2 pl-3 pr-10 text-left shadow focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm">
                <span class="block truncate">{{ state.group.Name }}</span>
                <span class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                    <ChevronUpDownIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
                </span>
            </ListboxButton>
            <transition leave-active-class="transition duration-100 ease-in" leave-from-class="opacity-100"
                leave-to-class="opacity-0">
                <ListboxOptions
                    class="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                    <ListboxOption v-slot="{ active, selected }" :value="defaultGroup">
                        <li :class="[
                          active ? 'bg-blue-100 text-blue-900' : 'text-gray-900',
                          'relative cursor-default select-none py-2 pl-10 pr-4',
                        ]">
                            <span :class="[
                              selected ? 'font-medium' : 'font-normal',
                              'block truncate',
                            ]">
                                不发送到任何圈子
                            </span>
                            <span v-if="selected"
                                class="absolute inset-y-0 left-0 flex items-center pl-3 text-blue-600">
                                <CheckIcon class="h-5 w-5" aria-hidden="true" />
                            </span>
                        </li>
                    </ListboxOption>
                    <ListboxOption v-slot="{ active, selected }" v-for="it in profile?.GroupSet" :key="it.ID"
                        :value="it">
                        <li :class="[
                          active ? 'bg-blue-100 text-blue-900' : 'text-gray-900',
                          'relative cursor-default select-none py-2 pl-10 pr-4',
                        ]">
                            <span :class="[
                              selected ? 'font-medium' : 'font-normal',
                              'block truncate',
                            ]">
                                {{ it.Name }}
                            </span>
                            <span v-if="selected"
                                class="absolute inset-y-0 left-0 flex items-center pl-3 text-blue-600">
                                <CheckIcon class="h-5 w-5" aria-hidden="true" />
                            </span>
                        </li>
                    </ListboxOption>
                </ListboxOptions>
            </transition>
        </Listbox>
        <div class=" flex items-center justify-center">
            <button type="submit"
                class=" min-w-[12rem] bg-blue-500 hover:bg-blue-400 text-white p-3 rounded">发布</button>
        </div>
    </form>
</template>
