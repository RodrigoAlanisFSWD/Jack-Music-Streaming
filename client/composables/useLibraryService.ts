import { Library } from "~/models/library"
import api from "./api"

export const useLibraryService = () => {

    const libraryStore = useLibraryStore()

    const state = storeToRefs(libraryStore)

    const getLibrary = async () => {
        const { data } = await api.get("/library/")

        libraryStore.setLibrary(data)
    }

    const addSong = async (song: number) => {
        try {
            const { data } = await api.post("/library/addSong/" + song)

            libraryStore.setLibrary(data)
        } catch (error) {
            console.log(error)
        }
    }

    const addPlaylist = async (playlist: number) => {
        try {
            const { data } = await api.post("/library/addPlaylist/" + playlist)

            libraryStore.setLibrary(data)
        } catch (error) {
            console.log(error)
        }
    }

    const addAlbum = async (album: number) => {
        try {
            const { data } = await api.post("/library/addAlbum/" + album)

            libraryStore.setLibrary(data)
        } catch (error) {
            console.log(error)
        }
    }

    

    return {
        getLibrary,
        addAlbum,
        addSong,
        addPlaylist,
        libraryStore,
        state
    }
}