/*
Package catfacts implements Echo middleware that to add cat facts in headers of HTTP responses.
Example:
  package main

  import (
    "net/http"

    "github.com/labstack/echo/v4"
    catfacts "github.com/regentag/cat-facts"
  )

  func main() {
    e := echo.New()

    e.Use(catfacts.CatFactsMiddleware)

    e.GET("/", func(c echo.Context) error {
      return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
  }
*/
package catfacts

import (
	"math/rand"

	"github.com/labstack/echo/v4"
)

// getCatFactsConst holds cat facts text data.
func getCatFactsConst() []string {
	return []string{
		// from http://maxellah.tripod.com/catfacts.htm
		"Cats can't taste sweets.",
		"A cat's tongue consists of small \"hooks,\" which come in handy when tearing up food.",
		"Americans spend more annually on cat food than on baby food.",
		"In 1987 cats overtook dogs as the number one pet in America.",
		"A group of youngsters (kittens) is called a kindle; those old-timers (adult cats) form a clowder.",
		"Black cat superstitions are as American as apple pie. In Asia and England, black cats are considered lucky.",
		"Cats have five toes on each front paw, but only four toes on each back paw.",
		"Cats have true fur, in that they have both an undercoat and an outer coat.",
		"When a domestic cat goes after mice, about one pounce in three results in a catch.",
		"The cheetah is the only cat in the world that can't retract it's claws.",
		"A cat has 32 muscles in each ear.",
		"Neutering a cat extends it's life span by two or three years.",
		"Cats must have fat in their diet because they can't produce it on their own.",
		"Cat's urine glows under a black light.",
		"The heaviest cat ever recorded was 46 lbs.",
		"Cats have a third eyelid called a haw and you will probably only see it when kitty isn't feeling well.",
		"Adult cats with no health problems are in deep sleep 15 percent of their lives. They are in light sleep 50 percent of the time.",
		"Cats are the only animal that walk on their claws, not the pads of their feet.",
		"Most cats have no eyelashes.",
		"A cat cannot see directly under its nose. This is why the cat cannot seem to find tidbits on the floor.",
		"It is a common belief that cats are color blind. However, recent studies have shown that cats can see blue, green and red.",
		"Cats with white fur and skin on their ears are very prone to sunburn.",
		"Abraham Lincoln loved cats. He had four of them while he lived in the White House.",
		"Napoleon was terrified of cats.",
		"Mother cats teach their kittens to use the litter box.",
		"The way you treat kittens in the early stages of it's life will render it's personality traits later in life.",
		"Tylenol and chocolate are both poisionous to cats.",
		"The average cat food meal is the equivalent to about five mice.",
		"Cats have AB blood groups just like people.",
		"A form of AIDS exists in cats.",
		"The ancestor of all domestic cats is the African Wild Cat which still exists today.",
		"In ancient Egypt, killing a cat was a crime punishable by death.",
		"In the Middle Ages, during the Festival of Saint John, cats were burned alive in town squares.",
		"Today there are about 100 distinct breeds of the domestic cat.",
		"A cat has four rows of whiskers on each side.",
		"A cat's jaws cannot move sideways.",
		"Cats have over one hundred vocal sounds, while dogs only have about ten.",
		"A cat can jump even seven times as high as it is tall.",
		"A cat is pregnant for about 58-65 days.",
		"In ancient Egypt, entire families would shave their eyebrows as a sign of mourning when the family cat died.",
	}
}

// CatFactsMiddleware is a middleware that to add cat facts
// in headers of HTTP responses.
func CatFactsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	facts := getCatFactsConst()
	count := len(facts)

	return func(c echo.Context) error {
		fact := facts[rand.Intn(count)]
		c.Response().Header().Set("X-Cat-Facts", fact)
		return next(c)
	}
}
