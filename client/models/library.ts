import { Song } from "./song"
import { Playlist } from "./playlist"
import { Album } from "./album"
import { Model } from ".";

export interface Library extends Model {
    songs: Song[];
    playlists: Playlist[];
    albums: Album[];
    user_id: number;
}