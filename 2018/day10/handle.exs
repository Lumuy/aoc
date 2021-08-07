defmodule Message do
  def parse_inputs(file) do
    file
    |> File.stream!()
    |> Enum.map(fn line ->
      [[x], [y], [vx], [vy]] = Regex.scan(~r/(-?\d+)/, line, capture: :all_but_first)
      {{String.to_integer(x), String.to_integer(y)}, {String.to_integer(vx), String.to_integer(vy)}}
    end)
  end

  def step({{x, y}, {dx, dy}}), do: {{x + dx, y + dy}, {dx, dy}}

  def process(stars, seconds, true) do
    IO.inspect(seconds: seconds)

    stars
    |> Enum.map(fn {p, _} -> p end)
    |> print_msg()
  end
  def process(stars, seconds, false) do
    stars = Enum.map(stars, &step/1)
    process(stars, seconds + 1, is_msg(stars))
  end

  defp is_msg(stars) do
    points =
      stars
      |> Enum.map(fn {p, _} -> p end)
      |> Enum.sort(&(elem(&1, 0) < elem(&2, 0)))

    y_group_by_x = Enum.group_by(points, &elem(&1, 0), &elem(&1, 1))

    heights =
      points
      |> Enum.map(fn {x, _} -> x end)
      |> Enum.uniq()
      |> group_x()
      |> Enum.map(fn list ->
        list
        |> Enum.reduce([], fn x, acc ->
          acc ++ y_group_by_x[x]
        end)
      end)
      |> Enum.map(fn list ->
        Enum.max(list) - Enum.min(list)
      end)

    length(heights) > 1 && heights |> Enum.uniq() |> length() == 1
  end

  defp group_x(list) do
    list
    |> Enum.reduce([[]], fn x, acc ->
      acc
      |> Enum.at(0)
      |> Enum.at(0)
      |> case do
        nil ->
          [[x]]

        prx ->
          if x == prx + 1 do
            [p | n] = acc
            [[x | p] | n]
          else
            [[x] | acc]
          end
      end
    end)
  end

  defp print_msg(points) do
    dxs = Enum.map(points, fn {x, _} -> x end)
    dys = Enum.map(points, fn {_, y} -> y end)

    (Enum.min(dys)..Enum.max(dys))
    |> Enum.each(fn y ->
      (Enum.min(dxs)..Enum.max(dxs))
      |> Stream.map(fn x ->
        if Enum.any?(points, fn p -> p == {x, y} end) do
          "#"
        else
          "."
        end
      end)
      |> Enum.join("")
      |> IO.puts()
    end)

    IO.puts("\n")
  end
end

case System.argv() do
  [input_file] ->
    input_file
    |> Message.parse_inputs()
    |> Message.process(0, false)

  _ ->
    IO.puts(:stderr, "usage: elixir handle.exs filename")
end
