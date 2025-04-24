import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

const PollVote = () => {
  const { id } = useParams();
  const [poll, setPoll] = useState(null);
  const [selectedOption, setSelectedOption] = useState('');

  useEffect(() => {
    const fetchPoll = async () => {
      try {
        const response = await axios.get(`/api/polls/${id}`);
        setPoll(response.data);
      } catch (error) {
        console.error('Error fetching poll:', error);
      }
    };
    fetchPoll();
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post(`/api/polls/${id}/vote`, { option: selectedOption });
      alert('Vote submitted successfully!');
    } catch (error) {
      console.error('Error submitting vote:', error);
    }
  };

  if (!poll) return <p className="text-center mt-8">Loading...</p>;

  return (
    <div className="p-8 bg-gray-100 min-h-screen">
      <form onSubmit={handleSubmit} className="max-w-md mx-auto bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-2xl font-bold mb-4">{poll.question}</h2>
        {poll.options.map((option, index) => (
          <div key={index} className="mb-4">
            <label className="inline-flex items-center">
              <input
                type="radio"
                name="option"
                value={option}
                checked={selectedOption === option}
                onChange={(e) => setSelectedOption(e.target.value)}
                className="form-radio h-5 w-5 text-indigo-600"
              />
              <span className="ml-2 text-gray-700">{option}</span>
            </label>
          </div>
        ))}
        <button
          type="submit"
          className="w-full inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Submit Vote
        </button>
      </form>
    </div>
  );
};

export default PollVote;