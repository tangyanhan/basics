# Design principles

## Single Responsibility Principle

## Liskov Substitution Principle
Let q(x) be a property provable about objects x of type T. Then q(y) should
be true for objects y of type S where S is a subtype of T.
## Dependency Inversion Principle
Depend upon abstractions, do not depend upon concrete classes.
## Coding to an Interface
## Least Knowledge Principle
Only talk to your immediate friends.

## Open-Close Principle
Software entities like classes, modules and functions should be open for extension  but closed for modifications.

# Singleton Pattern

# Template Pattern

Template pattern create template methods and called them in super classes, and the these templates are implemented in
 sub-classes. In this way, a super class is capable to make calls to methods implemented in sub-classes.
 
What details this design pattern hide:

1. It encapsulates the relevant calls of actual template implementations.
2. It opens for extension and closes for changes of algorithm

For example, we can encapsulate a sorting algorithm into a template class, with compare() function exposed for 
customization.

# Builder Pattern

Separate the construction of a complex object from its representation so that the same construction process can 
create different representations.

