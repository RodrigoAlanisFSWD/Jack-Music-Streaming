<template>
    <div class="w-screen h-screen bg-secondary">
        <div class="w-screen h-[calc(100vh-80px)] grid grid-cols-8">
            <app-sidebar />
            <div class="col-span-6">
                <slot />
            </div>
        </div>
        <audio-player></audio-player>
    </div>
</template>