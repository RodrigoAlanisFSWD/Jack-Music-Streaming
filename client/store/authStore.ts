import { Role, User } from '~~/models/auth';

export enum AuthStatus {
    UNAUTHENTICATED = 'UNAUTHENTICATED ',
    AUTHENTICATED = 'AUTHENTICATED',
    AUTH_ERROR = 'AUTH_ERROR',
    AUTH_INITIAL = 'AUTH_INITIAL',
}

export interface AuthState {
    isAuth: boolean;
    user: User | null;
    role: Role | null;
    status: AuthStatus;
}


export const useAuthStore = defineStore('AuthStore', {
    state: (): AuthState => ({ isAuth: false, role: null, user: null, status: AuthStatus.AUTH_INITIAL}),
    actions: {
        authenticateUser(user: User) {
            this.isAuth = true;
            this.role = user.role;
            this.user = user;
            this.status = AuthStatus.AUTHENTICATED
        },
        authInitial() {
            this.isAuth = false;
            this.role = null;
            this.user = null;
            this.status = AuthStatus.AUTH_INITIAL
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

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useAuthStore, import.meta.hot))
}