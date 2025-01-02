# HireEase : On-Campus Recruitment Management System (OCRMS)

### This prototype has been made as a closed-sourced application by the team-members of 'BroCode' as part of the 'Go for GoFr' Hackathon under 'Open Innovation' track. ###

[![Watch the video](https://github.com/joyprakashk/HireEase-team/Demo-Video-Thumbnail.png)](https://github.com/joyprakashk/HireEase-team/blob/0b02311646c10c966a187d74f0e50ece090c5e42/Demo_video.mp4)

## Overview
**HireEase** is a web application designed to streamline and manage on-campus recruitment processes. The application provides separate interfaces for students and recruiters, allowing them to interact with the system based on their roles. Key features include profile management, job postings, placement exam scheduling, and real-time notifications.

---

## Features
### **Student Interface**
  - **Profile**: View and update personal details.
  - **Dashboard**: Track application status and job opportunities.
  - **Placement Exams**: View and manage upcoming placement exams.
  - **Job Details**: Browse job postings and details.
  - **Notifications**: Receive updates on job openings, registration status and exams.
  - **Home Page**: Navigate back to the home page.

### **Recruiter Interface**
  - **Student Search**: Search for suitable candidates.
  - **Create Job Opening**: Post new job opportunities.
  - **Create Placement Exam**: Schedule exams for recruitment.
  - **Notifications**: Get updates on candidate responses and exams.
  - **Home Page**: Navigate back to the home page.

### **Common Features**
- **Authentication**: Separate sign-in/sign-up forms for students and recruiters.
  - Students: Use college email and password.
  - Recruiters: Use recruiter ID and password.
  - **Phone Authentication**: Firebase API used for phone number verification during the first login.
- **Responsive Design**: Side Menu-bars are used for good UI/UX.

---

## **Tech Stack**
- **Back-end**: Go
- **Microservices**: GoFr framework
- **Front-end**: HTML, CSS, Vanilla JavaScript, React.js (Running on Node.js)
- **AI/ML Models**: Python
- **Deployment**: zop.dev
- **Authentication**: Firebase API for phone verification

---

## **Setup Instructions**
### Prerequisites
1. Install Go (`>=1.20`).
2. Install Python (`>=3.8`) for AI/ML models.
3. Install Node.js and npm.
4. Firebase project for phone authentication.
5. Zopdev CLI for deployment.
6. Web browser (Chrome, Firefox are usually recommended) for front-end testing.

### Steps to Run Locally
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/joyprakashk/HireEase-team.git
   cd HireEase-team
2. **Set Up Backend**:
  Navigate to the backend folder and install dependencies:
   ```bash
   cd backend
   go mod tidy
3. **Run the Go server:**
   ```bash
   go run main.go
4. **Set Up Front-end:**
   Navigate to the frontend folder:
      ```bash
   cd ../frontend
5. **Setup NodeJS environment with no Build Configuration**
   _([npx](https://medium.com/@maybekatz/introducing-npx-an-npm-package-runner-55f7d4bd282b) comes with npm 5.2+ and higher, see [instructions for older npm versions](https://gist.github.com/gaearon/4064d3c23a77c74a3614c498a8bb1c5f))_

Then open [http://localhost:3000/](http://localhost:3000/) to see Sign-in page of the web-app running.

You **don’t** need to install or configure tools like webpack or Babel. They are preconfigured and hidden so that you can focus on the code.

**You’ll need to have Node 14.0.0 or later version on your local development machine** (but it’s not required on the server).

   
---

## Future Scope and Enhancements
# Our app is fully scalable and new features and functionalities will be added in subsequent releases.
- We plan to integrate real-time chat for student-recruiter communication.
- We also plan to use analytics for job applications and placement exam performance.
- We also plan to expand AI/ML models for candidate-job matching.

---

## Developers:
- Joyprakash Kalita [joyprakashk28@gmail.com](mailto:joyprakashk28@gmail.com)
- Satyam Gupta [issatyamgupta@gmail.com](mailto:issatyamgupta@gmail.com)
- Srijan Ratogi 
- P. Lohith [lohithcseng@gmail.com](mailto:lohithcseng@gmail.com)

