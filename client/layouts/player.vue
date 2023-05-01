<template>
    <div class="w-screen h-screen bg-secondary">
        <app-navbar></app-navbar>
        <div class="w-screen h-[calc(100vh-155px)]">
            <slot />
        </div>
        <audio-player></audio-player>
    </div>
</template>