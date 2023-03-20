import { Model } from ".";
import { User } from "./auth";

export interface Song extends Model {
    name: string;
    author: User;
    author_id: number;
    media: Media;
    media_id: number;
    logo: Media;
    logo_id: number;
}

export interface Media extends Model {
    src: string;
    ext: string;
    name: string;
}