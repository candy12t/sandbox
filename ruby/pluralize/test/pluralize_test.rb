require 'minitest/autorun'
require './lib/pluralize'

class PluralizeTest < MiniTest::Test
  def test_pluralize
    assert_equal 'buses', 'bus'.pluralize
    assert_equal 'dishes', 'dish'.pluralize
    assert_equal 'churches', 'church'.pluralize
    assert_equal 'boxes', 'box'.pluralize
    assert_equal 'tomatoes', 'tomato'.pluralize
    assert_equal 'cities', 'city'.pluralize
    assert_equal 'shelves', 'shelf'.pluralize
    assert_equal 'knives', 'knife'.pluralize
    assert_equal 'cats', 'cat'.pluralize
  end
end
