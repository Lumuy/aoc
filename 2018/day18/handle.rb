class Handle
  def initialize(file = 'data')
    @data = File.read(file).split("\n").map { |line| line.split('') }
  end

  # Part 1
  # Part 2
  def total_resouce_value(minutes = 10)
    changes, area, holds, cycles = {}, @data, [], []
    max_x, max_y = area[0].size, area.size

    # print_area(area, 0)

    (1..minutes).each do |time|
      (0..max_x - 1).each do |x|
        (0..max_y - 1).each do |y|
          adjacents = [
            [x - 1, y - 1],
            [x - 1, y],
            [x - 1, y + 1],
            [x, y - 1],
            [x, y + 1],
            [x + 1, y - 1],
            [x + 1, y],
            [x + 1, y + 1]
          ].select { |k, v| k >= 0 && v >= 0 && k <= max_x - 1 && v <= max_y - 1 }
           .map { |dx, dy| area[dx][dy] }.compact

          case area[x][y]
          when '.'
            changes.merge!([x, y] => '|') if adjacents.count('|') >= 3
          when '|'
            changes.merge!([x, y] => '#') if adjacents.count('#') >= 3
          when '#'
            changes.merge!([x, y] => '.') unless adjacents.include?('|') && adjacents.include?('#')
          end
        end
      end

      changes.each { |(cx, cy), value| area[cx][cy] = value }

      resource = [area.flatten.count('|'), area.flatten.count('#')].reduce(&:*)

      if holds.include?(resource)
        unless cycles.include?(resource)
          cycles << resource
        else
          if holds.join.include? cycles.join
            puts "Part 2: #{cycles[(minutes - holds.size) % cycles.size - 1]}"
            return
          else
            cycles = []
          end
        end
      end

      holds << resource

      # print_area(area, time)
    end

    puts "Part 1: #{area.flatten.count('|') * area.flatten.count('#')}"
  end

  private

  def print_area(area, time)
    puts
    puts "After #{time} minutes"
    area.each { |line| puts line.join }
  end
end

Handle.new.total_resouce_value
Handle.new.total_resouce_value(1_000_000_000)
