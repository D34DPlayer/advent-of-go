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
      happen with the current and previous line, so we only care about those two lines
    - This does have one caveat, a number could be counted multiple times if it touches multiple symbols, but apparently
      this doesn't happen in the input provided 🤷
    - For part B, I think a solution similar to part A's, with three lines instead of two would have worked, but tbh I
      was too lazy and wondered if the performance gain would be worth it.
- Day 4
  - [x] A
  - [x] B
  - Comments
    - Parsing was supposed to be easy, and then I hit a huge wall? It seems my parsing would skip the last number if it was
      one char wide. That's where parsing prerequisites come to the rescue and now the int parsing function assumes that
      the first character has to be alphanumeric and won't check the "ok" for it.
    - Part B scared me at first but then realized that the only information we need is how many matches each card had
- Day 5
  - [x] A
  - [x] B
  - Comments
    - Part A was a breeze now that we are master parsers, and it was just a matter of finding the intersection between one
      element and a range
    - Part B was harder than expected, but I'm not sure if it is the fault of going for the more complicated solution
    - A simple solution would be to expand the ranges from the inputs and use the same algorithm as part A, but I thought
      it wouldn't be optimized enough
    - My solution does range intersections with 5 different cases:
      - A recipe for that range has already been applied or the recipe doesn't apply => same range
      - The range fits entirely in the recipe => apply full range
      - There's a partial intersection => apply on a sub range, can result on up to 3 new ranges
- Day 6
  - [x] A
  - [x] B
  - Comments
    - More of a math challenge than a programming one (if we ignore the parsing). Thx Wolfram alpha for the help to
      visualize this problem.
    - Distance eq: $(T-x)*x$ <=> $-x^2+Tx$
    - Beat the record eq: $-x^2+Tx-R$
    - Solutions: ${T±\sqrt{T^2-4R}}\over 2$
    - $Rad = \sqrt{B^2-4AC}$
    - $Amount = ceil(\frac{T+Rad}{2})-floor(\frac{T-Rad}{2})-1$
- Day 7
  - [x] A
  - [x] B
  - Comments
    - The difficulty is going up, isn't it. The hardest part was to calculate the hands. There's probably a smartest
      algorithm but that's all I could come up with, and it works
    - I'm not proud but for part B too much debugging went into a small issue. My code didn't count JJJJJ as five of a kind 
- Day 8
  - [x] A
  - [x] B
  - Comments
    - Part A was quite simple, a regular node traversal problem
    - Part B had me stomped for quite a bit, because as far as I know the general problem requires bruteforce, but after
      leaving the program run for 15mins I realized this wouldn't be possible
    - Luckily for us, the input isn't completely random and after inspecting a few of the starting node possible solutions
      it became clear that the distance between possible solutions was constant
    - This simplifies the problem by a lot because we just need to find the smallest number that's a multiple of those constants

Total: 16/50