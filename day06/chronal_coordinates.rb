class ChronalCoordinates
  def initialize(file = 'data')
    @data = File.read(file).split("\n").map do |e|
      e.split(',').map(&:to_i)
    end
  end

  def min_distance(point)
    distances = @data.map { |e| (point[0] - e[0]).abs + (point[1] - e[1]).abs }
    min_distance = distances.min
    distances.map.with_index { |d, i| i if d == min_distance }.compact
  end

  # Part 1
  def largest_area_size(data = @data)
    xs, ys = data.map(&:first), data.map(&:last)
    x_min, x_max = xs.min, xs.max
    y_min, y_max = ys.min, ys.max
    infinite, points = [], []

    data.each_with_index do |(x, y), index|
      infinitely = false
      [[x_min, y], [x_max, y], [x, y_min], [x, y_max]].each { |p|
        if min_distance(p) == [index]
          infinitely = true
          break
        end
      }
      infinite << index if infinitely
    end

    (x_min..x_max).each do |x|
      (y_min..y_max).each do |y|
        md = min_distance([x, y])
        points += md if md.count == 1
      end
    end

    points.reject { |i| infinite.include?(i) }.group_by(&:itself).map { |_, points| points.count }.max
  end

  # Part 2
  def dist_sum(p)
    @data.reduce(0) { |r, e| r + (e[0] - p[0]).abs + (e[1] - p[1]).abs }
  end

  def region_size(data = @data)
    xs, ys = data.map(&:first), data.map(&:last)
    x_min, x_max = xs.min, xs.max
    y_min, y_max = ys.min, ys.max
    size = 0

    (x_min..x_max).each do |x|
      (y_min..y_max).each do |y|
        size += 1 if dist_sum([x, y]) < 10000
      end
    end
    size
  end
end

puts ChronalCoordinates.new.largest_area_size
puts ChronalCoordinates.new.region_size
