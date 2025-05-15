import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './pages/homepage';
import CreatePollPage from './pages/createPollPage';
import DashboardPage from './pages/dashboardPage';
import PollDetailPage from './pages/pollDetailPage';
import NotFoundPage from './pages/notFoundPage';

const App = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/create" element={<CreatePollPage />} />
                <Route path="/dashboard" element={<DashboardPage />} />
                <Route path="/polls/:id" element={<PollDetailPage />} />
                <Route path="/api/polls/:id/vote" element={<PollDetailPage />} />
                <Route path="*" element={<NotFoundPage />} />
            </Routes>
        </Router>
    );
};

export default App;