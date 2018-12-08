require 'set'

class Day22
  attr_reader :data, :result

  def initialize(file = 'data2')
    @data = File.read(file).split("\n")
  end

  def length
    @length ||= @data[0].length
  end

  def get(n)
    arr = @data.map do |str|
      if n == length - 1
        str[0, n]
      else
        str[0, n] + str[n + 1, length - 1]
      end
    end

    if arr.count > arr.uniq.count
      arr.each do |e|
        @result = e if arr.count(e) != 1
      end
    end
  end

  def perform
    (0..(length - 1)).each do |n|
      get(n)
    end

    @result
  end
end

puts Day22.new.perform

