require 'benchmark'

class Handle
  def initialize(file = 'data')
    @data = File.read(file).scan(/(\d+)[a-z;\s]+(\d+).*/).flatten.map(&:to_i)
  end

  # Part 1
  def winning_score(data = @data)
    players, points = data
    hash, current, scores = {0 => 0}, 0, Array.new(players, 0)

    (1..points).each do |marble|
      if marble % 23 == 0
        player = marble % players
        7.times { current = hash.key(current) }
        removed = current
        current = hash[removed]
        before = hash.key(removed)
        hash[before] = hash[removed]
        hash.delete(removed)

        scores[player] += marble + removed
      else
        replace = hash[hash[current]]
        hash[hash[current]] = marble
        hash[marble] = replace
        current = marble
      end
    end

    scores.max
  end

  # Part 2
  def larger_marble(data = @data)
    players, points = data
    current, scores = Marble.new(0), Array.new(players, 0)

    (1..points * 100).each do |point|
      if point % 23 == 0
        7.times { current = current.prev }
        removed = current
        current = removed.next
        removed.destroy

        scores[point % players] += removed.value + point
      else
        insert = Marble.new(point)
        insert.insert_after(current.next)
        current = insert
      end
    end

    scores.max
  end

  class Marble
    attr_accessor :value, :next, :prev

    def initialize(value)
      @value = value
      @next = self
      @prev = self
    end

    def destroy
      @prev.next = @next
      @next.prev = @prev
    end

    def insert_after(marble)
      @prev = marble
      @next = marble.next
      marble.next.prev = self
      marble.next = self
    end
  end
end

# puts Benchmark.measure { Handle.new.winning_score }
puts Handle.new.winning_score
# puts Benchmark.measure { Handle.new.larger_marble }
puts Handle.new.larger_marble
