# Wordle Guessser

## How it works

Pretty straightforward: Run the program, put in what word you guessed, then put in the hint pattern

The hint pattern is a 5 letter string of the 3 states:

* **a:** absent (white/grey). The letter is not in the word
* **p:** present (yellow). The letter is in the word, but in a different place
* **c:** correct (green). The letter is in the word and in the correct location

For example:

```
Guess: roate
Pattern: ppaaa
```

## Setup

### Compile the program

[install Go](https://go.dev/) and run ```go build``` in the src folder

### Calculate first guess (optional)

You can re-calculate the best first guess if you want. This may be desirable if you want to make modifications to ```answers.txt``` (i.e. excluding past answers).
Run ```./wordle -rough```, this will produce 2 JSON files, one with letter values and one with sorted rough guess values  
Next, run ```./wordle -detailed -start 0 -end 10```. This will take guesses start-end from the output of ```-rough``` and run them through the detailed scoring method.  
Note that this step can take a long time. Running it on the entire list of guesses can take many hours.  
Also note that the LOWER scores from this output are the best ones. They represent the total number of remaining answers after generating a guess on every possible answer.

You may also run ```./wordle -score adieu``` to score an individual guess and print the result to command line (replace adieu with whatever guess you want to score)
