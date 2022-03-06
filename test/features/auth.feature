Feature: Api auth validation
  Scenario: User without auth requests are declined
    Given User registers via OTP without auth header
    Then User request is unauthorized
