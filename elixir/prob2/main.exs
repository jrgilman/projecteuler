defmodule EulerTwo do
  def main() do
    even_fibonacci_sum(1, 2, 2, 4000000)
  end

  def even_fibonacci_sum(fib_one, fib_two, sum, end_value) do
    fib_three = fib_one + fib_two
    cond do
      fib_three > end_value ->
        sum
      true ->
        sum = if even?(fib_three) do
          sum + fib_three
        else
          sum
        end

        even_fibonacci_sum(fib_two, fib_three, sum, end_value)
    end
  end

  def even?(value) do
    cond do
      rem(value, 2) == 0 ->
        true
      true ->
        false
    end
  end
end
