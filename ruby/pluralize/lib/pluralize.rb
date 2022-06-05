# TODO: supoort exception word (man, child ect...)
class String
  def pluralize
    case self
    when /(s$|sh$|ch$|x$|[^aeiou]{1}o$)/
      return self + 'es'
    when /[^aeiou]{1}y$/
      return self[0..-2] + 'ies'
    when /f$/
      return self[0..-2] + 'ves'
    when /fe$/
      return self[0..-3] + 'ves'
    else
      return self + 's'
    end
  end

  def pluralize!
    return self.replace(pluralize)
  end
end
