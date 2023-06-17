import { Library } from "~/models/library"
import api from "./api"

export const useLibraryService = () => {

    const getLibrary = async (): Promise<Library> => {
        const { data } = await api.get("/library/")

        return data
    }

    const addSong = async (song: number) => {
        try {
            await api.post("/library/addSong/" + song)
        } catch (error) {
            console.log(error)
        }
    }

    const addPlaylist = async (playlist: number) => {
        try {
            await api.post("/library/addPlaylist/" + playlist)
        } catch (error) {
            console.log(error)
        }
    }

    const addAlbum = async (album: number) => {
        try {
            await api.post("/library/addAlbum/" + album)
        } catch (error) {
            console.log(error)
        }
    }

    

    return {
        getLibrary,
        addAlbum,
        addSong,
        addPlaylist
    }
}