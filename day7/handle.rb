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
    data.each { |d| orders[d[1]] = orders[d[1]] << d[0] }
    result = ''

    while orders.size > 0 do
      key, _ = orders.find do |step, requirements|
        requirements.empty? || requirements.all? { |c| result.include?(c) }
      end
      result += key
      orders.delete(key)
    end

    result
  end

  # Part 2
  def time_to_complete(data = @data, count = 5)
    step_times = ('A'..'Z').reduce({}) { |r, c| r.merge!({c => c.ord - 4}) }
    orders = ('A'..'Z').reduce({}) { |h, c| h.merge!({c => []}) }
    data.each { |d| orders[d[1]] = orders[d[1]] << d[0] }
    result, time, process = '', 0, {}

    while orders.size > 0 || process.size > 0 do
      count.times do
        next unless process.size < count

        pairs = orders.select do |step, requirements|
          requirements.empty? || requirements.all? { |c| result.include?(c) }
        end

        if pairs.size > 0
          key = pairs.keys.min
          process[key] = step_times[key]
          orders.delete(key)
        end
      end

      process.keys.each { |key| process[key] -= 1 }
      overflows = process.select { |_, v| v == 0 }
      overflows.each do |key, _|
        result += key
        process.delete(key)
      end

      time += 1
    end

    time
  end
end

puts Handle.new.get_order
puts Handle.new.time_to_complete

