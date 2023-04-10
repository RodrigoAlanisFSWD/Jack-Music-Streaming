import { Howl } from "howler";
import { Song } from "~~/models/song";

export interface PlayerState {
    songMedia: Howl;
    song: Song | null;
    duration: string;
    currentTime: string;
    barWidth: string;
    playing: boolean;
    loading: boolean;
}

export const usePlayerStore = defineStore("PlayerStore", {
    state: (): PlayerState => ({
        songMedia: new Howl({
            src: [],
        }),
        duration: "0:00",
        currentTime: "0:00",
        barWidth: "0%",
        playing: false,
        loading: true,
        song: null,
    }),
    actions: {
        setSongMedia(media: Howl) {
            this.songMedia = media;
        },
        resetPlayer() {
            this.songMedia = new Howl({
                src: [],
            });
            this.duration = "0:00"
            this.currentTime = "0:00"
            this.barWidth = "0%"
            this.loading = true
            this.song = null
        },
        setDuration(duration: string) {
            this.duration = duration
        },
        setTime(currentTime: string) {
            this.currentTime = currentTime
        },
        setBarWidth(barWidth: string) {
            this.barWidth = barWidth
        },
        setStatus(playing: boolean, loading: boolean = false) {
            this.playing = playing;
            this.loading = loading;
        },
        setSong(song: Song) {
            this.song = song;
        }
    }
})

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(usePlayerStore, import.meta.hot))
}