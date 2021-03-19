# Factory Method Design Pattern

See [Factory Method Pattern](https://en.wikipedia.org/wiki/Factory_method_pattern)

The Factory Method design pattern calls a function which creates an object.
The object can be returned, a pointer to the object can be returned
or an interface that the object satisfies can be returned. 

There's a mantra for Go which is "Accept interfaces return structs".

It's also possible to have a factory which makes factories, this is called a factory generator.