import { Song } from "~~/models/song"

export const useSongsService = () => {

    const access_token = useCookie('jack_access_token')

    const songsStore = useSongsStore()

    const state = storeToRefs(songsStore)

    const player = usePlayer()

    const getSongs = async (page: number) => {
        const { data } = await api.get("/songs/" + page)

        songsStore.setSongs(data)

        return data
    }

    const getSongsByAuthor = async (authorID: number) => {
        const { data } = await api.get("/songs/user/" + authorID)
        
        songsStore.setSongs(data)

        return data
    }

    const getSong = async (id: any) => {
        const { data } = await api.get("/songs/getOne/" + id)
        
        return data
    }

    const createSong = async (song: any) => {
        try {
            const { data } = await api.post<Song>("/songs/", song)

            return data
        } catch (error) {
            songsStore.setError("Error At Creating. The Name Already Exists")
        }

    }

    const updateSong = async (song: any) => {
        try {
            const { data } = await api.put<Song>("/songs/update", song)

            return data
        } catch (error) {
            songsStore.setError("Error At Creating. The Name Already Exists")
        }

    }

    const uploadSongMedia = async (songID: number, formData: FormData) => {
        try {
            const { data } = await api.post("/songs/uploadMedia/" + songID, formData)

            return data
        } catch (error) {
            songsStore.setError("Error At Uploading. Try Again Later By Editing Your New Song")
        }
    }

    const deleteSong = async (songID: number) => {
        try {
            const { data } = await api.delete("/songs/" + songID)

            songsStore.deleteSong(data)

            if (player.state.song.value?.id === songID) {
                player.pause()
                player.playerStore.setStatus(false)
                player.playerStore.resetPlayer()
            }

        } catch (error) {
            songsStore.setError("Error At Deleting. Try Again Later")
        }
    }

    return {
        getSongs,
        createSong,
        uploadSongMedia,
        songsStore,
        state,
        getSongsByAuthor,
        getSong,
        updateSong,
        deleteSong
    }
}