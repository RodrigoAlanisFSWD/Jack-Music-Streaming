import { Model } from ".";
import { User } from "./auth";
import { Media, Song } from "./song";

export interface Album extends Model {
    name: string;
    author: User;
    author_id: number;
    logo: Media;
    logo_id: number;
    songs: Song[];
}