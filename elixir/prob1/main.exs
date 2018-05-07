defmodule EulerOne do
  def main() do
    sum_of_multiples(1, 1000, 0)
  end

  def multiple?(value) do
    cond do
      rem(value, 3) == 0 ->
        true
      rem(value, 5) == 0 ->
        true
      true ->
        false
    end
  end

  def sum_of_multiples(current_value, end_value, sum) do
    sum = if multiple?(current_value) do
      sum + current_value
    else
      sum
    end

    if current_value < end_value do
      sum_of_multiples(current_value + 1, end_value, sum)
    else
      sum
    end
  end
end
