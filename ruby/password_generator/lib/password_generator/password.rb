module PasswordGenerator
  class Password

    ATTRIBUTES = %w(size number symbol)

    def initialize(options = {})
      ATTRIBUTES.each do |attr|
        instance_variable_set("@#{attr}", options[attr.to_sym]) if options[attr.to_sym]
      end
    end

    def generate
      base = base_strings
      password = @size.times.reduce('') do |result, _|
        result + base[rand(base.length)].to_s
      end
      return password
    end

    private

    def base_strings
      return @base unless @base.nil?

      @base = ('a'..'z').to_a + ('A'..'Z').to_a
      @base += (0..9).to_a if number?
      @base += %w(! " # $ % & ' ( ) * + , - . / : ; < = > ? [ \ ] ^ _ ` { | } ~) if symbol?
      return @base
    end

    def number?
      return @number
    end

    def symbol?
      return @symbol
    end

  end
end
