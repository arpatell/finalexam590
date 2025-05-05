defmodule AnimalTest do
  use ExUnit.Case
  doctest Animal

  test "greets the world" do
    assert Animal.hello() == :world
  end
end
