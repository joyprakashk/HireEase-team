import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Form, Alert } from "react-bootstrap";
import { Button } from "react-bootstrap";
import { useUserAuth } from "../context/UserAuthContext";

const Login = () => {
  // States for Student Login
  const [studentEmail, setStudentEmail] = useState("");
  const [studentPassword, setStudentPassword] = useState("");
  const [studentError, setStudentError] = useState("");
  const { logIn, googleSignIn } = useUserAuth();
  const navigate = useNavigate();

  // States for Recruiter Login
  const [recruiterEmail, setRecruiterEmail] = useState("");
  const [recruiterPassword, setRecruiterPassword] = useState("");
  const [recruiterError, setRecruiterError] = useState("");

  const handleStudentSubmit = async (e) => {
    e.preventDefault();
    setStudentError("");
    try {
      await logIn(studentEmail, studentPassword);
      navigate("/home");
    } catch (err) {
      setStudentError(err.message);
    }
  };

  const handleRecruiterSubmit = async (e) => {
    e.preventDefault();
    setRecruiterError("");
    try {
      await logIn(recruiterEmail, recruiterPassword);
      navigate("/recruiter-dashboard");
    } catch (err) {
      setRecruiterError(err.message);
    }
  };

  return (
    <>
      <div className="d-flex justify-content-center align-items-center container mt-5">
        {/* Student Login Card */}
        <div className="p-4 box mx-2" style={{ width: "350px" }}>
          <h2 className="mb-3">Student Login</h2>
          {studentError && <Alert variant="danger">{studentError}</Alert>}
          <Form onSubmit={handleStudentSubmit}>
            <Form.Group className="mb-3" controlId="formBasicEmail">
              <Form.Control
                type="email"
                placeholder="College Email address"
                onChange={(e) => setStudentEmail(e.target.value)}
              />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicPassword">
              <Form.Control
                type="password"
                placeholder="Password"
                onChange={(e) => setStudentPassword(e.target.value)}
              />
            </Form.Group>

            <div className="d-grid gap-2">
              <Button variant="primary" type="Submit">
                Log In
              </Button>
            </div>
          </Form>
          <hr />
          <Link to="/phonesignup">
            <div className="d-grid gap-2 mt-3">
              <Button variant="success" type="button">
                Sign in with Phone
              </Button>
            </div>
          </Link>
          <div className="p-4 box mt-3 text-center">
            Don't have an account? <Link to="/signup">Sign up</Link>
          </div>
        </div>

        {/* Recruiter Login Card */}
        <div className="p-4 box mx-2" style={{ width: "350px" }}>
          <h2 className="mb-3">Recruiter Login</h2>
          {recruiterError && <Alert variant="danger">{recruiterError}</Alert>}
          <Form onSubmit={handleRecruiterSubmit}>
            <Form.Group className="mb-3" controlId="formBasicEmail">
              <Form.Control
                type="email"
                placeholder="Recruiter Id"
                onChange={(e) => setRecruiterEmail(e.target.value)}
              />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicPassword">
              <Form.Control
                type="password"
                placeholder="Password"
                onChange={(e) => setRecruiterPassword(e.target.value)}
              />
            </Form.Group>

            <div className="d-grid gap-2">
              <Button variant="primary" type="Submit">
                Log In
              </Button>
            </div>
          </Form>
          <hr />
          <Link to="/phonesignup">
            <div className="d-grid gap-2 mt-3">
              <Button variant="success" type="button">
                Sign in with Phone
              </Button>
            </div>
          </Link>
          <div className="p-4 box mt-3 text-center">
            Don't have an account? <Link to="/recruiter-signup">Sign up</Link>
          </div>
        </div>
      </div>
    </>
  );
};

export default Login;
