class Delay
  def initialize(&func)
    @func = func
    @evaluated = false
    @value = nil
  end

  def force
    return @value if @evaluated

    @value = @func.call
    @evaluated = true
    return @value
  end
end

d = Delay.new { puts 'hoge'; 10+20 }
p d.force
p d.force
