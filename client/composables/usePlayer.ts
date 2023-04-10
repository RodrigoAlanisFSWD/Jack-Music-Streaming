import { Howl } from "howler"
import { Ref } from "nuxt/dist/app/compat/capi"
import { Song } from "~~/models/song"

export const usePlayer = () => {

    const playerStore = usePlayerStore()

    const state = storeToRefs(playerStore)

    const { songMedia, playing } = state

    // HTML Refs
    const mainBar: Ref<HTMLDivElement | null> = ref(null)

    // Set Song Store Method
    const setSong = (song: Song) => {
        if (playing.value) {
            songMedia.value.stop()
        }

        playerStore.resetPlayer()
        const src = "http://localhost:8080/api/file/" + song.media_id;

        const newSongMedia = new Howl({
            src: [src],
            format: ["mp3", "mp3", "m4a"],
            preload: true,
            onplay: function () {
                requestAnimationFrame(step.bind(this))
            },
            onseek: function () {
                // Start updating the progress of the track.
                requestAnimationFrame(step.bind(self));
            },
            onload: function () {
                playerStore.setDuration(formatTime(newSongMedia.duration()))
                newSongMedia.play()
                playerStore.setStatus(true)
            },
            onend: function () {
                playerStore.setStatus(false)
            }
        })

        playerStore.setSongMedia(newSongMedia)
        playerStore.setSong(song)
    }

    const formatTime = (secs: number) => {
        const m = Math.floor(secs % 3600 / 60).toString().padStart(2),
            s = Math.floor(secs % 60).toString().padStart(2, '0');

        return m + ':' + s;
    }

    function step() {
        const seek = songMedia.value.seek()
        playerStore.setTime(formatTime(Math.round(seek)))

        playerStore.setBarWidth((((seek / songMedia.value.duration()) * 100) || 0) + '%')

        if (songMedia.value.playing()) {
            requestAnimationFrame(step);
        }
    }

    const seek = (e: any) => {
        if (mainBar.value) {
            var per = e.offsetX / mainBar.value.scrollWidth

            songMedia.value.seek(songMedia.value.duration() * per)
        }
    }

    const play = () => {
        playerStore.setStatus(true)
        songMedia.value.play()
    }

    const pause = () => {
        playerStore.setStatus(false)
        songMedia.value.pause()
    }

    return {
        state,
        playerStore,
        play,
        pause,
        seek,
        mainBar,
        setSong
    }
}