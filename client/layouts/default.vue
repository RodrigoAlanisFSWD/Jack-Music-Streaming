<template>
    <div class="w-screen h-screen bg-secondary">
        <app-navbar></app-navbar>
        <div class="w-screen h-[calc(100vh-70px)]">
            <slot />
        </div>
    </div>
</template>