import React, { useState } from 'react';
import axios from 'axios';

const PollCreation = () => {
  const [question, setQuestion] = useState('');
  const [options, setOptions] = useState(['', '']);
  const [expiry, setExpiry] = useState('');

  const handleAddOption = () => {
    setOptions([...options, '']);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('/api/polls', { question, options, expiry });
      alert('Poll created successfully!');
      console.log(response.data);
    } catch (error) {
      console.error('Error creating poll:', error);
    }
  };

  return (
    <div className="p-8 bg-gray-100 min-h-screen">
      <form onSubmit={handleSubmit} className="max-w-md mx-auto bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-2xl font-bold mb-4">Create a New Poll</h2>
        <label className="block mb-4">
          <span className="text-gray-700">Question:</span>
          <input
            type="text"
            value={question}
            onChange={(e) => setQuestion(e.target.value)}
            required
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
          />
        </label>
        <label className="block mb-4">
          <span className="text-gray-700">Options:</span>
          {options.map((option, index) => (
            <input
              key={index}
              type="text"
              value={option}
              onChange={(e) => {
                const newOptions = [...options];
                newOptions[index] = e.target.value;
                setOptions(newOptions);
              }}
              placeholder={`Option ${index + 1}`}
              required
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            />
          ))}
          <button
            type="button"
            onClick={handleAddOption}
            className="mt-2 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Add Option
          </button>
        </label>
        <label className="block mb-4">
          <span className="text-gray-700">Expiry Time (hours):</span>
          <input
            type="number"
            value={expiry}
            onChange={(e) => setExpiry(e.target.value)}
            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
          />
        </label>
        <button
          type="submit"
          className="w-full inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Create Poll
        </button>
      </form>
    </div>
  );
};

export default PollCreation;