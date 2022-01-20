# Wordle Hacking

Analysis of the data in [https://www.powerlanguage.co.uk/wordle/](https://www.powerlanguage.co.uk/wordle/)

Extraction of word lists from that page into files `La.json` and `Ta.json` is left as an exercise to the reader.

## The Best Start Word: raise

There are two word lists of 5 letter words; one with 2315 words and one with 10657 words. Others have noted that the shorter list is probably words that could ever be a solution to wordle, and the longer list is things that will also be accepted as guesses. Linux /usr/share/dict/words contains only 6112 words that match `^[a-zA-Z]{5}$` .

Doing letter frequency analysis on the union of these lists leads to the impression that the best word (with the most frequent letters) is "[soare](https://www.thefreedictionary.com/soare)" (an obsolete word for a young hawk). Doing this just on the solution list yields "later" "alter" or "alert".

But we can do better. [wordlehack.go](wordlehack.go) does the following:

```
for each guess word:
  for each target word:
    for some {target} and {guess}, how many possible words remain?
  get the sum of remaning words for all targets
```

So, what word has on average the fewest remaining words after it has been guessed, across all possible targets? "raise"

This takes into account letter position, not just letter frequency. "arise" is not quite as good.

And if "raise" gets you nothing, the best next word is "clout".

So, there you go, go forth and be more reputable among your peers. raise your clout.