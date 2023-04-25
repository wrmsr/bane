//
/*
(package clj)

(func add1 [x int] int
  (return (+ x 1)))

(type bar struct
  [i int])

(func [b bar] add1 [] int
  (return (+ (.i b) 1)))

(func mapThing () (map int string)
  (return ((map int string) 1 "one" 2 "two"

==

func mapThing() map[int]string {
  return
*/
package clj
