class Handle
  def initialize(file = 'inputs')
    @data = File.read(file).split("\n").map(&:to_i)
  end

  def frequency(data = @data)
    # Part 1
    puts data.reduce(:+)

    size, arr = 0, [0]
    loop do
      elem = data[size] + arr[-1]
      if arr.include?(elem)
        # Part 2
        puts elem
        break
      else
        arr << elem
      end

      size += 1
      size = 0 if size == data.size
    end
  end
end

Handle.new.frequency
