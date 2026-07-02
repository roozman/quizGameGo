# Game Application

# use-case

## User use-case
### Register
users can register using phone number

### Login
users can login using phone number and password

## Game use-case
### Each game has a given number of questions
### Three levels of difficulty "easy, medium, hard"
### Winner is determined by count of the correct answers
### Each game belongs to a specific category

# entity

## User
- ID
- Phone Number
- Avatar
- Name

## Game
- ID
- Category
- Questions List
- Players

## Questions
- ID
- Questions
- Answers List
- Correct Answer
- Difficulty
- Category