import axios, { AxiosRequestConfig } from 'axios'

const api = axios.create({
    baseURL: "http://localhost:8080/api"
})

api.interceptors.request.use(
    async (config: any) => {
        const access_token = useCookie('jack_access_token')
        config.headers = {
            'Authorization': `Bearer ${access_token.value}`,
        }
        return config;
    },
    error => {
        Promise.reject(error)
    });

// Response interceptor for API calls
api.interceptors.response.use((response) => {
    return response
}, async function (error) {
    const originalRequest = error.config;
    if (error.response.status === 403 && error.response.data === "TOKEN_ERROR" && error.config.url !== "/auth/refresh" && !originalRequest._retry) {
        try {
            originalRequest._retry = true;
        const { refresh } = useAuthService()
        await refresh();
        const access_token = useCookie('jack_access_token')
        axios.defaults.headers.common['Authorization'] = 'Bearer ' + access_token;
        return api(originalRequest);
        } catch (error) {
            return Promise.reject(error); 
        }
    }
    return Promise.reject(error);
});

export default api