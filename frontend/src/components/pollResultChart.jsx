import React from 'react';

const PollResultChart = ({ results }) => {
    if (!results || results.length === 0) {
        return <p>Belum ada suara.</p>;
    }

    const totalVotes = results.reduce((sum, r) => sum + r.total_votes, 0);

    return (
        <div className="result-section mt-6">
            <h2 className="text-lg font-semibold mb-4">Hasil Polling</h2>
            <ul>
                {results.map((res, i) => (
                    <li key={i} className="mb-2">
                        <div className="flex justify-between mb-1">
                            <span>{res.option_text}</span>
                            <span>{res.total_votes} suara</span>
                        </div>
                        <div className="w-full bg-gray-200 h-2 rounded overflow-hidden">
                            <div
                                style={{
                                    width: `${(res.total_votes / totalVotes) * 100 || 0}%`,
                                    backgroundColor: '#3b82f6',
                                    height: '100%'
                                }}
                            ></div>
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default PollResultChart;