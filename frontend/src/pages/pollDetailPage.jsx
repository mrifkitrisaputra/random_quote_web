import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getPollDetails, exportPollResults } from '../service/api';
import PollResultChart from '../components/pollResultChart';

const PollDetailPage = () => {
    const { id } = useParams();
    const [pollData, setPollData] = useState(null);
    const [canSeeResults, setCanSeeResults] = useState(false);
    const [results, setResults] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data = await getPollDetails(id);
                setPollData(data.poll);
                setCanSeeResults(data.can_see_results);

                if (data.can_see_results) {
                    const resultData = await exportPollResults(id);
                    setResults(resultData);
                }
            } catch (err) {
                alert("Polling tidak ditemukan");
                window.location.href = "/";
            }
        };
        fetchData();
    }, [id]);

    if (!pollData) return <p>Loading...</p>;

    return (
        <div className="min-h-screen bg-white p-6">
            <div className="max-w-3xl mx-auto">
                <h1 className="text-2xl font-bold mb-4">{pollData.poll.title}</h1>
                <p className="mb-6">{pollData.poll.description}</p>

                {/* Voting Form */}
                <div className="mb-8">
                    <h2 className="text-lg font-medium mb-4">Pilih Jawaban:</h2>
                 {pollData.options.map((opt) => (
                  <div key={opt.id} className="mb-2">
                    <label className="flex items-center gap-2">
                      <input type="radio" name="vote" value={opt.id} />
                      {opt.text}
                    </label>
                  </div>
                ))}
                    <button
                        type="submit"
                        className="mt-4 bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
                    >
                        Kirim Suara
                    </button>
                </div>

                {/* Hasil Polling */}
                {canSeeResults && <PollResultChart results={results} />}
            </div>
        </div>
    );
};

export default PollDetailPage;