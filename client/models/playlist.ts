import { Model } from ".";
import { User } from "./auth";
import { Media, Song } from "./song";

export interface Playlist extends Model {
    name: string;
    type: string;
    author: User;
    author_id: number;
    logo: Media;
    logo_id: number;
    songs: Song[];
}