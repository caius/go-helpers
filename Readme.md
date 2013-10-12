# Go-Helpers

Collection of helper methods written in go that aren't substantial enough to place in their own package.

## `Ordinalize()`

Given an integer, returns it with the appropriate suffix for placing in an english sentence.

    Ordinalize(1) # => "1st"
    Ordinalize(2) # => "2nd"
    Ordinalize(3) # => "3rd"
    Ordinalize(4) # => "4th"
    Ordinalize(21) # => "21st"

## License

MIT License, see LICENSE file.
