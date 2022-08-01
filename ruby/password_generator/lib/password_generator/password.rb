module PasswordGenerator
  class Password

    @@base_chars = ('a'..'z').to_a.join + ('A'..'Z').to_a.join

    def initialize(length, options)
      @length = length
      @@base_chars += (0..9).to_a.join if options[:number]
      @@base_chars += %w(! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \\ ] ^ _ ` { | } ~).join if options[:symbol]
    end

    def generate
      password = @length.times.reduce('') do |result, _|
        result + @@base_chars[rand(@@base_chars.length)].to_s
      end
      return password
    end

  end
end
