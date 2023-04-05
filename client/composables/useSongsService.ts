import { Song } from "~~/models/song"

export const useSongsService = () => {

    const access_token = useCookie('jack_access_token')

    const songsStore = useSongsStore()

    const state = storeToRefs(songsStore)

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
            const { data } = await api.post<Song>("/songs/", song, {
                headers: {
                    "Authorization": "Bearer " + access_token.value
                }
            })

            return data
        } catch (error) {
            songsStore.setError("Error At Creating. The Name Already Exists")
        }

    }

    const updateSong = async (song: any) => {
        try {
            const { data } = await api.put<Song>("/songs/update", song, {
                headers: {
                    "Authorization": "Bearer " + access_token.value
                }
            })

            return data
        } catch (error) {
            songsStore.setError("Error At Creating. The Name Already Exists")
        }

    }

    const uploadSongMedia = async (songID: number, formData: FormData) => {
        try {
            const { data } = await api.post("/songs/uploadMedia/" + songID, formData, {
                headers: {
                    "Authorization": "Bearer " + access_token.value
                }
            })

            return data
        } catch (error) {
            songsStore.setError("Error At Uploading. Try Again Later By Editing Your New Song")
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
        updateSong
    }
}