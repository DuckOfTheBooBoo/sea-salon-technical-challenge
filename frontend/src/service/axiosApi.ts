import axios, { type AxiosInstance, type AxiosResponse, type AxiosError } from 'axios';
import * as constants from '@/constants';
import router from '@/router';

const apiClient: AxiosInstance = axios.create({
    baseURL: 'http://localhost:3000/api',
    headers: {
        'Content-Type': 'application/json',
    },
})

apiClient.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers = {
            ...config.headers,
            Authorization: `Bearer ${token}`,
        };
    }
    return config;
}, (error) => Promise.reject(error));

apiClient.interceptors.response.use((response: AxiosResponse) => {
    return response;
}, (error) => {
    const err: AxiosError = error as AxiosError;
    if (err.response?.status === constants.HTTP_UNAUTHORIZED) {
        router.push({ name: 'login' });
    }

    return Promise.reject(error);
});

export default apiClient;