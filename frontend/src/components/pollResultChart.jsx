import React from 'react';

const PollResultChart = ({ results }) => {
    if (!results || results.length === 0) {
        return <p className="text-gray-500">Belum ada suara.</p>;
    }

    const totalVotes = results.reduce((acc, res) => acc + res.total_votes, 0);

    return (
        <div className="mt-6">
            <h3 className="font-semibold mb-2">Hasil Voting:</h3>
            <ul>
                {results.map((res, index) => (
                    <li key={index} className="mb-3">
                        <div className="flex justify-between mb-1">
                            <span>{res.option_text}</span>
                            <span>{res.total_votes} suara</span>
                        </div>
                        <div className="w-full bg-gray-200 h-2 rounded">
                            <div
                                style={{
                                    width: `${totalVotes > 0 ? (res.total_votes / totalVotes) * 100 : 0}%`,
                                    backgroundColor: '#3b82f6',
                                    height: '100%',
                                    borderRadius: '999px'
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