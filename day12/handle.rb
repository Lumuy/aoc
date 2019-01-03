class Handle
  def initialize(file = 'data')
    data = File.read(file).split("\n")
    @status = data[0].gsub('initial state: ', '').chars.map.with_index { |char, index| index if char == '#' }.compact
    @rules = data[2..-1].map { |e| e.scan(/([#.]+) => ([#.]{1})/).flatten }.to_h
      .select { |_, v| v == '#' }.keys
  end

  def one_generation(state, rules = @rules)
    res = []
    (state.min - 2..state.max + 2).each do |index|
      res << index if rules.include?((index - 2..index + 2).map { |i| state.include?(i) ? '#' : '.' }.join)
    end
    res
  end

  def sum_of_plant_pots(n)
    status = @status
    n.times { status = one_generation(status) }
    status.sum
  end

  # Part 1
  def twenty_generation
    sum_of_plant_pots(20)
  end

  def get_circle_num(status = @status)
    n, base = 0, []
    while true do
      n += 1
      status = one_generation(status)
      min = status.min
      break if status.sort.map { |e| e - min } == base
      base = status.map { |e| e - min }
    end
    [n, status]
  end

  # Part 2
  def fifty_billion_generation(times = 50000000000)
    n, arr = get_circle_num
    arr.reduce(0) { |r, e| r + e + times - n }
  end
end

puts Handle.new.twenty_generation
puts Handle.new.fifty_billion_generation

