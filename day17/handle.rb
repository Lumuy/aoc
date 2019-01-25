class Handle
  def initialize(file = 'data')
    @area = Hash.new('.')
    File.read(file).split("\n").map { |e| e.split(', ').map { |e| e.split('=') } }.map(&:to_h).each do |hash|
      convert_cordin(hash['x']).each { |x| convert_cordin(hash['y']).each { |y| @area[[x, y]] = '#' } }
    end

    @min_y = @area.keys.map(&:last).min
    @max_y = @area.keys.map(&:last).max

    waters
  end

  def waters(spring_x = 500, min_y = @min_y)
    points = [[spring_x , min_y - 1]]
    until points.count == 0
      points = fall(points)
      points = spread(points)
    end
  end

  def fall(points)
    new_points = []
    points.each do |x, y|
      while @area[[x, y + 1]] == '.'
        y += 1
        break if y > @max_y

        @area[[x, y]] = '|'
      end

      new_points << [x, y] if y <= @max_y && @area[[x, y + 1]] != '|'
    end
    new_points
  end

  def spread(points)
    new_points = []
    points.each do |ox, y|
      bounds_x = []
      [1, -1].each do |step|
        x = ox
        while '.|'.include?(@area[[x + step, y]])
          x += step
          @area[[x, y]] = '|'
          break if '.|'.include?(@area[[x , y + 1]])
        end
        '.|'.include?(@area[[x, y + 1]]) ? new_points << [x, y] : bounds_x << x
      end
      next unless bounds_x.count == 2

      (bounds_x.min..bounds_x.max).each { |rx| @area[[rx, y]] = '~' }
      @area[[ox, y - 1]] = '|'
      new_points << [ox, y - 1]
    end
    new_points.uniq
  end

  # Part 1
  # Part 2
  def water_tiles
    puts @area.select { |_, v| '~|'.include?(v) }.keys.select { |_, y| y <= @max_y }.count
    puts @area.select { |_, v| v == '~' }.keys.select { |_, y| y <= @max_y }.count
  end

  private

  def convert_cordin(str)
    if str.include?('..')
      min, max = str.split('..').map(&:to_i)
      Range.new(min, max).to_a
    else
      [str.to_i]
    end
  end

end

Handle.new.water_tiles
