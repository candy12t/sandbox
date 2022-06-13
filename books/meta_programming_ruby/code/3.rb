module V1
  class Computer
    def initialize(computer_id, data_source)
      @id = computer_id
      @data_source = data_source
    end

    def mouse
      info = @data_source.get_mouse_info(@id)
      price = @data_source.get_mouse_price(@id)
      result = "Mouse: #{info} ($#{price})"
      return "* #{result}" if price >= 100
      return result
    end

    def cpu
      info = @data_source.get_cpu_info(@id)
      price = @data_source.get_cpu_price(@id)
      result = "Cpu: #{info} ($#{price})"
      return "* #{result}" if price >= 100
      return result
    end

    def keyboard
      info = @data_source.get_keyboard_info(@id)
      price = @data_source.get_keyboard_price(@id)
      result = "Keyboard: #{info} ($#{price})"
      return "* #{result}" if price >= 100
      return result
    end
  end
end

module V2
  class Computer
    def initialize(computer_id, data_source)
      @id = computer_id
      @data_source = data_source
    end

    def mouse
      compoent :mouse
    end

    def cpu
      compoent :cpu
    end

    def keyboard
      compoent :keyboard
    end

    def compoent(name)
      info = @data_source.send "get_#{name}_info", @id
      price = @data_source.send "get_#{name}_price", @id
      result = "#{name.capitalize}: #{info} ($#{price})"
      return "* #{result}" if price >= 100
      return result
    end
  end
end

module V3
  class Computer
    def initialize(computer_id, data_source)
      @id = computer_id
      @data_source = data_source
    end

    def self.define_component(name)
      define_method(name) do
        info = @data_source.send "get_#{name}_info", @id
        price = @data_source.send "get_#{name}_price", @id
        result = "#{name.capitalize}: #{info} ($#{price})"
        return "* #{result}" if price >= 100
        return result
      end
    end

    define_component :mouse
    define_component :cpu
    define_component :keyboard
  end
end

module V4
  class Computer
    def initialize(computer_id, data_source)
      @id = computer_id
      @data_source = data_source
      data_source.methods.grep(/^get_(.*)_info$/) {Computer.define_component $1}
    end

    def self.define_component(name)
      define_method(name) do
        info = @data_source.send "get_#{name}_info", @id
        price = @data_source.send "get_#{name}_price", @id
        result = "#{name.capitalize}: #{info} ($#{price})"
        return "* #{result}" if price >= 100
        return result
      end
    end
  end
end

module MethodMissing
  class Computer
    def initialize(computer_id, data_source)
      @id = computer_id
      @data_source = data_source
    end

    def method_missing(name)
      super if !@data_source.respond_to? "get_#{name}_info"

      info = @data_source.send "get_#{name}_info", @id
      price = @data_source.send "get_#{name}_price", @id
      result = "#{name.capitalize}: #{info} ($#{price})"
      return "* #{result}" if price >= 100
      return result
    end

    def respond_to_missing?(method, include_private = false)
      @data_source.respond_to? "get_#{method}_info" || super
    end
  end
end

class DS
  def initialize
    @data = {one: {mouse: {info: "hoge", price: 50}, cpu: {info: "fuga", price: 200}, keyboard: {info: "hhkb", price: 300}}}
  end

  %i(mouse cpu keyboard).each do |part|
    define_method("get_#{part}_info") do |id|
      @data[id.to_sym][part][:info]
    end

    define_method("get_#{part}_price") do |id|
      @data[id.to_sym][part][:price]
    end
  end
end

ds = DS.new
c = MethodMissing::Computer.new("one", ds)
p c.respond_to? :mouse
p c.mouse
p c.cpu
p c.keyboard
