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
        }
    }
})