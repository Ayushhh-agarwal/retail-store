#!/usr/bin/env bash
set -x

echo "\n Clearing any existing mocks \n"
rm -rf github.com/razorpay/retail-store/internal/**/*_mock.go

mockgen -destination=internal/customers/repo_mock.go -package=customers github.com/razorpay/retail-store/internal/customers IRepo
mockgen -destination=internal/customers/core_mock.go -package=customers github.com/razorpay/retail-store/internal/customers ICore

mockgen -destination=internal/orders/repo_mock.go -package=orders github.com/razorpay/retail-store/internal/orders IRepo
mockgen -destination=internal/orders/core_mock.go -package=orders github.com/razorpay/retail-store/internal/orders ICore

mockgen -destination=internal/products/repo_mock.go -package=products github.com/razorpay/retail-store/internal/products IRepo
mockgen -destination=internal/products/core_mock.go -package=products github.com/razorpay/retail-store/internal/products ICore

mockgen -destination=internal/transactions/repo_mock.go -package=transactions github.com/razorpay/retail-store/internal/transactions IRepo
mockgen -destination=internal/transactions/core_mock.go -package=transactions github.com/razorpay/retail-store/internal/transactions ICore