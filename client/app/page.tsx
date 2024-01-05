"use client"

import React from 'react';
import FileUpload from '../components/FileUpload';

const Home: React.FC = () => {
  const handleUpload = (response: any) => {
    console.log(response);
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded shadow-md max-w-md w-full">
        <h1 className="text-3xl font-semibold mb-6">MyFitnessPal CSV to Grafana</h1>
        <FileUpload onUpload={handleUpload} />
      </div>
    </div>
  );
};

export default Home;
