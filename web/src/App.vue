<script setup lang="ts">
import { createAuthStore } from './stores/auth';
import { UserCircleIcon, LockOpenIcon, CircleStackIcon, PaperAirplaneIcon } from '@heroicons/vue/24/outline';
import { Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue'

const auth = createAuthStore()
</script>

<template>
    <div class=" bg-gray-100">
        <div class=" w-full xl:max-w-screen-xl min-h-screen mx-auto flex flex-col">
            <nav class=" flex flex-row items-center p-5 gap-4 md:gap-12 text-slate-900">
                <router-link :to="{name: 'Home'}" class=" font-semibold font-serif text-2xl select-none text-slate-900">
                    WETALK
                </router-link>
                <div class=" flex-1"></div>
                <template v-if="auth.token">
                    <router-link :to="{ name: 'ForceMarket' }">权柄</router-link>
                    <router-link :to="{ name: 'Write' }" class=" bg-blue-500 text-white px-4 py-2 rounded-2xl">
                        <div class=" flex items-center gap-1">
                            <PaperAirplaneIcon class=" w-4 h-4" />
                            发布
                        </div>
                    </router-link>
                    <Menu as="div" class=" relative">
                        <MenuButton>
                            <UserCircleIcon class=" w-8 aspect-square rounded-full" />
                        </MenuButton>
                        <transition enter-active-class="transition duration-100 ease-out"
                            enter-from-class="transform scale-95 opacity-0"
                            enter-to-class="transform scale-100 opacity-100"
                            leave-active-class="transition duration-75 ease-out"
                            leave-from-class="transform scale-100 opacity-100"
                            leave-to-class="transform scale-95 opacity-0">
                            <MenuItems
                                class=" absolute z-50 right-0 min-w-[8rem] flex flex-col p-1 origin-top-right divide-y divide-gray-100 rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                                <!-- <MenuItem v-slot="{ active }">
                                <router-link :to="{ name: 'Home' }" :class="[
                                  active ? 'bg-blue-500 text-white' : 'text-gray-900',
                                  'flex gap-1 items-center rounded-md p-3 text-base',
                                ]">
                                    <InformationCircleIcon class=" w-5 aspect-square" />
                                    共识
                                </router-link>
                                </MenuItem>
                                <MenuItem v-slot="{ active }">
                                <router-link :to="{ name: 'Home' }" :class="[
                                  active ? 'bg-blue-500 text-white' : 'text-gray-900',
                                  'flex gap-1 items-center rounded-md p-3 text-base',
                                ]">
                                    <InformationCircleIcon class=" w-5 aspect-square" />
                                    价值
                                </router-link>
                                </MenuItem> -->
                                <MenuItem v-slot="{ active }">
                                <router-link :to="{ name: 'AccountInfo'}" :class="[
                                  active ? 'bg-blue-500 text-white' : 'text-gray-900',
                                  'flex gap-1 items-center rounded-md p-3 text-base',
                                ]">
                                    <CircleStackIcon class=" w-5 aspect-square" />
                                    个人资料
                                </router-link>
                                </MenuItem>
                                <MenuItem v-slot="{ active }">
                                <button :class="[
                                  active ? 'bg-blue-500 text-white' : 'text-gray-900',
                                  'flex gap-1 items-center rounded-md p-3 text-base',
                                ]" @click="auth.logout">
                                    <LockOpenIcon class=" w-5 aspect-square" />
                                    退出
                                </button>
                                </MenuItem>
                            </MenuItems>
                        </transition>
                    </Menu>
                </template>
                <template v-else>
                    <router-link :to="{ name: 'Login' }">登录</router-link>
                    <router-link :to="{ name: 'CreateAccount' }"
                        class=" p-2 px-4 bg-slate-900 text-white hover:bg-slate-700 rounded">
                        注册
                    </router-link>
                </template>
            </nav>
            <hr>
            <router-view></router-view>
        </div>
    </div>
</template>

<style>
@tailwind base;
@tailwind components;
@tailwind utilities;

:root {
    --oruga-variant-primary: rgb(14, 137, 209);
    --oruga-variant-danger: rgb(216, 60, 60);
    --oruga-tabs-link-active-color: rgb(14, 137, 209);
    --oruga-tabs-boxed-link-active-border-color: transparent;
    --oruga-tabs-border-bottom-color: transparent;
}
</style>
