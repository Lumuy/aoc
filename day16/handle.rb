class Handle
  OPS = %w[addr addi mulr muli banr bani borr bori setr seti gtir gtri gtrr eqir eqri eqrr]

  def initialize(file = 'data')
    data = File.read(file).split("\n").select { |e| e != '' }
    len  = data.size
    dlen = 0
    data.each_with_index { |item, index| dlen = index if item.match? /^After.*/ }

    l_data = data[0..dlen]
    @l_data = []
    (0..((dlen - 2) / 3)).to_a.each do |n|
      i = n * 3
      before = l_data[i].scan(/\d+[\,\]]/).map(&:to_i)
      state  = l_data[i + 1].split(' ').map(&:to_i)
      after  = l_data[i + 2].scan(/\d+[\,\]]/).map(&:to_i)

      @l_data << [before, state, after]
    end

    @r_data = data[dlen + 1..len].map { |e| e.split(' ').map(&:to_i) }
  end

  # Part 1
  def opcodes_count(ldata = @l_data)
    ldata.map do |before, state, after|
      OPS.map { |op| send(op, before, state) }
         .count(after)
    end.select { |e| e >= 3 }.count
  end

  # Part 2
  def last_registers(registers = [0, 0, 0, 0], rdata = @r_data)
    hash = get_code_op_map
    rdata.each do |state|
      code = state[0]
      registers = send(hash[code], registers, state)
    end

    registers[0]
  end

  def get_code_op_map(ldata = @l_data, map = [])
    while map.size <= 15
      before_count = map.size
      ldata.each do |before, state, after|
        opcode   = state[0]
        possible = []
        next if map.map(&:first).include?(opcode)
        (OPS - map.map(&:last)).each{ |op| possible << op if send(op, before, state) == after }

        map << [opcode, possible[0]] if possible.size == 1
      end

      break if map.size != before_count + 1
      map = map.uniq
    end

    map.to_h
  end

  # operate
  def addr(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] + r[b] }
  end
  def addi(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] + b }
  end
  def mulr(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] * r[b] }
  end
  def muli(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] * b }
  end
  def banr(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] & r[b] }
  end
  def bani(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] & b }
  end
  def borr(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] | r[b] }
  end
  def bori(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] | b }
  end
  def setr(before, op)
    op(before, op) { |r, a, _, c| r[c] = r[a] }
  end
  def seti(before, op)
    op(before, op) { |r, a, _, c| r[c] = a }
  end
  def gtir(before, op)
    op(before, op) { |r, a, b, c| r[c] = a > r[b] ? 1 : 0 }
  end
  def gtri(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] > b ? 1 : 0 }
  end
  def gtrr(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] > r[b] ? 1 : 0 }
  end
  def eqir(before, op)
    op(before, op) { |r, a, b, c| r[c] = a == r[b] ? 1 : 0 }
  end
  def eqri(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] == b ? 1 : 0 }
  end
  def eqrr(before, op)
    op(before, op) { |r, a, b, c| r[c] = r[a] == r[b] ? 1 : 0 }
  end

  private

  def op(before, op, &block)
    registers = before.dup
    _, a, b, c = op
    block.call(registers, a, b, c)
    registers
  end
end

handle = Handle.new
puts handle.opcodes_count
puts handle.last_registers

