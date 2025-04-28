import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

export const handleTypingRequest = async (data) => {
    try {
        const response = await axios.post(`${API_BASE_URL}/typing`, data);
        return response.data;
    } catch (error) {
        console.error('Error handling typing request:', error);
        throw error;
    }
};