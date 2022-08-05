require 'thor'

module PasswordGenerator
  class CLI < Thor

    default_command :help

    desc 'version', 'Show `password_generator` version.'
    def version
      $stdout.puts PasswordGenerator::VERSION
    end

    desc 'generate', 'Generate password.'
    option :length, aliases: '-l', type: :numeric, default: 18,    desc: 'word length'
    option :number, aliases: '-n', type: :boolean, default: true,  desc: 'include a number'
    option :symbol, aliases: '-s', type: :boolean, default: false, desc: 'include a symbol'
    def generate
      $stdout.puts PasswordGenerator::generate(options.dup)
    end

  end
end
