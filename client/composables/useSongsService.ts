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

    const createSong = async (song: any) => {
        try {
            console.log(access_token)
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
        getSongsByAuthor
    }
}