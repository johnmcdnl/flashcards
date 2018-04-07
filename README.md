# Flashcards

[![Go Report Card](https://goreportcard.com/badge/github.com/johnmcdnl/flashcards)](https://goreportcard.com/report/github.com/johnmcdnl/flashcards)

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/0ff3da47e33146029b5b5fba2bc21510)](https://www.codacy.com/app/johnmcdnl/flashcards?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=johnmcdnl/flashcards&amp;utm_campaign=Badge_Grade)

Simulate a set of flashcards to help learn a new language

Automatically adjusts cards returns to factor in what you've previously learned i.e. focus on what you don't know rather than on what you've already learned

For now just terminal based so something a little bit like

``` text
10
Translate: кукуруза  kukuruza
    1) birthday
    2) baby
    3) corn
    4) snow
Type your answer: corn
Correct!


11
Translate: пол  pol
    1) floor
    2) ring
    3) party
    4) flower
Type your answer: floor
Correct!


12
Translate: земля  zemlya
    1) ground
    2) hand
    3) letter
    4) doll
Type your answer: hand
ERRO[0037] Incorrect!!   земля  zemlya    ground
```

Updates to come that

* add ability to render actual pictures as flashcards
* expand phrase set and translations
* auto translations
* groups of nouns/verbs etc.
* adding better cmd args
