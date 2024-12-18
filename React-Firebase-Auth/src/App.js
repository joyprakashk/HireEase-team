import { Container, Row, Col } from "react-bootstrap";
import { Routes, Route } from "react-router-dom";
import "./App.css";
import Home from "./components/Home";
import Login from "./components/Login";
import Signup from "./components/Signup";
import ProtectedRoute from "./components/ProtectedRoute";
import { UserAuthContextProvider } from "./context/UserAuthContext";
import PhoneSignUp from "./components/PhoneSignUp";

import RecruiterSignIn from './components/RecruiterSignIn';
import RecruiterSignUp from './components/RecruiterSignUp';

function App() {
  return (
    <Container style={{ width: "400px" }}>
      <Row>
        <Col>
          <UserAuthContextProvider>
            <Routes>
              <Route
                path="/home"
                element={
                  <ProtectedRoute>
                    <Home />
                  </ProtectedRoute>
                }
              />
              <Route path="/" element={<Login />} />
              <Route path="/phonesignup" element={<PhoneSignUp />} />
              <Route path="/signup" element={<Signup />} />

              <Route path="/recruiter-signin" element={<RecruiterSignIn />} />
              <Route path="/recruiter-signup" element={<RecruiterSignUp />} />
            </Routes>
          </UserAuthContextProvider>
        </Col>
      </Row>
    </Container>
  );
}

export default App;
