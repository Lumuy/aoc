# frozen_string_literal: true

require 'pry'

# Go with the flow
class Flow
  INSTRUCTIONS = {
    addr: ->(r, a, b, c) { r[c] = r[a] + r[b] },
    addi: ->(r, a, b, c) { r[c] = r[a] + b },
    mulr: ->(r, a, b, c) { r[c] = r[a] * r[b] },
    muli: ->(r, a, b, c) { r[c] = r[a] * b },
    banr: ->(r, a, b, c) { r[c] = r[a] & r[b] },
    bani: ->(r, a, b, c) { r[c] = r[a] & b },
    borr: ->(r, a, b, c) { r[c] = r[a] | r[b] },
    bori: ->(r, a, b, c) { r[c] = r[a] | b },
    setr: ->(r, a, _, c) { r[c] = r[a] },
    seti: ->(r, a, _, c) { r[c] = a },
    gtir: ->(r, a, b, c) { r[c] = a > r[b] ? 1 : 0 },
    gtri: ->(r, a, b, c) { r[c] = r[a] > b ? 1 : 0 },
    gtrr: ->(r, a, b, c) { r[c] = r[a] > r[b] ? 1 : 0 },
    eqir: ->(r, a, b, c) { r[c] = a == r[b] ? 1 : 0 },
    eqri: ->(r, a, b, c) { r[c] = r[a] == b ? 1 : 0 },
    eqrr: ->(r, a, b, c) { r[c] = r[a] == r[b] ? 1 : 0 }
  }.freeze

  def initialize(file = 'input')
    data = File.read(file).split("\n")

    @modify_register = data[0].gsub(/[^\d]+/, '').to_i
    @instructions    = data[1..data.size].map do |e|
      e.split(' ').map.with_index { |i, n| n.zero? ? i.to_sym : i.to_i }
    end
  end

  def process(registers = Array.new(6, 0))
    ip           = 0
    instructions = @instructions
    max_size     = instructions.size

    loop do
      ip = instruction_step(ip, instructions[ip], registers)

      break if ip >= max_size
    end

    registers[0]
  end

  def instruction_step(ip, instruction, registers)
    mri = @modify_register
    op, a, b, c = instruction

    registers[mri] = ip
    INSTRUCTIONS[op][registers, a, b, c]

    puts registers.join(', ')
    registers[mri] + 1
  end
end

flow = Flow.new
# Part 1
# puts flow.process

# Part 2
puts flow.process([1, 0, 0, 0, 0, 0])
