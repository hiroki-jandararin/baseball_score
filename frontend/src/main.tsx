import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';
import App from './App.tsx';
import MatchResultPage from './pages/MatchResultPage';
import NewsPreviewPage from './pages/NewsPreviewPage';

function resolvePage(pathname: string) {
  if (pathname === '/entry') return <MatchResultPage />;
  if (pathname === '/news') return <NewsPreviewPage />;
  return <App />;
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    {resolvePage(window.location.pathname)}
  </StrictMode>,
);
