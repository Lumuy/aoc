require 'pry'

class Handle
  def initialize(file = 'data')
    @data = File.read(file).split("\n").map do |line|
      line.scan(/[a-zA-z]+ ([A-Z]{1}) [a-z\s]+ ([A-Z]{1}) .*/).flatten
    end
  end

  # Part 1
  def get_order(data = @data)
    orders = ('A'..'Z').reduce({}) { |h, c| h.merge!({c => []}) }
    data.each { |d| orders[d[0]] = orders[d[0]] << d[1] }
    binding.pry
  end
end

Handle.new.get_order

