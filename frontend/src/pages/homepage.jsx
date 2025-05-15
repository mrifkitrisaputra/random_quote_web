import React from 'react';

const HomePage = () => {
    return (
        <div className="min-h-screen bg-gray-50">
            <div className="container mx-auto px-4 py-12 text-center">
                <h1 className="text-4xl font-bold mb-4">Selamat Datang di QuickPoll</h1>
                <p className="text-lg text-gray-700 mb-8">
                    Buat polling instan, bagikan tautan, kumpulkan suara dalam hitungan detik.
                </p>

                <div className="flex flex-col md:flex-row gap-4 justify-center">
                    <a href="/create" className="bg-blue-600 hover:bg-blue-700 text-white py-3 px-6 rounded text-lg">
                        Buat Polling Baru
                    </a>
                    <a href="/dashboard" className="bg-green-600 hover:bg-green-700 text-white py-3 px-6 rounded text-lg">
                        Dashboard Saya
                    </a>
                </div>
            </div>
        </div>
    );
};

export default HomePage;