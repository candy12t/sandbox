Dir.glob('password_generator/*.rb', base: File.dirname(__FILE__)).each do |module_file|
  require module_file
end

module PasswordGenerator
  def self.generate(options)
    base = ('a'..'z').to_a + ('A'..'Z').to_a
    base += (0..9).to_a if options[:number]
    base += %w(! " # $ % & ' ( ) * + , - . / : ; < = > ? [ \ ] ^ _ ` { | } ~) if options[:symbol]

    password = options[:size].times.reduce('') do |result, _|
      result + base[rand(base.length)].to_s
    end

    return password
  end
end
