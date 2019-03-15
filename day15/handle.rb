require 'pry'

class Handle
  def initialize(file = 'sample_data')
    @data = {}
    File.read(file).split("\n").map(&:chars).each.with_index do |arr, y|
      arr.each_with_index { |e, x| @data[[x, y]] = e }
    end
  end

  # Part 1
  def the_outcome_of_combat(data = @data)
    goblins = @data.select { |_, v| v == 'G' }
    elves   = @data.select { |_, v| v == 'E' }
    binding.pry
  end
end

Handle.new.the_outcome_of_combat
