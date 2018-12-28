class Handle
  def initialize
    @serial_number = 7315
  end

  def power_level(x, y, serial_number = @serial_number)
    rack_id = x + 10
    power = (rack_id * y + serial_number) * rack_id
    hundred_digit = power.to_s.reverse[2].to_i
    hundred_digit - 5
  end

  def fuel_cells
    points = {}

    (1..300).each do |x|
      (1..300).each do |y|
        points.merge!({ "#{x},#{y}" => power_level(x, y) })
      end
    end

    points
  end

  # Part 1
  def right_coordinate
    cells, squares = fuel_cells, {}
    (1..298).each do |x|
      (3..300).each do |y|
        sum = 0
        (x..x + 2).each do |sx|
          (y - 2..y).each do |sy|
            sum += cells["#{sx},#{sy}"]
          end
        end

        squares.merge!({ "#{x},#{y - 2}" => sum })
      end
    end

    squares.key(squares.values.max)
  end

  # Part 2
  def sum_cells
    cells, hash = fuel_cells, {}

    (1..300).each do |x|
      sum = 0
      (1..300).each do |y|
        sum += cells["#{x},#{y}"]
        left = hash[[x - 1, y]] || 0
        hash[[x, y]] = sum + left
      end
    end

    hash
  end

  def identifier_of_square
    cells, res = sum_cells, {}

    (1..299).each do |x|
      (2..300).each do |y|
        max = [300 - x, y - 1].min
        (1..max).each do |i|
          sum = 0
          sum += cells[[x, y - i]] + cells[[x + i, y]]
          sum -= cells[[x, y]] + cells[[x + i, y - i]]
          res.merge!([x + 1, y - i + 1, i] => sum)
        end
      end
    end

    res.key(res.values.max).join(?,)
  end
end

# puts Handle.new.right_coordinate
puts Handle.new.identifier_of_square

