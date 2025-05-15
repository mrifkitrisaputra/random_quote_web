import React, { useState } from 'react';

const VoteForm = ({ options, onVote }) => {
    const [selectedOption, setSelectedOption] = useState('');
    const [userKey, setUserKey] = useState(localStorage.getItem('user_key') || '');

    const handleSubmit = () => {
        if (!selectedOption) return alert("Pilih salah satu opsi");
        onVote(selectedOption);
    };

    return (
        <form onSubmit={handleSubmit}>
            {options.map((opt) => (
                <div key={opt.id} className="mb-2">
                    <label className="flex items-center gap-2">
                        <input
                            type="radio"
                            name="vote"
                            value={opt.id}
                            onChange={(e) => setSelectedOption(parseInt(e.target.value))}
                        />
                        <span>{opt.text}</span>
                    </label>
                </div>
            ))}

            <button
                type="button"
                onClick={handleSubmit}
                disabled={!selectedOption}
                className="mt-4 w-full py-2 px-4 bg-blue-600 text-white rounded hover:bg-blue-700 transition"
            >
                Kirim Suara
            </button>
        </form>
    );
};

export default VoteForm;