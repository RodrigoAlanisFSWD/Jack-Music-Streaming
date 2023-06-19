import { Song } from "~~/models/song";

export interface SongState {
    songs: Song[],
    error: string,
}

export const useSongsStore = defineStore("SongStore", {
    state: (): SongState => ({
        songs: [],
        error: "",
    }),
    actions: {
        setError(error: string) {
            this.error = error;
        },
        setSongs(songs: Song[]) {
            this.songs = songs
        },
        deleteSong(song: Song) {
            this.songs = this.songs.filter((s: Song) => s.id !== song.id)
        }
    }
})

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useSongsStore, import.meta.hot))
}