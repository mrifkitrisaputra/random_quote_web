import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getPollDetails, castVote } from '../service/api';
import PollResultChart from '../components/pollResultChart';
import VoteForm from '../components/voteForm';

const PollDetailPage = () => {
    const { id } = useParams();
    const [pollData, setPollData] = useState(null);
    const [results, setResults] = useState([]);
    const [canSeeResults, setCanSeeResults] = useState(false);
    const [voted, setVoted] = useState(false);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data = await getPollDetails(id);
                setPollData(data.poll);
                setCanSeeResults(data.can_see_results);

                if (data.results) {
                    setResults(data.results);
                }
            } catch (err) {
                alert("Polling tidak ditemukan");
                window.location.href = "/";
            }
        };

        fetchData();
    }, [id]);

    const isParticipant = window.location.pathname.includes('api/polls');

    const handleVote = async (optionID) => {
        const userKey = localStorage.getItem(`user_key_${id}`) || `user_${Date.now()}`;
        localStorage.setItem(`user_key_${id}`, userKey);

        try {
            await castVote(id, optionID, userKey);
            setVoted(true);
            alert("Suara berhasil dikirim!");
            window.location.reload();
        } catch (err) {
            alert(err.response?.data?.error || "Gagal mengirim suara");
        }
    };

    if (!pollData) return <p>Loading...</p>;

    return (
        <div className="min-h-screen bg-white p-6">
            <div className="max-w-3xl mx-auto border rounded shadow-sm p-6">
                <h1 className="text-2xl font-bold">{pollData.poll.title}</h1>
                <p className="mb-4">{pollData.poll.description}</p>

                {/* Jika participant */}
                {isParticipant ? (
                    <>
                        <VoteForm options={pollData.options} onVote={handleVote} />
                        
                        {pollData.can_see_results && (
                            <PollResultChart results={results} />
                        )}
                    </>
                ) : (
                    /* Jika poll creator */
                    <>
                        <h2 className="mt-6 text-xl">Tools Admin</h2>
                        <label className="flex items-center gap-2 mb-4">
                            <input
                                type="checkbox"
                                defaultChecked={pollData.poll.allow_realtime}
                                onChange={(e) => toggleRealtimeVisibility(id, e.target.checked)}
                            />
                            Izinkan lihat hasil saat aktif
                        </label>

                        <button onClick={() => closePoll(id)}>Tutup Polling</button>
                        <button onClick={() => exportPollResults(id)}>Ekspor Hasil</button>

                        <PollResultChart results={results} />
                    </>
                )}
            </div>
        </div>
    );
};

export default PollDetailPage;