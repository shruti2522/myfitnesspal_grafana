"use client"

import React from 'react';
import FileUpload from '../components/FileUpload';

const Home = () => {
  const handleUpload = (response: any) => {
    console.log(response);
  };

  return (
    <div>
      <h1>MyFitnessPal csv to Grafana</h1>
      <FileUpload onUpload={handleUpload} />
    </div>
  );
};

export default Home;

