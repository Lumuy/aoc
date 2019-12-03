require 'pry'

class Handle
  def initialize(file = 'data')
    @data = []
    File.read(file).split("\n").each_with_index do |line, index|
      next if index == 0
      @data << line.split(' ')
    end
  end

  # Part 1
  def lowest_no_negative_integer(data = @data)
  end
end

Handle.new.lowest_no_negative_integer
