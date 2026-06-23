What you are building

Go Reloaded is a text-editing tool written in Go. It reads an input file, applies a series of transformations, and writes the result to an output file.

The program is driven by markers embedded in the text, such as (hex), (bin), (up), (low), (cap), and optional counts like (up, 2).
Core transformations

(hex) replaces the word before it with the decimal form of a hexadecimal number. Example: `1E (hex)` becomes `30`.

(bin) replaces the word before it with the decimal form of a binary number. Example: `10 (bin)` becomes `2`.

(up), (low), and (cap) change the casing of the word before them, or of the previous N words when a count is provided.

Punctuation spacing, apostrophe pairs, and article correction (`a` to `an`) are also part of the final output.
What to practice here

Break the project into small pure functions first: number conversion, casing helpers, article fixes, and punctuation cleanup.

Write unit tests for each helper before wiring them together with Go's file system APIs.