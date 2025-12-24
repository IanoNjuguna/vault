import type { NextConfig } from "next";
import type { Configuration } from "webpack";

const nextConfig: NextConfig = {
  reactStrictMode: true,
  
  // 1. Resolve the "Does not provide an export" error
  transpilePackages: ['@fabianbormann/cardano-peer-connect'],

  // 2. Properly typed Webpack function
  webpack: (config: Configuration) => {
    config.experiments = {
      ...config.experiments,
      asyncWebAssembly: true,
      layers: true,
    };

    return config;
  },
};

export default nextConfig;
