# Prototype Pattern

See [Prototype Pattern](https://en.wikipedia.org/wiki/Prototype_pattern)

Prototype pattern is a creational pattern where new objects are cloned from a prototypical instance.

The base object will have a clone method that generates a new instance of the base object without any coupling to the base object.

Use this pattern when:
- creating a new object in the standard way, e.g. `NewThing()` is prohibitively expensive. e.g. the thing contains tree shaped data, like a file structure.
