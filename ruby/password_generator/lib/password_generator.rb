Dir.glob('password_generator/*.rb', base: File.dirname(__FILE__)).each do |module_file|
  require module_file
end

module PasswordGenerator
  def self.generate(options)
    length = options.delete(:size)
    password = PasswordGenerator::Password.new(length, options).generate
    return password
  end
end
