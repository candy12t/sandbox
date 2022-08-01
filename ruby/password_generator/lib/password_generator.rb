Dir.glob('password_generator/*.rb', base: File.dirname(__FILE__)).each do |module_file|
  require module_file
end

module PasswordGenerator
  def self.generate(options)
    password = PasswordGenerator::Password.new(options).generate
    return password
  end
end
