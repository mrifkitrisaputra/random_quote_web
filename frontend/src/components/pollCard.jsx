import React, { useState } from 'react';
import { toggleRealtimeVisibility, closePoll, deletePoll } from '../service/api';

const PollCard = ({ poll, onRefresh }) => {
    const [allowRealtime, setAllowRealtime] = useState(poll.allow_realtime_results);

    const handleToggleVisibility = async () => {
        const newVisibility = !allowRealtime;
        await toggleRealtimeVisibility(poll.id, newVisibility);
        setAllowRealtime(newVisibility);
        onRefresh();
    };

    const handleClosePoll = async () => {
        if (window.confirm("Tutup polling ini?")) {
            await closePoll(poll.id);
            onRefresh();
        }
    };

    const handleDeletePoll = async () => {
        if (window.confirm("Hapus polling ini?")) {
            await deletePoll(poll.id);
            onRefresh();
        }
    };

    return (
        <div className="border p-4 mb-4 bg-white rounded shadow-sm">
            <h2 className="font-semibold">{poll.title}</h2>
            <p>{poll.description}</p>
            <p>Status: <strong>{poll.status}</strong></p>
            <p>Jumlah Suara: <strong>{poll.vote_count}</strong></p>

            <label className="flex items-center gap-2 mt-2">
                <input
                    type="checkbox"
                    checked={allowRealtime}
                    onChange={handleToggleVisibility}
                />
                <span>Izinkan lihat hasil saat aktif</span>
            </label>

            <div className="mt-4 flex gap-2">
                <button
                    onClick={handleClosePoll}
                    disabled={poll.status === 'closed'}
                    className="bg-yellow-500 text-white px-3 py-1 rounded hover:bg-yellow-600"
                >
                    {poll.status === 'closed' ? 'Polling Ditutup' : 'Tutup Polling'}
                </button>
                <button
                    onClick={handleDeletePoll}
                    className="bg-red-600 text-white px-3 py-1 rounded hover:bg-red-700"
                >
                    Hapus Polling
                </button>
            </div>
        </div>
    );
};

export default PollCard;