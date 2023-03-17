export const useSongsService = () => {

    const getSongs = async (page: number) => {
        const { data } = await api.get("/songs/" + page)

        return data
    }

    return {
        getSongs
    }
}