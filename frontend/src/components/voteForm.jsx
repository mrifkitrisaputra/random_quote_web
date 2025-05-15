import React, { useState } from 'react';
import { castVote } from '../service/api';

const VoteForm = ({ pollID }) => {
    const [optionId, setOptionId] = useState('');
    const [userKey, setUserKey] = useState(localStorage.getItem(`user_key_${pollID}`) || '');
    const [voted, setVoted] = useState(false);
    const [error, setError] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!optionId) {
            setError("Silakan pilih salah satu opsi");
            return;
        }

        try {
            await castVote(pollID, parseInt(optionId), userKey || `user_${Date.now()}`);
            setVoted(true);
            alert("Terima kasih atas suaramu!");
        } catch (err) {
            setError(err.response?.data?.error || "Gagal mengirim suara");
        }
    };

    const options = [
        { id: 1, text: "Sangat Baik" },
        { id: 2, text: "Baik" },
        { id: 3, text: "Biasa Saja" },
        { id: 4, text: "Buruk" }
    ];

    return (
        <div className="max-w-md mx-auto mt-6 p-4 bg-white shadow rounded">
            <h2 className="text-xl font-bold mb-4">Pilih Jawaban</h2>

            <form onSubmit={handleSubmit}>
                {options.map((opt) => (
                    <div key={opt.id} className="mb-2">
                        <label className="flex items-center gap-2">
                            <input
                                type="radio"
                                name="vote"
                                value={opt.id}
                                onChange={(e) => {
                                    setOptionId(e.target.value);
                                    setError("");
                                }}
                                disabled={voted}
                            />
                            <span>{opt.text}</span>
                        </label>
                    </div>
                ))}

                {error && <p className="mt-2 text-red-500">{error}</p>}

                <button
                    type="submit"
                    disabled={voted}
                    className={`mt-4 w-full py-2 px-4 ${
                        voted ? 'bg-gray-400' : 'bg-blue-600 hover:bg-blue-700'
                    } text-white rounded transition`}
                >
                    {voted ? "Suara Telah Dikirim" : "Kirim Suara"}
                </button>
            </form>
        </div>
    );
};

export default VoteForm;