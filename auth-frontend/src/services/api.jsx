import axios from 'axios';

const API = axios.create({
  baseURL: 'http://localhost:8000/api',
  withCredentials: true,
});

export const signup = (data) => API.post('/user/signup', data);
export const login = (data) => API.post('/user/login', data);
export const googleOAuth = () => window.location.href = 'http://localhost:8000/api/oauth/google';
export const getSession = () => API.get('/session');
