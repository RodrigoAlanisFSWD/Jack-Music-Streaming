import { Playlist } from "~~/models/playlist"
import { Song } from "~~/models/song"

export const usePlaylistsService = () => {

    const playlistsStore = usePlaylistsStore()

    const state = storeToRefs(playlistsStore)

    const player = usePlayer()

    const getPlaylists = async (page: number) => {
        const { data } = await api.get("/playlist/all/" + page)

        playlistsStore.setPlaylists(data)

        return data
    }

    const getPlaylistsByAuthor = async (authorID: number) => {
        const { data } = await api.get("/playlist/author/" + authorID)

        playlistsStore.setPlaylists(data)

        return data
    }

    const getPlaylist = async (id: any) => {
        const { data } = await api.get("/playlist/" + id)

        return data
    }

    const createPlaylist = async (playlist: any) => {
        try {
            const { data } = await api.post<Playlist>("/playlist/", playlist)

            return data
        } catch (error) {
            playlistsStore.setError("Error At Creating. The Name Already Exists")
        }

    }

    const updatePlaylist = async (playlist: any) => {
        try {
            const { data } = await api.put<Playlist>("/playlist/update", playlist)

            return data
        } catch (error) {
            playlistsStore.setError("Error At Creating. The Name Already Exists")
        }

    }

    const uploadPlaylistLogo = async (playlistID: number, formData: FormData) => {
        try {
            const { data } = await api.post("/playlist/logo/" + playlistID, formData)

            return data
        } catch (error) {
            playlistsStore.setError("Error At Uploading. Try Again Later By Editing Your New Song")
        }
    }

    const deletePlaylist = async (playlistID: number) => {
        try {
            const { data } = await api.delete("/playlist/" + playlistID)

            playlistsStore.deletePlaylist(data)

            player.pause()
            player.playerStore.setStatus(false)
            player.playerStore.resetPlayer()
        } catch (error) {
            playlistsStore.setError("Error At Deleting. Try Again Later")
        }
    }

    const addSong = async (songID: number, playlistID: number) => {
        try {
            await api.put(`/playlist/addSong/${playlistID}/${songID}`)
        } catch (error) {
            playlistsStore.setError("Error At Deleting. Try Again Later")
        }
    }

    return {
        getPlaylists,
        createPlaylist,
        uploadPlaylistLogo,
        playlistsStore,
        state,
        getPlaylistsByAuthor,
        getPlaylist,
        updatePlaylist,
        deletePlaylist,
        addSong
    }
}