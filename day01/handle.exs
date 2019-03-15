defmodule Handle do
  def get_inputs(file \\ "inputs") do
    File.read(file)
    |> elem(1)
    |> String.trim
    |> String.split("\n")
    |> Enum.map(&String.to_integer/1)
  end

  # Part 1
  def sum do
    get_inputs()
    |> Enum.reduce(0, &+/2)
    |> IO.puts
  end

  # Part 2
  def frequency do
    get_inputs()
    |> sum_list([0])
    |> Enum.reverse
    |> reaches_twice([], :false)
    |> IO.inspect
  end

  defp sum_list([], result), do: result
  defp sum_list([head | tail], result) do
    sum_list(tail, [List.first(result) + head | result])
  end

  defp reaches_twice([n | _], _, :true), do: n
  defp reaches_twice([h1 | t1], [], :false), do: reaches_twice(t1, [h1], :false)
  defp reaches_twice([h1 | t1], [h2 | t2], :false) do
    if h1 in [h2 | t2] do
      reaches_twice([h1 | t1], [h2 | t2], :true)
    else
      reaches_twice(t1, [h1] ++ [h2 | t2], :false)
    end
  end
  defp reaches_twice(_, _, _), do: raise("error")

end

Handle.sum
Handle.frequency
