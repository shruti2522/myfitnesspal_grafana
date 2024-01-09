/** @type {import('next').NextConfig} */
const nextConfig = {
    publicRuntimeConfig : {
        // other public runtime configurations
        NEXT_PUBLIC_BACKEND_URL: process.env.NEXT_APP_BACKEND_URL,
      },
};



module.exports = nextConfig;
