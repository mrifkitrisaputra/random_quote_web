import React from 'react';

const NotFoundPage = () => {
    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-50">
            <div className="text-center">
                <h1 className="text-5xl font-bold text-red-600 mb-2">404</h1>
                <p className="text-xl mb-6">Halaman tidak ditemukan</p>
                <a href="/" className="text-blue-600 underline">← Kembali ke Beranda</a>
            </div>
        </div>
    );
};

export default NotFoundPage;