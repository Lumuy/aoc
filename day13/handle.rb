require 'pry'

class Handle
  def initialize(file = 'data')
    @data, @types = {}, %w[< > ^ v]
    File.read(file).split("\n").each_with_index do |str, x|
      str.chars.each_with_index { |char, y| @data.merge!([y, x] => char) unless char == ' ' }
    end
  end

  # Part 1
  def location_of_first_crash(area = @data.dup, types = @types, origin = @data.dup)
    cars = area.select { |_, v| types.include?(v) }.to_a
    intersection_points  = area.select { |_, v| v == '+' }.keys
    intersection_counts  = Array.new(cars.size, 0)
    next_directions      = [:left, :straight, :right]
    points               = area.keys
    status               = true

    loop do
      temp_area = area.dup
      new_cars = []
      crash = []
      next_positions = step_position(cars, area.dup)

      0.upto(cars.size - 1).each do |index|
        next_position = next_positions[index]
        current_position, current_direction = cars[index]
        next_direction = nil

        state = area[next_position]
        state = origin[next_position] if types.include?(state)
        case state
        when '|'
          next_direction = current_direction if ['^', 'v'].include?(current_direction)
        when '-'
          next_direction = current_direction if ['<', '>'].include?(current_direction)
        when '/'
          next_direction = 'v' if current_direction == '<'
          next_direction = '>' if current_direction == '^'
          next_direction = '^' if current_direction == '>'
          next_direction = '<' if current_direction == 'v'
        when "\\"
          next_direction = 'v' if current_direction == '>'
          next_direction = '<' if current_direction == '^'
          next_direction = '>' if current_direction == 'v'
          next_direction = '^' if current_direction == '<'
        when '+'
          next_direction = next_intersection_direction(
            current_direction,
            next_directions[intersection_counts[index] % 3]
          )

          intersection_counts[index] += 1
        end

        if intersection_points.include?(current_position)
          temp_area[current_position] = '+'
        else
          x, y = current_position
          case current_direction
          when '^'
            if points.include?([x, y + 1])
              temp_area[current_position] = '|'
            else
              temp_area[current_position] = "/"  if next_direction == '^'
            end
          when 'v'
            if points.include?([x, y - 1])
              temp_area[current_position] = '|'
            else
              temp_area[current_position] = "\\" if next_direction == 'v'
            end
          when '<'
            if points.include?([x + 1, y])
              temp_area[current_position] = '-'
            else
              temp_area[current_position] = "\\" if next_direction == '<'
            end
          when '>'
            if points.include?([x - 1, y])
              temp_area[current_position] = '-'
            else
              temp_area[current_position] = "/"  if next_direction == '>'
            end
          end
        end
        temp_area[next_position] = next_direction

        new_cars << [next_position, next_direction]
      end

      cars = new_cars
      area = temp_area
      
      # crash
      cars.map(&:first).each_with_index do |e, index|
        crash << index if next_positions.count(e) != 1
      end
      if crash.size != 0
        ca = []
        cars.each_with_index { |e, index| ca << e unless crash.include? index }

        if status
          puts 'Part 1: ' + cars[crash[0]][0].join(',')
          status = false
        end

        crash.map { |n| cars[n][0] }.uniq.each do |point|
          area[point] = origin[point]
        end

        if ca.size == 1
          puts 'Part 2: ' + ca[0][0].join(',')
          break
        end

        cars = ca
      end

      # map
      if cars.include?([[39, 110], nil])
        p_area area
        binding.pry
      end

    end
  end

  def p_area(area)
    points = area.keys
    max_x = points.map(&:first).max
    max_y = points.map(&:last).max
    file  = 'map'

    (0..max_y).each do |y|
      line = ''
      (0..max_x).each do |x|
        if points.include? [x, y]
          line += area[[x, y]] || 'X'
        else
          line += ' '
        end
      end
      File.write(file, line + "\n", mode: 'a')
    end
  end

  def next_intersection_direction(direction, turn)
    case direction
    when '^'
      case turn
      when :left
        '<'
      when :straight
        '^'
      when :right
        '>'
      end
    when 'v'
      case turn
      when :left
        '>'
      when :straight
        'v'
      when :right
        '<'
      end
    when '<'
      case turn
      when :left
        'v'
      when :straight
        '<'
      when :right
        '^'
      end
    when '>'
      case turn
      when :left
        '^'
      when :straight
        '>'
      when :right
        'v'
      end
    end
  end

  def step_position(cars, map)
    result = []
    dh = {
      '^' => [0, -1],
      'v' => [0, 1],
      '<' => [-1, 0],
      '>' => [1, 0],
    }

    cars.each_with_index do |(point, direction), index|
      x, y = point
      case map[point]
      when "/"
        result << [x + 1, y] if direction == '^'
        result << [x - 1, y] if direction == 'v'
      when "\\"
        result << [x - 1, y] if direction == '^'
        result << [x + 1, y] if direction == 'v'
      else
        dx, dy = dh[direction]
        result << [x + dx, y + dy]
      end
    end
    
    result
  rescue TypeError
    binding.pry
  end
end

puts Handle.new.location_of_first_crash

