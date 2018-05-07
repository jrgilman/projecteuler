defmodule EulerThree do

  def main() do
    find_largest_prime_factor(600851475143, 2, 1)
  end

  def factor?(num, divisor) do
    cond do
      rem(num, divisor) == 0 ->
        true
      true ->
        false
    end
  end

  def find_largest_prime_factor(num_to_factor, current_divisor, current_largest) do
    cond do
      num_to_factor == 1 ->
        current_largest
      factor?(num_to_factor, current_divisor) and current_divisor > current_largest ->
        find_largest_prime_factor(
          trunc(num_to_factor / current_divisor),
          2,
          current_divisor
        )
      factor?(num_to_factor, current_divisor) ->
        find_largest_prime_factor(
          trunc(num_to_factor / current_divisor),
          2,
          current_largest
        )
      true ->
        find_largest_prime_factor(num_to_factor, current_divisor + 1, current_largest)
    end
  end

end
