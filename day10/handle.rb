require 'pry'

class Handle
  def initialize(file = 'data')
    @data = File.read(file).split("\n").map do |line|
      line.scan(/-?\d+/).map(&:to_i)
    end
  end

  # Part 1
  def message_in_sky(data = @data)
    binding.pry
  end
end

puts Handle.new.message_in_sky

