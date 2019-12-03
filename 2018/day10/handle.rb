class Array
  def split_sorted
    dup, res = self, []
    while dup.size > 1 do
      index, tmp = 0, []
      while dup.include?(dup[index] + 1) do
        index += 1
      end
      (index + 1).times { tmp << dup.shift }
      res << tmp
    end

    res
  end
end

class Handle
  def initialize(file = 'data')
    @data = File.read(file).split("\n").map do |line|
      line.scan(/-?\d+/).map(&:to_i)
    end
  end

  # Part 1
  # Part 2
  def message_in_sky(data = @data)
    seconds = 0
    while true do
      seconds += 1
      points = data.map { |e| [e[0] + e[2] * seconds, e[1] + e[3] * seconds] }
      group_x = points.group_by(&:first).map { |k, v| [k, v.map(&:last)] }.to_h
      group_y = points.group_by(&:last).map { |k, v| [k, v.map(&:first)] }.to_h

      if is_msg?(group_x) || is_msg?(group_y)
        puts seconds
        print_msg points
        break
      end
    end
  end

  def is_msg?(group)
    res = group.keys.sort.split_sorted.map do |arr|
     arr.reduce([]) { |r, x| r + group[x] }
    end.map { |e| [e.min, e.max] }
    res.size > 1 && res.uniq.size == 1
  end

  def print_msg(points)
    xs, ys = points.map(&:first), points.map(&:last)
    (ys.min..ys.max).each do |y|
      (xs.min..xs.max).each do |x|
        if points.include? [x, y]
          print '#'
        else
          print ' '
        end
      end
      print "\n"
    end
  end
end

puts Handle.new.message_in_sky

