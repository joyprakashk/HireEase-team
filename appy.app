import streamlit as st

# Title
st.title("Aptitude Test")

# Questions and answers
questions = [
    {
        "question": "1. What is 5 + 3?",
        "options": ["6", "7", "8", "9"],
        "answer": "8"
    },
    {
        "question": "2. Which is the largest planet in our solar system?",
        "options": ["Earth", "Mars", "Jupiter", "Venus"],
        "answer": "Jupiter"
    },
    {
        "question": "3. What is the capital of France?",
        "options": ["Berlin", "Paris", "Madrid", "Lisbon"],
        "answer": "Paris"
    },
    {
        "question": "4. What is 12 * 2?",
        "options": ["22", "24", "26", "28"],
        "answer": "24"
    },
    {
        "question": "5. Which is the smallest prime number?",
        "options": ["0", "1", "2", "3"],
        "answer": "2"
    }
]

# Initialize session state
if "current_question" not in st.session_state:
    st.session_state.current_question = 0
    st.session_state.score = 0
    st.session_state.answers = []

# Function to handle the "Next" button click
def next_question(selected_option):
    current_index = st.session_state.current_question
    if selected_option == questions[current_index]["answer"]:
        st.session_state.score += 1
    st.session_state.answers.append(selected_option)
    st.session_state.current_question += 1

# Display the current question
current_index = st.session_state.current_question
if current_index < len(questions):
    current_question = questions[current_index]
    st.write(current_question["question"])
    selected_option = st.radio(
        "Options", current_question["options"], key=f"q{current_index}"
    )

    if st.button("Next"):
        if selected_option:  # Ensure an option is selected
            next_question(selected_option)
        else:
            st.warning("Please select an option before proceeding.")
else:
    # Display the final score
    st.success(f"Test completed! Your total score is: {st.session_state.score}/{len(questions)}")
    st.write("Your answers:")
    for i, answer in enumerate(st.session_state.answers):
        st.write(f"Q{i + 1}: {answer} (Correct: {questions[i]['answer']})")
