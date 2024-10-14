import os
import sys

def check_structure(directory):
    errors = []
    for root, dirs, files in os.walk(directory):
        for dir in dirs:
            if not dir.islower() and '_' not in dir:
                errors.append(f"Directory name should be snake_case: {os.path.join(root, dir)}")

        for file in files:
            if file.endswith('.go'):
                base_name = file[:-3]
                if not base_name.islower() or '_' not in base_name:
                    errors.append(f"File name should be snake_case: {os.path.join(root, file)}")

                test_file = f"{base_name}_test.go"
                if test_file not in files:
                    errors.append(f"Missing test file for: {os.path.join(root, file)}")

    return errors

if __name__ == "__main__":
    errors = check_structure(".")
    for error in errors:
        print(error)

    if errors:
        sys.exit(1)
    else:
        print("Structure check passed!")
