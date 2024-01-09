

import React, { useState, ChangeEvent } from 'react';
import axios from 'axios';

interface FileUploadProps {
  onUpload: (response: any) => void;
}

const FileUpload: React.FC<FileUploadProps> = ({ onUpload }) => {
  const [file, setFile] = useState<File | null>(null);

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      setFile(event.target.files[0]);
    }
  };

  const handleUpload = async () => {
    console.log("hello entered handle upload")
    if (!file) {
      console.error('No file selected');
      return;
    }
    console.log('Uploading file:', file);
    console.log('Backend URL:', process.env.NEXT_PUBLIC_BACKEND_URL);


    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await axios.post(process.env.NEXT_PUBLIC_BACKEND_URL+'/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      onUpload(response.data);
    } catch (error) {
      console.error('Error uploading file:', error);
    }
  };

  return (
    <div>
      <input type="file" accept=".csv" onChange={handleFileChange} className="mb-4 p-2 border border-gray-300 rounded"/>
      <button onClick={handleUpload} 
      className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Upload</button>
    </div>
  );
};

export default FileUpload;
