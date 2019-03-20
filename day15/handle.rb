require 'pry'

class Handle
  def initialize(file = 'sample_data')
    @data = {}
    File.read(file).split("\n").map(&:chars).each.with_index do |arr, y|
      arr.each_with_index { |e, x| @data[[x, y]] = e }
    end
  end

  # Part 1
  def the_outcome_of_combat(area = @data)
    hps = area.select { |_, v| ['G', 'E'].include?(v) }.keys
              .reduce({}) { |r, k| r.merge!(k => 200) }
    binding.pry
  end

  def round_step(area, hps)
    goblins = area.select { |_, v| v == 'G' }.keys
    elves   = area.select { |_, v| v == 'E' }.keys
    walls   = area.select { |_, v| v == '#' }.keys
    cavern  = area.select { |_, v| v == '.' }.keys
    temp    = area

    reading_order(goblins + elves).each do |point|
      enemies   = elves.include?(point) ? goblins : evles
      reachable = reachable_points(enemies, area, point)
      nearest   = nearest_points(reachable, point)
      chosen    = reading_order(nearest).first

      # move
      x1, y1 = point
      x2, y2 = chosen

      goal = [x1 + 1, y1] if x2 > x1 && y2 >= y1
      goal = [x1 - 1, y1] if x2 < x1 && y2 >= y1
      goal = [x1, y1 - 1] if y2 < y1

      # replace
      type = elves.include?(point) ? 'E' : 'G'
      temp[chosen] = type
      temp[point]  = '.'

      # attach
    end
  end

  def reachable_points(points, area, spoint)
    sx, sy = spoint
    adjacents = points.reduce([]) do |r, (x , y)|
      r << [x - 1, y]
      r << [x + 1, y]
      r << [x, y - 1]
      r << [x, y + 1]
    end.select { |point| area[point] == '.' }
    adjacents.select do |dx, dy|
      x_min, x_max = [sx, dx].min, [sx, dx].max
      y_min, y_max = [sy, dy].min, [sy, dy].max
      top = (x_min..x_max).map { |bx| [bx, y_min] } +
            (y_min..y_max).map { |by| [x_max, by] } -
            [sx, sy] - [dx, dy]
      bom = (x_min..x_max).map { |bx| [bx, y_max] } +
            (y_min..y_max).map { |by| [x_min, by] } -
            [sx, sy] - [dx, dy]
      top.all? { |e| area[e] == '.' } ||
      bom.all? { |e| area[e] == '.' }
    end
  end

  def nearest_points(points, spoint)
    hash = points.reduce({}) do |r, dpoint|
      r.merge!(dpoint => point_distance(spoint, dpoint))
    end
    min_distance = hash.values.min
    # attack
    hash.select { |_, v| v == min_distance }.keys
  end

  def point_distance(spoint, dpoint)
    sx, sy = spoint
    dx, dy = dpoint
    (sx - dx).abs + (sy - dy).abs
  end

  # top-to-bottom then left-to-right
  def reading_order(points)
    points.sort_by { |_, y| y }
          .sort_by { |x, _| x }
  end
end

Handle.new.the_outcome_of_combat
