class Handle
  def initialize(file = 'data')
    @data = File.read(file).split(' ').map(&:to_i)
  end

  def get_tree(data = @data)
    # data = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2".split(' ').map(&:to_i)
    type, tree, parent = 0, {}, 1

    while data.size > 0 do
      header, type = data.shift(2), type + 1

      tree[type] = [Array.new(header[0], -1), Array.new(header[1], -1)]

      parent = type - 1
      until parent == 0 || tree[parent][0].include?(-1) do
        parent, = tree.find { |_, v| v[0].include?(parent) }
      end
      if parent != 0
        tree[parent][0].shift
        tree[parent][0] << type
      end

      tree[type][1] = data.shift(header[1]) if header[0] == 0

      current = type
      while parent && tree[current].flatten.all? { |e| e != -1 } && tree[parent][0].all? { |e| e != -1 } do
        tree[parent][1] = data.shift(tree[parent][1].size)

        current = parent
        parent, = tree.find { |_, v| v[0].include?(current) }
      end
    end

    tree
  end

  # Part 1
  def metadata_entries_sum(data = @data)
    get_tree.values.map(&:last).flatten.reduce(&:+)
  end

  # Part 2
  def value_of_root_node
    tree = get_tree
    tree.each do |k, (child, metadata)|
      if child.empty?
        tree[k] = metadata.sum
      else
        indexes = metadata.select { |e| e > 0 && e <= child.count }
        tree[k] = indexes.empty? ? 0 : indexes.map { |index| child[index - 1] }
      end
    end

    root, value = tree[1], 0
    while root.size > 0 do
      root = root.reduce([]) do |r, type|
        if tree[type].is_a?(Array)
          r + tree[type]
        elsif tree[type].is_a?(Integer)
          value += tree[type]
          r + []
        end
      end
    end

    value
  end
end

puts Handle.new.metadata_entries_sum
puts Handle.new.value_of_root_node
