import { Model } from "./";

export interface User extends Model {
    name: string;
    password: string;
    email: string;
    role: Role
    role_id: number;
}

export interface Role extends Model {
    name: string
}