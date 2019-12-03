require 'pry'
require 'test/unit'
require_relative 'handle.rb'

class Handle
  def load(content)
    content.gsub(/ /, '').split("\n")
  end
end

class Day15Test < Test::Unit::TestCase
  def test_case_one
    input = """
    #######
    #.G...#
    #...EG#
    #.#.#G#
    #..G#E#
    #.....#
    #######
    """

    assert_equal 27730, Handle.new(input).the_outcome_of_combat
  end

  def test_case_two
    input = """
    #######
    #G..#E#
    #E#E.E#
    #G.##.#
    #...#E#
    #...E.#
    #######
    """

    assert_equal 36334, Handle.new(input).the_outcome_of_combat
  end

  def test_case_three
    input = """
    #######
    #E..EG#
    #.#G.E#
    #E.##E#
    #G..#.#
    #..E#.#
    #######
    """

    assert_equal 39514, Handle.new(input).the_outcome_of_combat
  end

  def test_case_four
    input = """
    #######
    #E..EG#
    #.#G.E#
    #E.##E#
    #G..#.#
    #..E#.#
    #######
    """

    assert_equal 27755, Handle.new(input).the_outcome_of_combat
  end

  def test_case_five
    input = """
    #######
    #.E...#
    #.#..G#
    #.###.#
    #E#G#G#
    #...#G#
    #######
    """

    assert_equal 28944, Handle.new(input).the_outcome_of_combat
  end

  def test_case_six
    input = """
    #########
    #G......#
    #.E.#...#
    #..##..G#
    #...##..#
    #...#...#
    #.G...G.#
    #.....G.#
    #########
    """

    assert_equal 18740, Handle.new(input).the_outcome_of_combat
  end
end

