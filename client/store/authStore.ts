import { defineStore } from 'pinia'
import { Role, User } from '~~/models/auth';

export enum AuthStatus {
    UNAUTHENTICATED,
    AUTHENTICATED,
    AUTH_ERROR,
    AUTH_INITIAL,
}

export interface AuthState {
    isAuth: boolean;
    user: User | null;
    role: Role | null;
    status: AuthStatus;
}

const initialState: AuthState = { isAuth: false, role: null, user: null, status: AuthStatus.AUTH_INITIAL}

export const useAuthStore = defineStore('auth', {
    state: () => initialState,
    actions: {
        authenticateUser(user: User) {
            this.isAuth = true;
            this.role = user.role;
            this.user = user;
            this.status = AuthStatus.AUTHENTICATED
        },
        authInitial() {
            return initialState
        },
        authError() {
            this.status = AuthStatus.AUTH_ERROR
        },
        unaunthenticated() {
            this.isAuth = false;
            this.status = AuthStatus.UNAUTHENTICATED
        }
    }
})