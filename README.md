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

Total: 4/50