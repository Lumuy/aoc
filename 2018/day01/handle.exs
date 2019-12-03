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
    |> reaches_twice
    |> IO.inspect
  end

  defp reaches_twice(list), do: _reaches_twice(list, list, [0], [])

  defp _reaches_twice(_, _, _, result) when result != [] do
    List.first(result)
  end
  defp _reaches_twice(base, [], list, []) do
    _reaches_twice(base, base, list, [])
  end
  defp _reaches_twice(base, [h | t], list, []) do
    e = List.first(list) + h
    if e in list do
      _reaches_twice(base, t, list, [e])
    else
      _reaches_twice(base, t, [e | list], [])
    end
  end

end

Handle.sum
Handle.frequency
