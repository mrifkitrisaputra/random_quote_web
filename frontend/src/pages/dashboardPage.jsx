import React, { useEffect, useState } from 'react';
import { fetchDashboardPolls, closePoll, deletePoll } from '../service/api';
import PollCard from '../components/pollCard';

const DashboardPage = () => {
    const [polls, setPolls] = useState([]);

    useEffect(() => {
        const loadPolls = async () => {
            const data = await fetchDashboardPolls();
            setPolls(data || []);
        };

        loadPolls();
    }, []);

    const refresh = async () => {
        const data = await fetchDashboardPolls();
        setPolls(data || []);
    };

    return (
        <div className="min-h-screen bg-gray-50 p-6">
            <div className="max-w-4xl mx-auto bg-white p-6 rounded shadow-md">
                <h1 className="text-2xl font-bold mb-6">Dashboard Poll Creator</h1>

                {polls.length === 0 ? (
                    <p>Belum ada polling</p>
                ) : (
                    polls.map(poll => (
                        <PollCard key={poll.id} poll={poll} onRefresh={refresh} />
                    ))
                )}

                <a href="/create" className="inline-block mt-4 bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
                    Buat Polling Baru
                </a>
            </div>
        </div>
    );
};

export default DashboardPage;