# Guess the Number

**Guess the Number** is a simple command-line game where the player tries to guess a randomly generated number within a specific range. This game is perfect for testing your luck and logic skills! 

## Features

- Randomly generated number for every new game.
- Configurable range for the number to guess.
- Feedback on whether the guess is too high or too low.
- Victory message when the correct number is guessed.

## Prerequisites
- Go 1.24.1

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/earik87/guess-number.git
   cd guess-number
   ```

2. Make sure you have [Go](https://golang.org/) installed on your system. 

## Usage

Run the game using:
```bash
go run .
```

### How to Play

1. The game will randomly generate a number within a predefined range (e.g., 1 to 100).
2. You will be prompted to guess the number.
3. After each guess, you'll receive feedback:
   - `Too high!` if your guess is greater than the number.
   - `Too low!` if your guess is less than the number.
4. Keep guessing until you find the correct number, and celebrate your win!

### Example Session

```bash
➜  guess-number git:(main) go run guess-number game

Welcome to the Number Guessing Game!
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice (1-3): 1

Great! You have selected the Easy difficulty level!

I'm thinking of a number between 1 and 100. You have 10 chances.

Chances left: 10
Enter your guess: 10
Too low! Try a higher number.

Chances left: 9
Enter your guess: 50
Too low! Try a higher number.

Chances left: 8
Enter your guess: 70
Too low! Try a higher number.

Chances left: 7
Enter your guess: 85
Too low! Try a higher number.

Chances left: 6
Enter your guess: 92
Too low! Try a higher number.

Chances left: 5
Enter your guess: 95
Too low! Try a higher number.

Chances left: 4
Enter your guess: 99
Too low! Try a higher number.

Chances left: 3
Enter your guess: 100

Congratulations! You've guessed the number in 8 attempts!
```

## Configuration

You can configure the range of the random number by modifying the source code or providing parameters (if implemented).

## Contributing

Feel free to fork this repository, make improvements, and submit a pull request! Your contributions are welcome.

---

Let me know if you'd like to add custom instructions or features! Note that the idea of the game is from https://roadmap.sh/projects/number-guessing-game.
