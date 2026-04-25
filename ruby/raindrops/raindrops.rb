# frozen_string_literal: true

# Convert an integer to Pling, Plang, Plong
class Raindrops
  def self.convert(int)
    result = ''
    result += 'Pling' if (int % 3).zero?
    result += 'Plang' if (int % 5).zero?
    result += 'Plong' if (int % 7).zero?
    result = int.to_s if result == ''
    result = 'Input must be a positive integer' if int < 1

    result
  end
end
