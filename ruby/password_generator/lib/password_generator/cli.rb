require 'thor'

module PasswordGenerator
  class CLI < Thor

    default_command :help

    desc 'version', 'Show `password_generator` version.'
    def version
      $stdout.puts PasswordGenerator::VERSION
    end

    desc 'generate', 'Generate password.'
    option :size,      aliases: '-s', type: :numeric, default: 18,    desc: 'word size'
    option :number,    aliases: '',   type: :boolean, default: true,  desc: ''
    option :symbol,    aliases: '',   type: :boolean, default: false, desc: ''
    def generate
      $stdout.puts PasswordGenerator::generate(options)
    end

  end
end
