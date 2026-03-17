import { useEffect, useState } from 'react';

type HealthResponse = {
  status: string;
  message: string;
};

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

function App() {
  const [message, setMessage] = useState<string>('Loading...');
  const [error, setError] = useState<string>('');

  useEffect(() => {
    const fetchHealth = async (): Promise<void> => {
      try {
        const response = await fetch(`${apiBaseUrl}/health`);
        if (!response.ok) {
          throw new Error(`HTTP ${response.status}`);
        }

        const data: HealthResponse = await response.json();
        setMessage(data.message);
      } catch (err) {
        if (err instanceof Error) {
          setError(err.message);
          return;
        }
        setError('Unknown error');
      }
    };

    void fetchHealth();
  }, []);

  return (
    <main style={{ padding: '24px', fontFamily: 'sans-serif' }}>
      <h1>Baseball Score App</h1>
      <p>Frontend is running with React + TypeScript.</p>
      <p>API Base URL: {apiBaseUrl}</p>
      {error ? <p>Error: {error}</p> : <p>Backend response: {message}</p>}
    </main>
  );
}

export default App;
