my_var = "success"

MyClass = Class.new do
  puts "#{my_var} in defined class"

  define_method :my_method do
    "#{my_var} in defined method"
  end
end

puts MyClass.new.my_method
