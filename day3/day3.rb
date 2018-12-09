class Day3
  attr_reader :points, :data, :overlap

  def initialize(file = 'data')
    @overlap = 0
    @points = Array.new(1000) { Array.new(1000, 0) }
    @data = File.read(file).split("\n").map do |str|
      str.scan(/#\d+\s@\s(\d+),(\d+):\s(\d+)x(\d+)/).flatten.map(&:to_i)
    end
  end

  def perform
    @data.each do |rec|
      x, y, w, h = rec
      w.times do |dx|
        h.times do |dy|
          @overlap += 1 if @points[x + dx][y + dy] == 1
          @points[x + dx][y + dy] += 1
        end
      end
    end

    @overlap
  end

end

puts Day3.new.perform

