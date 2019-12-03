class Handle
  attr_accessor :cars, :map, :directions, :intersections, :intersection_counts

  def initialize(file = 'data')
    @map, @directions = {}, %w[^ v > <]
    File.read(file).split("\n").each_with_index do |str, x|
      str.chars.each_with_index { |char, y| @map.merge!([y, x] => char) unless char == ' ' }
    end
    @cars = @map.select { |_, v| @directions.include?(v) }
    @intersections = @map.select { |_, v| v == '+' }.keys
    @intersection_counts = @cars.reduce({}) { |res, (p, _)| res.merge(p => 0) }

    @cars.each do |point, direction|
      @map[point] = '-' if ['>', '<'].include?(direction)
      @map[point] = '|' if ['^', 'v'].include?(direction)
    end
  end

  def run
    echo_first_crash = true

    while @cars.count > 1
      cars, skip_points = {}, []
      points = @cars.keys

      @cars.dup.each do |point, direction|
        next if skip_points.include?(point)

        new_point = car_next_point(point, direction)

        # crash
        if points.include?(new_point)
          # Part 1
          puts new_point.join(',') if echo_first_crash
          echo_first_crash = false

          skip_points << new_point

          @intersection_counts.delete(point)
          points.delete(point)
          points.delete(new_point)
          if cars.keys.include?(new_point)
            @intersection_counts.delete(new_point)
            cars.delete(new_point)
          end
        else
          new_direction = car_next_direction(point, new_point, direction)
          cars[new_point] = new_direction

          points.append(new_point)
        end

        points.delete(point)
      end

      @cars = sort_cars(cars)
    end

    # Part 2
    puts @cars.keys.join(',')
  end

  private

  # sort from top to bottom, then left to right
  def sort_cars(cars)
    cars.to_a
        .group_by { |e| e[0][1] }
        .reduce({}) { |res, (k, v)| res.merge(k => v.sort_by { |e| e[0][0] }) }
        .sort.to_h.values
        .reduce([]) { |res, e| res + e }
        .to_h
  end

  # next car point and direction
  def car_next_point(point, direction)
    x, y = point

    {
      '<' => [x - 1, y],
      '>' => [x + 1, y],
      '^' => [x, y - 1],
      'v' => [x, y + 1]
    }[direction]
  end

  def car_next_direction(p, np, direction)
    directions = %w[< ^ > v]
    count = @intersection_counts[p]
    state = @map[np]

    @intersection_counts.delete(p)
    @intersection_counts[np] = count if state != '+'
    @intersection_counts[np] = count + 1 if state == '+'

    case state
    when '-', '|'
      direction
    when "\\"
      @directions[-1 - @directions.index(direction)]
    when '/'
      @directions[@directions.index(direction) - 2]
    when '+'
      case count % 3
      # turn left
      when 0
        directions[directions.index(direction) - 1]
      # turn straignt
      when 1
        direction
      # turn right
      when 2
        directions[directions.index(direction) - 3]
      end
    end
  end
end

Handle.new.run
