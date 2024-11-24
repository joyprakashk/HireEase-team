// Import the required functions from the SDKs
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

// TODO: Add SDKs for Firebase products that is to be used
// https://firebase.google.com/docs/web/setup#available-libraries

// Web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyAGVCFjGgpSvEkpXwCYDcLajZUqDMrz5O0",
  authDomain: "phone-auth-70e5d.firebaseapp.com",
  projectId: "phone-auth-70e5d",
  storageBucket: "phone-auth-70e5d.firebasestorage.app",
  messagingSenderId: "938158760788",
  appId: "1:938158760788:web:a238aba503e0bf709a5ec1"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export default app;
