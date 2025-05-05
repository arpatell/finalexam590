defmodule Animal do
  # base processes for abstract animal, with a common encapsulated name field
  def start(name) do
    spawn_link(fn -> loop(name) end) # use basic spawn_link
  end
  defp loop(name) do
    receive do
      {:get_name, caller} ->
        send(caller, name)
        loop(name)
    end
  end
  def get_name(animal) do # encapsulation
    send(animal, {:get_name, self()})
    receive do
      name -> name
    end
  end
end

defmodule Dog do
  def start(name) do
    dog = Animal.start(name) # reuse animal process (inheritance equivalent in elixir)
    speak = fn -> n = Animal.get_name(dog) # call get_name to retrieve encapsulated name
      IO.puts("#{n} says: Woof!") # module specific implementation for speak (abstraction + polymorphism)
    end
    spawn_link(fn ->
      receive do
        {:speak, caller} -> speak.(); send(caller, :ok)
      end
    end)
  end
end

defmodule Cat do
  def start(name) do
    cat = Animal.start(name)
    speak = fn -> n = Animal.get_name(cat)
      IO.puts("#{n} says: Meow!")
    end
    spawn_link(fn ->
      receive do
        {:speak, caller} -> speak.(); send(caller, :ok)
      end
    end)
  end
end

defmodule Main do
  def run do
    a1 = Dog.start("Rex")
    a2 = Cat.start("Whiskers")
    send(a1, {:speak, self()}) # polymorphism
    receive do _ -> :ok end
    send(a2, {:speak, self()})
    receive do _ -> :ok end
  end
end