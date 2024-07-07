import request from './api';

export const addWord = (word) => request('/word', 'POST', { word });
