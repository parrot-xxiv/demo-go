# Todo
To make the password validation system more scalable and maintainable as you add more requirements (e.g., 20 or more), we can adopt a more **declarative** and **extensible** approach by organizing validation logic into a set of reusable components. This will keep the code clean, modular, and easy to extend in the future.

Here’s how you can approach it:

### Key Steps to Scale:
1. **Use a List of Validators**: Instead of having multiple individual checks in the `isValidPassword` function, you can define a slice of validation functions. Each validation function will return a specific error message if it fails.
2. **Allow Adding New Validators Dynamically**: This way, to add a new requirement, you just add a new function to the list of validators.
3. **Error Handling**: Collect all error messages and return them together, which is especially useful if you want to display multiple error messages at once.
4. **Custom Validator Types**: You can use a more generic approach where validators are passed as function pointers, making the system extensible for any future requirements.

### Refactored Scalable Solution

Here’s how you can refactor the code:

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Validator function type, each validator returns an error message if it fails.
type Validator func(password string) string

// Password Validators
func lengthValidator(password string) string {
	if len(password) < 8 {
		return "Password must be at least 8 characters long."
	}
	return ""
}

func uppercaseValidator(password string) string {
	for _, c := range password {
		if unicode.IsUpper(c) {
			return ""
		}
	}
	return "Password must contain at least one uppercase letter."
}

func specialCharValidator(password string) string {
	re := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	if !re.MatchString(password) {
		return "Password must include at least one special character."
	}
	return ""
}

func digitValidator(password string) string {
	for _, c := range password {
		if unicode.IsDigit(c) {
			return ""
		}
	}
	return "Password must contain at least one digit."
}

func spaceValidator(password string) string {
	if strings.Contains(password, " ") {
		return "Password must not contain spaces."
	}
	return ""
}

func commonPasswordValidator(password string) string {
	commonPasswords := []string{
		"password", "123456", "qwerty", "letmein", "welcome", "admin", "abc123",
	}
	for _, common := range commonPasswords {
		if password == common {
			return "Password is too common."
		}
	}
	return ""
}

func repeatingCharsValidator(password string) string {
	for i := 2; i < len(password); i++ {
		if password[i] == password[i-1] && password[i-1] == password[i-2] {
			return "Password must not have more than 2 repeated characters in a row."
		}
	}
	return ""
}

// Add more validators here...

// Function to validate the password
func validatePassword(password string, validators []Validator) []string {
	var errors []string
	for _, validator := range validators {
		if err := validator(password); err != "" {
			errors = append(errors, err)
		}
	}
	return errors
}

func main() {
	// List of validators (add new ones here as needed)
	validators := []Validator{
		lengthValidator,
		uppercaseValidator,
		specialCharValidator,
		digitValidator,
		spaceValidator,
		commonPasswordValidator,
		repeatingCharsValidator,
		// Add new validators here...
	}

	for {
		var password string
		fmt.Print("Enter password: ")
		fmt.Scanln(&password)

		// Validate the password
		errors := validatePassword(password, validators)

		if len(errors) == 0 {
			fmt.Println("Password is valid!")
			break
		} else {
			// Print all the error messages
			for _, err := range errors {
				fmt.Println(err)
			}
		}
	}
}
```

### Key Improvements:
1. **Validators as Functions**: Each validation rule is encapsulated as a function that returns an error message if the password fails the check. This allows you to easily add new validators or modify existing ones.
   
2. **Scalable Structure**: The `validators` slice holds all the validation functions, making it easy to add new ones. You just append a new validator function to this slice, and the validation system will automatically handle it.

3. **Collecting All Errors**: Instead of stopping at the first error, this implementation collects all the error messages and prints them at once. This is more user-friendly, especially when you have multiple requirements.

4. **Dynamic and Flexible**: You can dynamically add any number of validators to the list. For example, adding a new requirement like "must not contain the username" would just require creating a new function and adding it to the `validators` list.

### Adding More Validators:
To add new password validation requirements, simply create a new validator function and append it to the `validators` slice.

For instance:
- **Check if password contains a specific word:**
  ```go
  func noSpecificWordValidator(password string) string {
      if strings.Contains(password, "password") {
          return "Password must not contain the word 'password'."
      }
      return ""
  }
  ```

- **Check if the password matches a specific regex pattern:**
  ```go
  func regexPatternValidator(password string) string {
      re := regexp.MustCompile(`^[A-Za-z0-9!@#]*$`) // Example: Only letters, numbers, and special characters
      if !re.MatchString(password) {
          return "Password contains invalid characters."
      }
      return ""
  }
  ```

You can then add the new validator like this:
```go
validators = append(validators, noSpecificWordValidator)
validators = append(validators, regexPatternValidator)
```

### Example Output:

```text
Enter password: password123
Password must be at least 8 characters long.
Password must contain at least one uppercase letter.
Password must include at least one special character.

