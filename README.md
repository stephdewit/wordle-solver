# Wordle Solver

A fast Wordle assistant written in Go. Available as an interactive CLI,
an HTTP API, and a web UI you can open directly in your browser.

Given the letters you've tried and their results, it narrows down the
remaining candidates and suggests the best probes to uncover unknown letters.

## How it works

After each guess, encode the result character by character:

| Symbol | Meaning |
| ------ | ------- |
| `!` | Correct letter, correct position (green) |
| `:` | Correct letter, wrong position (yellow) |
| `-` | Letter not in the word (grey) |

For example, guessing `crane` where `r` is misplaced and `e` is correct:
`--:-!`

The solver returns two lists:

- **Candidates** — words that could still be the answer
- **Probes** — words (not necessarily candidates) that best reveal unknown
  letters based on their frequency in the remaining candidates

## Project structure

```text
wordle-solver/
├── cli/        # Interactive command-line interface
├── api/        # HTTP JSON API
├── ui/         # Vue.js web interface
└── solver/     # Core logic (filtering, scoring, probes)
```

## Requirements

- Go 1.25+
- A word list (default: `/usr/share/dict/words`)

## Building

```sh
make          # build both binaries
make test     # run tests with coverage
make install  # install CLI to $PREFIX/bin (default: /usr/bin)
```

## CLI

```sh
./wordle-solver
```

Each round displays candidates and probes, then prompts for your guess
and its result:

```text
Found 4591 5-letter words
Candidates: [saner stare snare laser lares ...]
Probes:     [alien arise stale crane trace ...]
  Word: crane
Result: --:-!
```

The word list path can be overridden:

```sh
WORDLIST=/usr/share/dict/american-english ./wordle-solver
```

## API

```sh
./wordle-solver-api          # listen on :8080
./wordle-solver-api -cors    # enable CORS (required for browser clients)
PORT=3000 ./wordle-solver-api -cors
```

### `POST /suggestions`

#### Request

```json
{
  "g": [
    { "w": "crane", "r": "--:-!" },
    { "w": "stole", "r": "-!-!-" }
  ]
}
```

#### Response

```json
{
  "s": ["stone", "atone", "ozone"],
  "p": ["sloth", "droit", "pilot"],
  "t": 12
}
```

| Field | Description |
| ----- | ----------- |
| `s` | Top 10 candidate words |
| `p` | Top 10 probe words |
| `t` | Total number of remaining candidates |

## Web UI

Open `ui/index.html` in a browser while the API is running with `-cors`:

```sh
./wordle-solver-api -cors
open ui/index.html
```

Click any suggestion or probe to auto-fill it as your next guess. Click the
letter tiles to cycle through grey / yellow / green before submitting.
