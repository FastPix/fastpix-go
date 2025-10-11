#!/bin/bash

# FastPix Go SDK Test Runner
# This script runs the test suite with proper configuration

set -e

echo "🚀 FastPix Go SDK Test Runner"
echo "=============================="

# Check if required environment variables are set
if [ -z "$FASTPIX_USERNAME" ]; then
    echo "❌ Error: FASTPIX_USERNAME environment variable is not set"
    echo "Please set your FastPix username:"
    echo "export FASTPIX_USERNAME=\"your-access-token\""
    exit 1
fi

if [ -z "$FASTPIX_PASSWORD" ]; then
    echo "❌ Error: FASTPIX_PASSWORD environment variable is not set"
    echo "Please set your FastPix password:"
    echo "export FASTPIX_PASSWORD=\"your-secret-key\""
    exit 1
fi

echo "✅ Environment variables configured"
echo "Username: $FASTPIX_USERNAME"
echo "Base URL: ${FASTPIX_BASE_URL:-https://api.fastpix.io/v1}"

# Parse command line arguments
VERBOSE=false
COVERAGE=false
PARALLEL=1
TIMEOUT="10m"
TEST_PATTERN=""

while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--verbose)
            VERBOSE=true
            shift
            ;;
        -c|--coverage)
            COVERAGE=true
            shift
            ;;
        -p|--parallel)
            PARALLEL="$2"
            shift 2
            ;;
        -t|--timeout)
            TIMEOUT="$2"
            shift 2
            ;;
        --pattern)
            TEST_PATTERN="$2"
            shift 2
            ;;
        -h|--help)
            echo "Usage: $0 [OPTIONS]"
            echo ""
            echo "Options:"
            echo "  -v, --verbose     Run tests in verbose mode"
            echo "  -c, --coverage    Run tests with coverage report"
            echo "  -p, --parallel N  Run tests in parallel (default: 1)"
            echo "  -t, --timeout D   Set test timeout (default: 10m)"
            echo "  --pattern PATTERN Run tests matching pattern"
            echo "  -h, --help        Show this help message"
            echo ""
            echo "Examples:"
            echo "  $0                                    # Run all tests"
            echo "  $0 -v -c                            # Run with verbose output and coverage"
            echo "  $0 --pattern TestMedia               # Run only media tests"
            echo "  $0 -p 4 -t 5m                       # Run with 4 parallel workers, 5min timeout"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Build test command
TEST_CMD="go test"

if [ "$VERBOSE" = true ]; then
    TEST_CMD="$TEST_CMD -v"
fi

if [ "$COVERAGE" = true ]; then
    TEST_CMD="$TEST_CMD -cover"
fi

if [ "$PARALLEL" != "1" ]; then
    TEST_CMD="$TEST_CMD -parallel $PARALLEL"
fi

TEST_CMD="$TEST_CMD -timeout $TIMEOUT"

if [ -n "$TEST_PATTERN" ]; then
    TEST_CMD="$TEST_CMD -run $TEST_PATTERN"
fi

TEST_CMD="$TEST_CMD ./..."

echo ""
echo "🧪 Running tests..."
echo "Command: $TEST_CMD"
echo ""

# Run the tests
eval $TEST_CMD

TEST_EXIT_CODE=$?

echo ""
if [ $TEST_EXIT_CODE -eq 0 ]; then
    echo "✅ All tests passed!"
else
    echo "❌ Some tests failed (exit code: $TEST_EXIT_CODE)"
fi

echo ""
echo "📊 Test Summary:"
echo "- Username: $FASTPIX_USERNAME"
echo "- Base URL: ${FASTPIX_BASE_URL:-https://api.fastpix.io/v1}"
echo "- Parallel workers: $PARALLEL"
echo "- Timeout: $TIMEOUT"
echo "- Coverage: $COVERAGE"
echo "- Verbose: $VERBOSE"

exit $TEST_EXIT_CODE
