# Abstract Factory Design Pattern

See [Abstract Factory](https://en.wikipedia.org/wiki/Abstract_factory_pattern)

This is a creational design pattern that lets you create a family of related objects. 
It solves the problem of creating entire product families without specifying concrete types.

Use the Abstract Factory pattern when:
- A family of related product objects is designed to be used together and this constraint needs to be enforced.
- An application uses products from one family at a time
- There is a need to easily interchange product objects

The example in this folder is a foodfactory that can create a meat factory or a vegan factory. 

Both meat factories and vegan factories can create burgers as they both return data that satisfies the burger interface.

Vegan burgers and meat burgers both wrap the burger type.

Vegan factories know how to create burgers which are considered vegan.
Meat factories know how to create burgers which contain meat.

Working with this pattern, it would be easy to expand the code to include a vegetarian burger factory which makes vegetarian burgers. This would be done by adding a field to the burger type `containsDairy` and setting up a vegetarian factory which knows how to create burgers which contain dairy but do not contain meat.

It would also be possible to easily expand the code if a new product line of sausages was added, the meat and vegan factories would be able to support a `makeSausage` constructor.

