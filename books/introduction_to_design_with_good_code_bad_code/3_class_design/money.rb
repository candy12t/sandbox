class Money
  attr_reader :amount, :currency

  def initialize(amount, currency)
    if amount < 0
      raise ArgumentError.new("specify amount more than 0.")
    end
    if currency.nil? || currency.empty?
      raise ArgumentError.new("specify currency.")
    end

    @amount = amount
    @currency = currency
    self.freeze
  end

  def add(other)
    if @currency != other.currency
      raise ArgumentError.new("different currency.")
    end
    added = @amount + other.amount
    return Money.new(added, @currency)
  end
end

one_handread = Money.new(100, "en")
two_handread = one_handread.add(Money.new(100, "en"))
puts one_handread.amount, one_handread.currency
puts two_handread.amount, two_handread.currency
