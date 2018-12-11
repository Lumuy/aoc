class Polymer
  attr_accessor :polymer

  def initialize(file = 'data')
    @polymer = File.read(file).gsub("\n", "")
  end

  # Part 1
  def remain_units_count(polymer = @polymer)
    remain = ""
    while polymer.length > 0 do
      keep   = true
      index  = 0
      first  = remain[-1] || 'remain last'
      second = polymer[0] || 'left  first'
      if first == second.swapcase
        remain = remain[0, remain.length - 1]
        keep   = false
      elsif polymer.length > index  && polymer[index + 1] == polymer[index].swapcase
         index  += 1
         keep = false if keep
      end
      remain += polymer[0] if keep
      polymer = polymer[index + 1, polymer.length - index - 1]
    end
    remain.length
  end

  # Part 2
  def shortest_polymer_length
    chars = ('a'..'z').to_a
    counts = chars.map do |lower|
      upper = lower.upcase
      polymer = @polymer.gsub(/[#{lower}#{upper}]/, '')
      remain_units_count(polymer)
    end
    counts.min
  end
end

puts Polymer.new.remain_units_count
puts Polymer.new.shortest_polymer_length
