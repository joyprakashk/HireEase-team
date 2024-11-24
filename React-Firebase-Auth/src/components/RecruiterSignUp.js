import React, { useState } from 'react';
import { createUserWithEmailAndPassword } from 'firebase/auth';
import { auth } from '../firebase';

const RecruiterSignUp = () => {
  const [recruiterId, setRecruiterId] = useState('');
  const [password, setPassword] = useState('');

  const handleSignUp = async (e) => {
    e.preventDefault();
    try {
      await createUserWithEmailAndPassword(auth, recruiterId, password);
      alert('Sign-up Successful');
    } catch (error) {
      alert('Error: ' + error.message);
    }
  };

  return (
    <div className="container mt-5">
      <h2>Recruiter Sign-Up</h2>
      <form onSubmit={handleSignUp}>
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
        <button type="submit" className="btn btn-success">Sign Up</button>
      </form>
    </div>
  );
};

export default RecruiterSignUp;
