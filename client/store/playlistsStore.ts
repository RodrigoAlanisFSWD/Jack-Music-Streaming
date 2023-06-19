import { Playlist } from "~~/models/playlist";

export interface PlaylistsState {
    playlists: Playlist[],
    error: string,
}

export const usePlaylistsStore = defineStore("PlaylistsStore", {
    state: (): PlaylistsState => ({
        playlists: [],
        error: "",
    }),
    actions: {
        setError(error: string) {
            this.error = error;
        },
        setPlaylists(playlists: Playlist[]) {
            this.playlists = playlists
        },
        deletePlaylist(playlist: Playlist) {
            this.playlists = this.playlists.filter((p: Playlist) => p.id !== playlist.id)
        },
        addPlaylist(playlist: Playlist) {
            this.playlists = [
                ...this.playlists,
                playlist
            ]
        }
    }
})

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(usePlaylistsStore, import.meta.hot))
}