package tsak

import (
  "time"
  "math/rand"
)

var wisdom = []string{
  "Wear the tsak and go to pepelac!",
  "Please, wear a tsak, my dear...",
  "Fiddler is no use...",
  "Ku !",
  "Ku Ku !",
  "Ku Ku Ku ...",
  "Kju !",
  "Show me the man who thinks truth on Plyuk.",
  "Are you colorblind? Can't you even tell orange from green?",
  "Plyuk is a chatlanian planet, so we patsaks must wear tsaks and do ku.",
  "You on Earth have a very primitive society, you don't even make differences by the colour of one's pants!",
  "Yellow pants! Ku twice!",
  "Let's go, fiddler, to outer space.",
  "Comrade, there's a man there who says he's an alien. We have to do something!",
  "A-ah, see that. You've got the same raging racism as here, on Plyuk, only power's taken not by Chatlans, but by Patsaks; the ones like you and your friend Nightingale.",
  "How can I brake after you've drunken all the brake fluid?",
  "Either you're giving us this Ketse right now, or we won't lay you on Earth for less than 8 matchboxes.",
  "Strainger in the ku ....",
  "They took a women away and replace he with an automaton",
  "What is your planet number in tentura ? What's yours galaxy number in spiral ? Think !",
  "If there is no color differentiation of one's pants, there is no future!",
  "Time is relative ! Understand ?",
}

func RunVerification() string {
  rand.Seed(time.Now().Unix())
  n := rand.Int() % len(wisdom)
  return wisdom[n]
}
