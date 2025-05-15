import React, { useState } from 'react';
import CreatePollForm from '../components/createPollForm';

const CreatePollPage = () => {
    const [pollID, setPollID] = useState(null);
    const [pollURL, setPollURL] = useState('');

    const handleCreate = (id, url) => {
        setPollID(id);
        setPollURL(url);
    };

    return (
        <div className="min-h-screen bg-white p-6">
            <div className="max-w-3xl mx-auto">
                <h1 className="text-2xl font-bold mb-6">Buat Polling Baru</h1>
                <CreatePollForm onCreate={handleCreate} />

                {pollID && (
                    <div className="mt-6 bg-blue-50 border-l-4 border-blue-500 p-4 rounded">
                        <p><strong>Poll ID:</strong> {pollID}</p>
                        <p><strong>Tautan polling:</strong> <a href={pollURL} className="text-blue-600 underline">{pollURL}</a></p>
                    </div>
                )}
            </div>
        </div>
    );
};

export default CreatePollPage;