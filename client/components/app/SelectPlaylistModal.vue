<script lang="ts" setup>
import { Ref } from 'vue'
import { Playlist } from '~~/models/playlist'

const { show } = defineProps({
    show: Boolean
})

const { getPlaylistsByAuthor } = usePlaylistsService()
const user = useUser()

const playlists: Ref<Playlist[]> = ref([])

const fetchData = async (page: number) => {
    console.log(user.value)
    if (user.value) {
        playlists.value = await getPlaylistsByAuthor(user.value?.id as number)
    }
}

onMounted(async () => {
    await fetchData(1)
})

const emit = defineEmits(['hide'])

const handleClickOutside = () => {
    if (show) {
        emit('hide', false)
    }
}
</script>

<template>
    <Transition enter-active-class="transition ease-out duration-200 transform" enter-from-class="opacity-0"
        enter-to-class="opacity-100" leave-active-class="transition ease-in duration-200 transform"
        leave-from-class="opacity-100" leave-to-class="opacity-0">
        <Teleport to="body">
            <div v-if="show" v-click-outside="handleClickOutside" class="grid grid-cols-8 w-screen h-[calc(100vh-80px)] top-0 left-0 absolute">
                <div class="col-start-3 col-end-9 w-full h-full bg-[#111] bg-opacity-75 z-50">

                </div>
            </div>
        </Teleport>
    </Transition>
</template>