import pandas as pd
import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

# Load student data
FILE_PATH = "selected.xlsx"  # Path to your Excel file
students = pd.read_excel(FILE_PATH)

# SMTP Configuration (Update these with your credentials)
SMTP_SERVER = 'smtp.gmail.com'
SMTP_PORT = 587
SMTP_EMAIL = 'godgenoside@gmail.com'  # Replace with your email
SMTP_PASSWORD = 'Nmit@123'        # Replace with your email's password

# Function to send email
def send_email(to_email, subject, message):
    try:
        msg = MIMEMultipart()
        msg['From'] = SMTP_EMAIL
        msg['To'] = to_email
        msg['Subject'] = subject
        msg.attach(MIMEText(message, 'plain'))
        
        with smtplib.SMTP(SMTP_SERVER, SMTP_PORT) as server:
            server.starttls()
            server.login(SMTP_EMAIL, SMTP_PASSWORD)
            server.send_message(msg)
        print(f"Email sent to {to_email} successfully.")
    except Exception as e:
        print(f"Failed to send email to {to_email}. Error: {e}")

# Function to approve a student
def approve_student(student_name):
    student = students[students['Name'] == student_name]
    if not student.empty:
        email = student.iloc[0]['mail']
        send_email(
            email,
            "Application Approved",
            f"Dear {student_name},\n\nCongratulations! Your application has been approved."
        )
        print(f"Student {student_name} has been approved.")
    else:
        print(f"Student {student_name} not found.")

# Function to reject a student
def reject_student(student_name):
    student = students[students['Name'] == student_name]
    if not student.empty:
        email = student.iloc[0]['mail']
        send_email(
            email,
            "Application Rejected",
            f"Dear {student_name},\n\nWe regret to inform you that your application has been rejected."
        )
        print(f"Student {student_name} has been rejected.")
    else:
        print(f"Student {student_name} not found.")

# Example Usage
if __name__ == "__main__":
    print("Student Management System")
    print("1. Approve Student")
    print("2. Reject Student")
    print("3. Exit")
    
    while True:
        choice = input("Enter your choice: ")
        if choice == '1':
            name = input("Enter the student's name to approve: ")
            approve_student(name)
        elif choice == '2':
            name = input("Enter the student's name to reject: ")
            reject_student(name)
        elif choice == '3':
            print("Exiting...")
            break
        else:
            print("Invalid choice. Please try again.")
