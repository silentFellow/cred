#!/bin/sh

# creating vendor dependencies
# Check if go.mod or go.sum changed
git diff --cached --name-only | rg '^(go\.mod|go\.sum)$' > /dev/null
if [ $? -eq 0 ]; then
    echo "Detected changes in go.mod or go.sum. Running 'go mod tidy' and 'go mod vendor'..."

    # Run go mod tidy to clean up any unused dependencies
    go mod tidy
    
    # Run go mod vendor to update vendor directory with new dependencies
    go mod vendor

    # Optionally, you can check if the vendor directory was actually updated
    git diff --cached --name-only | rg '^vendor/' > /dev/null
    if [ $? -eq 0 ]; then
        echo "Vendor directory updated, staging changes."
        git add vendor/
    fi
fi
