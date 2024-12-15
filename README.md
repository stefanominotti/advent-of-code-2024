# Advent of Code 2024

**Go** solutions for the [Advent of Code 2024](https://adventofcode.com/2024).

## Setup instructions

To run the solutions, you'll need to set up an environment variable with your Advent of Code session token.

### Step 1: Get your session token
1. Log in to the [Advent of Code](https://adventofcode.com) website.
2. Open your browser's developer tools (usually accessible via `F12` or `Ctrl+Shift+I` / `Cmd+Option+I`).
3. Navigate to the **Application** tab (or **Storage** tab in some browsers).
4. Look for your cookies under `https://adventofcode.com`.
5. Find the `session` cookie, and copy its value.

### Step 2: Set the environment variable
Set the session token as an environment variable named `AOC_SESSION`. For example:
- On Linux/macOS:
  ```bash
  export AOC_SESSION=your-session-token
  ```
- On Windows (PowerShell):
  ```powershell
  $env:AOC_SESSION="your-session-token"
  ```

## Running the solutions

Once the environment variable is set, you can run the code in two ways:

### 1. Run all solutions
To run all solutions for the available days:
```bash
go run main.go -all
```

### 2. Run a specific solution
To run the solution for a specific day (replace `X` with the day number, e.g., `1` for Day 1):
```bash
go run main.go -solution=X
```

## Solutions summary

Here is a summary of the solutions so far, including their approximate execution times and links to the respective Advent of Code pages:

| Day | Link                                           | Part A   | Part B    |
|----:|------------------------------------------------|---------:|----------:|
| 1   | [Day 1](https://adventofcode.com/2024/day/1)   | 1.567ms  | 0.396ms   |
| 2   | [Day 2](https://adventofcode.com/2024/day/2)   | 1.288ms  | 1.013ms   |
| 3   | [Day 3](https://adventofcode.com/2024/day/3)   | 1.404ms  | 1.092ms   |
| 4   | [Day 4](https://adventofcode.com/2024/day/4)   | 3.252ms  | 0.640ms   |
| 5   | [Day 5](https://adventofcode.com/2024/day/5)   | 3.162ms  | 31.514ms  |
| 6   | [Day 6](https://adventofcode.com/2024/day/6)   | 0.655ms  | 891.980ms |
| 7   | [Day 7](https://adventofcode.com/2024/day/7)   | 17.889ms | 592.452ms |
| 8   | [Day 8](https://adventofcode.com/2024/day/8)   | 2.086ms  | 0.969ms   |
| 9   | [Day 9](https://adventofcode.com/2024/day/9)   | 1.226ms  | 114.803ms |
| 10  | [Day 10](https://adventofcode.com/2024/day/10) | 1.465ms  | 0.383ms   |
| 11  | [Day 11](https://adventofcode.com/2024/day/11) | 0.724ms  | 7.651ms   |
| 12  | [Day 12](https://adventofcode.com/2024/day/12) | 1.136ms  | 1.608ms   |
| 13  | [Day 13](https://adventofcode.com/2024/day/13) | 0.775ms  | 0.317ms   |
| 14  | [Day 14](https://adventofcode.com/2024/day/14) | 0.658ms  | 63.671ms  
