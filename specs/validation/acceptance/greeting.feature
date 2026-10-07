Feature: Greeting

  @story-1
  Rule: Calling /hello with a name returns a greeting that includes it

    Scenario: Greeting a named caller
      Given the greeter service is running
      When an API Consumer calls "GET /hello?name=Alice"
      Then the response is a JSON greeting containing "Alice"

  @story-2
  Rule: Calling /hello without a name still returns a generic greeting

    Scenario: Greeting with no name provided
      Given the greeter service is running
      When an API Consumer calls "GET /hello" with no "name" parameter
      Then the response is a JSON greeting that does not depend on any name

    Scenario: Greeting with an empty name provided
      Given the greeter service is running
      When an API Consumer calls "GET /hello?name="
      Then the response is a JSON greeting that does not depend on any name
