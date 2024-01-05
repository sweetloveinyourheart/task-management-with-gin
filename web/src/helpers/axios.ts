import axios from 'axios'
import { BASE_URL } from '../configs/constants'

const axiosInstance = axios.create({
    baseURL: BASE_URL
});

// Add a response interceptor
axiosInstance.interceptors.response.use(
    (response) => {
        // Handle successful responses here
        return response;
    },
    (error) => {
        // Handle error responses here
        if (error.response.status === 401) {
            console.log("Need to refresh token");
        }
        return Promise.reject(error);
    }
);

export default axiosInstance