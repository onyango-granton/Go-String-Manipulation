## Project Title: Go Text Editing Tool
### Introduction

This project aims to create a text completion, editing, and auto-correction tool using Go. The tool will receive as input a text file that requires modifications and output the modified text to another file. The modifications include various transformations and formatting rules.

### Requirements

    Language: Go
    Coding Practices: Adherence to good practices
    Testing: Test files for unit testing

### Functionality

The tool will perform the following modifications to the text:

    Hexadecimal Conversion: Replace hexadecimal numbers with their decimal equivalents.

    - Example: "1E (hex) files were added" -> "30 files were added"

    Binary Conversion: Replace binary numbers with their decimal equivalents.

    - Example: "It has been 10 (bin) years" -> "It has been 2 years"

    Uppercase Conversion: Convert words to uppercase.

    - Example: "Ready, set, go (up) !" -> "Ready, set, GO !"

    Lowercase Conversion: Convert words to lowercase.

    - Example: "I should stop SHOUTING (low)" -> "I should stop shouting"

    Capitalize Conversion: Capitalize words.

    - Example: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge"

    Custom Case Conversion: Convert specified number of words to lowercase, uppercase, or capitalized.

    - Example: "This is so exciting (up, 2)" -> "This is SO EXCITING"

    Punctuation Formatting: Ensure proper spacing around punctuation marks and handle special cases like ellipses and consecutive punctuation marks.

    - Example: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!"

    Quote Handling: Place single quotation marks correctly around words or phrases.

    - Example: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'"

    Indefinite Article Correction: Change "a" to "an" before words starting with a vowel or "h".

    - Example: "There it was. A amazing rock!" -> "There it was. An amazing rock!"

### Contribution Guidelines

    Fork the repository.
    Create a feature branch (git checkout -b feature/MyFeature).
    Commit changes (git commit -am 'Add some feature').
    Push to the branch (git push origin feature/MyFeature).
    Create a new Pull Request.

### Testing

Test your modifications thoroughly, considering edge cases and expected behavior for each transformation and formatting rule.

### License

This project is licensed under the MIT License. See the LICENSE file for details.