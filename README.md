# Advent of Go 2023

Solutions for the 2023 edition, written in Go.

Maybe this being my very first Golang experience wasn't
the best idea...

Tasks:
- Day 1
  - [x] A
  - [x] B
  - Comments
    - Nice first challenge as an introduction to Go, the biggest hurdle was figuring out how mutability works
    - I did come up with a very easy regex solution, but I found more value in doing the parsing manually
    - I'm quite proud that this solution only reads each character once!
- Day 2
  - [x] A
  - [x] B
  - Comments
    - Nice refresh to string parsing, I did skip the Token step, but it didn't seem necessary
    - Only a few more lines were needed for part B, once parsing was figured
    - After the fact I noticed that separating the different sets isn't necessary, as we only care about the minimal set
      for each game (Part A comparing it to the provided one, Part B calculating its power)
    - Even though it was overkill, it sure was fun
- Day 3
  - [x] A
  - [x] B
  - Comments
    - With the regained confidence in my parsing from day 2, the parsing for this one was a breeze
    - For part A, I had first thought of collecting all the symbols and numbers and do matching with two big for loops
      (spoiler alert, I did that for part B)
    - I wondered if there was a way to reuse the parsing loop, and there was, for every single number, touches can only
      happen with the current and previous line so we only care about those two lines
    - This does have one caveat, a number could be counted multiple times if it touches multiple symbols, but apparently
      this doesn't happen in the input provided 🤷
    - For part B, I think a solution similar to part A's, with three lines instead of two would have worked, but tbh I
      was too lazy and wondered if the performance gain would be worth it.

Total: 6/50