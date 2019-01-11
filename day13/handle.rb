require 'pry'

class Handle
  def initialize(file = 'sample_data')
    @data, @types = {}, %w[< > ^ v]
    File.read(file).split("\n").each_with_index do |str, x|
      str.chars.each_with_index { |char, y| @data.merge!([x, y] => char) unless char == ' ' }
    end
  end

  # Part 1
  def location_of_first_crash(data = @data, types = @types)
    intersections = 
    while true do
      cars = data.select { |_, v| types.include?(v) }
      corners = Array.new(cars.count, 0)
      cars.each_with_index do |((x, y), type), index|
        case type
        when '<'
          next = [x - 1, y]
        when '>'
          next = [x + 1, y]
        when '^'
          next = [x, y - 1]
        when 'v'
          next = [x - 1, y + 1]
        end
      end
    end
  end
end

puts Handle.new.location_of_first_crash

