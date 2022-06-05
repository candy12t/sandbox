class String
  VOWEL = %w(a e i o u)
  def pluralize
    case
    when end_s?, end_sh?, end_ch?, end_x?, end_consonant_o?
      return self + 'es'
    when end_consonant_y?
      return self[0..-2] + 'ies'
    when end_f?
      return self[0..-2] + 'ves'
    when end_fe?
      return self[0..-3] + 'ves'
    else
      return self + 's'
    end
  end

  private
  def end_s?
    return self[-1] == 's'
  end

  def end_sh?
    return self[-2..-1] == 'sh'
  end

  def end_ch?
    return self[-2..-1] == 'ch'
  end

  def end_x?
    return self[-1] == 'x'
  end

  def end_consonant_o?
    return VOWEL.none?(self[-2]) && self[-1] == 'o'
  end

  def end_consonant_y?
    return VOWEL.none?(self[-2]) && self[-1] == 'y'
  end

  def end_f?
    return self[-1] == 'f'
  end

  def end_fe?
    return self[-2..-1] == 'fe'
  end
end
