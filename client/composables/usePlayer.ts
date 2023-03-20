import { Howl } from "howler"
import { Ref } from "nuxt/dist/app/compat/capi"
import { Song } from "~~/models/song"

export const usePlayer = () => {

    const playerStore = usePlayerStore()

    const state = storeToRefs(playerStore)

    const { song } = state

    // HTML Refs
    const mainBar: Ref<HTMLDivElement | null> = ref(null)

    // Set Song Store Method
    const setSong = (song: Song) => {
        const src = "http://localhost:8080/api/songs/media/" + song.id;

        const songMedia = new Howl({
            src: [src],
            format: "mp4",
            preload: true,
            onplay: function () {
                requestAnimationFrame(step.bind(this))
            },
            onseek: function () {
                // Start updating the progress of the track.
                requestAnimationFrame(step.bind(self));
            },
            onload: function () {
                playerStore.setDuration(formatTime(songMedia.duration()))
            },
            onend: function () {
                playerStore.setStatus(false)
            }
        })

        playerStore.setSong(songMedia)
    }

    const formatTime = (secs: number) => {
        var minutes = Math.floor(secs / 60) || 0;
        var seconds = (secs - minutes * 60) || 0;

        return minutes + ':' + (seconds < 10 ? '0' : '') + seconds;
    }

    function step() {
        const seek = song.value.seek()
        playerStore.setTime(formatTime(Math.round(seek)))

        playerStore.setBarWidth((((seek / song.value.duration()) * 100) || 0) + '%')

        if (song.value.playing()) {
            requestAnimationFrame(step);
        }
    }

    const seek = (e: any) => {
        if (mainBar.value) {
            var per = e.offsetX / mainBar.value.scrollWidth

            song.value.seek(song.value.duration() * per)
        }
    }

    const play = () => {
        playerStore.setStatus(true)
        song.value.play()
    }

    const pause = () => {
        playerStore.setStatus(false)
        song.value.pause()
    }

    return {
        state,
        playerStore,
        setSong,
        play,
        pause,
        seek,
        mainBar
    }
}