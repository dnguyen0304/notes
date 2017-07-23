#include <iostream>
#include <string>

using namespace std;

bool IsUnique(string s)
{
    // The string type determines the character set size and therefore the
    // array length.
    int ascii_size = 128;
    // The default values of a boolean array within the scope of a function
    // are random. The array should be initialized.
    bool x[ascii_size] = { false };

    // Empty strings are always unique.
    if (s.size() == 0) {
        return true;
    }
    // Strings longer than the character set are never unique.
    if (s.size() > ascii_size) {
        return false;
    }

    for (int i = 0; i < s.size(); i++) {
        // Use a signed or unsigned char rather than a char because the ASCII
        // value is needed. Use a signed char rather than an unsigned char
        // because only the values between 0 and 127 are needed.
        signed char character = s[i];
        // The character has been seen at least once before.
        if (x[character]) {
            return false;
        // The character has not been seen before.
        } else {
            x[character] = true;
        }
    }

    return true;
}

void TestIsUnique(string s)
{
    string message = "\"" + s + "\" contains ";
    if (IsUnique(s)) {
        message += "only unique characters.";
    } else {
        message += "duplicated characters.";
    }
    cout << message + "\n";
}

int main()
{
    TestIsUnique("bar");
    TestIsUnique("foobar");

    TestIsUnique("");

    TestIsUnique(string(128 + 1, '.'));
}
