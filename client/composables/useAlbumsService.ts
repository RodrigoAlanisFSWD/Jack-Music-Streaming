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

    return {
        getAlbum,
        getAlbums
    }
}