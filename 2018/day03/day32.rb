class FindId
  attr_reader :data, :points

  def initialize(file = 'data2')
    @points = Array.new(1000) { Array.new(1000, -1) }
    @data = File.read(file).split("\n").map do |str|
      str.scan(/#\d+ @ (\d+),(\d+): (\d+)x(\d+)/).flatten.map(&:to_i)
    end
    @overlap = Array.new(@data.size) { false }
  end

  def perform
    @data.each_with_index do |rec, i|
      x, y, w, h = rec
      w.times do |dx|
        h.times do |dy|
          if @points[x + dx][y + dy] != -1
            @overlap[i] = true
            @overlap[@points[x + dx][y + dy]] = true
          end
          @points [x + dx][y + dy] = i
        end
      end
    end
    @overlap.index(false) + 1
  end
end

puts FindId.new.perform

