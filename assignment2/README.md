# Architect a solution which satisfies the following set of requirements:

- Create at least one implementation for each of the abstractions listed above.
- You should write a Fake implementation of the Persistence layer which can be used to read/write data into memory.
- The TodoDAO implementation(s) will be responsible for communicating with the persistence layer.
  - It should check for errors and retry the request in the event of failure.
    - How often it will attempt to reissue a failed request should be configurable.
    - It should only retry the request when the error returned is of type TemporaryError.
  - It should be able to return a cached instance of the Todo object.
    - We should be able to enable/disable caching via the environment variable.
    - It should be able to flush the appropriate TODO cache after a write operation has occurred.
- You are free to use libraries but your solution must include unit tests.
