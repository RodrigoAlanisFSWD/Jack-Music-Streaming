import { Model } from "./";
import { Media, Song } from './song'

export interface User extends Model {
    name: string;
    password: string;
    email: string;
    role: Role
    role_id: number;
    profile: Profile;
}

export interface Profile extends Model {
    user_id: number;
    songs: Song[];
    logo: Media;
    logo_id: number;
}

export interface Role extends Model {
    name: string
}