const questions = [
    {
        question: "What is the primary objective of a business?",
        options: ["Make profit", "Expand globally", "Provide customer service", "Create jobs"],
        correct: 0
    },
    {
        question: "Which of the following is considered an asset in business?",
        options: ["Cash", "Debt", "Rent", "Expenses"],
        correct: 0
    },
    
    {
        question: "What does ROI stand for in business?",
        options: ["Rate of Interest", "Return on Investment", "Reinvestment Opportunity", "Revenue Over Income"],
        correct: 1
    },
    {
        question: "Which business structure allows for limited liability?",
        options: ["Sole Proprietorship", "Partnership", "Corporation", "Freelance"],
        correct: 2
    },
    {
        question: "What is the term for a business's ongoing expenses?",
        options: ["Capital", "Liabilities", "Assets", "Operating Costs"],
        correct: 3
    }
];

let currentQuestionIndex = 0;
let score = 0;

function loadQuestion() {
    const question = questions[currentQuestionIndex];
    document.getElementById('question').textContent = question.question;
    
    const optionsContainer = document.getElementById('options');
    optionsContainer.innerHTML = '';  // This is previous options 
    
    question.options.forEach((option, index) => {
        const optionElement = document.createElement('div');
        optionElement.classList.add('option');
        optionElement.textContent = option;
        optionElement.onclick = () => checkAnswer(index);
        optionsContainer.appendChild(optionElement);
    });
}

function checkAnswer(selectedIndex) {
    const correctIndex = questions[currentQuestionIndex].correct;
    if (selectedIndex === correctIndex) {
        score++;
    }
    document.getElementById('next-btn').disabled = false;  // yeh next button enable karega
}

function nextQuestion() {
    currentQuestionIndex++;
    if (currentQuestionIndex < questions.length) {
        loadQuestion();
        document.getElementById('next-btn').disabled = true;
    } else {
        showResult();
    }
}

function showResult() {
    document.getElementById('quiz-container').style.display = 'none';
    document.getElementById('result-container').style.display = 'block';
    document.getElementById('score').textContent = `${score} / ${questions.length}`;
}

function restartQuiz() {
    currentQuestionIndex = 0;
    score = 0;
    document.getElementById('quiz-container').style.display = 'block';
    document.getElementById('result-container').style.display = 'none';
    loadQuestion();
    document.getElementById('next-btn').disabled = true;
}

window.onload = loadQuestion;
