import axios from 'axios';

const apiClient = axios.create({
    baseURL: 'http://localhost:8080/api',
    headers: {
        'Content-Type': 'application/json'
    }
});

// 1. Membuat polling baru
export const createPoll = async (data) => {
    const res = await apiClient.post('/polls', data);
    return res.data;
};

// 2. Ambil semua polling untuk dashboard
export const fetchDashboardPolls = async () => {
    const res = await apiClient.get('/dashboard');
    return res.data.polls || [];
};

// 3. Voting
export const castVote = async (pollID, optionID, userKey) => {
    const res = await apiClient.post(`/polls/${pollID}/vote`, {
        option_id: optionID,
        user_key: userKey
    });
    return res.data;
};

// 4. Ambil detail polling
export const getPollDetails = async (pollID) => {
    const res = await apiClient.get(`/polls/${pollID}`);
    return res.data;
};

// 🔁 Tambahkan fungsi tambahan dari routes.go

// 5. Tutup polling
export const closePoll = async (pollID) => {
    const res = await apiClient.put(`/polls/${pollID}/close`, {});
    return res.data;
};

// 6. Hapus polling
export const deletePoll = async (pollID) => {
    const res = await apiClient.delete(`/polls/${pollID}`);
    return res.data;
};

// 7. Ekspor hasil polling
export const exportPollResults = async (pollID) => {
    const res = await apiClient.get(`/polls/${pollID}/export`);
    return res.data;
};

export const toggleRealtimeVisibility = async (pollID, allow) => {
    const res = await apiClient.put(`/polls/${pollID}/toggle-realtime`, {
        allow_realtime: allow
    });
    return res.data;
};