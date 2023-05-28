import { Model } from ".";
import { Album } from "./album";
import { User } from "./auth";

export interface Song extends Model {
    name: string;
    duration: string;
    author: User;
    author_id: number;
    media: Media;
    media_id: number;
    logo: Media;
    logo_id: number;
    album_id: number;
    album: Album;
}

export interface Media extends Model {
    src: string;
    ext: string;
    name: string;
}