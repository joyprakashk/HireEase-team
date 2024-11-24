import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { signInWithEmailAndPassword } from 'firebase/auth';
import { auth } from '../firebase'; // Assuming firebase.js is in the root directory

const RecruiterSignIn = () => {
  const [recruiterId, setRecruiterId] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleSignIn = async (e) => {
    e.preventDefault();
    try {
      await signInWithEmailAndPassword(auth, recruiterId, password);
      alert('Sign-in Successful');
      navigate('/dashboard'); // Redirect to dashboard
    } catch (error) {
      alert('Error: ' + error.message);
    }
  };

  return (
    <div className="container mt-5">
      <h2>Recruiter Sign-In</h2>
      <form onSubmit={handleSignIn}>
        <div className="mb-3">
          <label className="form-label">Recruiter ID</label>
          <input
            type="email"
            className="form-control"
            value={recruiterId}
            onChange={(e) => setRecruiterId(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label">Password</label>
          <input
            type="password"
            className="form-control"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">Sign In</button>
      </form>
    </div>
  );
};

export default RecruiterSignIn;
