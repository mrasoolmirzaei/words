import request from './api';

export const addSynonym = (word, synonym) => request(`/synonym/${word}`, 'POST', { 'synonym':synonym });
export const searchSynonym = (query) => request(`/synonyms/${query}`, 'GET');