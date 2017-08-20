# modules

Available commands can be added by creating new modules.

## Commands

### whitespace merge

To merge visible code with whitespace code, pass in two parameters:

- file containing standard code
- file containing whitespace code

`go run steganosaurus.go whitespace merge test/figlet.js test/count.ws`

### zero-width encode

Encodes plaintext into binary using zero-width unicode codepoints.

`go run steganosaurus.go zero-width encode < LICENSE > ENCODED`

### zero-width decode

Decodes zero-width encoded text into plaintext.

`go run steganosaurus.go zero-width decode < ENCODED`
