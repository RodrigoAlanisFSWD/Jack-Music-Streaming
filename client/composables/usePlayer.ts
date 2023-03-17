import { Song } from "~~/models/song"

export const usePlayer = () => {

    const playerStore = usePlayerStore()

    const state = storeToRefs(playerStore)

    const setSong = (song: Song) => {
        playerStore.setSong(song, "http://localhost:8080/api/songs/media/" + song.id)
    }

    const resetPlayer = () => {
        playerStore.resetPlayer()
    }

    return {
        state,
        playerStore,
        setSong,
        resetPlayer
    }
}