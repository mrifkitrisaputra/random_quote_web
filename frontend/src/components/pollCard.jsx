import React, { useState } from 'react';
import { toggleRealtimeVisibility, closePoll, deletePoll } from '../service/api';

const PollCard = ({ poll, onRefresh }) => {
    const [loading, setLoading] = useState(false);

    const handleToggleVisibility = async () => {
        setLoading(true);
        try {
            await toggleRealtimeVisibility(poll.id, !poll.allow_realtime_results);
            onRefresh();
        } catch (err) {
            alert("Gagal mengubah izin realtime");
        } finally {
            setLoading(false);
        }
    };

    const handleCloseOperation = async () => {
        if (!window.confirm("Tutup polling ini?")) return;

        setLoading(true);
        try {
            await closePoll(poll.id);
            onRefresh();
        } catch (err) {
            alert("Gagal menutup polling");
        } finally {
            setLoading(false);
        }
    };

    const handleDelete = async () => {
        if (!window.confirm("Hapus polling ini?")) return;

        setLoading(true);
        try {
            await deletePoll(poll.id);
            onRefresh();
        } catch (err) {
            alert("Gagal menghapus polling");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="border p-4 mb-4 rounded shadow-sm bg-white">
            <h3 className="font-semibold">{poll.title}</h3>
            <p>{poll.description}</p>
            <p>Status: <strong>{poll.status}</strong></p>
            <p>Jumlah Suara: <strong>{poll.vote_count}</strong></p>

            <label className="flex items-center mt-2">
                <input
                    type="checkbox"
                    checked={poll.allow_realtime_results || false}
                    onChange={handleToggleVisibility}
                    disabled={poll.status === "closed"}
                />
                <span className="ml-2">Lihat hasil saat aktif</span>
            </label>

            <div className="mt-4 flex gap-2">
                <button
                    onClick={handleCloseOperation}
                    disabled={poll.status === "closed" || loading}
                    className="bg-yellow-500 text-white px-3 py-1 rounded hover:bg-yellow-600"
                >
                    {poll.status === "closed" ? "Sudah Ditutup" : "Tutup"}
                </button>
                <button
                    onClick={handleDelete}
                    disabled={loading}
                    className="bg-red-600 text-white px-3 py-1 rounded hover:bg-red-700"
                >
                    Hapus
                </button>
            </div>
        </div>
    );
};

export default PollCard;