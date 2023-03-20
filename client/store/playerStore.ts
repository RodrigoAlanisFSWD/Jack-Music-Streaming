import { Howl } from "howler";

export interface PlayerState {
    song: Howl;
    duration: string;
    currentTime: string;
    barWidth: string;
    playing: boolean;
    loading: boolean;
}

export const usePlayerStore = defineStore("PlayerStore", {
    state: (): PlayerState => ({
        song: new Howl({
            src: [],
        }),
        duration: "0:00",
        currentTime: "0:00",
        barWidth: "0%",
        playing: false,
        loading: true,
    }),
    actions: {
        setSong(song: Howl) {
            this.song = song;
        },
        resetPlayer() {
            this.song = new Howl({
                src: [],
            });
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
        }
    }
})

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(usePlayerStore, import.meta.hot))
}