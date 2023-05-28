<template>
    <app-song-list :playlist="false" :songs="songs"></app-song-list>
</template>

<script lang="ts" setup>
import { Song } from '~/models/song';

const user = useUser()

const router = useRouter()

const { getSongsByAuthor } = useSongsService()

const songs: Ref<Song[]> = ref([])

onMounted(async () => {
    if (user.value) {
        songs.value = await getSongsByAuthor(user.value.id)
    }
})

</script>