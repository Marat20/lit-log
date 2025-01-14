import react from '@vitejs/plugin-react';
import path from 'path';
import { defineConfig } from 'vite';

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const isDev = mode === 'development';
  const API = isDev ? 'http://localhost:8080' : 'https://litlog.shop';

  return {
    plugins: [react()],
    resolve: {
      alias: [{ find: '@', replacement: path.resolve(__dirname, 'src') }],
    },
    define: {
      __API__: JSON.stringify(API),
      __IS_DEV__: JSON.stringify(isDev),
    },
  };
});
