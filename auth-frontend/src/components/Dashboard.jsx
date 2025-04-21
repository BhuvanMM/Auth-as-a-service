import React, { useEffect, useState } from 'react';
import { getSession } from '../services/api';

function Dashboard() {
  const [session, setSession] = useState(null);

  useEffect(() => {
    getSession()
      .then(res => setSession(res.data))
      .catch(() => alert('Session fetch failed'));
  }, []);

  if (!session) return <p>Loading...</p>;

  return (
    <div>
      <h2>Dashboard</h2>
      <p>Welcome, {session.email}</p>
      <p>Session ID: {session.sessionId}</p>
    </div>
  );
}

export default Dashboard;
