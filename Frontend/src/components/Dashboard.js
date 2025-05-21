import React, { useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import axios from 'axios';

const Dashboard = () => {
  const { user, logout } = useContext(AuthContext);

  const handleAdminAction = async () => {
    try {
      const res = await axios.get('http://localhost:8080/api/admin/dashboard');
      alert(res.data.message);
    } catch (error) {
      alert('Access denied');
    }
  };

  return (
    <div>
      <h2>Welcome, {user.username}</h2>
      <button onClick={handleAdminAction}>Admin Action</button>
      <button onClick={logout}>Logout</button>
    </div>
  );
};

export default Dashboard;