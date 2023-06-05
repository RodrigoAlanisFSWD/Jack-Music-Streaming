export const useAlbumsService = () => {

    const getAlbums = async (authorId: number) => {
        try {
            const { data } = await api.get("/albums/author/" + authorId)

            return data
        } catch (error) {
            console.log(error)
        }
    }

    const getAlbum = async (id: number) => {
        try {
            const { data } = await api.get("/albums/getOne/" + id)

            return data
        } catch (error) {
            console.log(error)
        }
    }

    const addSong = async (songID: number, albumID: number) => {
        try {
            await api.put(`/albums/addSong/${albumID}/${songID}`)
        } catch (error) {
            console.log(error)
        }
    }

    const removeSong = async (songID: number, albumID: number) => {
        try {
            await api.delete(`/albums/removeSong/${albumID}/${songID}`)
        } catch (error) {
            console.log(error)
        }
    }

    return {
        getAlbum,
        getAlbums,
        addSong,
        removeSong
    }
}