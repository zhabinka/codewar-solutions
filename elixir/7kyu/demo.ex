defmodule Demo do
  def factorial(0), do: 1
  def factorial(num) when num > 0, do: num * factorial(num - 1)
end

Demo.factorial(5) |> io.inspect
Demo.factorial(-5) |> IO.inspect
