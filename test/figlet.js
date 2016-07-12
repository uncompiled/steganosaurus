const figlet = require('figlet');
const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.on('line', (input) => {
  if (input == 'quit') {
    rl.close()
  } else {
    figlet(input,
      function(err, data) {
        if (err) {
          console.log(err);
          return;
        }
        console.log(data)
      }
    );
  }
});