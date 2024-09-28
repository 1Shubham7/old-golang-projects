# Test-Driven-Development

### This repo contains multiple projects on go testing. I have used testing and assertion packages like:

1. basic unit testing
2. Unit Testing using testify
3. Parameterized testing
4. Testing using Ginkgo and Gomega

## Some notes for Ginkgo

### Describe

- This is the top-level component of a Ginkgo test suite, and is used to group together a set of related tests. A Describe block typically describes a specific aspect of the system being tested, such as a package or module.

### Context

- Within a Describe block, you can use context blocks to further group related tests together. A Context block typically describes a specific set of conditions or inputs for the tests being run.

### It

- Finally, within a Context block, you can define individual test cases using It blocks. An It block describes a specific behavior or feature of the system being tested and contains the actual test code.