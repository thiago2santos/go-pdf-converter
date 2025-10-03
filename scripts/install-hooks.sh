#!/bin/bash
# Install git hooks for the project

HOOK_DIR=".git/hooks"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "📦 Installing git hooks..."

# Create pre-commit hook
cat > "$HOOK_DIR/pre-commit" << 'EOF'
#!/bin/bash
# Pre-commit hook to run Go linting and formatting checks

set -e

echo "🔍 Running pre-commit checks..."

# Check if golangci-lint is installed
if ! command -v golangci-lint &> /dev/null; then
    echo "⚠️  golangci-lint not found. Installing..."
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

# Run go fmt check
echo "📝 Checking code formatting..."
UNFORMATTED=$(gofmt -l .)
if [ -n "$UNFORMATTED" ]; then
    echo "❌ The following files are not formatted:"
    echo "$UNFORMATTED"
    echo ""
    echo "Please run: go fmt ./..."
    exit 1
fi

# Run golangci-lint
echo "🔎 Running golangci-lint..."
if ! golangci-lint run --timeout=5m; then
    echo ""
    echo "❌ Linting failed. Please fix the issues above."
    echo "💡 Tip: You can skip this hook with 'git commit --no-verify'"
    exit 1
fi

echo "✅ All checks passed!"
exit 0
EOF

chmod +x "$HOOK_DIR/pre-commit"

echo "✅ Git hooks installed successfully!"
echo ""
echo "The following hooks are now active:"
echo "  • pre-commit: Runs go fmt and golangci-lint"
echo ""
echo "💡 To skip hooks, use: git commit --no-verify"

