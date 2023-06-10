require "ffi"

module Fib
  extend FFI::Library
  ffi_lib "fib.so"
  attach_function :fib, [:uint], :uint
end

Fib.fib(40)
