<script lang="ts" setup>
import UserInfo from './account-info/UserInfo.vue';
import Secure from './account-info/Secure.vue';
</script>

<template>
    <div class=" flex-1 flex justify-center items-center w-full max-w-screen-sm mx-auto">
        <o-tabs type="boxed" vertical expanded contentClass=" bg-white rounded-r p-6" tabItemWrapperClass=" !grow-0"
            :animated="false">
            <o-tab-item label="个人资料">
                <Suspense>
                    <UserInfo></UserInfo>
                </Suspense>
            </o-tab-item>
            <o-tab-item label="安全中心">
                <Suspense>
                    <Secure></Secure>
                </Suspense>
            </o-tab-item>
            <!-- <o-tab-item label="社交账号">
                ccc
            </o-tab-item> -->
        </o-tabs>
    </div>
</template>