const startButton = document.getElementById('start-btn');
const welcomePage = document.getElementById('quiz-fp');
const finalScore = document.getElementById('quiz-lp');
const quizQuestions = document.getElementById('qst-cont');
let q1 = document.getElementById('qsts');
let btn1 = document.getElementById('a');
let btn2 = document.getElementById('b');
let btn3 = document.getElementById('c');
let btn4 = document.getElementById('d');
const nextButton = document.getElementById('nxt-btn');
const allOption = document.getElementById('opt-cont').children;
let scoreRec = document.getElementsByClassName('result');

startButton.addEventListener('click', startQuiz);
nextButton.addEventListener('click', nextQuestion);

// Tracks the index of the 'question' property in the 'questions' array of objects
let questionIndex = 0;
let score = 0;


// For start button to start quiz
function startQuiz() {
    console.log("start Quiz!");
    welcomePage.classList.add('hide');
    quizQuestions.classList.remove('hide');
    nextQuestion();
}

// For next button to go to next quiz
function nextQuestion() {
    // shows the result page after the fifth question
    if (questionIndex == 5) {
        quizQuestions.classList.add('hide');
        finalScore.classList.remove('hide');
    }
    showQuestions();
    reset();
    questionIndex++;
}

// Shows the questions
function showQuestions() {
    let q = questions[questionIndex];
    q1.innerHTML = q.question;
    btn1.innerHTML = q.options[0];
    btn2.innerHTML = q.options[1];
    btn3.innerHTML = q.options[2];
    btn4.innerHTML = q.options[3];

}

// To verify the chosen option
function checkAnswer(option) {
    // To make the questionIndex start from 0 again
    let q = questions[questionIndex - 1];
    if (option.id == q.answer) {
        option.classList.add('correct');
        score++;
        scoreKeeper();
    } else {
        option.classList.add('wrong');
    }
    optionInvalid()
}

//To mmake other options unclickable
function optionInvalid() {
    let q = questions[questionIndex - 1];
    for (let i = 0; i < allOption.length; i++) {
        allOption[i].classList.add('invalid');
        if (allOption[i].id == q.answer) {
            allOption[i].classList.add('correct');
        }
    }
}

// resets the class of the buttons
function reset() {
    for (let i = 0; i < allOption.length; i++) {
        allOption[i].classList.remove('correct', 'wrong', 'invalid');
    }
}

// Keeps record of the questions gotten correctly
function scoreKeeper() {
    // iterates the modification of scores on question page and 
    // final result page
    for (let i = 0; i < scoreRec.length; i++) {
        scoreRec[i].innerHTML = `Your Score: ${score}`;
    }
}

const questions = [
    {
        question: "Which of these snakes could hypothetically be a developer?",
        options: [
            "Boa", "Viper", "Python", "Mamba"
        ],
        answer: 'c'
    },
    {
        question: "She found the language as precious as the stone.",
        options: [
            "Diamond", "Emerald", "Sapphire", "Ruby"
        ],
        answer: 'd'
    },
    {
        question: "A certain programmer in Paris hated a prominent structure but preferred working in it.",
        options: [
            "Eiffel", "Notre-Dame Cathedral", "Louvre Museum", "Arc de Triomphe"
        ],
        answer: 'a'
    },
    {
        question: "When this language is used, the PC is as a volcano.",
        options: [
            "Magma", "Cinder", "Lava", "Rocks"
        ],
        answer: 'c'
    },
    {
        question: "You are the language.",
        options: [
            "You", "Me", "Self", "U"
        ],
        answer: 'c'
    }

]