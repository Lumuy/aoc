class Handle
  class << self
    def step
      scoreboard, first, second = [3, 7], 0, 1

      loop do
        sum = scoreboard[first] + scoreboard[second]
        (scoreboard << sum / 10; yield scoreboard) if sum > 9
        scoreboard << sum % 10
        yield scoreboard

        first = (first + 1 + scoreboard[first]) % scoreboard.size
        second = (second + 1 + scoreboard[second]) % scoreboard.size
      end
    end

    def count_recipes(input = 880751)
      start, len, digits = nil, 0, input.to_s.chars.map(&:to_i)

      step do |scoreboard|
        if scoreboard.size == input + 10
          puts "Part 1: #{scoreboard[input, 10].join}"
        end

        unless scoreboard.last == digits[len]
          len = 0
          start = nil
        end

        if scoreboard.last == digits[len]
          len += 1
          start = scoreboard.size - 1 if start.nil?
          break if len == digits.size
        end
      end

      puts "Part 2: #{start}"
    end
  end
end

Handle.count_recipes
