# COMP 590-059 Final Exam - Aaron Patel

## One: Elixir OO Problem

In this Elixir solution, the implementation of the 4 pillars of OOP is attempted in the context of the actor model and higher-order functions.

1. **Encapsulation**

   In the specifc module implementations for Dog and Cat, we access the name in a way that promotes encapsulation by calling the `get_name()` method in `Animal`. This is achieved because the internal states of elixir processes are private and can only be accessed through message passing (`:get_name` in `Animal`). It is accessed like: `Animal.get_name(a1)`, where a1 is a specific instance of an animal started up (`a1 = Dog.start("Rex")`).
   - Java mapping: In the Java class, `private String name` in the Animal class is only accessible through the `getName()` method. Instead of directly accessing the name through a start function or other similar methods, name is accessed through `Animal.get_name()` to prevent leaks, promoting strong encapsulation.

2. **Abstraction**

   From outside functions like `main()`, the processes that power the atoms `:speak` and `:get_name` for `Animal` behavior don't need to be understood. Callers can simply use `{:speak, self()}`, `{:get_name, self()}`, or the `get_name()` to interact with Animal behavior without needing to know the internal mechanics.
   - Java mapping: The methods in the Java class (`speak()`, `getName()`) are called similarly in Elixir as: `send(a1, {:speak, self()})` or `send(a1, {:get_name, self()})` or `Animal.get_name(a1)`

3. **Inheritance**

   In this solution, code reuse is leveraged to define the `Animal` module with a base process (internal loop that initializes behavior through `start()`) and a common name getter, which are accessed in specific animal modules to initialize behavior. However, this is only function reuse and does not include the subclassing and extension that inheritance involves. Thus, this solution does not have inheritance.
   - Java Mapping: These functions are reused in specific Dog and Cat modules (`Animal.start()` and `Animal.get_name()`) to be able to access the encapsulated name field and define the speak function. This appears similar to inheritance, though not functionally equivalent.

4. **Polymorphism**

   The call to `:speak` for both animals happens in the same structure, but the actual output depends on what type of animal it is. This is because the injected function in Dog and Cat each define a module specific implementation for the speak function. This indicates polymorphic behavior, similar to how Java has subclass specifc implementations. 
   - Java mapping: This is equivalent to calling `a1.speak()` in Java, but instead we do `send(a1, {:speak, self()})`, which returns the respective message based on what type of animal `a1` is.

## Running the Program
1. **Navigation**

   ```bash
    cd one/animal
   ```
2. **Start the Elixir environment and generate application**

   MacOS:
   ```bash
    iex -S mix
   ```
   Windows:
   ```bash
    iex.bat -S mix
   ```
3. **Run the program in the Elixir shell**

   ```elixir
    Main.run()
   ```
