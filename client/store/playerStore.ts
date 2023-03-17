import { Song } from "~~/models/song"

export interface PlayerState {
    song: Song | null;
    mediaUrl: string;
}

export const usePlayerStore = defineStore("PlayerStore", {
    state: (): PlayerState => ({
        song: null,
        mediaUrl: ""
    }),
    actions: {
        setSong(song: Song, mediaUrl: string) {
            this.song = song;
            this.mediaUrl = mediaUrl;
        },
        resetPlayer() {
            this.song = null;
            this.mediaUrl = "";
        }
    }
})

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(usePlayerStore, import.meta.hot))
}