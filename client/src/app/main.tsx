import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './App.tsx';
import './style/index.scss';

const root = document.getElementById('root');

if (!root) {
  throw new Error('root is not found');
}

createRoot(root).render(
  <StrictMode>
    <App />
  </StrictMode>,
);
