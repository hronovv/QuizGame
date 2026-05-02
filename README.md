# QuizGame

A CLI quiz: questions and answers are loaded from a CSV file; you type an answer for each prompt. A single **timer** applies to the whole run (default: 30 seconds).

## Requirements

- Go 1.26+ (see `go.mod`)

## Build

```bash
make
```

or

```bash
go build -o quizgame ./cmd
```

The built binary is `quizgame` (listed in `.gitignore`).

## Run

```bash
./quizgame
```

### Flags

| Flag     | Default                 | Description |
|----------|-------------------------|-------------|
| `-csv`   | `problems/problems.csv` | Path to a CSV with `question,answer` columns |
| `-limit` | `30`                    | Time limit for the whole quiz, in **seconds** |

Example:

```bash
./quizgame -csv=problems/problems.csv -limit=60
```

**Paths:** relative paths in `-csv` are resolved from the **current working directory** (where you run the command)

## CSV format

First column: question. Second column: expected answer. Example: `problems/problems.csv`.

## Repository layout

```
.
├── cmd/
│   └── main.go          # entrypoint
├── internal/
│   ├── csv/             # map CSV rows to problems
│   └── problem/         # problem type
├── problems/
│   └── problems.csv     # default dataset
├── Makefile
├── go.mod
└── README.md
```
