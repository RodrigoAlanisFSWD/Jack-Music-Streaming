import { Album } from "~/models/album"
import { Library } from "~/models/library"
import { Playlist } from "~/models/playlist"
import { Song } from "~/models/song"

export interface LibraryState {
    songs: Song[]
    playlists: Playlist[]
    albums: Album[]
    loaded: Boolean
}

export const useLibraryStore = defineStore('LibraryStore', {
    state: (): LibraryState => ({
        songs: [],
        playlists: [],
        albums: [],
        loaded: false
    }),
    actions: {
        setLibrary(library: Library) {
            console.log(library)
            this.songs = library.songs ? library.songs : [];
            this.playlists = library.playlists ? library.playlists : [];
            this.albums = library.albums ? library.albums : []
            this.loaded = true
        },
    }
})

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useLibraryStore, import.meta.hot))
}