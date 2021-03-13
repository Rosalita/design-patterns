# Builder Design Pattern

See [Builder Pattern](https://en.wikipedia.org/wiki/Builder_pattern)

This is a creational design pattern that separates the creation of an object from its representation.
It gives control over the construction process allowing creation of different variations of an object.
For example a vehicle builder could build cars and motorbikes, or a configuration builder could build json or yaml files.

There are a few different builder patterns

## Method Chaining - Common in Java, but not widely adopted in Go.
This pattern appears to be one of the oldest and has the downside that it's not great at coping with optional parameters.
This builder pattern calls a number of methods which slowly build up the object. 
Then a .build() method is called which returns the finished object.
It's not considered idiomatic Go, however for reference I have created an example in Go. 
The main reason for this is to be able to spot and identify the pattern when it is encountered.
There are a number of articles which advocate for using functional options instead of method chaining.

## Creation Parameters in a struct - The AWS Go SDK is full of this pattern.
This pattern uses a struct to hold all the creation parameters.
This struct is often referred to as a fooInput where it is used to create a foo.
It accomplishes the same task as method chaining in approx half as many lines of code as it doesn't rely on setter functions.