Enter password: Password123!
Password is valid!
```

### Advantages:
- **Easier to Extend**: To add a new requirement, simply write a new function and add it to the `validators` list. No need to modify the core logic.
- **Better Error Handling**: All errors are collected and displayed together, so the user gets clear feedback on all the issues with their password.
- **Modular**: Each validation function is independent, making the code modular and easier to maintain or update.
- **Flexible**: You can add as many validators as you need, without needing to change the core logic of your program.

### Conclusion:
This design allows you to easily scale your password validation system as your requirements grow. Each new condition is just a small function that fits into the broader validation system, making the code maintainable and scalable even as you add many more rules.

# Using package

Yes, **separating validators into a separate package** is a **great idea** and follows good software engineering practices. It makes your code **more modular**, **reusable**, and **maintainable**, especially as your codebase grows and you add more requirements. Having a dedicated package for password validation keeps the `main` function clean and focused on its core responsibility, which is interacting with the user and handling the flow.

### Benefits of Separating Validators into a Package:
1. **Modularity**: It keeps the password validation logic separate from the application logic, making it easier to modify or extend the validators without affecting other parts of your program.
2. **Reusability**: Once you have a dedicated package for password validation, you can reuse it across different projects or parts of your codebase.
3. **Maintainability**: If you need to add, remove, or modify validation logic, you can do so in one place. This is especially beneficial as the number of validation rules increases.
4. **Testing**: It’s easier to write unit tests for validators when they are in a separate package, as the validators will be independent of the rest of the application logic.

### How to Structure the Project

Here's how you can refactor your code into multiple packages:

### Project Structure:

```
myapp/
│
├── main.go           # Main application logic
├── password/         # Password validation package
│   ├── validators.go # Contains validation functions
│   └── validators_test.go # Unit tests for validators
└── go.mod            # Go modules file
```

### Step-by-Step Guide:

#### 1. Create the `password` package

- In your project directory, create a folder named `password`.
- Inside the `password` folder, create a file `validators.go` to define the validators.

#### `password/validators.go`:

```go
package password

import (
	"regexp"
	"strings"
	"unicode"
)

// Validator function type, each validator returns an error message if it fails.
type Validator func(password string) string

// Password Validators

func LengthValidator(password string) string {
	if len(password) < 8 {
		return "Password must be at least 8 characters long."
	}
	return ""
}

func UppercaseValidator(password string) string {
	for _, c := range password {
		if unicode.IsUpper(c) {
			return ""
		}
	}
	return "Password must contain at least one uppercase letter."
}

func SpecialCharValidator(password string) string {
	re := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	if !re.MatchString(password) {
		return "Password must include at least one special character."
	}
	return ""
}

func DigitValidator(password string) string {
	for _, c := range password {
		if unicode.IsDigit(c) {
			return ""
		}
	}
	return "Password must contain at least one digit."
}

func SpaceValidator(password string) string {
	if strings.Contains(password, " ") {
		return "Password must not contain spaces."
	}
	return ""
}

func CommonPasswordValidator(password string) string {
	commonPasswords := []string{
		"password", "123456", "qwerty", "letmein", "welcome", "admin", "abc123",
	}
	for _, common := range commonPasswords {
		if password == common {
			return "Password is too common."
		}
	}
	return ""
}

func RepeatingCharsValidator(password string) string {
	for i := 2; i < len(password); i++ {
		if password[i] == password[i-1] && password[i-1] == password[i-2] {
			return "Password must not have more than 2 repeated characters in a row."
		}
	}
	return ""
}

// Add more validators here...
```

#### 2. Use the `password` package in `main.go`

Now that your validators are separated into their own package, you can import and use them in the `main.go` file.

#### `main.go`:

```go
package main

import (
	"fmt"
	"myapp/password" // Import the password package
)

func validatePassword(password string, validators []password.Validator) []string {
	var errors []string
	for _, validator := range validators {
		if err := validator(password); err != "" {
			errors = append(errors, err)
		}
	}
	return errors
}

func main() {
	// List of validators (you can add more from the password package here)
	validators := []password.Validator{
		password.LengthValidator,
		password.UppercaseValidator,
		password.SpecialCharValidator,
		password.DigitValidator,
		password.SpaceValidator,
		password.CommonPasswordValidator,
		password.RepeatingCharsValidator,
		// Add more validators here...
	}

	for {
		var password string
		fmt.Print("Enter password: ")
		fmt.Scanln(&password)

		// Validate the password
		errors := validatePassword(password, validators)

		if len(errors) == 0 {
			fmt.Println("Password is valid!")
			break
		} else {
			// Print all the error messages
			for _, err := range errors {
				fmt.Println(err)
			}
		}
	}
}
```

### Explanation:
1. **The `password` package**: Contains all the validation functions. You can add new validators to `validators.go` as your requirements grow.
2. **The `main.go` file**: Now the `main.go` file is clean, and its only responsibility is orchestrating user input and validating the password using the validators from the `password` package.
3. **Reusability**: If you need to use this set of password validators in another project, you can simply import the `password` package and reuse the validation logic.

### 3. Unit Testing the Validators

A separate package also makes it easier to write unit tests for your validators. You can create a `validators_test.go` file inside the `password` package.

#### `password/validators_test.go`:

```go
package password

import "testing"

func TestValidators(t *testing.T) {
	// Test LengthValidator
	tests := []struct {
		password  string
		shouldFail bool
	}{
		{"short", true},
		{"longenough", false},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			result := LengthValidator(test.password)
			if test.shouldFail && result == "" {
				t.Errorf("Expected failure for password %q", test.password)
			}
			if !test.shouldFail && result != "" {
				t.Errorf("Expected success for password %q, but got: %v", test.password, result)
			}
		})
	}
	// Additional tests for other validators can be added here
}
```

You can run the tests using `go test` to ensure all your validators are working correctly.

### 4. Running the Program

To run your program, you can simply execute the `main.go` file:

```bash
go run main.go
```

### Benefits of Separating the Package:

1. **Modular Code**: You can easily add or modify validators without touching the main application logic.
2. **Reusability**: Validators can be reused across different projects or applications.
3. **Maintainability**: It’s easier to maintain and extend your password validation logic by keeping it in a dedicated package.
4. **Testability**: With the validators isolated in a separate package, you can write focused unit tests to ensure each validator works as expected.

### Conclusion:

Yes, separating validators into a separate package is not only a good idea but also follows best practices in software development. It makes your code **cleaner**, **more modular**, and **easier to maintain and extend** as you continue to add more validation rules. Additionally, it improves the **testability** of your code, allowing for more efficient unit testing.