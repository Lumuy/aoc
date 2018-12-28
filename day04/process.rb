require 'time'
require 'pry'

class Find
  def initialize(file = 'data')
    @data = File.read(file).split("\n").map do |str|
      str.scan(/\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] (.*)/).flatten
    end.sort_by { |e| Time.parse(e[0]).to_i }
    File.write('inputs', @data.map { |e| e[0].to_s + e[1].to_s }.join("\n"))
  end

  def grouped_data
    hash = {}
    id = 0
    @data.each do |line|
      if line[1].include?('begins shift')
        id = line[1].scan(/(\d+)/).flatten.first
        hash[id] = [] if hash[id].nil?
      else
        hash[id] << line
      end
    end
    hash
  end

  # Part 1
  def find_guard_most_sleep
    inputs = grouped_data
    ids, times = inputs.keys, inputs.values
    minutes = ->(str) { str.scan(/[\d\s-]+\:(\d{2})/).flatten.first.to_i }
    sleep_times = times.map do |time_array|
      if time_array.count > 1
        0.upto(time_array.length/2 - 1).reduce(0) do |r, i|
          r += (minutes.call(time_array[2*i+1][0]) - minutes.call(time_array[2*i][0]))
        end
      else
        0
      end
    end
    index = sleep_times.index(sleep_times.max)
    id = ids[index]
    arr = times[index]
    len = arr.length/2

    res = []
    fall_asleep_times = (0..len-1).map { |i| arr[2*i][0] }
    wake_up_times     = (0..len-1).map { |i| arr[2*i+1][0] }
    keep_day = ""
    (0..len-1).each do |i|
      day = fall_asleep_times[i].split(" ")[0]
      if day != keep_day
        stats = Array.new(60, false)
        keep_day = day
      else
        stats = res.pop
      end
      sleep  = fall_asleep_times[i].scan(/[\d\s-]+:(\d{2})/).flatten.first.to_i
      wake   = wake_up_times[i].scan(/[\d\s-]+:(\d{2})/).flatten.first.to_i
      (sleep..wake).each { |i| stats[i] = true }
      res.push stats
    end

    base = Array.new(60, 0)
    res.each do |i|
      i.each_with_index do |j, index|
        base[index] += 1 if j
      end
    end
    shared_min = base.index(base.max)
    puts shared_min * id.to_i
  end

  # Part 2
  def find_guard_most_frequently_sleep
    inputs = grouped_data
    ids, guards_times = inputs.keys.map(&:to_i), inputs.values
    len = ids.length

    res = Array.new(60) { Array.new(len, 0) }
    guards_times.each_with_index do |guard_times, i|
      half_len = guard_times.count/2
      fall_asleep_times = (0..half_len-1).map { |n| guard_times[2*n][0] }
      wake_up_times     = (0..half_len-1).map { |n| guard_times[2*n+1][0] }
      (0..half_len-1).each do |j|
        sleep  = fall_asleep_times[j].scan(/[\d\s-]+:(\d{2})/).flatten.first.to_i
        wake   = wake_up_times[j].scan(/[\d\s-]+:(\d{2})/).flatten.first.to_i
        (sleep..wake).each { |minute| res[minute][i] += 1 }
      end
    end

    max = res.flatten.max
    minute = index = 0
    res.each_with_index do |e, i|
      break if minute!= 0
      e.each_with_index do |c, j|
        if c == max
          minute = i
          index = j
          break
        end
      end
    end
    puts minute * ids[index]
  end

end

puts Find.new.find_guard_most_sleep
puts Find.new.find_guard_most_frequently_sleep

