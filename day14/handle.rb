require 'pry'
require 'matrix'

class Handle
  def initialize(n = 880751)
    @recipes = n
  end

  # Part 1
  def the_scores_of_the_ten_recipes(recipes = @recipes)
    scoreboard, first, second = '37', 0, 1

    while scoreboard.size <= recipes + 10 do
      scoreboard += (scoreboard[first].to_i + scoreboard[second].to_i).to_s

      first  = (first  + 1 + scoreboard[first].to_i)  % scoreboard.size
      second = (second + 1 + scoreboard[second].to_i) % scoreboard.size
    end

    scoreboard[recipes, 10]
  end

  # Part 2
  def appear_after_recipes(recipes = 880751.to_s)
    scoreboard, first, second = '37', 0, 1

    until scoreboard.match(/#{recipes}/) do
      scoreboard += (scoreboard[first].to_i + scoreboard[second].to_i).to_s

      first  = (first  + 1 + scoreboard[first].to_i)  % scoreboard.size
      second = (second + 1 + scoreboard[second].to_i) % scoreboard.size
    end

    scoreboard.index(recipes)
  end
end

# puts Handle.new.the_scores_of_the_ten_recipes
puts Handle.new.appear_after_recipes
