class MyClass
  def my_method
    @value = 1
  end
end

obj1 = MyClass.new
obj1.my_method
p obj1.instance_variables # => [:@value]

obj2 = MyClass.new
p obj2.instance_variables # => []
