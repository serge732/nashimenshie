import axios from 'axios';

export const moyskladInstance = axios.create({
    baseURL: 'http://localhost:3000/moysklad/',
    withCredentials: true,
});
