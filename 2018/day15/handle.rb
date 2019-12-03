require 'pry'
require 'set'

class Handle
  def load(file)
    File.read(file).split("\n")
  end
  def initialize(file = 'test_data1')
    @data   = {}
    content = load(file)
    content.map(&:chars).each.with_index do |arr, y|
      arr.each_with_index { |e, x| @data[[x, y]] = e }
    end
    @width  = content[0].size
    @height = content.size
  end

  def the_outcome_of_combat(area = @data.dup)
    rounds = 0
    hps    = area.select { |_, v| 'GE'.include?(v) }.keys
                 .reduce({}) { |r, p| r.merge(p => 200) }

    loop do
      # one round
      reading_order(hps).each do |current, hp|
        enemies = area.select { |_, v| v == 'GE'.delete(area[current]) }
        ranges  = reading_order(enemies).map { |enemy, _| next_step(current, enemy, area) }.compact
        target  = ranges.sort_by { |e| e[1] }[0][0] unless ranges.empty?

        if target && enemy_adjacent?(current, area)
          move(current, target, area, hps)
          current = target
        end
        attack(current, area, hps)
      end

      # end
      rounds += 1
      # test
      print_area(rounds, area)
      puts hps
      break if hps.keys.map { |p| area[p] }.uniq.size == 1
    end

    # Part 1
    puts rounds
    puts hps.values.sum
    rounds * hps.values.sum
  end

  def enemy_adjacent?(current, area)
    type = 'GE'.delete(area[current])
    adjacents(current).select { |p| area[p] == type } == []
  end

  def move(current, target, area, hps)
    hp = hps[current]
    hps[target] = hp
    hps.delete(current)

    type = area[current]
    area[target] = type
    area[current] = '.'
  end

  def attack(current, area, hps)
    type = 'GE'.delete(area[current])
    enemies = adjacents(current).select { |p| area[p] == type }
    e_hps = hps.select { |p, _| enemies.include?(p) }
    e_hps = reading_order(e_hps).sort_by { |k, v| v }

    unless e_hps.empty?
      e_point, e_hp = e_hps[0]
      new_hp = e_hp -3
      hps[e_point] = e_hp - 3

      if new_hp < 1
        hps.delete(e_point)
        area[e_point] = '.'
      else
        hps[e_point] = new_hp
      end
    end
  end

  def next_step(current, target, area)
    possible = adjacents(current).select { |p| area[p] == '.' }

    options = [target]
    seen = []

    distance = 0
    next_point = nil

    loop do
      reachable = options.reduce([]) do |r, p|
        r + adjacents(p)
      end.select { |p| area[p] == '.' } - seen
      reachable = reachable.uniq

      break if reachable.empty?

      options = reachable
      seen |= reachable
      distance += 1
      reached = possible & reachable

      unless reached.empty?
        next_point = reached[0]
        break
      end
    end

    [next_point, distance] if next_point
  end

  def adjacents(point)
      x, y   = point

      [[x, y - 1], [x - 1, y], [x + 1, y], [x, y + 1]].select do |dx, dy|
        dx > 0 && dx < @width - 1 && dy > 0 && dy < @height - 1
      end
  end

  def can_attack(spoint, dpoint)
    sx, sy = spoint
    dx, dy = dpoint

    (sx - dx).abs + (sy - dy).abs == 1
  end

  # top-to-bottom then left-to-right
  def reading_order(points)
    points.to_a
          .group_by { |e| e[0][1] }
          .reduce({}) { |res, (k, v)| res.merge(k => v.sort_by { |e| e[0][0] }) }
          .sort.to_h.values
          .reduce([]) { |res, e| res + e }
          .to_h
  end

  def print_area(rounds, area)
    puts "After #{rounds} rounds"
    w, h = area.to_a[-1][0]
    (0..h).each do |y|
      (0..w).each do |x|
        print area[[x, y]]
      end
      puts
    end
  end
end

puts Handle.new.the_outcome_of_combat
